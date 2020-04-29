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
	initGlossariesList()
	initGlossaryCreate()
	initGlossaryDelete()
	initGlossaryShow()
	initGlossaryUpdate()

	rootCmd.AddCommand(glossaryApiCmd)
}

var glossaryApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("glossaryapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("GlossaryApi", "Api"), "API"}, " "),
}


func initGlossariesList() {
	params := viper.New()
	var glossariesList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GlossariesList", strings.TrimSuffix("GlossaryApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GlossaryApi", "Api"), "s"))),
		Short: "List glossaries",
		Long:  `List all glossaries the current user has access to.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.GetString("access_token"),
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossariesListOpts{}

			
			accountId := params.GetString("accountId")
			

			data, api_response, err := client.GlossaryApi.GlossariesList(auth, accountId, &localVarOptionals)

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

	glossaryApiCmd.AddCommand(glossariesList)

	
	AddFlag(glossariesList, "string", "accountId", "", "ID")
	

	params.BindPFlags(glossariesList.Flags())
}

func initGlossaryCreate() {
	params := viper.New()
	var glossaryCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GlossaryCreate", strings.TrimSuffix("GlossaryApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GlossaryApi", "Api"), "s"))),
		Short: "Create a glossary",
		Long:  `Create a new glossary.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.GetString("access_token"),
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryCreateOpts{}

			
			accountId := params.GetString("accountId")
			
			glossaryCreateParameters := api.GlossaryCreateParameters{}
			

			api_response, err := client.GlossaryApi.GlossaryCreate(auth, accountId, glossaryCreateParameters, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	glossaryApiCmd.AddCommand(glossaryCreate)

	
	AddFlag(glossaryCreate, "string", "accountId", "", "ID")
	
	// glossaryCreateParameters := api.GlossaryCreateParameters{}
	

	params.BindPFlags(glossaryCreate.Flags())
}

func initGlossaryDelete() {
	params := viper.New()
	var glossaryDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GlossaryDelete", strings.TrimSuffix("GlossaryApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GlossaryApi", "Api"), "s"))),
		Short: "Delete a glossary",
		Long:  `Delete an existing glossary.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.GetString("access_token"),
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryDeleteOpts{}

			
			accountId := params.GetString("accountId")
			
			id := params.GetString("id")
			

			api_response, err := client.GlossaryApi.GlossaryDelete(auth, accountId, id, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	glossaryApiCmd.AddCommand(glossaryDelete)

	
	AddFlag(glossaryDelete, "string", "accountId", "", "ID")
	
	AddFlag(glossaryDelete, "string", "id", "", "ID")
	

	params.BindPFlags(glossaryDelete.Flags())
}

func initGlossaryShow() {
	params := viper.New()
	var glossaryShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GlossaryShow", strings.TrimSuffix("GlossaryApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GlossaryApi", "Api"), "s"))),
		Short: "Get a single glossary",
		Long:  `Get details on a single glossary.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.GetString("access_token"),
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryShowOpts{}

			
			accountId := params.GetString("accountId")
			
			id := params.GetString("id")
			

			data, api_response, err := client.GlossaryApi.GlossaryShow(auth, accountId, id, &localVarOptionals)

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

	glossaryApiCmd.AddCommand(glossaryShow)

	
	AddFlag(glossaryShow, "string", "accountId", "", "ID")
	
	AddFlag(glossaryShow, "string", "id", "", "ID")
	

	params.BindPFlags(glossaryShow.Flags())
}

func initGlossaryUpdate() {
	params := viper.New()
	var glossaryUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GlossaryUpdate", strings.TrimSuffix("GlossaryApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GlossaryApi", "Api"), "s"))),
		Short: "Update a glossary",
		Long:  `Update an existing glossary.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.GetString("access_token"),
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryUpdateOpts{}

			
			accountId := params.GetString("accountId")
			
			id := params.GetString("id")
			
			glossaryUpdateParameters := api.GlossaryUpdateParameters{}
			

			data, api_response, err := client.GlossaryApi.GlossaryUpdate(auth, accountId, id, glossaryUpdateParameters, &localVarOptionals)

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

	glossaryApiCmd.AddCommand(glossaryUpdate)

	
	AddFlag(glossaryUpdate, "string", "accountId", "", "ID")
	
	AddFlag(glossaryUpdate, "string", "id", "", "ID")
	
	// glossaryUpdateParameters := api.GlossaryUpdateParameters{}
	

	params.BindPFlags(glossaryUpdate.Flags())
}

