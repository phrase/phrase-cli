package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/antihax/optional"
	"github.com/mitchellh/mapstructure"
	"github.com/phrase/phrase-cli/cmd/internal/print"
	"github.com/phrase/phrase-cli/cmd/internal/prompt"
	"github.com/phrase/phrase-cli/cmd/internal/shared"
	"github.com/phrase/phrase-go/v4"
	"github.com/spf13/viper"
)

var Debug bool

type SourcesOrTargets interface {
	// returns a list of LocalesCacheKeys (ProjectId, Branch) for all targets
	GetAllLocalesCacheKeys() []LocalesCacheKey
}

type LocalesCacheKey struct {
	ProjectID string
	Branch    string
}

type LocaleCache map[LocalesCacheKey][]*phrase.Locale

// for every source or target, retrieves and caches the list of locales
func GetLocalesCache(client *phrase.APIClient, sourcesOrTargets SourcesOrTargets, branch string) (LocaleCache, error) {
	localesCache := LocaleCache{}

	for _, localesCacheKey := range sourcesOrTargets.GetAllLocalesCacheKeys() {
		branchToUse := localesCacheKey.Branch
		if branch != "" {
			branchToUse = branch
		}
		key := LocalesCacheKey{
			ProjectID: localesCacheKey.ProjectID,
			Branch:    branchToUse,
		}

		if _, ok := localesCache[key]; !ok {

			remoteLocales, http_response, err := RemoteLocales(client, key)
			if err != nil {
				if http_response != nil && http_response.StatusCode == 404 && branchToUse != "" {
					// skip this key if we targeted a branch in
					// a project which does not exist
					continue
				}
				return nil, err
			}

			localesCache[key] = remoteLocales
		}
	}
	return localesCache, nil
}

func RemoteLocales(client *phrase.APIClient, key LocalesCacheKey) ([]*phrase.Locale, *phrase.APIResponse, error) {
	page := 1

	localVarOptionals := phrase.LocalesListOpts{
		Page:    optional.NewInt32(int32(page)),
		PerPage: optional.NewInt32(100),
	}

	if key.Branch != "" {
		localVarOptionals.Branch = optional.NewString(key.Branch)
	}

	locales, http_response, err := client.LocalesApi.LocalesList(Auth, key.ProjectID, &localVarOptionals)
	if err != nil {
		return nil, http_response, err
	}
	result := locales
	for http_response.NextPage > 0 {
		page = page + 1
		localVarOptionals.Page = optional.NewInt32(int32(page))

		locales, http_response, err = client.LocalesApi.LocalesList(Auth, key.ProjectID, &localVarOptionals)
		if err != nil {
			return nil, http_response, err
		}
		result = append(result, locales...)
	}

	var data []*phrase.Locale

	for i := 0; i < len(result); i++ {
		data = append(data, &result[i])
	}

	return data, nil, nil
}

func Projects(client *phrase.APIClient) ([]phrase.Project, *phrase.APIResponse, error) {
	page := 1
	projectVarOptionals := phrase.ProjectsListOpts{
		Page:    optional.NewInt32(int32(page)),
		PerPage: optional.NewInt32(100),
	}

	projects, http_response, err := client.ProjectsApi.ProjectsList(Auth, &projectVarOptionals)
	if err != nil {
		return nil, http_response, err
	}

	result := projects
	for http_response.NextPage > 0 {
		page = page + 1
		projectVarOptionals.Page = optional.NewInt32(int32(page))

		projects, http_response, err = client.ProjectsApi.ProjectsList(Auth, &projectVarOptionals)
		if err != nil {
			return nil, http_response, err
		}
		result = append(result, projects...)
	}

	var data []phrase.Project

	for i := 0; i < len(result); i++ {
		data = append(data, result[i])
	}

	return data, nil, nil
}

func ViperStructTag() viper.DecoderConfigOption {
	return func(c *mapstructure.DecoderConfig) {
		c.TagName = "json"
		c.Squash = true
		c.DecodeHook = mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			StringToOptionalString(),
			StringToOptionalBool(),
			StringToInterface(),
		)
	}
}

// StringToOptionalString returns a DecodeHookFunc that converts
// strings to optional.String.
func StringToOptionalString() mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}
		if t != reflect.TypeOf(optional.String{}) {
			return data, nil
		}

		return optional.NewString(data.(string)), nil
	}
}

// StringToOptionalBool returns a DecodeHookFunc that converts
// strings to optional.Bool.
func StringToOptionalBool() mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		if f.Kind() != reflect.Bool {
			return data, nil
		}
		if t != reflect.TypeOf(optional.Bool{}) {
			return data, nil
		}

		return optional.NewBool(data.(bool)), nil
	}
}

// StringToInterface returns a DecodeHookFunc that converts
// strings to optional.Interface.
func StringToInterface() mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		if f.Kind() != reflect.Map {
			return data, nil
		}
		if t != reflect.TypeOf(optional.Interface{}) {
			return data, nil
		}

		var params map[string]interface{}
		config := &mapstructure.DecoderConfig{
			Result: &params,
			DecodeHook: mapstructure.ComposeDecodeHookFunc(
				StringToOptionalString(),
				StringToOptionalBool(),
			),
		}

		decoder, err := mapstructure.NewDecoder(config)
		if err != nil {
			return data, nil
		}

		err = decoder.Decode(data)
		if err != nil {
			return data, nil
		}

		return optional.NewInterface(params), nil
	}
}

func UploadCleanup(client *phrase.APIClient, confirm bool, ids []string, branch string, projectId string) error {
	if !shared.BatchMode {
		fmt.Println("Keys not mentioned in the following uploads will be deleted:")
		fmt.Println(strings.Join(ids, "\n"))
	}
	if !confirm {
		if shared.BatchMode {
			return errors.New("Can't ask for confirmation in batch mode. Aborting")
		}
		confirmation := ""
		err := prompt.WithDefault("Are you sure you want to continue? (y/n)", &confirmation, "n")
		if err != nil {
			return err
		}

		if strings.ToLower(confirmation) != "y" {
			fmt.Println("Clean up aborted")
			return nil
		}
	}

	q := "unmentioned_in_upload:" + strings.Join(ids, ",")
	optionalBranch := optional.String{}
	if branch != "" {
		optionalBranch = optional.NewString(branch)
	}
	keysDeletelocalVarOptionals := phrase.KeysDeleteCollectionOpts{
		Q:      optional.NewString(q),
		Branch: optionalBranch,
	}
	affected, _, err := client.KeysApi.KeysDeleteCollection(Auth, projectId, &keysDeletelocalVarOptionals)

	if err != nil {
		return err
	}
	if shared.BatchMode {
		jsonBuf, jsonErr := json.MarshalIndent(affected, "", " ")
		if jsonErr != nil {
			print.Error(jsonErr)
		}
		fmt.Printf("%s\n", string(jsonBuf))
	} else {
		print.Success("%d key(s) successfully deleted.\n", affected.RecordsAffected)
	}
	return nil
}
