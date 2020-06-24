package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/antihax/optional"
	helpers "github.com/phrase/phrase-cli/helpers"
	api "github.com/phrase/phrase-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initAuthorizationCreate()
	initAuthorizationDelete()
	initAuthorizationShow()
	initAuthorizationUpdate()
	initAuthorizationsList()

	rootCmd.AddCommand(AuthorizationsApiCmd)
}

var AuthorizationsApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("Authorizations"),
	Short: "Authorizations API",
}

func initAuthorizationCreate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("authorization/create", "/")[1:], "_")
	var AuthorizationCreate = &cobra.Command{
		Use:   use,
		Short: "Create an authorization",
		Long:  `Create a new authorization.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.AuthorizationCreateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			authorizationCreateParameters := api.AuthorizationCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &authorizationCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", authorizationCreateParameters)
			}
			data, api_response, err := client.AuthorizationsApi.AuthorizationCreate(auth, authorizationCreateParameters, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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

	AuthorizationsApiCmd.AddCommand(AuthorizationCreate)
	AddFlag(AuthorizationCreate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(AuthorizationCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(AuthorizationCreate.Flags())
}
func initAuthorizationDelete() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("authorization/delete", "/")[1:], "_")
	var AuthorizationDelete = &cobra.Command{
		Use:   use,
		Short: "Delete an authorization",
		Long:  `Delete an existing authorization. API calls using that token will stop working.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.AuthorizationDeleteOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.AuthorizationsApi.AuthorizationDelete(auth, id, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
				os.Stdout.Write(data)
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	AuthorizationsApiCmd.AddCommand(AuthorizationDelete)
	AddFlag(AuthorizationDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(AuthorizationDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(AuthorizationDelete.Flags())
}
func initAuthorizationShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("authorization/show", "/")[1:], "_")
	var AuthorizationShow = &cobra.Command{
		Use:   use,
		Short: "Get a single authorization",
		Long:  `Get details on a single authorization.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.AuthorizationShowOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.AuthorizationsApi.AuthorizationShow(auth, id, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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

	AuthorizationsApiCmd.AddCommand(AuthorizationShow)
	AddFlag(AuthorizationShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(AuthorizationShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(AuthorizationShow.Flags())
}
func initAuthorizationUpdate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("authorization/update", "/")[1:], "_")
	var AuthorizationUpdate = &cobra.Command{
		Use:   use,
		Short: "Update an authorization",
		Long:  `Update an existing authorization.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.AuthorizationUpdateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			id := params.GetString(helpers.ToSnakeCase("Id"))

			authorizationUpdateParameters := api.AuthorizationUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &authorizationUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", authorizationUpdateParameters)
			}
			data, api_response, err := client.AuthorizationsApi.AuthorizationUpdate(auth, id, authorizationUpdateParameters, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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

	AuthorizationsApiCmd.AddCommand(AuthorizationUpdate)
	AddFlag(AuthorizationUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(AuthorizationUpdate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(AuthorizationUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(AuthorizationUpdate.Flags())
}
func initAuthorizationsList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("authorizations/list", "/")[1:], "_")
	var AuthorizationsList = &cobra.Command{
		Use:   use,
		Short: "List authorizations",
		Long:  `List all your authorizations.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.AuthorizationsListOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("page")) {
				localVarOptionals.Page = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("Page")))
			}
			if params.IsSet(helpers.ToSnakeCase("perPage")) {
				localVarOptionals.PerPage = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("PerPage")))
			}

			data, api_response, err := client.AuthorizationsApi.AuthorizationsList(auth, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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

	AuthorizationsApiCmd.AddCommand(AuthorizationsList)
	AddFlag(AuthorizationsList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(AuthorizationsList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(AuthorizationsList, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 10 by default", false)

	params.BindPFlags(AuthorizationsList.Flags())
}
