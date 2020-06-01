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
	initGlossariesList()
	initGlossaryCreate()
	initGlossaryDelete()
	initGlossaryShow()
	initGlossaryUpdate()

	rootCmd.AddCommand(GlossaryApiCmd)
}

var GlossaryApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("Glossary"),
	Short: "Glossary API",
}

func initGlossariesList() {
	params := viper.New()
	var GlossariesList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GlossariesList", strings.TrimSuffix("GlossaryApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GlossaryApi", "Api"), "s"))),
		Short: "List glossaries",
		Long:  `List all glossaries the current user has access to.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossariesListOpts{}

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

			data, api_response, err := client.GlossaryApi.GlossariesList(auth, accountId, &localVarOptionals)

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

	GlossaryApiCmd.AddCommand(GlossariesList)

	AddFlag(GlossariesList, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(GlossariesList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(GlossariesList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(GlossariesList, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 10 by default", false)
	params.BindPFlags(GlossariesList.Flags())
}
func initGlossaryCreate() {
	params := viper.New()
	var GlossaryCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GlossaryCreate", strings.TrimSuffix("GlossaryApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GlossaryApi", "Api"), "s"))),
		Short: "Create a glossary",
		Long:  `Create a new glossary.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryCreateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			glossaryCreateParameters := api.GlossaryCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &glossaryCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", glossaryCreateParameters)
			}
			data, api_response, err := client.GlossaryApi.GlossaryCreate(auth, accountId, glossaryCreateParameters, &localVarOptionals)

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

	GlossaryApiCmd.AddCommand(GlossaryCreate)

	AddFlag(GlossaryCreate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(GlossaryCreate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(GlossaryCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(GlossaryCreate.Flags())
}
func initGlossaryDelete() {
	params := viper.New()
	var GlossaryDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GlossaryDelete", strings.TrimSuffix("GlossaryApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GlossaryApi", "Api"), "s"))),
		Short: "Delete a glossary",
		Long:  `Delete an existing glossary.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryDeleteOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.GlossaryApi.GlossaryDelete(auth, accountId, id, &localVarOptionals)

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

	GlossaryApiCmd.AddCommand(GlossaryDelete)

	AddFlag(GlossaryDelete, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(GlossaryDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(GlossaryDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(GlossaryDelete.Flags())
}
func initGlossaryShow() {
	params := viper.New()
	var GlossaryShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GlossaryShow", strings.TrimSuffix("GlossaryApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GlossaryApi", "Api"), "s"))),
		Short: "Get a single glossary",
		Long:  `Get details on a single glossary.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryShowOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.GlossaryApi.GlossaryShow(auth, accountId, id, &localVarOptionals)

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

	GlossaryApiCmd.AddCommand(GlossaryShow)

	AddFlag(GlossaryShow, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(GlossaryShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(GlossaryShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(GlossaryShow.Flags())
}
func initGlossaryUpdate() {
	params := viper.New()
	var GlossaryUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GlossaryUpdate", strings.TrimSuffix("GlossaryApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GlossaryApi", "Api"), "s"))),
		Short: "Update a glossary",
		Long:  `Update an existing glossary.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryUpdateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			glossaryUpdateParameters := api.GlossaryUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &glossaryUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", glossaryUpdateParameters)
			}
			data, api_response, err := client.GlossaryApi.GlossaryUpdate(auth, accountId, id, glossaryUpdateParameters, &localVarOptionals)

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

	GlossaryApiCmd.AddCommand(GlossaryUpdate)

	AddFlag(GlossaryUpdate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(GlossaryUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(GlossaryUpdate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(GlossaryUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(GlossaryUpdate.Flags())
}
