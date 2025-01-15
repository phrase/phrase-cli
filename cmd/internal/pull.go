package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/phrase/phrase-cli/cmd/internal/paths"
	"github.com/phrase/phrase-cli/cmd/internal/placeholders"
	"github.com/phrase/phrase-cli/cmd/internal/print"

	"github.com/antihax/optional"
	"github.com/phrase/phrase-go/v4"
)

const (
	timeoutInMinutes = 30 * time.Minute
	asyncWaitTime    = 5 * time.Second
	asyncRetryCount  = 360 // 30 minutes
)

var Config *phrase.Config

type PullCommand struct {
	phrase.Config
	Branch             string
	UseLocalBranchName bool
	Async              bool
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

	localesCache, err := GetLocalesCache(client, targets, cmd.Branch)
	if err != nil {
		return err
	}

	for _, target := range targets {
		if cmd.Branch != "" {
			target.Params.Branch = optional.NewString(cmd.Branch)
		}

		val, ok := localesCache[LocalesCacheKey{target.ProjectID, target.GetBranch()}]
		if !ok || len(val) == 0 {
			if cmd.Branch != "" {
				return fmt.Errorf("Branch '%s' does not exist in project '%s'", cmd.Branch, target.ProjectID)
			}
			return fmt.Errorf("Could not find any locales for project %q", target.ProjectID)
		}
		target.RemoteLocales = val
	}

	for _, target := range targets {
		err := target.Pull(client, cmd.Async)
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

func (target *Target) Pull(client *phrase.APIClient, async bool) error {
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

		err = target.DownloadAndWriteToFile(client, localeFile, async)
		if err != nil {
			if openapiError, ok := err.(phrase.GenericOpenAPIError); ok {
				print.Warn("API response: %s", openapiError.Body())
			}
			return fmt.Errorf("%s for %s", err, localeFile.Path)
		} else {
			print.Success("Downloaded %s to %s", localeFile.Message(), localeFile.RelPath())
		}
		debugFprintln(strings.Repeat("-", 10))
	}

	return nil
}

func (target *Target) DownloadAndWriteToFile(client *phrase.APIClient, localeFile *LocaleFile, async bool) error {
	localVarOptionals := phrase.LocaleDownloadOpts{}

	if target.Params != nil {
		localVarOptionals = target.Params.LocaleDownloadOpts
		translationKeyPrefix, err := placeholders.ResolveTranslationKeyPrefix(target.Params.TranslationKeyPrefix, localeFile.Path)
		if err != nil {
			return err
		}
		localVarOptionals.TranslationKeyPrefix = translationKeyPrefix
	}

	if localVarOptionals.FileFormat.Value() == "" {
		localVarOptionals.FileFormat = optional.NewString(localeFile.FileFormat)
	}

	if localeFile.Tag != "" {
		localVarOptionals.Tags = optional.NewString(localeFile.Tag)
		localVarOptionals.Tag = optional.EmptyString()
	}

	debugFprintln("Target file pattern:", target.File)
	debugFprintln("Actual file path", localeFile.Path)
	debugFprintln("LocaleID", localeFile.ID)
	debugFprintln("ProjectID", target.ProjectID)
	debugFprintln("FileFormat", localVarOptionals.FileFormat)
	debugFprintln("ConvertEmoji", localVarOptionals.ConvertEmoji)
	debugFprintln("IncludeEmptyTranslations", localVarOptionals.IncludeEmptyTranslations)
	debugFprintln("KeepNotranslateTags", localVarOptionals.KeepNotranslateTags)
	debugFprintln("Tags", localVarOptionals.Tags)
	debugFprintln("Branch", localVarOptionals.Branch)
	debugFprintln("FormatOptions", localVarOptionals.FormatOptions)
	debugFprintln("TranslationKeyPrefix", localVarOptionals.TranslationKeyPrefix)

	if async {
		return target.downloadAsynchronously(client, localeFile, localVarOptionals)
	} else {
		return target.downloadSynchronously(client, localeFile, localVarOptionals)
	}
}

func (target *Target) downloadAsynchronously(client *phrase.APIClient, localeFile *LocaleFile, downloadOpts phrase.LocaleDownloadOpts) error {
	localeDownloadCreateParams := asyncDownloadParams(downloadOpts)

	localVarOptionals := phrase.LocaleDownloadCreateOpts{}
	debugFprintln("Initiating async download...")
	asyncDownload, _, err := client.LocaleDownloadsApi.LocaleDownloadCreate(Auth, target.ProjectID, localeFile.ID, localeDownloadCreateParams, &localVarOptionals)
	if err != nil {
		return err
	}

	for i := 0; asyncDownload.Status == "processing"; i++ {
		debugFprintln("Waiting for the files to be exported...")
		time.Sleep(asyncWaitTime)
		debugFprintln("Checking if the download is ready...")
		localVarOptionals := phrase.LocaleDownloadShowOpts{}
		asyncDownload, _, err = client.LocaleDownloadsApi.LocaleDownloadShow(Auth, target.ProjectID, localeFile.ID, asyncDownload.Id, &localVarOptionals)
		if err != nil {
			return err
		}
		if i > asyncRetryCount {
			return fmt.Errorf("download is taking too long")
		}
	}
	if asyncDownload.Status == "completed" {
		return downloadExportedLocale(asyncDownload.Result.Url, localeFile.Path)
	}
	return fmt.Errorf("download failed: %s", asyncDownload.Error)
}

func (target *Target) downloadSynchronously(client *phrase.APIClient, localeFile *LocaleFile, downloadOpts phrase.LocaleDownloadOpts) error {
	file, response, err := client.LocalesApi.LocaleDownload(Auth, target.ProjectID, localeFile.ID, &downloadOpts)
	if err != nil {
		if response.Rate.Remaining == 0 {
			waitForRateLimit(response.Rate)
			file, _, err = client.LocalesApi.LocaleDownload(Auth, target.ProjectID, localeFile.ID, &downloadOpts)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return copyToDestination(file, localeFile.Path)
}

func copyToDestination(file *os.File, path string) error {
	destFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer destFile.Close()
	if file != nil {
		defer file.Close()
		_, err = io.Copy(destFile, file)
		return err
	}
	return nil
}

func downloadExportedLocale(url string, localName string) error {
	debugFprintln("Downloading file from ", url)
	file, err := os.Create(localName)
	if err != nil {
		return err
	}
	defer file.Close()
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	request.Header.Set("Authorization", "Bearer "+Config.Credentials.Token)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	io.Copy(file, response.Body)
	return nil
}

func asyncDownloadParams(localVarOptionals phrase.LocaleDownloadOpts) phrase.LocaleDownloadCreateParameters {
	sourceFields := reflect.VisibleFields(reflect.TypeOf(localVarOptionals))
	localeDownloadCreateParams := phrase.LocaleDownloadCreateParameters{}
	targetFields := reflect.VisibleFields(reflect.TypeOf(localeDownloadCreateParams))

	for i, targetField := range targetFields {
		for _, sourceField := range sourceFields {
			if targetField.Name == sourceField.Name {
				sourceValue := reflect.ValueOf(localVarOptionals).FieldByName(sourceField.Name)
				if sourceValue.MethodByName("IsSet").Call([]reflect.Value{})[0].Interface().(bool) {
					targetValue := reflect.ValueOf(&localeDownloadCreateParams).Elem().Field(i)
					sourceOptionalValue := sourceValue.MethodByName("Value").Call([]reflect.Value{})[0]
					switch sourceField.Type {
					case reflect.TypeOf((*optional.String)(nil)).Elem():
						targetValue.Set(sourceOptionalValue)
					case reflect.TypeOf((*optional.Bool)(nil)).Elem():
						boolValue := sourceOptionalValue.Interface().(bool)
						targetValue.Set(reflect.ValueOf(&boolValue))
					case reflect.TypeOf((*optional.Interface)(nil)).Elem():
						jsonValue, _ := json.Marshal(sourceOptionalValue.Interface())
						json.Unmarshal(jsonValue, targetValue.Addr().Interface())
					}
				}
				break
			}
		}
	}
	return localeDownloadCreateParams
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

func debugFprintln(a ...any) {
	if Debug {
		fmt.Fprintln(os.Stderr, a...)
	}
}
