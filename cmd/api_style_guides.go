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
	initStyleguideCreate()
	initStyleguideDelete()
	initStyleguideShow()
	initStyleguideUpdate()
	initStyleguidesList()

	rootCmd.AddCommand(styleGuidesApiCmd)
}

var styleGuidesApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("styleguidesapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("StyleGuidesApi", "Api"), "API"}, " "),
}


func initStyleguideCreate() {
	params := viper.New()
	var styleguideCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("StyleguideCreate", strings.TrimSuffix("StyleGuidesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("StyleGuidesApi", "Api"), "s"))),
		Short: "Create a style guide",
		Long:  `Create a new style guide.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.StyleguideCreateOpts{}

			
			projectId := params.GetString("projectId")

			

			styleguideCreateParameters := api.StyleguideCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &styleguideCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", styleguideCreateParameters)
			}
			

			data, api_response, err := client.StyleGuidesApi.StyleguideCreate(auth, projectId, styleguideCreateParameters, &localVarOptionals)

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

	styleGuidesApiCmd.AddCommand(styleguideCreate)

	
	AddFlag(styleguideCreate, "string", "projectId", "", "Project ID", true)
	
	AddFlag(styleguideCreate, "string", "data", "d", "payload in JSON format", true)
	// styleguideCreateParameters := api.StyleguideCreateParameters{}
	

	params.BindPFlags(styleguideCreate.Flags())
}

func initStyleguideDelete() {
	params := viper.New()
	var styleguideDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("StyleguideDelete", strings.TrimSuffix("StyleGuidesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("StyleGuidesApi", "Api"), "s"))),
		Short: "Delete a style guide",
		Long:  `Delete an existing style guide.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.StyleguideDeleteOpts{}

			
			projectId := params.GetString("projectId")

			
			id := params.GetString("id")

			

			data, api_response, err := client.StyleGuidesApi.StyleguideDelete(auth, projectId, id, &localVarOptionals)

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

	styleGuidesApiCmd.AddCommand(styleguideDelete)

	
	AddFlag(styleguideDelete, "string", "projectId", "", "Project ID", true)
	
	AddFlag(styleguideDelete, "string", "id", "", "ID", true)
	

	params.BindPFlags(styleguideDelete.Flags())
}

func initStyleguideShow() {
	params := viper.New()
	var styleguideShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("StyleguideShow", strings.TrimSuffix("StyleGuidesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("StyleGuidesApi", "Api"), "s"))),
		Short: "Get a single style guide",
		Long:  `Get details on a single style guide.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.StyleguideShowOpts{}

			
			projectId := params.GetString("projectId")

			
			id := params.GetString("id")

			

			data, api_response, err := client.StyleGuidesApi.StyleguideShow(auth, projectId, id, &localVarOptionals)

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

	styleGuidesApiCmd.AddCommand(styleguideShow)

	
	AddFlag(styleguideShow, "string", "projectId", "", "Project ID", true)
	
	AddFlag(styleguideShow, "string", "id", "", "ID", true)
	

	params.BindPFlags(styleguideShow.Flags())
}

func initStyleguideUpdate() {
	params := viper.New()
	var styleguideUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("StyleguideUpdate", strings.TrimSuffix("StyleGuidesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("StyleGuidesApi", "Api"), "s"))),
		Short: "Update a style guide",
		Long:  `Update an existing style guide.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.StyleguideUpdateOpts{}

			
			projectId := params.GetString("projectId")

			
			id := params.GetString("id")

			

			styleguideUpdateParameters := api.StyleguideUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &styleguideUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", styleguideUpdateParameters)
			}
			

			data, api_response, err := client.StyleGuidesApi.StyleguideUpdate(auth, projectId, id, styleguideUpdateParameters, &localVarOptionals)

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

	styleGuidesApiCmd.AddCommand(styleguideUpdate)

	
	AddFlag(styleguideUpdate, "string", "projectId", "", "Project ID", true)
	
	AddFlag(styleguideUpdate, "string", "id", "", "ID", true)
	
	AddFlag(styleguideUpdate, "string", "data", "d", "payload in JSON format", true)
	// styleguideUpdateParameters := api.StyleguideUpdateParameters{}
	

	params.BindPFlags(styleguideUpdate.Flags())
}

func initStyleguidesList() {
	params := viper.New()
	var styleguidesList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("StyleguidesList", strings.TrimSuffix("StyleGuidesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("StyleGuidesApi", "Api"), "s"))),
		Short: "List style guides",
		Long:  `List all styleguides for the given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.StyleguidesListOpts{}

			
			projectId := params.GetString("projectId")

			

			data, api_response, err := client.StyleGuidesApi.StyleguidesList(auth, projectId, &localVarOptionals)

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

	styleGuidesApiCmd.AddCommand(styleguidesList)

	
	AddFlag(styleguidesList, "string", "projectId", "", "Project ID", true)
	

	params.BindPFlags(styleguidesList.Flags())
}

