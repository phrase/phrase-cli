package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

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

	rootCmd.AddCommand(webhooksApiCmd)
}

var webhooksApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("webhooksapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("WebhooksApi", "Api"), "API"}, " "),
}


func initWebhookCreate() {
	params := viper.New()
	var webhookCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("WebhookCreate", strings.TrimSuffix("WebhooksApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("WebhooksApi", "Api"), "s"))),
		Short: "Create a webhook",
		Long:  `Create a new webhook.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.WebhookCreateOpts{}

			
			projectId := params.GetString("projectId")
			
			webhookCreateParameters := api.WebhookCreateParameters{}
			

			data, api_response, err := client.WebhooksApi.WebhookCreate(auth, projectId, webhookCreateParameters, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	webhooksApiCmd.AddCommand(webhookCreate)

	
	AddFlag(webhookCreate, "string", "projectId", "", "ID")
	
	// webhookCreateParameters := api.WebhookCreateParameters{}
	

	params.BindPFlags(webhookCreate.Flags())
}

func initWebhookDelete() {
	params := viper.New()
	var webhookDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("WebhookDelete", strings.TrimSuffix("WebhooksApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("WebhooksApi", "Api"), "s"))),
		Short: "Delete a webhook",
		Long:  `Delete an existing webhook.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.WebhookDeleteOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			

			data, api_response, err := client.WebhooksApi.WebhookDelete(auth, projectId, id, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	webhooksApiCmd.AddCommand(webhookDelete)

	
	AddFlag(webhookDelete, "string", "projectId", "", "ID")
	
	AddFlag(webhookDelete, "string", "id", "", "ID")
	

	params.BindPFlags(webhookDelete.Flags())
}

func initWebhookShow() {
	params := viper.New()
	var webhookShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("WebhookShow", strings.TrimSuffix("WebhooksApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("WebhooksApi", "Api"), "s"))),
		Short: "Get a single webhook",
		Long:  `Get details on a single webhook.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.WebhookShowOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			

			data, api_response, err := client.WebhooksApi.WebhookShow(auth, projectId, id, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	webhooksApiCmd.AddCommand(webhookShow)

	
	AddFlag(webhookShow, "string", "projectId", "", "ID")
	
	AddFlag(webhookShow, "string", "id", "", "ID")
	

	params.BindPFlags(webhookShow.Flags())
}

func initWebhookTest() {
	params := viper.New()
	var webhookTest = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("WebhookTest", strings.TrimSuffix("WebhooksApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("WebhooksApi", "Api"), "s"))),
		Short: "Test a webhook",
		Long:  `Perform a test request for a webhook.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.WebhookTestOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			

			data, api_response, err := client.WebhooksApi.WebhookTest(auth, projectId, id, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	webhooksApiCmd.AddCommand(webhookTest)

	
	AddFlag(webhookTest, "string", "projectId", "", "ID")
	
	AddFlag(webhookTest, "string", "id", "", "ID")
	

	params.BindPFlags(webhookTest.Flags())
}

func initWebhookUpdate() {
	params := viper.New()
	var webhookUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("WebhookUpdate", strings.TrimSuffix("WebhooksApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("WebhooksApi", "Api"), "s"))),
		Short: "Update a webhook",
		Long:  `Update an existing webhook.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.WebhookUpdateOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			
			webhookUpdateParameters := api.WebhookUpdateParameters{}
			

			data, api_response, err := client.WebhooksApi.WebhookUpdate(auth, projectId, id, webhookUpdateParameters, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	webhooksApiCmd.AddCommand(webhookUpdate)

	
	AddFlag(webhookUpdate, "string", "projectId", "", "ID")
	
	AddFlag(webhookUpdate, "string", "id", "", "ID")
	
	// webhookUpdateParameters := api.WebhookUpdateParameters{}
	

	params.BindPFlags(webhookUpdate.Flags())
}

func initWebhooksList() {
	params := viper.New()
	var webhooksList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("WebhooksList", strings.TrimSuffix("WebhooksApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("WebhooksApi", "Api"), "s"))),
		Short: "List webhooks",
		Long:  `List all webhooks for the given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.WebhooksListOpts{}

			
			projectId := params.GetString("projectId")
			

			data, api_response, err := client.WebhooksApi.WebhooksList(auth, projectId, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	webhooksApiCmd.AddCommand(webhooksList)

	
	AddFlag(webhooksList, "string", "projectId", "", "ID")
	

	params.BindPFlags(webhooksList.Flags())
}

