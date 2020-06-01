package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/antihax/optional"
	helpers "github.com/phrase/phrase-cli/helpers"
	api "github.com/phrase/phrase-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initKeyCreate()
	initKeyDelete()
	initKeyShow()
	initKeyUpdate()
	initKeysDelete()
	initKeysList()
	initKeysSearch()
	initKeysTag()
	initKeysUntag()

	rootCmd.AddCommand(KeysApiCmd)
}

var KeysApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("Keys"),
	Short: "Keys API",
}

func initKeyCreate() {
	params := viper.New()
	var KeyCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("KeyCreate", strings.TrimSuffix("KeysApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("KeysApi", "Api"), "s"))),
		Short: "Create a key",
		Long:  `Create a new key.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.KeyCreateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			keyCreateParameters := api.KeyCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &keyCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", keyCreateParameters)
			}
			data, api_response, err := client.KeysApi.KeyCreate(auth, projectId, keyCreateParameters, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	KeysApiCmd.AddCommand(KeyCreate)

	AddFlag(KeyCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(KeyCreate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(KeyCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(KeyCreate.Flags())
}
func initKeyDelete() {
	params := viper.New()
	var KeyDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("KeyDelete", strings.TrimSuffix("KeysApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("KeysApi", "Api"), "s"))),
		Short: "Delete a key",
		Long:  `Delete an existing key.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.KeyDeleteOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.KeysApi.KeyDelete(auth, projectId, id, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	KeysApiCmd.AddCommand(KeyDelete)

	AddFlag(KeyDelete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(KeyDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(KeyDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(KeyDelete, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)
	params.BindPFlags(KeyDelete.Flags())
}
func initKeyShow() {
	params := viper.New()
	var KeyShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("KeyShow", strings.TrimSuffix("KeysApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("KeysApi", "Api"), "s"))),
		Short: "Get a single key",
		Long:  `Get details on a single key for a given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.KeyShowOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.KeysApi.KeyShow(auth, projectId, id, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	KeysApiCmd.AddCommand(KeyShow)

	AddFlag(KeyShow, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(KeyShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(KeyShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(KeyShow, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)
	params.BindPFlags(KeyShow.Flags())
}
func initKeyUpdate() {
	params := viper.New()
	var KeyUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("KeyUpdate", strings.TrimSuffix("KeysApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("KeysApi", "Api"), "s"))),
		Short: "Update a key",
		Long:  `Update an existing key.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.KeyUpdateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			keyUpdateParameters := api.KeyUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &keyUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", keyUpdateParameters)
			}
			data, api_response, err := client.KeysApi.KeyUpdate(auth, projectId, id, keyUpdateParameters, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	KeysApiCmd.AddCommand(KeyUpdate)

	AddFlag(KeyUpdate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(KeyUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(KeyUpdate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(KeyUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(KeyUpdate.Flags())
}
func initKeysDelete() {
	params := viper.New()
	var KeysDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("KeysDelete", strings.TrimSuffix("KeysApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("KeysApi", "Api"), "s"))),
		Short: "Delete collection of keys",
		Long:  `Delete all keys matching query. Same constraints as list. Please limit the number of affected keys to about 1,000 as you might experience timeouts otherwise.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.KeysDeleteOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}
			if params.IsSet(helpers.ToSnakeCase("q")) {
				localVarOptionals.Q = optional.NewString(params.GetString(helpers.ToSnakeCase("Q")))
			}
			if params.IsSet(helpers.ToSnakeCase("localeId")) {
				localVarOptionals.LocaleId = optional.NewString(params.GetString(helpers.ToSnakeCase("LocaleId")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			data, api_response, err := client.KeysApi.KeysDelete(auth, projectId, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	KeysApiCmd.AddCommand(KeysDelete)

	AddFlag(KeysDelete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(KeysDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(KeysDelete, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)
	AddFlag(KeysDelete, "string", helpers.ToSnakeCase("Q"), "", "q_description_placeholder", false)
	AddFlag(KeysDelete, "string", helpers.ToSnakeCase("LocaleId"), "", "Locale used to determine the translation state of a key when filtering for untranslated or translated keys.", false)
	params.BindPFlags(KeysDelete.Flags())
}
func initKeysList() {
	params := viper.New()
	var KeysList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("KeysList", strings.TrimSuffix("KeysApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("KeysApi", "Api"), "s"))),
		Short: "List keys",
		Long:  `List all keys for the given project. Alternatively you can POST requests to /search.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.KeysListOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("page")) {
				localVarOptionals.Page = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("Page")))
			}
			if params.IsSet(helpers.ToSnakeCase("perPage")) {
				localVarOptionals.PerPage = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("PerPage")))
			}
			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}
			if params.IsSet(helpers.ToSnakeCase("sort")) {
				localVarOptionals.Sort = optional.NewString(params.GetString(helpers.ToSnakeCase("Sort")))
			}
			if params.IsSet(helpers.ToSnakeCase("order")) {
				localVarOptionals.Order = optional.NewString(params.GetString(helpers.ToSnakeCase("Order")))
			}
			if params.IsSet(helpers.ToSnakeCase("q")) {
				localVarOptionals.Q = optional.NewString(params.GetString(helpers.ToSnakeCase("Q")))
			}
			if params.IsSet(helpers.ToSnakeCase("localeId")) {
				localVarOptionals.LocaleId = optional.NewString(params.GetString(helpers.ToSnakeCase("LocaleId")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			data, api_response, err := client.KeysApi.KeysList(auth, projectId, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	KeysApiCmd.AddCommand(KeysList)

	AddFlag(KeysList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(KeysList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(KeysList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(KeysList, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 10 by default", false)
	AddFlag(KeysList, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)
	AddFlag(KeysList, "string", helpers.ToSnakeCase("Sort"), "", "Sort by field. Can be one of: name, created_at, updated_at.", false)
	AddFlag(KeysList, "string", helpers.ToSnakeCase("Order"), "", "Order direction. Can be one of: asc, desc.", false)
	AddFlag(KeysList, "string", helpers.ToSnakeCase("Q"), "", "q_description_placeholder", false)
	AddFlag(KeysList, "string", helpers.ToSnakeCase("LocaleId"), "", "Locale used to determine the translation state of a key when filtering for untranslated or translated keys.", false)
	params.BindPFlags(KeysList.Flags())
}
func initKeysSearch() {
	params := viper.New()
	var KeysSearch = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("KeysSearch", strings.TrimSuffix("KeysApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("KeysApi", "Api"), "s"))),
		Short: "Search keys",
		Long:  `Search keys for the given project matching query.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.KeysSearchOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("page")) {
				localVarOptionals.Page = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("Page")))
			}
			if params.IsSet(helpers.ToSnakeCase("perPage")) {
				localVarOptionals.PerPage = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("PerPage")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			keysSearchParameters := api.KeysSearchParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &keysSearchParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", keysSearchParameters)
			}
			data, api_response, err := client.KeysApi.KeysSearch(auth, projectId, keysSearchParameters, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	KeysApiCmd.AddCommand(KeysSearch)

	AddFlag(KeysSearch, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(KeysSearch, "string", "data", "d", "payload in JSON format", true)

	AddFlag(KeysSearch, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(KeysSearch, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(KeysSearch, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 10 by default", false)
	params.BindPFlags(KeysSearch.Flags())
}
func initKeysTag() {
	params := viper.New()
	var KeysTag = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("KeysTag", strings.TrimSuffix("KeysApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("KeysApi", "Api"), "s"))),
		Short: "Add tags to collection of keys",
		Long:  `Tags all keys matching query. Same constraints as list.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.KeysTagOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			keysTagParameters := api.KeysTagParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &keysTagParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", keysTagParameters)
			}
			data, api_response, err := client.KeysApi.KeysTag(auth, projectId, keysTagParameters, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	KeysApiCmd.AddCommand(KeysTag)

	AddFlag(KeysTag, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(KeysTag, "string", "data", "d", "payload in JSON format", true)

	AddFlag(KeysTag, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(KeysTag.Flags())
}
func initKeysUntag() {
	params := viper.New()
	var KeysUntag = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("KeysUntag", strings.TrimSuffix("KeysApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("KeysApi", "Api"), "s"))),
		Short: "Remove tags from collection of keys",
		Long:  `Removes specified tags from keys matching query.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.KeysUntagOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			keysUntagParameters := api.KeysUntagParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &keysUntagParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", keysUntagParameters)
			}
			data, api_response, err := client.KeysApi.KeysUntag(auth, projectId, keysUntagParameters, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	KeysApiCmd.AddCommand(KeysUntag)

	AddFlag(KeysUntag, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(KeysUntag, "string", "data", "d", "payload in JSON format", true)

	AddFlag(KeysUntag, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(KeysUntag.Flags())
}
