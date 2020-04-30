package pull

import (
	"github.com/antihax/optional"
	"github.com/phrase/phrase-go"
)

var Debug bool

type ProjectLocales interface {
	ProjectIds() []string
}

type LocaleCacheKey struct {
	ProjectID string
	Branch    string
}

type LocaleCache map[LocaleCacheKey][]*phrase.Locale

func LocalesForProjects(client *phrase.APIClient, projectLocales ProjectLocales, branch string) (LocaleCache, error) {
	projectIdToLocales := LocaleCache{}

	for _, pid := range projectLocales.ProjectIds() {
		key := LocaleCacheKey{
			ProjectID: pid,
			Branch:    branch,
		}

		if _, ok := projectIdToLocales[key]; !ok {

			remoteLocales, http_response, err := RemoteLocales(client, key)
			if err != nil {
				if http_response.StatusCode == 404 && branch != "" {
					// skip this key if we targeted a branch in
					// a project which does not exist
					continue
				}
				return nil, err
			}

			projectIdToLocales[key] = remoteLocales
		}
	}
	return projectIdToLocales, nil
}

func RemoteLocales(client *phrase.APIClient, key LocaleCacheKey) ([]*phrase.Locale, *phrase.APIResponse, error) {
	page := 1

	localVarOptionals := phrase.LocalesListOpts{
		Page:    optional.NewInt32(int32(page)),
		PerPage: optional.NewInt32(int32(25)),
	}

	if key.Branch != "" {
		localVarOptionals.Branch = optional.NewString(key.Branch)
	}

	locales, http_response, err := client.LocalesApi.LocalesList(Auth, key.ProjectID, &localVarOptionals)
	if err != nil {
		return nil, http_response, err
	}
	result := locales
	for len(locales) == 25 {
		page = page + 1
		localVarOptionals.Page = optional.NewInt32(int32(page))

		locales, http_response, err := client.LocalesApi.LocalesList(Auth, key.ProjectID, &localVarOptionals)
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
