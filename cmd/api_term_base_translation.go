package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/antihax/optional"
	helpers "github.com/phrase/phrase-cli/helpers"
	api "github.com/phrase/phrase-go/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initGlossaryTermTranslationDelete()

	rootCmd.AddCommand(TermBaseTranslationApiCmd)
}

var TermBaseTranslationApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("TermBaseTranslation"),
	Short: "TermBaseTranslation API",
}

func initGlossaryTermTranslationDelete() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("glossary_term_translation/delete", "/")[1:], "_")
	var GlossaryTermTranslationDelete = &cobra.Command{
		Use:   use,
		Short: "Delete a translation for a term",
		Long:  `Delete an existing translation of a term in a term base (previously: glossary).`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

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

			data, api_response, err := client.TermBaseTranslationApi.GlossaryTermTranslationDelete(auth, accountId, glossaryId, termId, id, &localVarOptionals)

			if err != nil {
				switch castedError := err.(type) {
				case api.GenericOpenAPIError:
					fmt.Printf("\n%s\n\n", string(castedError.Body()))
					HandleError(castedError)

				default:
					HandleError(castedError)
				}
			} else if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
				os.Stdout.Write(data)

				if Config.Debug {
					fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
				}
			}
		},
	}

	TermBaseTranslationApiCmd.AddCommand(GlossaryTermTranslationDelete)
	AddFlag(GlossaryTermTranslationDelete, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(GlossaryTermTranslationDelete, "string", helpers.ToSnakeCase("GlossaryId"), "", "Glossary ID", true)
	AddFlag(GlossaryTermTranslationDelete, "string", helpers.ToSnakeCase("TermId"), "", "Term ID", true)
	AddFlag(GlossaryTermTranslationDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(GlossaryTermTranslationDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(GlossaryTermTranslationDelete.Flags())
}
