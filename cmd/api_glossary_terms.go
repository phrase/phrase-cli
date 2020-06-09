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
	initGlossaryTermCreate()
	initGlossaryTermDelete()
	initGlossaryTermShow()
	initGlossaryTermUpdate()
	initGlossaryTermsList()

	rootCmd.AddCommand(GlossaryTermsApiCmd)
}

var GlossaryTermsApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("GlossaryTerms"),
	Short: "GlossaryTerms API",
}

func initGlossaryTermCreate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("glossary_term/create", "/")[1:], "_")
	var GlossaryTermCreate = &cobra.Command{
		Use:   use,
		Short: "Create a glossary term",
		Long:  `Create a new glossary term.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryTermCreateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			glossaryId := params.GetString(helpers.ToSnakeCase("GlossaryId"))

			glossaryTermCreateParameters := api.GlossaryTermCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &glossaryTermCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", glossaryTermCreateParameters)
			}
			data, api_response, err := client.GlossaryTermsApi.GlossaryTermCreate(auth, accountId, glossaryId, glossaryTermCreateParameters, &localVarOptionals)

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

	GlossaryTermsApiCmd.AddCommand(GlossaryTermCreate)

	AddFlag(GlossaryTermCreate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(GlossaryTermCreate, "string", helpers.ToSnakeCase("GlossaryId"), "", "Glossary ID", true)
	AddFlag(GlossaryTermCreate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(GlossaryTermCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(GlossaryTermCreate.Flags())
}
func initGlossaryTermDelete() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("glossary_term/delete", "/")[1:], "_")
	var GlossaryTermDelete = &cobra.Command{
		Use:   use,
		Short: "Delete a glossary term",
		Long:  `Delete an existing glossary term.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryTermDeleteOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			glossaryId := params.GetString(helpers.ToSnakeCase("GlossaryId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.GlossaryTermsApi.GlossaryTermDelete(auth, accountId, glossaryId, id, &localVarOptionals)

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

	GlossaryTermsApiCmd.AddCommand(GlossaryTermDelete)

	AddFlag(GlossaryTermDelete, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(GlossaryTermDelete, "string", helpers.ToSnakeCase("GlossaryId"), "", "Glossary ID", true)
	AddFlag(GlossaryTermDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(GlossaryTermDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(GlossaryTermDelete.Flags())
}
func initGlossaryTermShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("glossary_term/show", "/")[1:], "_")
	var GlossaryTermShow = &cobra.Command{
		Use:   use,
		Short: "Get a single glossary term",
		Long:  `Get details on a single glossary term.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryTermShowOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			glossaryId := params.GetString(helpers.ToSnakeCase("GlossaryId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.GlossaryTermsApi.GlossaryTermShow(auth, accountId, glossaryId, id, &localVarOptionals)

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

	GlossaryTermsApiCmd.AddCommand(GlossaryTermShow)

	AddFlag(GlossaryTermShow, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(GlossaryTermShow, "string", helpers.ToSnakeCase("GlossaryId"), "", "Glossary ID", true)
	AddFlag(GlossaryTermShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(GlossaryTermShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(GlossaryTermShow.Flags())
}
func initGlossaryTermUpdate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("glossary_term/update", "/")[1:], "_")
	var GlossaryTermUpdate = &cobra.Command{
		Use:   use,
		Short: "Update a glossary term",
		Long:  `Update an existing glossary term.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryTermUpdateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			glossaryId := params.GetString(helpers.ToSnakeCase("GlossaryId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			glossaryTermUpdateParameters := api.GlossaryTermUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &glossaryTermUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", glossaryTermUpdateParameters)
			}
			data, api_response, err := client.GlossaryTermsApi.GlossaryTermUpdate(auth, accountId, glossaryId, id, glossaryTermUpdateParameters, &localVarOptionals)

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

	GlossaryTermsApiCmd.AddCommand(GlossaryTermUpdate)

	AddFlag(GlossaryTermUpdate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(GlossaryTermUpdate, "string", helpers.ToSnakeCase("GlossaryId"), "", "Glossary ID", true)
	AddFlag(GlossaryTermUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(GlossaryTermUpdate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(GlossaryTermUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(GlossaryTermUpdate.Flags())
}
func initGlossaryTermsList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("glossary_terms/list", "/")[1:], "_")
	var GlossaryTermsList = &cobra.Command{
		Use:   use,
		Short: "List glossary terms",
		Long:  `List all glossary terms the current user has access to.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryTermsListOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("page")) {
				localVarOptionals.Page = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("Page")))
			}
			if params.IsSet(helpers.ToSnakeCase("perPage")) {
				localVarOptionals.PerPage = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("PerPage")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			glossaryId := params.GetString(helpers.ToSnakeCase("GlossaryId"))

			data, api_response, err := client.GlossaryTermsApi.GlossaryTermsList(auth, accountId, glossaryId, &localVarOptionals)

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

	GlossaryTermsApiCmd.AddCommand(GlossaryTermsList)

	AddFlag(GlossaryTermsList, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(GlossaryTermsList, "string", helpers.ToSnakeCase("GlossaryId"), "", "Glossary ID", true)
	AddFlag(GlossaryTermsList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(GlossaryTermsList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(GlossaryTermsList, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 10 by default", false)
	params.BindPFlags(GlossaryTermsList.Flags())
}
