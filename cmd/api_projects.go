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
	initProjectCreate()
	initProjectDelete()
	initProjectShow()
	initProjectUpdate()
	initProjectsList()

	rootCmd.AddCommand(projectsApiCmd)
}

var projectsApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("projectsapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("ProjectsApi", "Api"), "API"}, " "),
}


func initProjectCreate() {
	params := viper.New()
	var projectCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ProjectCreate", strings.TrimSuffix("ProjectsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ProjectsApi", "Api"), "s"))),
		Short: "Create a project",
		Long:  `Create a new project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ProjectCreateOpts{}

			
			projectCreateParameters := api.ProjectCreateParameters{}
			

			api_response, err := client.ProjectsApi.ProjectCreate(auth, projectCreateParameters, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	projectsApiCmd.AddCommand(projectCreate)

	
	// projectCreateParameters := api.ProjectCreateParameters{}
	

	params.BindPFlags(projectCreate.Flags())
}

func initProjectDelete() {
	params := viper.New()
	var projectDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ProjectDelete", strings.TrimSuffix("ProjectsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ProjectsApi", "Api"), "s"))),
		Short: "Delete a project",
		Long:  `Delete an existing project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ProjectDeleteOpts{}

			
			id := params.GetString("id")
			

			api_response, err := client.ProjectsApi.ProjectDelete(auth, id, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	projectsApiCmd.AddCommand(projectDelete)

	
	AddFlag(projectDelete, "string", "id", "", "ID")
	

	params.BindPFlags(projectDelete.Flags())
}

func initProjectShow() {
	params := viper.New()
	var projectShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ProjectShow", strings.TrimSuffix("ProjectsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ProjectsApi", "Api"), "s"))),
		Short: "Get a single project",
		Long:  `Get details on a single project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ProjectShowOpts{}

			
			id := params.GetString("id")
			

			data, api_response, err := client.ProjectsApi.ProjectShow(auth, id, &localVarOptionals)

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

	projectsApiCmd.AddCommand(projectShow)

	
	AddFlag(projectShow, "string", "id", "", "ID")
	

	params.BindPFlags(projectShow.Flags())
}

func initProjectUpdate() {
	params := viper.New()
	var projectUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ProjectUpdate", strings.TrimSuffix("ProjectsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ProjectsApi", "Api"), "s"))),
		Short: "Update a project",
		Long:  `Update an existing project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ProjectUpdateOpts{}

			
			id := params.GetString("id")
			
			projectUpdateParameters := api.ProjectUpdateParameters{}
			

			data, api_response, err := client.ProjectsApi.ProjectUpdate(auth, id, projectUpdateParameters, &localVarOptionals)

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

	projectsApiCmd.AddCommand(projectUpdate)

	
	AddFlag(projectUpdate, "string", "id", "", "ID")
	
	// projectUpdateParameters := api.ProjectUpdateParameters{}
	

	params.BindPFlags(projectUpdate.Flags())
}

func initProjectsList() {
	params := viper.New()
	var projectsList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ProjectsList", strings.TrimSuffix("ProjectsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ProjectsApi", "Api"), "s"))),
		Short: "List projects",
		Long:  `List all projects the current user has access to.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ProjectsListOpts{}

			

			data, api_response, err := client.ProjectsApi.ProjectsList(auth, &localVarOptionals)

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

	projectsApiCmd.AddCommand(projectsList)

	

	params.BindPFlags(projectsList.Flags())
}

