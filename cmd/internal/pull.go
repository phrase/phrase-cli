package internal

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/phrase/phrase-cli/cmd/internal/paths"
	"github.com/phrase/phrase-cli/cmd/internal/placeholders"
	"github.com/phrase/phrase-cli/cmd/internal/print"

	"github.com/antihax/optional"
	"github.com/phrase/phrase-go/v2"
)

const (
	timeoutInMinutes = 30 * time.Minute
)

var Config *phrase.Config

type PullCommand struct {
	phrase.Config
	Branch             string
	UseLocalBranchName bool
}

var Auth context.Context

func (cmd *PullCommand) Run(config *phrase.Config) error {
	Config = config

	if Config.Debug {
		// suppresses content output
		Config.Debug = false
		Debug = true
	}
	client := newClient()

	targets, err := TargetsFromConfig(*Config)
	if err != nil {
		return err
	}

	branchName, err := usedBranchName(cmd.UseLocalBranchName, cmd.Branch)
	if err != nil {
		return err
	}
	cmd.Branch = branchName

	projectIdToLocales, err := LocalesForProjects(client, targets, cmd.Branch)
	if err != nil {
		return err
	}

	for _, target := range targets {
		val, ok := projectIdToLocales[LocaleCacheKey{target.ProjectID, cmd.Branch}]
		if !ok || len(val) == 0 {
			if cmd.Branch != "" {
				continue
			}
			return fmt.Errorf("Could not find any locales for project %q", target.ProjectID)
		}
		target.RemoteLocales = val
	}

	for _, target := range targets {
		err := target.Pull(client, cmd.Branch)
		if err != nil {
			return err
		}
	}

	return nil
}

func newClient() *phrase.APIClient {
	Auth = context.WithValue(context.Background(), phrase.ContextAPIKey, phrase.APIKey{
		Key:    Config.Credentials.Token,
		Prefix: "token",
	})

	cfg := phrase.NewConfiguration()
	cfg.SetUserAgent(Config.UserAgent)
	if Config.Credentials.Host != "" {
		cfg.BasePath = Config.Credentials.Host
	}
	return phrase.NewAPIClient(cfg)
}

type PullParams struct {
	phrase.LocaleDownloadOpts `json:",squash" mapstructure:",squash"`
	LocaleID                  string `json:"locale_id"`
}

func (target *Target) Pull(client *phrase.APIClient, branch string) error {
	if err := target.CheckPreconditions(); err != nil {
		return err
	}

	localeFiles, err := target.LocaleFiles()
	if err != nil {
		return err
	}

	startedAt := time.Now()
	for _, localeFile := range localeFiles {
		if time.Since(startedAt) >= timeoutInMinutes {
			return fmt.Errorf("Timeout of %d minutes exceeded", timeoutInMinutes)
		}

		err := createFile(localeFile.Path)
		if err != nil {
			return err
		}

		err = target.DownloadAndWriteToFile(client, localeFile, branch)
		if err != nil {
			return fmt.Errorf("%s for %s", err, localeFile.Path)
		} else {
			print.Success("Downloaded %s to %s", localeFile.Message(), localeFile.RelPath())
		}
		if Debug {
			fmt.Fprintln(os.Stderr, strings.Repeat("-", 10))
		}
	}

	return nil
}

func (target *Target) DownloadAndWriteToFile(client *phrase.APIClient, localeFile *LocaleFile, branch string) error {
	localVarOptionals := phrase.LocaleDownloadOpts{}

	if target.Params != nil {
		localVarOptionals = target.Params.LocaleDownloadOpts
	}

	if localVarOptionals.FileFormat.Value() == "" {
		localVarOptionals.FileFormat = optional.NewString(localeFile.FileFormat)
	}

	if branch != "" {
		localVarOptionals.Branch = optional.NewString(branch)
	}

	if localeFile.Tag != "" {
		localVarOptionals.Tags = optional.NewString(localeFile.Tag)
		localVarOptionals.Tag = optional.EmptyString()
	}

	if Debug {
		fmt.Fprintln(os.Stderr, "Target file pattern:", target.File)
		fmt.Fprintln(os.Stderr, "Actual file path", localeFile.Path)
		fmt.Fprintln(os.Stderr, "LocaleID", localeFile.ID)
		fmt.Fprintln(os.Stderr, "ProjectID", target.ProjectID)
		fmt.Fprintln(os.Stderr, "FileFormat", localVarOptionals.FileFormat)
		fmt.Fprintln(os.Stderr, "ConvertEmoji", localVarOptionals.ConvertEmoji)
		fmt.Fprintln(os.Stderr, "IncludeEmptyTranslations", localVarOptionals.IncludeEmptyTranslations)
		fmt.Fprintln(os.Stderr, "KeepNotranslateTags", localVarOptionals.KeepNotranslateTags)
		fmt.Fprintln(os.Stderr, "Tags", localVarOptionals.Tags)
		fmt.Fprintln(os.Stderr, "Branch", localVarOptionals.Branch)
		fmt.Fprintln(os.Stderr, "FormatOptions", localVarOptionals.FormatOptions)
	}

	file, response, err := client.LocalesApi.LocaleDownload(Auth, target.ProjectID, localeFile.ID, &localVarOptionals)
	if err != nil {
		if response.Rate.Remaining == 0 {
			waitForRateLimit(response.Rate)
			file, _, err = client.LocalesApi.LocaleDownload(Auth, target.ProjectID, localeFile.ID, &localVarOptionals)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	var data []byte
	if file != nil {
		data, err = ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		file.Close()
		os.Remove(file.Name())
	}

	err = ioutil.WriteFile(localeFile.Path, data, 0644)
	return err
}

func (target *Target) LocaleFiles() (LocaleFiles, error) {
	files := []*LocaleFile{}

	if target.GetLocaleID() != "" {
		// a specific locale was requested
		remoteLocale, err := target.localeForRemote()
		if err != nil {
			return nil, err
		}

		localeFiles, err := target.createLocaleFiles(remoteLocale)
		if err != nil {
			return nil, err
		}

		files = append(files, localeFiles...)
	} else if placeholders.ContainsLocalePlaceholder(target.File) {
		// multiple locales were requested
		for _, remoteLocale := range target.RemoteLocales {
			localesFiles, err := target.createLocaleFiles(remoteLocale)
			if err != nil {
				return nil, err
			}

			files = append(files, localesFiles...)
		}
	} else {
		// no local files match remote locale
		return nil, fmt.Errorf("could not find any files on your system that matches the locales for project %q", target.ProjectID)
	}

	return files, nil
}

func (target *Target) createLocaleFiles(remoteLocale *phrase.Locale) (LocaleFiles, error) {
	files := []*LocaleFile{}
	tags := target.GetTags()
	if len(tags) > 0 && placeholders.ContainsTagPlaceholder(target.File) {
		for _, tag := range tags {
			localeFile, err := createLocaleFile(target, remoteLocale, tag)
			if err != nil {
				return nil, err
			}

			files = append(files, localeFile)
		}
	} else {
		localeFile, err := createLocaleFile(target, remoteLocale, "")
		if err != nil {
			return nil, err
		}

		files = append(files, localeFile)
	}
	return files, nil
}

func waitForRateLimit(rate phrase.Rate) {
	endTime := rate.Reset.Time.Add(5 * time.Second)
	duration := time.Until(endTime)
	message := "Rate limit exceeded. Download will resume in %d seconds"
	seconds := int64(time.Until(endTime).Seconds())
	print.Warn(message, seconds)
	counter := int64(1)
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		counter++
		seconds = int64(time.Until(endTime).Seconds())
		print.Warn(message, seconds)
		if counter > int64(duration/time.Second) {
			ticker.Stop()
			break
		}
	}
}

func createLocaleFile(target *Target, remoteLocale *phrase.Locale, tag string) (*LocaleFile, error) {
	localeFile := &LocaleFile{
		Name:       remoteLocale.Name,
		ID:         remoteLocale.Id,
		Code:       remoteLocale.Code,
		Tag:        tag,
		FileFormat: target.GetFormat(),
		Path:       target.File,
	}

	absPath, err := target.ReplacePlaceholders(localeFile)
	if err != nil {
		return nil, err
	}

	localeFile.Path = absPath
	return localeFile, nil
}

func createFile(path string) error {
	err := paths.Exists(path)
	if err != nil {
		absDir := filepath.Dir(path)
		err := paths.Exists(absDir)
		if err != nil {
			_ = os.MkdirAll(absDir, 0755)
		}

		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
	}

	return nil
}
