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
	initUploadCreate()
	initUploadShow()
	initUploadsList()

	rootCmd.AddCommand(uploadsApiCmd)
}

var uploadsApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("uploadsapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("UploadsApi", "Api"), "API"}, " "),
}


func initUploadCreate() {
	params := viper.New()
	var uploadCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("UploadCreate", strings.TrimSuffix("UploadsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("UploadsApi", "Api"), "s"))),
		Short: "Upload a new file",
		Long:  `Upload a new language file. Creates necessary resources in your project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.UploadCreateOpts{}

			
			projectId := params.GetString("projectId")

			

			uploadCreateParameters := api.UploadCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &uploadCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", uploadCreateParameters)
			}
			

			data, api_response, err := client.UploadsApi.UploadCreate(auth, projectId, uploadCreateParameters, &localVarOptionals)

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

	uploadsApiCmd.AddCommand(uploadCreate)

	
	AddFlag(uploadCreate, "string", "projectId", "", "Project ID", true)
	
	AddFlag(uploadCreate, "string", "data", "d", "payload in JSON format", true)
	// uploadCreateParameters := api.UploadCreateParameters{}
	

	params.BindPFlags(uploadCreate.Flags())
}

func initUploadShow() {
	params := viper.New()
	var uploadShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("UploadShow", strings.TrimSuffix("UploadsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("UploadsApi", "Api"), "s"))),
		Short: "View upload details",
		Long:  `View details and summary for a single upload.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.UploadShowOpts{}

			
			projectId := params.GetString("projectId")

			
			id := params.GetString("id")

			

			data, api_response, err := client.UploadsApi.UploadShow(auth, projectId, id, &localVarOptionals)

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

	uploadsApiCmd.AddCommand(uploadShow)

	
	AddFlag(uploadShow, "string", "projectId", "", "Project ID", true)
	
	AddFlag(uploadShow, "string", "id", "", "ID", true)
	

	params.BindPFlags(uploadShow.Flags())
}

func initUploadsList() {
	params := viper.New()
	var uploadsList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("UploadsList", strings.TrimSuffix("UploadsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("UploadsApi", "Api"), "s"))),
		Short: "List uploads",
		Long:  `List all uploads for the given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.UploadsListOpts{}

			
			projectId := params.GetString("projectId")

			

			data, api_response, err := client.UploadsApi.UploadsList(auth, projectId, &localVarOptionals)

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

	uploadsApiCmd.AddCommand(uploadsList)

	
	AddFlag(uploadsList, "string", "projectId", "", "Project ID", true)
	

	params.BindPFlags(uploadsList.Flags())
}

