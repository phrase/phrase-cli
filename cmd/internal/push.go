package internal

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/antihax/optional"
	ct "github.com/daviddengcn/go-colortext"
	"github.com/jpillora/backoff"
	"github.com/phrase/phrase-cli/cmd/internal/paths"
	"github.com/phrase/phrase-cli/cmd/internal/placeholders"
	"github.com/phrase/phrase-cli/cmd/internal/print"
	"github.com/phrase/phrase-cli/cmd/internal/spinner"
	"github.com/phrase/phrase-go/v2"
)

type PushCommand struct {
	phrase.Config
	Wait               bool
	Branch             string
	UseLocalBranchName bool
}

func (cmd *PushCommand) Run() error {
	if cmd.Config.Debug {
		// suppresses content output
		cmd.Config.Debug = false
		Debug = true
	}
	Config = &cmd.Config

	client := newClient()

	sources, err := SourcesFromConfig(cmd.Config)
	if err != nil {
		return err
	}

	if err := sources.Validate(); err != nil {
		return err
	}

	formatMap, err := formatsByApiName(client)
	if err != nil {
		return fmt.Errorf("Error retrieving format list from Phrase: %s", err)
	}

	for _, source := range sources {
		formatName := source.GetFileFormat()
		if val, ok := formatMap[formatName]; ok {
			source.Format = val
		}

		if source.Format == nil {
			return fmt.Errorf("Format %q of source %q is not supported by Phrase!", formatName, source.File)
		}
	}

	projectsAffected := map[string]bool{}
	for _, source := range sources {
		projectsAffected[source.ProjectID] = true
	}

	branchName, err := usedBranchName(cmd.UseLocalBranchName, cmd.Branch)
	if err != nil {
		return err
	}
	cmd.Branch = branchName

	if cmd.Branch != "" {
		for projectId := range projectsAffected {
			_, _, err := client.BranchesApi.BranchShow(Auth, projectId, cmd.Branch, nil)
			if err != nil {
				if useLocalBranchName(cmd.UseLocalBranchName) {
					printCreateBranchQuestion(cmd.Branch)
					text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

					if !isYes(strings.TrimSpace(text)) {
						return nil
					}
				}

				branchParams := &phrase.BranchCreateParameters{Name: cmd.Branch}
				branch, _, err := client.BranchesApi.BranchCreate(Auth, projectId, *branchParams, nil)
				if err != nil {
					return err
				}

				fmt.Println()

				taskResult := make(chan string, 1)
				taskErr := make(chan error, 1)

				fmt.Printf("Waiting for branch %s is created!", branch.Name)
				spinner.While(func() {
					branchCreateResult, err := getBranchCreateResult(client, projectId, &branch)
					taskResult <- branchCreateResult
					taskErr <- err
				})
				fmt.Println()

				if err := <-taskErr; err != nil {
					return err
				}

				switch <-taskResult {
				case "success":
					print.Success("Successfully created branch %s", branch.Name)
				case "error":
					print.Failure("There was an error creating branch %s.", branch.Name)
				}
			}
		}
	}

	projectIdToLocales, err := LocalesForProjects(client, sources, cmd.Branch)
	if err != nil {
		return err
	}
	for _, source := range sources {
		val, ok := projectIdToLocales[LocaleCacheKey{source.ProjectID, cmd.Branch}]
		if ok {
			source.RemoteLocales = val
		}
	}

	for _, source := range sources {
		err := source.Push(client, cmd.Wait, cmd.Branch)
		if err != nil {
			return err
		}
	}

	return nil
}

func (source *Source) Push(client *phrase.APIClient, waitForResults bool, branch string) error {
	localeFiles, err := source.LocaleFiles()
	if err != nil {
		return err
	}

	for _, localeFile := range localeFiles {
		fmt.Printf("Uploading %s... ", localeFile.RelPath())

		if localeFile.shouldCreateLocale(source, branch) {
			localeDetails, err := source.createLocale(client, localeFile, branch)
			if err == nil {
				localeFile.ID = localeDetails.Id
				localeFile.Code = localeDetails.Code
				localeFile.Name = localeDetails.Name
			} else {
				fmt.Printf("failed to create locale: %s\n", err)
				continue
			}
		}

		upload, err := source.uploadFile(client, localeFile, branch)
		if err != nil {
			return err
		}

		if waitForResults {
			fmt.Println()

			taskResult := make(chan string, 1)
			taskErr := make(chan error, 1)

			fmt.Printf("Upload Id: %s, filename: %s succeeded. Waiting for your file to be processed... ", upload.Id, upload.Filename)
			spinner.While(func() {
				result, err := getUploadResult(client, source.ProjectID, upload, branch)
				taskResult <- result
				taskErr <- err
			})
			fmt.Println()

			if err := <-taskErr; err != nil {
				return err
			}

			switch <-taskResult {
			case "success":
				print.Success("Successfully uploaded and processed %s.", localeFile.RelPath())
			case "error":
				print.Failure("There was an error processing %s. Your changes were not saved online.", localeFile.RelPath())
			}
		} else {
			fmt.Println("done!")
			fmt.Printf("Check upload Id: %s, filename: %s for information about processing results.\n", upload.Id, upload.Filename)
		}

		if Debug {
			fmt.Fprintln(os.Stderr, strings.Repeat("-", 10))
		}
	}

	return nil
}

func formatsByApiName(client *phrase.APIClient) (map[string]*phrase.Format, error) {
	formats, _, err := client.FormatsApi.FormatsList(Auth, &phrase.FormatsListOpts{})
	if err != nil {
		return nil, err
	}
	formatMap := map[string]*phrase.Format{}
	for _, format := range formats {
		formatMap[format.ApiName] = &format
	}
	return formatMap, nil
}

// Return all locale files from disk that match the source pattern.
func (source *Source) LocaleFiles() (LocaleFiles, error) {
	sourcePattern := toOsSeparator(source.File)
	filePaths, err := paths.Glob(placeholders.ToGlobbingPattern(sourcePattern))
	if err != nil {
		return nil, err
	}

	var localeFiles LocaleFiles
	for _, path := range filePaths {
		if paths.IsPhraseAppYmlConfig(path) {
			continue
		}

		localeFile := new(LocaleFile)
		localeFile.fillFromPath(path, sourcePattern)

		localeFile.Path, err = filepath.Abs(path)
		if err != nil {
			return nil, err
		}

		locale := source.getRemoteLocaleForLocaleFile(localeFile)
		// TODO: sinnvoll?
		if locale != nil {
			localeFile.ExistsRemote = true
			localeFile.Code = locale.Code
			localeFile.Name = locale.Name
			localeFile.ID = locale.Id
		}

		if Debug {
			fmt.Printf(
				"Code:%q, Name:%q, Id:%q, Tag:%q\n",
				localeFile.Code, localeFile.Name, localeFile.ID, localeFile.Tag,
			)
		}

		localeFiles = append(localeFiles, localeFile)
	}

	if len(localeFiles) == 0 {
		abs, err := filepath.Abs(source.File)
		if err != nil {
			abs = source.File
		}
		return nil, fmt.Errorf("Could not find any files on your system that matches: '%s'", abs)
	}

	return localeFiles, nil
}

func (source *Source) getRemoteLocaleForLocaleFile(localeFile *LocaleFile) *phrase.Locale {
	candidates := source.RemoteLocales

	filterApplied := false

	filter := func(cands []*phrase.Locale, preCond string, pred func(cand *phrase.Locale) bool) []*phrase.Locale {
		if preCond == "" {
			return cands
		}
		filterApplied = true
		tmpCands := []*phrase.Locale{}
		for _, cand := range cands {
			if pred(cand) {
				tmpCands = append(tmpCands, cand)
			}
		}
		return tmpCands
	}

	localeName := source.replacePlaceholderInParams(localeFile)
	if localeName != "" {
		// This means the name can contain the value specified in LocaleId, with
		// `<locale_code>` being substituted by the value of the currently handled
		// localeFile (like push only locales with name `en-US`).
		candidates = filter(candidates, localeName, func(cand *phrase.Locale) bool {
			return strings.Contains(cand.Name, localeName)
		})
	} else {
		localeId := source.GetLocaleID()
		candidates = filter(candidates, localeId, func(cand *phrase.Locale) bool {
			return cand.Name == localeId || cand.Id == localeId
		})
	}

	candidates = filter(candidates, localeFile.Name, func(cand *phrase.Locale) bool {
		return cand.Name == localeFile.Name
	})

	candidates = filter(candidates, localeFile.Code, func(cand *phrase.Locale) bool {
		return cand.Code == localeFile.Code
	})

	// If no filter was applied the candidates list still contains all remote
	// locales, while actually nothing matches.
	if !filterApplied {
		return nil
	}

	switch len(candidates) {
	case 0:
		return nil
	case 1:
		return candidates[0]
	default:
		// TODO I guess this should return an error, as this is a problem.
		return candidates[0]
	}
}

func toOsSeparator(pattern string) string {
	if filepath.Separator == '/' {
		return strings.ReplaceAll(pattern, "\\", "/")
	} else {
		return filepath.FromSlash(pattern)
	}
}

func (localeFile *LocaleFile) fillFromPath(path, pattern string) {
	pathStart, patternStart, pathEnd, patternEnd, err := paths.SplitAtDirGlobOperator(path, pattern)
	if err != nil {
		print.Error(err)
		return
	}

	fillFrom := func(path, pattern string) {
		params, err := placeholders.Resolve(path, pattern)
		if err != nil {
			print.Error(err)
			return
		}

		for placeholder, value := range params {
			switch placeholder {
			case "locale_code":
				localeFile.Code = value
			case "locale_name":
				localeFile.Name = value
			case "tag":
				localeFile.Tag = value
			}
		}
	}

	fillFrom(pathStart, patternStart)
	fillFrom(pathEnd, patternEnd)
}

func (localeFile *LocaleFile) shouldCreateLocale(source *Source, branch string) bool {
	if localeFile.ExistsRemote {
		return false
	}

	if *source.Format.IncludesLocaleInformation {
		return false
	}

	// we could not find an existing locale in Phrase
	// if a locale_name or locale_code was provided by the placeholder logic
	// we assume that it should be created
	// every other source should be uploaded and validated in uploads#create
	return (localeFile.Name != "" || localeFile.Code != "")
}

func getUploadResult(client *phrase.APIClient, projectId string, upload *phrase.Upload, branch string) (result string, err error) {
	b := &backoff.Backoff{
		Min:    500 * time.Millisecond,
		Max:    10 * time.Second,
		Factor: 2,
		Jitter: true,
	}

	for ; result != "success" && result != "error"; result = upload.State {
		time.Sleep(b.Duration())
		uploadShowOpts := phrase.UploadShowOpts{
			Branch: optional.NewString(branch),
		}
		uploadhDetails, _, err := client.UploadsApi.UploadShow(Auth, projectId, upload.Id, &uploadShowOpts)
		upload = &uploadhDetails
		if err != nil {
			break
		}
	}

	return
}

func getBranchCreateResult(client *phrase.APIClient, projectId string, branch *phrase.Branch) (result string, err error) {
	b := &backoff.Backoff{
		Min:    500 * time.Millisecond,
		Max:    10 * time.Second,
		Factor: 2,
		Jitter: true,
	}

	for ; result != "success" && result != "error"; result = branch.State {
		time.Sleep(b.Duration())
		branchDetails, _, err := client.BranchesApi.BranchShow(Auth, projectId, branch.Name, nil)
		branch = &branchDetails
		if err != nil {
			break
		}
	}

	return
}

func printCreateBranchQuestion(branch string) {
	fmt.Printf("\nYou have currently checked out the branch '")
	ct.ChangeColor(ct.Green, false, ct.None, false)
	fmt.Printf("%s", branch)
	ct.ResetColor()
	fmt.Printf("'.\nThere currently is no branch in Phrase with this name.\n\n")
	fmt.Printf("Should we create a new branch in Phrase with the same name and push to it? [y/N]: ")
}

func isYes(text string) bool {
	return text == "y" ||
		text == "Y" ||
		text == "yes" ||
		text == "Yes" ||
		text == "YES"
}
