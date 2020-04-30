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
	initScreenshotMarkerCreate()
	initScreenshotMarkerDelete()
	initScreenshotMarkerShow()
	initScreenshotMarkerUpdate()
	initScreenshotMarkersList()

	rootCmd.AddCommand(screenshotMarkersApiCmd)
}

var screenshotMarkersApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("screenshotmarkersapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("ScreenshotMarkersApi", "Api"), "API"}, " "),
}


func initScreenshotMarkerCreate() {
	params := viper.New()
	var screenshotMarkerCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ScreenshotMarkerCreate", strings.TrimSuffix("ScreenshotMarkersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ScreenshotMarkersApi", "Api"), "s"))),
		Short: "Create a screenshot marker",
		Long:  `Create a new screenshot marker.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ScreenshotMarkerCreateOpts{}

			
			projectId := params.GetString("projectId")
			
			screenshotId := params.GetString("screenshotId")
			
			screenshotMarkerCreateParameters := api.ScreenshotMarkerCreateParameters{}
			

			api_response, err := client.ScreenshotMarkersApi.ScreenshotMarkerCreate(auth, projectId, screenshotId, screenshotMarkerCreateParameters, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	screenshotMarkersApiCmd.AddCommand(screenshotMarkerCreate)

	
	AddFlag(screenshotMarkerCreate, "string", "projectId", "", "ID")
	
	AddFlag(screenshotMarkerCreate, "string", "screenshotId", "", "ID")
	
	// screenshotMarkerCreateParameters := api.ScreenshotMarkerCreateParameters{}
	

	params.BindPFlags(screenshotMarkerCreate.Flags())
}

func initScreenshotMarkerDelete() {
	params := viper.New()
	var screenshotMarkerDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ScreenshotMarkerDelete", strings.TrimSuffix("ScreenshotMarkersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ScreenshotMarkersApi", "Api"), "s"))),
		Short: "Delete a screenshot marker",
		Long:  `Delete an existing screenshot marker.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ScreenshotMarkerDeleteOpts{}

			
			projectId := params.GetString("projectId")
			
			screenshotId := params.GetString("screenshotId")
			

			api_response, err := client.ScreenshotMarkersApi.ScreenshotMarkerDelete(auth, projectId, screenshotId, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	screenshotMarkersApiCmd.AddCommand(screenshotMarkerDelete)

	
	AddFlag(screenshotMarkerDelete, "string", "projectId", "", "ID")
	
	AddFlag(screenshotMarkerDelete, "string", "screenshotId", "", "ID")
	

	params.BindPFlags(screenshotMarkerDelete.Flags())
}

func initScreenshotMarkerShow() {
	params := viper.New()
	var screenshotMarkerShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ScreenshotMarkerShow", strings.TrimSuffix("ScreenshotMarkersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ScreenshotMarkersApi", "Api"), "s"))),
		Short: "Get a single screenshot marker",
		Long:  `Get details on a single screenshot marker for a given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ScreenshotMarkerShowOpts{}

			
			projectId := params.GetString("projectId")
			
			screenshotId := params.GetString("screenshotId")
			
			id := params.GetString("id")
			

			data, api_response, err := client.ScreenshotMarkersApi.ScreenshotMarkerShow(auth, projectId, screenshotId, id, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	screenshotMarkersApiCmd.AddCommand(screenshotMarkerShow)

	
	AddFlag(screenshotMarkerShow, "string", "projectId", "", "ID")
	
	AddFlag(screenshotMarkerShow, "string", "screenshotId", "", "ID")
	
	AddFlag(screenshotMarkerShow, "string", "id", "", "ID")
	

	params.BindPFlags(screenshotMarkerShow.Flags())
}

func initScreenshotMarkerUpdate() {
	params := viper.New()
	var screenshotMarkerUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ScreenshotMarkerUpdate", strings.TrimSuffix("ScreenshotMarkersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ScreenshotMarkersApi", "Api"), "s"))),
		Short: "Update a screenshot marker",
		Long:  `Update an existing screenshot marker.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ScreenshotMarkerUpdateOpts{}

			
			projectId := params.GetString("projectId")
			
			screenshotId := params.GetString("screenshotId")
			
			screenshotMarkerUpdateParameters := api.ScreenshotMarkerUpdateParameters{}
			

			data, api_response, err := client.ScreenshotMarkersApi.ScreenshotMarkerUpdate(auth, projectId, screenshotId, screenshotMarkerUpdateParameters, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	screenshotMarkersApiCmd.AddCommand(screenshotMarkerUpdate)

	
	AddFlag(screenshotMarkerUpdate, "string", "projectId", "", "ID")
	
	AddFlag(screenshotMarkerUpdate, "string", "screenshotId", "", "ID")
	
	// screenshotMarkerUpdateParameters := api.ScreenshotMarkerUpdateParameters{}
	

	params.BindPFlags(screenshotMarkerUpdate.Flags())
}

func initScreenshotMarkersList() {
	params := viper.New()
	var screenshotMarkersList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ScreenshotMarkersList", strings.TrimSuffix("ScreenshotMarkersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ScreenshotMarkersApi", "Api"), "s"))),
		Short: "List screenshot markers",
		Long:  `List all screenshot markers for the given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ScreenshotMarkersListOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			

			data, api_response, err := client.ScreenshotMarkersApi.ScreenshotMarkersList(auth, projectId, id, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	screenshotMarkersApiCmd.AddCommand(screenshotMarkersList)

	
	AddFlag(screenshotMarkersList, "string", "projectId", "", "ID")
	
	AddFlag(screenshotMarkersList, "string", "id", "", "ID")
	

	params.BindPFlags(screenshotMarkersList.Flags())
}

