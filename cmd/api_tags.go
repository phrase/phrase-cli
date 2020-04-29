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
				Key:    Config.AccessToken,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TagCreateOpts{}

			
			projectId := params.GetString("projectId")
			
			tagCreateParameters := api.TagCreateParameters{}
			

			api_response, err := client.TagsApi.TagCreate(auth, projectId, tagCreateParameters, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	tagsApiCmd.AddCommand(tagCreate)

	
	AddFlag(tagCreate, "string", "projectId", "", "ID")
	
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
				Key:    Config.AccessToken,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TagDeleteOpts{}

			
			projectId := params.GetString("projectId")
			
			name := params.GetString("name")
			
			tagDeleteParameters := api.TagDeleteParameters{}
			

			api_response, err := client.TagsApi.TagDelete(auth, projectId, name, tagDeleteParameters, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	tagsApiCmd.AddCommand(tagDelete)

	
	AddFlag(tagDelete, "string", "projectId", "", "ID")
	
	AddFlag(tagDelete, "string", "name", "", "ID")
	
	// tagDeleteParameters := api.TagDeleteParameters{}
	

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
				Key:    Config.AccessToken,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TagShowOpts{}

			
			projectId := params.GetString("projectId")
			
			name := params.GetString("name")
			
			tagShowParameters := api.TagShowParameters{}
			

			data, api_response, err := client.TagsApi.TagShow(auth, projectId, name, tagShowParameters, &localVarOptionals)

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

	tagsApiCmd.AddCommand(tagShow)

	
	AddFlag(tagShow, "string", "projectId", "", "ID")
	
	AddFlag(tagShow, "string", "name", "", "ID")
	
	// tagShowParameters := api.TagShowParameters{}
	

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
				Key:    Config.AccessToken,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TagsListOpts{}

			
			projectId := params.GetString("projectId")
			
			tagsListParameters := api.TagsListParameters{}
			

			data, api_response, err := client.TagsApi.TagsList(auth, projectId, tagsListParameters, &localVarOptionals)

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

	tagsApiCmd.AddCommand(tagsList)

	
	AddFlag(tagsList, "string", "projectId", "", "ID")
	
	// tagsListParameters := api.TagsListParameters{}
	

	params.BindPFlags(tagsList.Flags())
}

