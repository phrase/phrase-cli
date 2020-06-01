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
	initWebhookCreate()
	initWebhookDelete()
	initWebhookShow()
	initWebhookTest()
	initWebhookUpdate()
	initWebhooksList()

	rootCmd.AddCommand(WebhooksApiCmd)
}

var WebhooksApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("Webhooks"),
	Short: "Webhooks API",
}

func initWebhookCreate() {
	params := viper.New()
	var WebhookCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("WebhookCreate", strings.TrimSuffix("WebhooksApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("WebhooksApi", "Api"), "s"))),
		Short: "Create a webhook",
		Long:  `Create a new webhook.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.WebhookCreateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			webhookCreateParameters := api.WebhookCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &webhookCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", webhookCreateParameters)
			}
			data, api_response, err := client.WebhooksApi.WebhookCreate(auth, projectId, webhookCreateParameters, &localVarOptionals)

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

	WebhooksApiCmd.AddCommand(WebhookCreate)

	AddFlag(WebhookCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(WebhookCreate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(WebhookCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(WebhookCreate.Flags())
}
func initWebhookDelete() {
	params := viper.New()
	var WebhookDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("WebhookDelete", strings.TrimSuffix("WebhooksApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("WebhooksApi", "Api"), "s"))),
		Short: "Delete a webhook",
		Long:  `Delete an existing webhook.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.WebhookDeleteOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.WebhooksApi.WebhookDelete(auth, projectId, id, &localVarOptionals)

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

	WebhooksApiCmd.AddCommand(WebhookDelete)

	AddFlag(WebhookDelete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(WebhookDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(WebhookDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(WebhookDelete.Flags())
}
func initWebhookShow() {
	params := viper.New()
	var WebhookShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("WebhookShow", strings.TrimSuffix("WebhooksApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("WebhooksApi", "Api"), "s"))),
		Short: "Get a single webhook",
		Long:  `Get details on a single webhook.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.WebhookShowOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.WebhooksApi.WebhookShow(auth, projectId, id, &localVarOptionals)

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

	WebhooksApiCmd.AddCommand(WebhookShow)

	AddFlag(WebhookShow, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(WebhookShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(WebhookShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(WebhookShow.Flags())
}
func initWebhookTest() {
	params := viper.New()
	var WebhookTest = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("WebhookTest", strings.TrimSuffix("WebhooksApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("WebhooksApi", "Api"), "s"))),
		Short: "Test a webhook",
		Long:  `Perform a test request for a webhook.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.WebhookTestOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.WebhooksApi.WebhookTest(auth, projectId, id, &localVarOptionals)

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

	WebhooksApiCmd.AddCommand(WebhookTest)

	AddFlag(WebhookTest, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(WebhookTest, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(WebhookTest, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(WebhookTest.Flags())
}
func initWebhookUpdate() {
	params := viper.New()
	var WebhookUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("WebhookUpdate", strings.TrimSuffix("WebhooksApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("WebhooksApi", "Api"), "s"))),
		Short: "Update a webhook",
		Long:  `Update an existing webhook.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.WebhookUpdateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			webhookUpdateParameters := api.WebhookUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &webhookUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", webhookUpdateParameters)
			}
			data, api_response, err := client.WebhooksApi.WebhookUpdate(auth, projectId, id, webhookUpdateParameters, &localVarOptionals)

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

	WebhooksApiCmd.AddCommand(WebhookUpdate)

	AddFlag(WebhookUpdate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(WebhookUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(WebhookUpdate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(WebhookUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(WebhookUpdate.Flags())
}
func initWebhooksList() {
	params := viper.New()
	var WebhooksList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("WebhooksList", strings.TrimSuffix("WebhooksApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("WebhooksApi", "Api"), "s"))),
		Short: "List webhooks",
		Long:  `List all webhooks for the given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.WebhooksListOpts{}

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

			data, api_response, err := client.WebhooksApi.WebhooksList(auth, projectId, &localVarOptionals)

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

	WebhooksApiCmd.AddCommand(WebhooksList)

	AddFlag(WebhooksList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(WebhooksList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(WebhooksList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(WebhooksList, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 10 by default", false)
	params.BindPFlags(WebhooksList.Flags())
}
