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
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("glossary_term_translation/create", "/")[1:], "_")
	var GlossaryTermTranslationCreate = &cobra.Command{
		Use:   use,
		Short: "Create a glossary term translation",
		Long:  `Create a new glossary term translation.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.GlossaryTermTranslationCreateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

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

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}
				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				switch castedError := err.(type) {
				case api.GenericOpenAPIError:
					fmt.Printf("\n%s\n\n", string(castedError.Body()))
					HandleError(castedError)

				default:
					HandleError(castedError)
				}
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
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("glossary_term_translation/delete", "/")[1:], "_")
	var GlossaryTermTranslationDelete = &cobra.Command{
		Use:   use,
		Short: "Delete a glossary term translation",
		Long:  `Delete an existing glossary term translation.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.GlossaryTermTranslationDeleteOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			glossaryId := params.GetString(helpers.ToSnakeCase("GlossaryId"))
			termId := params.GetString(helpers.ToSnakeCase("TermId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.GlossaryTermTranslationsApi.GlossaryTermTranslationDelete(auth, accountId, glossaryId, termId, id, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
				os.Stdout.Write(data)
			}
			if err != nil {
				switch castedError := err.(type) {
				case api.GenericOpenAPIError:
					fmt.Printf("\n%s\n\n", string(castedError.Body()))
					HandleError(castedError)

				default:
					HandleError(castedError)
				}
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
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("glossary_term_translation/update", "/")[1:], "_")
	var GlossaryTermTranslationUpdate = &cobra.Command{
		Use:   use,
		Short: "Update a glossary term translation",
		Long:  `Update an existing glossary term translation.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.GlossaryTermTranslationUpdateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

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

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}
				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				switch castedError := err.(type) {
				case api.GenericOpenAPIError:
					fmt.Printf("\n%s\n\n", string(castedError.Body()))
					HandleError(castedError)

				default:
					HandleError(castedError)
				}
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
