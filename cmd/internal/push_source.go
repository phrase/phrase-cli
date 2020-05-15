package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/antihax/optional"
	"github.com/phrase/phrase-cli/cmd/internal/paths"
	"github.com/phrase/phrase-go"
	"github.com/spf13/viper"
)

func SourcesFromConfig(config phrase.Config) (Sources, error) {
	if config.Sources == nil || len(config.Sources) == 0 {
		return nil, fmt.Errorf("no sources for upload specified")
	}

	tmp := struct {
		Sources Sources
	}{}

	sources := viper.New()
	sources.SetConfigType("yaml")
	err := sources.ReadConfig(bytes.NewReader(config.Sources))

	if err != nil {
		return nil, err
	}

	err = sources.UnmarshalExact(&tmp, ViperStructTag())
	if err != nil {
		return nil, err
	}

	srcs := tmp.Sources

	projectId := config.DefaultProjectID
	fileFormat := config.DefaultFileFormat

	validSources := []*Source{}
	for _, source := range srcs {
		if source == nil {
			continue
		}
		if source.ProjectID == "" {
			source.ProjectID = projectId
		}
		if source.Params == nil {
			source.Params = new(phrase.UploadCreateOpts)
		}

		if !source.Params.FileFormat.IsSet() {
			switch {
			case source.FileFormat != "":
				source.Params.FileFormat = optional.NewString(source.FileFormat)
			case fileFormat != "":
				source.Params.FileFormat = optional.NewString(fileFormat)
			}
		}
		validSources = append(validSources, source)
	}

	if len(validSources) <= 0 {
		return nil, fmt.Errorf("no sources could be identified! Refine the sources list in your config")
	}

	return validSources, nil
}

type Sources []*Source

func (sources Sources) Validate() error {
	for _, source := range sources {
		if err := source.CheckPreconditions(); err != nil {
			return err
		}
	}
	return nil
}

type Source struct {
	File        string                   `json:"file"`
	ProjectID   string                   `json:"project_id"`
	Branch      string                   `json:"branch"`
	AccessToken string                   `json:"access_token"`
	FileFormat  string                   `json:"file_format"`
	Params      *phrase.UploadCreateOpts `json:"params,omitempty"`

	RemoteLocales []*phrase.Locale
	Format        *phrase.Format
}

func (source *Source) GetLocaleID() string {
	if source.Params != nil && !source.Params.LocaleId.IsSet() {
		return source.Params.LocaleId.Value()
	}
	return ""
}

func (source *Source) GetFileFormat() string {
	if source.Params != nil && source.Params.FileFormat.IsSet() {
		return source.Params.FileFormat.Value()
	}
	if source.FileFormat != "" {
		return source.FileFormat
	}
	return ""
}

func (source *Source) CheckPreconditions() error {
	if err := paths.Validate(source.File, source.FileFormat, ""); err != nil {
		return err
	}

	duplicatedPlaceholders := []string{}
	for _, name := range []string{"<locale_name>", "<locale_code>", "<tag>"} {
		if strings.Count(source.File, name) > 1 {
			duplicatedPlaceholders = append(duplicatedPlaceholders, name)
		}
	}

	starCount := strings.Count(source.File, "*")
	recCount := strings.Count(source.File, "**")

	// starCount contains the `**` so that must be taken into account.
	if starCount-(recCount*2) > 1 {
		duplicatedPlaceholders = append(duplicatedPlaceholders, "*")
	}

	if recCount > 1 {
		duplicatedPlaceholders = append(duplicatedPlaceholders, "**")
	}

	if len(duplicatedPlaceholders) > 0 {
		dups := strings.Join(duplicatedPlaceholders, ", ")
		return fmt.Errorf(fmt.Sprintf("%s can only occur once in a file pattern!", dups))
	}

	return nil
}

func (sources Sources) ProjectIds() []string {
	projectIds := []string{}
	for _, source := range sources {
		projectIds = append(projectIds, source.ProjectID)
	}
	return projectIds
}
func (source *Source) uploadFile(client *phrase.APIClient, localeFile *LocaleFile, branch string) (*phrase.Upload, error) {
	if Debug {
		fmt.Fprintln(os.Stdout, "Source file pattern:", source.File)
		fmt.Fprintln(os.Stdout, "Actual file location:", localeFile.Path)
	}

	params := new(phrase.UploadCreateOpts)
	*params = *source.Params

	var err error
	file, err := os.Open(localeFile.Path)
	params.File = optional.NewInterface(file)

	if err != nil {
		return nil, err
	}

	if params.LocaleId.IsSet() {
		switch {
		case localeFile.ID != "":
			params.LocaleId = optional.NewString(localeFile.ID)
		case localeFile.Code != "":
			params.LocaleId = optional.NewString(localeFile.Code)
		}
	}
	if localeFile.Tag != "" {
		var v string
		if params.Tags.IsSet() {
			v = params.Tags.Value() + ","
		}
		v += localeFile.Tag
		params.Tags = optional.NewString(v)
	}

	if branch != "" {
		params.Branch = optional.NewString(branch)
	}

	var upload *phrase.Upload
	data, _, err := client.UploadsApi.UploadCreate(Auth, source.ProjectID, params)

	if err := json.Unmarshal(data, &upload); err != nil {
		return nil, err
	}

	return upload, err
}

func (source *Source) createLocale(client *phrase.APIClient, localeFile *LocaleFile, branch string) (*phrase.LocaleDetails, error) {
	localeDetails, found, err := source.getLocaleIfExist(client, localeFile, branch)
	if err != nil {
		return nil, err
	} else if found {
		return localeDetails, nil
	}

	localeParams := new(phrase.LocaleCreateParameters)

	if localeFile.Name != "" {
		localeParams.Name = localeFile.Name
	} else if localeFile.Code != "" {
		localeParams.Name = localeFile.Code
	}

	if localeFile.Code == "" {
		localeFile.Code = localeFile.Name
	}

	localeName := source.replacePlaceholderInParams(localeFile)
	if localeName != "" && localeName != localeFile.Code {
		localeParams.Name = localeName
	}

	if localeFile.Code != "" {
		localeParams.Code = localeFile.Code
	}

	if branch != "" {
		localeParams.Branch = branch
	}

	var localeDetail *phrase.LocaleDetails
	data, _, err := client.LocalesApi.LocaleCreate(Auth, source.ProjectID, *localeParams, &phrase.LocaleCreateOpts{})

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &localeDetail); err != nil {
		return nil, err
	}

	return localeDetails, nil
}

func (source *Source) getLocaleIfExist(client *phrase.APIClient, localeFile *LocaleFile, branch string) (*phrase.LocaleDetails, bool, error) {
	identifier := localeIdentifier(source, localeFile)
	if identifier == "" {
		return nil, false, nil
	}

	localeShowParams := &phrase.LocaleShowOpts{
		Branch: optional.NewString(branch),
	}
	localeDetail, _, err := client.LocalesApi.LocaleShow(Auth, source.ProjectID, identifier, localeShowParams)

	if err != nil {
		return nil, false, err
	}

	return &localeDetail, true, nil
}

func localeIdentifier(source *Source, localeFile *LocaleFile) string {
	localeName := source.replacePlaceholderInParams(localeFile)
	if localeName != "" && localeName != localeFile.Code {
		return localeName
	}

	if localeFile.Name != "" {
		return localeFile.Name
	}

	if localeFile.Code != "" {
		return localeFile.Code
	}

	return ""
}

func (source *Source) replacePlaceholderInParams(localeFile *LocaleFile) string {
	if localeFile.Code != "" && strings.Contains(source.GetLocaleID(), "<locale_code>") {
		return strings.Replace(source.GetLocaleID(), "<locale_code>", localeFile.Code, 1)
	}
	return ""
}
