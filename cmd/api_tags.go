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
	initTagCreate()
	initTagDelete()
	initTagShow()
	initTagsList()

	rootCmd.AddCommand(tagsApiCmd)
}

var tagsApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("tagsapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("TagsApi", "Api"), "API"}, " "),
}


func initTagCreate() {
	params := viper.New()
	var tagCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("TagCreate", strings.TrimSuffix("TagsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("TagsApi", "Api"), "s"))),
		Short: "Create a tag",
		Long:  `Create a new tag.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TagCreateOpts{}

			
			projectId := params.GetString("projectId")

			

			tagCreateParameters := api.TagCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &tagCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", tagCreateParameters)
			}
			

			data, api_response, err := client.TagsApi.TagCreate(auth, projectId, tagCreateParameters, &localVarOptionals)

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

	tagsApiCmd.AddCommand(tagCreate)

	
	AddFlag(tagCreate, "string", "projectId", "", "Project ID", true)
	
	AddFlag(tagCreate, "string", "data", "d", "payload in JSON format", true)
	// tagCreateParameters := api.TagCreateParameters{}
	

	params.BindPFlags(tagCreate.Flags())
}

func initTagDelete() {
	params := viper.New()
	var tagDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("TagDelete", strings.TrimSuffix("TagsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("TagsApi", "Api"), "s"))),
		Short: "Delete a tag",
		Long:  `Delete an existing tag.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TagDeleteOpts{}

			
			projectId := params.GetString("projectId")

			
			name := params.GetString("name")

			

			data, api_response, err := client.TagsApi.TagDelete(auth, projectId, name, &localVarOptionals)

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

	tagsApiCmd.AddCommand(tagDelete)

	
	AddFlag(tagDelete, "string", "projectId", "", "Project ID", true)
	
	AddFlag(tagDelete, "string", "name", "", "name", true)
	

	params.BindPFlags(tagDelete.Flags())
}

func initTagShow() {
	params := viper.New()
	var tagShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("TagShow", strings.TrimSuffix("TagsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("TagsApi", "Api"), "s"))),
		Short: "Get a single tag",
		Long:  `Get details and progress information on a single tag for a given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TagShowOpts{}

			
			projectId := params.GetString("projectId")

			
			name := params.GetString("name")

			

			data, api_response, err := client.TagsApi.TagShow(auth, projectId, name, &localVarOptionals)

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

	tagsApiCmd.AddCommand(tagShow)

	
	AddFlag(tagShow, "string", "projectId", "", "Project ID", true)
	
	AddFlag(tagShow, "string", "name", "", "name", true)
	

	params.BindPFlags(tagShow.Flags())
}

func initTagsList() {
	params := viper.New()
	var tagsList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("TagsList", strings.TrimSuffix("TagsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("TagsApi", "Api"), "s"))),
		Short: "List tags",
		Long:  `List all tags for the given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TagsListOpts{}

			
			projectId := params.GetString("projectId")

			

			data, api_response, err := client.TagsApi.TagsList(auth, projectId, &localVarOptionals)

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

	tagsApiCmd.AddCommand(tagsList)

	
	AddFlag(tagsList, "string", "projectId", "", "Project ID", true)
	

	params.BindPFlags(tagsList.Flags())
}

