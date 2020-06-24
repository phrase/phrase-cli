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
	initBlacklistedKeyCreate()
	initBlacklistedKeyDelete()
	initBlacklistedKeyShow()
	initBlacklistedKeyUpdate()
	initBlacklistedKeysList()

	rootCmd.AddCommand(BlacklistedKeysApiCmd)
}

var BlacklistedKeysApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("BlacklistedKeys"),
	Short: "BlacklistedKeys API",
}

func initBlacklistedKeyCreate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("blacklisted_key/create", "/")[1:], "_")
	var BlacklistedKeyCreate = &cobra.Command{
		Use:   use,
		Short: "Create a blacklisted key",
		Long:  `Create a new rule for blacklisting keys.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.BlacklistedKeyCreateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			blacklistedKeyCreateParameters := api.BlacklistedKeyCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &blacklistedKeyCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", blacklistedKeyCreateParameters)
			}
			data, api_response, err := client.BlacklistedKeysApi.BlacklistedKeyCreate(auth, projectId, blacklistedKeyCreateParameters, &localVarOptionals)

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

	BlacklistedKeysApiCmd.AddCommand(BlacklistedKeyCreate)
	AddFlag(BlacklistedKeyCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(BlacklistedKeyCreate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(BlacklistedKeyCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(BlacklistedKeyCreate.Flags())
}
func initBlacklistedKeyDelete() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("blacklisted_key/delete", "/")[1:], "_")
	var BlacklistedKeyDelete = &cobra.Command{
		Use:   use,
		Short: "Delete a blacklisted key",
		Long:  `Delete an existing rule for blacklisting keys.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.BlacklistedKeyDeleteOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.BlacklistedKeysApi.BlacklistedKeyDelete(auth, projectId, id, &localVarOptionals)

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

	BlacklistedKeysApiCmd.AddCommand(BlacklistedKeyDelete)
	AddFlag(BlacklistedKeyDelete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(BlacklistedKeyDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(BlacklistedKeyDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(BlacklistedKeyDelete.Flags())
}
func initBlacklistedKeyShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("blacklisted_key/show", "/")[1:], "_")
	var BlacklistedKeyShow = &cobra.Command{
		Use:   use,
		Short: "Get a single blacklisted key",
		Long:  `Get details on a single rule for blacklisting keys for a given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.BlacklistedKeyShowOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.BlacklistedKeysApi.BlacklistedKeyShow(auth, projectId, id, &localVarOptionals)

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

	BlacklistedKeysApiCmd.AddCommand(BlacklistedKeyShow)
	AddFlag(BlacklistedKeyShow, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(BlacklistedKeyShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(BlacklistedKeyShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(BlacklistedKeyShow.Flags())
}
func initBlacklistedKeyUpdate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("blacklisted_key/update", "/")[1:], "_")
	var BlacklistedKeyUpdate = &cobra.Command{
		Use:   use,
		Short: "Update a blacklisted key",
		Long:  `Update an existing rule for blacklisting keys.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.BlacklistedKeyUpdateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			blacklistedKeyUpdateParameters := api.BlacklistedKeyUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &blacklistedKeyUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", blacklistedKeyUpdateParameters)
			}
			data, api_response, err := client.BlacklistedKeysApi.BlacklistedKeyUpdate(auth, projectId, id, blacklistedKeyUpdateParameters, &localVarOptionals)

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

	BlacklistedKeysApiCmd.AddCommand(BlacklistedKeyUpdate)
	AddFlag(BlacklistedKeyUpdate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(BlacklistedKeyUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(BlacklistedKeyUpdate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(BlacklistedKeyUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(BlacklistedKeyUpdate.Flags())
}
func initBlacklistedKeysList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("blacklisted_keys/list", "/")[1:], "_")
	var BlacklistedKeysList = &cobra.Command{
		Use:   use,
		Short: "List blacklisted keys",
		Long:  `List all rules for blacklisting keys for the given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.BlacklistedKeysListOpts{}

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

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			data, api_response, err := client.BlacklistedKeysApi.BlacklistedKeysList(auth, projectId, &localVarOptionals)

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

	BlacklistedKeysApiCmd.AddCommand(BlacklistedKeysList)
	AddFlag(BlacklistedKeysList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(BlacklistedKeysList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(BlacklistedKeysList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(BlacklistedKeysList, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 10 by default", false)

	params.BindPFlags(BlacklistedKeysList.Flags())
}
