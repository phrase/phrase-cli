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
	initGlossaryTermTranslationCreate()
	initGlossaryTermTranslationDelete()
	initGlossaryTermTranslationUpdate()

	rootCmd.AddCommand(GlossaryTermTranslationsApiCmd)
}

var GlossaryTermTranslationsApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("GlossaryTermTranslations"),
	Short: "GlossaryTermTranslations API",
}

func initGlossaryTermTranslationCreate() {
	params := viper.New()
	var GlossaryTermTranslationCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GlossaryTermTranslationCreate", strings.TrimSuffix("GlossaryTermTranslationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GlossaryTermTranslationsApi", "Api"), "s"))),
		Short: "Create a glossary term translation",
		Long:  `Create a new glossary term translation.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryTermTranslationCreateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			glossaryId := params.GetString(helpers.ToSnakeCase("GlossaryId"))
			termId := params.GetString(helpers.ToSnakeCase("TermId"))

			glossaryTermTranslationCreateParameters := api.GlossaryTermTranslationCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &glossaryTermTranslationCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", glossaryTermTranslationCreateParameters)
			}
			data, api_response, err := client.GlossaryTermTranslationsApi.GlossaryTermTranslationCreate(auth, accountId, glossaryId, termId, glossaryTermTranslationCreateParameters, &localVarOptionals)

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

	GlossaryTermTranslationsApiCmd.AddCommand(GlossaryTermTranslationCreate)

	AddFlag(GlossaryTermTranslationCreate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(GlossaryTermTranslationCreate, "string", helpers.ToSnakeCase("GlossaryId"), "", "Glossary ID", true)
	AddFlag(GlossaryTermTranslationCreate, "string", helpers.ToSnakeCase("TermId"), "", "Term ID", true)
	AddFlag(GlossaryTermTranslationCreate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(GlossaryTermTranslationCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(GlossaryTermTranslationCreate.Flags())
}
func initGlossaryTermTranslationDelete() {
	params := viper.New()
	var GlossaryTermTranslationDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GlossaryTermTranslationDelete", strings.TrimSuffix("GlossaryTermTranslationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GlossaryTermTranslationsApi", "Api"), "s"))),
		Short: "Delete a glossary term translation",
		Long:  `Delete an existing glossary term translation.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryTermTranslationDeleteOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			glossaryId := params.GetString(helpers.ToSnakeCase("GlossaryId"))
			termId := params.GetString(helpers.ToSnakeCase("TermId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.GlossaryTermTranslationsApi.GlossaryTermTranslationDelete(auth, accountId, glossaryId, termId, id, &localVarOptionals)

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

	GlossaryTermTranslationsApiCmd.AddCommand(GlossaryTermTranslationDelete)

	AddFlag(GlossaryTermTranslationDelete, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(GlossaryTermTranslationDelete, "string", helpers.ToSnakeCase("GlossaryId"), "", "Glossary ID", true)
	AddFlag(GlossaryTermTranslationDelete, "string", helpers.ToSnakeCase("TermId"), "", "Term ID", true)
	AddFlag(GlossaryTermTranslationDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(GlossaryTermTranslationDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(GlossaryTermTranslationDelete.Flags())
}
func initGlossaryTermTranslationUpdate() {
	params := viper.New()
	var GlossaryTermTranslationUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GlossaryTermTranslationUpdate", strings.TrimSuffix("GlossaryTermTranslationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GlossaryTermTranslationsApi", "Api"), "s"))),
		Short: "Update a glossary term translation",
		Long:  `Update an existing glossary term translation.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryTermTranslationUpdateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			glossaryId := params.GetString(helpers.ToSnakeCase("GlossaryId"))
			termId := params.GetString(helpers.ToSnakeCase("TermId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			glossaryTermTranslationUpdateParameters := api.GlossaryTermTranslationUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &glossaryTermTranslationUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", glossaryTermTranslationUpdateParameters)
			}
			data, api_response, err := client.GlossaryTermTranslationsApi.GlossaryTermTranslationUpdate(auth, accountId, glossaryId, termId, id, glossaryTermTranslationUpdateParameters, &localVarOptionals)

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

	GlossaryTermTranslationsApiCmd.AddCommand(GlossaryTermTranslationUpdate)

	AddFlag(GlossaryTermTranslationUpdate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(GlossaryTermTranslationUpdate, "string", helpers.ToSnakeCase("GlossaryId"), "", "Glossary ID", true)
	AddFlag(GlossaryTermTranslationUpdate, "string", helpers.ToSnakeCase("TermId"), "", "Term ID", true)
	AddFlag(GlossaryTermTranslationUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(GlossaryTermTranslationUpdate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(GlossaryTermTranslationUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(GlossaryTermTranslationUpdate.Flags())
}
