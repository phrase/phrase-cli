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
	initSpaceCreate()
	initSpaceDelete()
	initSpaceShow()
	initSpaceUpdate()
	initSpacesList()
	initSpacesProjectsCreate()
	initSpacesProjectsDelete()
	initSpacesProjectsList()

	rootCmd.AddCommand(spacesApiCmd)
}

var spacesApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("spacesapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("SpacesApi", "Api"), "API"}, " "),
}


func initSpaceCreate() {
	params := viper.New()
	var spaceCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("SpaceCreate", strings.TrimSuffix("SpacesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("SpacesApi", "Api"), "s"))),
		Short: "Create a Space",
		Long:  `Create a new Space.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.SpaceCreateOpts{}

			
			accountId := params.GetString("accountId")

			

			spaceCreateParameters := api.SpaceCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &spaceCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", spaceCreateParameters)
			}
			

			data, api_response, err := client.SpacesApi.SpaceCreate(auth, accountId, spaceCreateParameters, &localVarOptionals)

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

	spacesApiCmd.AddCommand(spaceCreate)

	
	AddFlag(spaceCreate, "string", "accountId", "", "Account ID", true)
	
	AddFlag(spaceCreate, "string", "data", "d", "payload in JSON format", true)
	// spaceCreateParameters := api.SpaceCreateParameters{}
	

	params.BindPFlags(spaceCreate.Flags())
}

func initSpaceDelete() {
	params := viper.New()
	var spaceDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("SpaceDelete", strings.TrimSuffix("SpacesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("SpacesApi", "Api"), "s"))),
		Short: "Delete Space",
		Long:  `Delete the specified Space.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.SpaceDeleteOpts{}

			
			accountId := params.GetString("accountId")

			
			id := params.GetString("id")

			

			data, api_response, err := client.SpacesApi.SpaceDelete(auth, accountId, id, &localVarOptionals)

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

	spacesApiCmd.AddCommand(spaceDelete)

	
	AddFlag(spaceDelete, "string", "accountId", "", "Account ID", true)
	
	AddFlag(spaceDelete, "string", "id", "", "ID", true)
	

	params.BindPFlags(spaceDelete.Flags())
}

func initSpaceShow() {
	params := viper.New()
	var spaceShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("SpaceShow", strings.TrimSuffix("SpacesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("SpacesApi", "Api"), "s"))),
		Short: "Get Space",
		Long:  `Show the specified Space.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.SpaceShowOpts{}

			
			accountId := params.GetString("accountId")

			
			id := params.GetString("id")

			

			data, api_response, err := client.SpacesApi.SpaceShow(auth, accountId, id, &localVarOptionals)

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

	spacesApiCmd.AddCommand(spaceShow)

	
	AddFlag(spaceShow, "string", "accountId", "", "Account ID", true)
	
	AddFlag(spaceShow, "string", "id", "", "ID", true)
	

	params.BindPFlags(spaceShow.Flags())
}

func initSpaceUpdate() {
	params := viper.New()
	var spaceUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("SpaceUpdate", strings.TrimSuffix("SpacesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("SpacesApi", "Api"), "s"))),
		Short: "Update Space",
		Long:  `Update the specified Space.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.SpaceUpdateOpts{}

			
			accountId := params.GetString("accountId")

			
			id := params.GetString("id")

			

			spaceUpdateParameters := api.SpaceUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &spaceUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", spaceUpdateParameters)
			}
			

			data, api_response, err := client.SpacesApi.SpaceUpdate(auth, accountId, id, spaceUpdateParameters, &localVarOptionals)

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

	spacesApiCmd.AddCommand(spaceUpdate)

	
	AddFlag(spaceUpdate, "string", "accountId", "", "Account ID", true)
	
	AddFlag(spaceUpdate, "string", "id", "", "ID", true)
	
	AddFlag(spaceUpdate, "string", "data", "d", "payload in JSON format", true)
	// spaceUpdateParameters := api.SpaceUpdateParameters{}
	

	params.BindPFlags(spaceUpdate.Flags())
}

func initSpacesList() {
	params := viper.New()
	var spacesList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("SpacesList", strings.TrimSuffix("SpacesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("SpacesApi", "Api"), "s"))),
		Short: "List Spaces",
		Long:  `List all Spaces for the given account.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.SpacesListOpts{}

			
			accountId := params.GetString("accountId")

			

			data, api_response, err := client.SpacesApi.SpacesList(auth, accountId, &localVarOptionals)

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

	spacesApiCmd.AddCommand(spacesList)

	
	AddFlag(spacesList, "string", "accountId", "", "Account ID", true)
	

	params.BindPFlags(spacesList.Flags())
}

func initSpacesProjectsCreate() {
	params := viper.New()
	var spacesProjectsCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("SpacesProjectsCreate", strings.TrimSuffix("SpacesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("SpacesApi", "Api"), "s"))),
		Short: "Add Project",
		Long:  `Adds an existing project to the space.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.SpacesProjectsCreateOpts{}

			
			accountId := params.GetString("accountId")

			
			spaceId := params.GetString("spaceId")

			

			spacesProjectsCreateParameters := api.SpacesProjectsCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &spacesProjectsCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", spacesProjectsCreateParameters)
			}
			

			data, api_response, err := client.SpacesApi.SpacesProjectsCreate(auth, accountId, spaceId, spacesProjectsCreateParameters, &localVarOptionals)

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

	spacesApiCmd.AddCommand(spacesProjectsCreate)

	
	AddFlag(spacesProjectsCreate, "string", "accountId", "", "Account ID", true)
	
	AddFlag(spacesProjectsCreate, "string", "spaceId", "", "Space ID", true)
	
	AddFlag(spacesProjectsCreate, "string", "data", "d", "payload in JSON format", true)
	// spacesProjectsCreateParameters := api.SpacesProjectsCreateParameters{}
	

	params.BindPFlags(spacesProjectsCreate.Flags())
}

func initSpacesProjectsDelete() {
	params := viper.New()
	var spacesProjectsDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("SpacesProjectsDelete", strings.TrimSuffix("SpacesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("SpacesApi", "Api"), "s"))),
		Short: "Remove Project",
		Long:  `Removes a specified project from the specified space.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.SpacesProjectsDeleteOpts{}

			
			accountId := params.GetString("accountId")

			
			spaceId := params.GetString("spaceId")

			
			id := params.GetString("id")

			

			data, api_response, err := client.SpacesApi.SpacesProjectsDelete(auth, accountId, spaceId, id, &localVarOptionals)

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

	spacesApiCmd.AddCommand(spacesProjectsDelete)

	
	AddFlag(spacesProjectsDelete, "string", "accountId", "", "Account ID", true)
	
	AddFlag(spacesProjectsDelete, "string", "spaceId", "", "Space ID", true)
	
	AddFlag(spacesProjectsDelete, "string", "id", "", "ID", true)
	

	params.BindPFlags(spacesProjectsDelete.Flags())
}

func initSpacesProjectsList() {
	params := viper.New()
	var spacesProjectsList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("SpacesProjectsList", strings.TrimSuffix("SpacesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("SpacesApi", "Api"), "s"))),
		Short: "List Projects",
		Long:  `List all projects for the specified Space.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.SpacesProjectsListOpts{}

			
			accountId := params.GetString("accountId")

			
			spaceId := params.GetString("spaceId")

			

			data, api_response, err := client.SpacesApi.SpacesProjectsList(auth, accountId, spaceId, &localVarOptionals)

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

	spacesApiCmd.AddCommand(spacesProjectsList)

	
	AddFlag(spacesProjectsList, "string", "accountId", "", "Account ID", true)
	
	AddFlag(spacesProjectsList, "string", "spaceId", "", "Space ID", true)
	

	params.BindPFlags(spacesProjectsList.Flags())
}

