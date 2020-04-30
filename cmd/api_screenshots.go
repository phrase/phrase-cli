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
	initScreenshotCreate()
	initScreenshotDelete()
	initScreenshotShow()
	initScreenshotUpdate()
	initScreenshotsList()

	rootCmd.AddCommand(screenshotsApiCmd)
}

var screenshotsApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("screenshotsapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("ScreenshotsApi", "Api"), "API"}, " "),
}


func initScreenshotCreate() {
	params := viper.New()
	var screenshotCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ScreenshotCreate", strings.TrimSuffix("ScreenshotsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ScreenshotsApi", "Api"), "s"))),
		Short: "Create a screenshot",
		Long:  `Create a new screenshot.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ScreenshotCreateOpts{}

			
			projectId := params.GetString("projectId")
			
			screenshotCreateParameters := api.ScreenshotCreateParameters{}
			

			api_response, err := client.ScreenshotsApi.ScreenshotCreate(auth, projectId, screenshotCreateParameters, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	screenshotsApiCmd.AddCommand(screenshotCreate)

	
	AddFlag(screenshotCreate, "string", "projectId", "", "ID")
	
	// screenshotCreateParameters := api.ScreenshotCreateParameters{}
	

	params.BindPFlags(screenshotCreate.Flags())
}

func initScreenshotDelete() {
	params := viper.New()
	var screenshotDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ScreenshotDelete", strings.TrimSuffix("ScreenshotsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ScreenshotsApi", "Api"), "s"))),
		Short: "Delete a screenshot",
		Long:  `Delete an existing screenshot.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ScreenshotDeleteOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			

			api_response, err := client.ScreenshotsApi.ScreenshotDelete(auth, projectId, id, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	screenshotsApiCmd.AddCommand(screenshotDelete)

	
	AddFlag(screenshotDelete, "string", "projectId", "", "ID")
	
	AddFlag(screenshotDelete, "string", "id", "", "ID")
	

	params.BindPFlags(screenshotDelete.Flags())
}

func initScreenshotShow() {
	params := viper.New()
	var screenshotShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ScreenshotShow", strings.TrimSuffix("ScreenshotsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ScreenshotsApi", "Api"), "s"))),
		Short: "Get a single screenshot",
		Long:  `Get details on a single screenshot for a given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ScreenshotShowOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			

			data, api_response, err := client.ScreenshotsApi.ScreenshotShow(auth, projectId, id, &localVarOptionals)

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

	screenshotsApiCmd.AddCommand(screenshotShow)

	
	AddFlag(screenshotShow, "string", "projectId", "", "ID")
	
	AddFlag(screenshotShow, "string", "id", "", "ID")
	

	params.BindPFlags(screenshotShow.Flags())
}

func initScreenshotUpdate() {
	params := viper.New()
	var screenshotUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ScreenshotUpdate", strings.TrimSuffix("ScreenshotsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ScreenshotsApi", "Api"), "s"))),
		Short: "Update a screenshot",
		Long:  `Update an existing screenshot.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ScreenshotUpdateOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			
			screenshotUpdateParameters := api.ScreenshotUpdateParameters{}
			

			data, api_response, err := client.ScreenshotsApi.ScreenshotUpdate(auth, projectId, id, screenshotUpdateParameters, &localVarOptionals)

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

	screenshotsApiCmd.AddCommand(screenshotUpdate)

	
	AddFlag(screenshotUpdate, "string", "projectId", "", "ID")
	
	AddFlag(screenshotUpdate, "string", "id", "", "ID")
	
	// screenshotUpdateParameters := api.ScreenshotUpdateParameters{}
	

	params.BindPFlags(screenshotUpdate.Flags())
}

func initScreenshotsList() {
	params := viper.New()
	var screenshotsList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ScreenshotsList", strings.TrimSuffix("ScreenshotsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ScreenshotsApi", "Api"), "s"))),
		Short: "List screenshots",
		Long:  `List all screenshots for the given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ScreenshotsListOpts{}

			
			projectId := params.GetString("projectId")
			

			data, api_response, err := client.ScreenshotsApi.ScreenshotsList(auth, projectId, &localVarOptionals)

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

	screenshotsApiCmd.AddCommand(screenshotsList)

	
	AddFlag(screenshotsList, "string", "projectId", "", "ID")
	

	params.BindPFlags(screenshotsList.Flags())
}

