package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/antihax/optional"
	helpers "github.com/phrase/phrase-cli/helpers"
	api "github.com/phrase/phrase-go/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initKeyFormatAnnotationsList()

	rootCmd.AddCommand(KeyFormatAnnotationsApiCmd)
}

var KeyFormatAnnotationsApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("KeyFormatAnnotations"),
	Short: "KeyFormatAnnotations API",
}

func initKeyFormatAnnotationsList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("key_format_annotations/list", "/")[1:], "_")
	var KeyFormatAnnotationsList = &cobra.Command{
		Use:   use,
		Short: "List format annotations for a key",
		Long:  `Returns the format annotations stored on a translation key. Format annotations capture file-format data recorded when the key was imported — for example, an ARB placeholder block or an XLIFF note.  Results are limited to 1,000 entries. `,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.KeyFormatAnnotationsListOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			id := params.GetString(helpers.ToSnakeCase("Id"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			data, api_response, err := client.KeyFormatAnnotationsApi.KeyFormatAnnotationsList(auth, projectId, id, &localVarOptionals)

			if err != nil {
				switch castedError := err.(type) {
				case api.GenericOpenAPIError:
					fmt.Printf("\n%s\n\n", string(castedError.Body()))
					HandleError(castedError)

				default:
					HandleError(castedError)
				}
			} else if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}
				fmt.Printf("%s\n", string(jsonBuf))

				if Config.Debug {
					fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
				}
			}
		},
	}

	KeyFormatAnnotationsApiCmd.AddCommand(KeyFormatAnnotationsList)
	AddFlag(KeyFormatAnnotationsList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(KeyFormatAnnotationsList, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(KeyFormatAnnotationsList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(KeyFormatAnnotationsList, "string", helpers.ToSnakeCase("Branch"), "", "Branch to use", false)

	params.BindPFlags(KeyFormatAnnotationsList.Flags())
}
