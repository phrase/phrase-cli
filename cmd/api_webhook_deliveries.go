package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/antihax/optional"
	helpers "github.com/phrase/phrase-cli/helpers"
	api "github.com/phrase/phrase-go/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initWebhookDeliveriesList()
	initWebhookDeliveriesRedeliver()
	initWebhookDeliveriesShow()

	rootCmd.AddCommand(WebhookDeliveriesApiCmd)
}

var WebhookDeliveriesApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("WebhookDeliveries"),
	Short: "WebhookDeliveries API",
}

func initWebhookDeliveriesList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("webhook_deliveries/list", "/")[1:], "_")
	var WebhookDeliveriesList = &cobra.Command{
		Use:   use,
		Short: "List webhook deliveries",
		Long:  `List all webhook deliveries for the given webhook_id.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.WebhookDeliveriesListOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			webhookId := params.GetString(helpers.ToSnakeCase("WebhookId"))
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("responseStatusCodes")) {
				localVarOptionals.ResponseStatusCodes = optional.NewString(params.GetString(helpers.ToSnakeCase("ResponseStatusCodes")))
			}

			data, api_response, err := client.WebhookDeliveriesApi.WebhookDeliveriesList(auth, projectId, webhookId, &localVarOptionals)

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

	WebhookDeliveriesApiCmd.AddCommand(WebhookDeliveriesList)
	AddFlag(WebhookDeliveriesList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(WebhookDeliveriesList, "string", helpers.ToSnakeCase("WebhookId"), "", "Webhook ID", true)
	AddFlag(WebhookDeliveriesList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(WebhookDeliveriesList, "string", helpers.ToSnakeCase("ResponseStatusCodes"), "", "List of Response Status Codes", false)

	params.BindPFlags(WebhookDeliveriesList.Flags())
}
func initWebhookDeliveriesRedeliver() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("webhook_deliveries/redeliver", "/")[1:], "_")
	var WebhookDeliveriesRedeliver = &cobra.Command{
		Use:   use,
		Short: "Redeliver a single webhook delivery",
		Long:  `Trigger an individual webhook delivery to be redelivered.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.WebhookDeliveriesRedeliverOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			webhookId := params.GetString(helpers.ToSnakeCase("WebhookId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.WebhookDeliveriesApi.WebhookDeliveriesRedeliver(auth, projectId, webhookId, id, &localVarOptionals)

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

	WebhookDeliveriesApiCmd.AddCommand(WebhookDeliveriesRedeliver)
	AddFlag(WebhookDeliveriesRedeliver, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(WebhookDeliveriesRedeliver, "string", helpers.ToSnakeCase("WebhookId"), "", "Webhook ID", true)
	AddFlag(WebhookDeliveriesRedeliver, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(WebhookDeliveriesRedeliver, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(WebhookDeliveriesRedeliver.Flags())
}
func initWebhookDeliveriesShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("webhook_deliveries/show", "/")[1:], "_")
	var WebhookDeliveriesShow = &cobra.Command{
		Use:   use,
		Short: "Get a single webhook delivery",
		Long:  `Get all information about a single webhook delivery for the given ID.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.WebhookDeliveriesShowOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			webhookId := params.GetString(helpers.ToSnakeCase("WebhookId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.WebhookDeliveriesApi.WebhookDeliveriesShow(auth, projectId, webhookId, id, &localVarOptionals)

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

	WebhookDeliveriesApiCmd.AddCommand(WebhookDeliveriesShow)
	AddFlag(WebhookDeliveriesShow, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(WebhookDeliveriesShow, "string", helpers.ToSnakeCase("WebhookId"), "", "Webhook ID", true)
	AddFlag(WebhookDeliveriesShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(WebhookDeliveriesShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(WebhookDeliveriesShow.Flags())
}
