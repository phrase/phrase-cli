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
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossariesListOpts{}

			
			accountId := params.GetString("accountId")

			

			data, api_response, err := client.GlossaryApi.GlossariesList(auth, accountId, &localVarOptionals)

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

	glossaryApiCmd.AddCommand(glossariesList)

	
	AddFlag(glossariesList, "string", "accountId", "", "Account ID", true)
	

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
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryCreateOpts{}

			
			accountId := params.GetString("accountId")

			

			glossaryCreateParameters := api.GlossaryCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &glossaryCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", glossaryCreateParameters)
			}
			

			data, api_response, err := client.GlossaryApi.GlossaryCreate(auth, accountId, glossaryCreateParameters, &localVarOptionals)

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

	glossaryApiCmd.AddCommand(glossaryCreate)

	
	AddFlag(glossaryCreate, "string", "accountId", "", "Account ID", true)
	
	AddFlag(glossaryCreate, "string", "data", "d", "payload in JSON format", true)
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
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryDeleteOpts{}

			
			accountId := params.GetString("accountId")

			
			id := params.GetString("id")

			

			data, api_response, err := client.GlossaryApi.GlossaryDelete(auth, accountId, id, &localVarOptionals)

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

	glossaryApiCmd.AddCommand(glossaryDelete)

	
	AddFlag(glossaryDelete, "string", "accountId", "", "Account ID", true)
	
	AddFlag(glossaryDelete, "string", "id", "", "ID", true)
	

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
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryShowOpts{}

			
			accountId := params.GetString("accountId")

			
			id := params.GetString("id")

			

			data, api_response, err := client.GlossaryApi.GlossaryShow(auth, accountId, id, &localVarOptionals)

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

	glossaryApiCmd.AddCommand(glossaryShow)

	
	AddFlag(glossaryShow, "string", "accountId", "", "Account ID", true)
	
	AddFlag(glossaryShow, "string", "id", "", "ID", true)
	

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
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryUpdateOpts{}

			
			accountId := params.GetString("accountId")

			
			id := params.GetString("id")

			

			glossaryUpdateParameters := api.GlossaryUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &glossaryUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", glossaryUpdateParameters)
			}
			

			data, api_response, err := client.GlossaryApi.GlossaryUpdate(auth, accountId, id, glossaryUpdateParameters, &localVarOptionals)

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

	glossaryApiCmd.AddCommand(glossaryUpdate)

	
	AddFlag(glossaryUpdate, "string", "accountId", "", "Account ID", true)
	
	AddFlag(glossaryUpdate, "string", "id", "", "ID", true)
	
	AddFlag(glossaryUpdate, "string", "data", "d", "payload in JSON format", true)
	// glossaryUpdateParameters := api.GlossaryUpdateParameters{}
	

	params.BindPFlags(glossaryUpdate.Flags())
}

