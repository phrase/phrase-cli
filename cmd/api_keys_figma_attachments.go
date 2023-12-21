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
	initFigmaAttachmentAttachToKey()
	initFigmaAttachmentDetachFromKey()

	rootCmd.AddCommand(KeysFigmaAttachmentsApiCmd)
}

var KeysFigmaAttachmentsApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("KeysFigmaAttachments"),
	Short: "KeysFigmaAttachments API",
}

func initFigmaAttachmentAttachToKey() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("figma_attachment/attach_to_key", "/")[1:], "_")
	var FigmaAttachmentAttachToKey = &cobra.Command{
		Use:   use,
		Short: "Attach the Figma attachment to a key",
		Long:  `Attach the Figma attachment to a key`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.FigmaAttachmentAttachToKeyOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			figmaAttachmentId := params.GetString(helpers.ToSnakeCase("FigmaAttachmentId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			// trigger API request using phrase-go client
			data, api_response, err := client.KeysFigmaAttachmentsApi.FigmaAttachmentAttachToKey(auth, projectId, figmaAttachmentId, id, &localVarOptionals)

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

	KeysFigmaAttachmentsApiCmd.AddCommand(FigmaAttachmentAttachToKey)
	AddFlag(FigmaAttachmentAttachToKey, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(FigmaAttachmentAttachToKey, "string", helpers.ToSnakeCase("FigmaAttachmentId"), "", "Figma attachment ID", true)
	AddFlag(FigmaAttachmentAttachToKey, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(FigmaAttachmentAttachToKey, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(FigmaAttachmentAttachToKey, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(FigmaAttachmentAttachToKey.Flags())
}
func initFigmaAttachmentDetachFromKey() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("figma_attachment/detach_from_key", "/")[1:], "_")
	var FigmaAttachmentDetachFromKey = &cobra.Command{
		Use:   use,
		Short: "Detach the Figma attachment from a key",
		Long:  `Detach the Figma attachment from a key`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.FigmaAttachmentDetachFromKeyOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			figmaAttachmentId := params.GetString(helpers.ToSnakeCase("FigmaAttachmentId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			// trigger API request using phrase-go client
			data, api_response, err := client.KeysFigmaAttachmentsApi.FigmaAttachmentDetachFromKey(auth, projectId, figmaAttachmentId, id, &localVarOptionals)

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

	KeysFigmaAttachmentsApiCmd.AddCommand(FigmaAttachmentDetachFromKey)
	AddFlag(FigmaAttachmentDetachFromKey, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(FigmaAttachmentDetachFromKey, "string", helpers.ToSnakeCase("FigmaAttachmentId"), "", "Figma attachment ID", true)
	AddFlag(FigmaAttachmentDetachFromKey, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(FigmaAttachmentDetachFromKey, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(FigmaAttachmentDetachFromKey, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(FigmaAttachmentDetachFromKey.Flags())
}
