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
	initGlossaryTermCreate()
	initGlossaryTermDelete()
	initGlossaryTermShow()
	initGlossaryTermUpdate()
	initGlossaryTermsList()

	rootCmd.AddCommand(glossaryTermsApiCmd)
}

var glossaryTermsApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("glossarytermsapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("GlossaryTermsApi", "Api"), "API"}, " "),
}


func initGlossaryTermCreate() {
	params := viper.New()
	var glossaryTermCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GlossaryTermCreate", strings.TrimSuffix("GlossaryTermsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GlossaryTermsApi", "Api"), "s"))),
		Short: "Create a glossary term",
		Long:  `Create a new glossary term.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryTermCreateOpts{}

			
			accountId := params.GetString("accountId")

			
			glossaryId := params.GetString("glossaryId")

			

			glossaryTermCreateParameters := api.GlossaryTermCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &glossaryTermCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", glossaryTermCreateParameters)
			}
			

			data, api_response, err := client.GlossaryTermsApi.GlossaryTermCreate(auth, accountId, glossaryId, glossaryTermCreateParameters, &localVarOptionals)

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

	glossaryTermsApiCmd.AddCommand(glossaryTermCreate)

	
	AddFlag(glossaryTermCreate, "string", "accountId", "", "Account ID", true)
	
	AddFlag(glossaryTermCreate, "string", "glossaryId", "", "Glossary ID", true)
	
	AddFlag(glossaryTermCreate, "string", "data", "d", "payload in JSON format", true)
	// glossaryTermCreateParameters := api.GlossaryTermCreateParameters{}
	

	params.BindPFlags(glossaryTermCreate.Flags())
}

func initGlossaryTermDelete() {
	params := viper.New()
	var glossaryTermDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GlossaryTermDelete", strings.TrimSuffix("GlossaryTermsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GlossaryTermsApi", "Api"), "s"))),
		Short: "Delete a glossary term",
		Long:  `Delete an existing glossary term.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryTermDeleteOpts{}

			
			accountId := params.GetString("accountId")

			
			glossaryId := params.GetString("glossaryId")

			
			id := params.GetString("id")

			

			data, api_response, err := client.GlossaryTermsApi.GlossaryTermDelete(auth, accountId, glossaryId, id, &localVarOptionals)

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

	glossaryTermsApiCmd.AddCommand(glossaryTermDelete)

	
	AddFlag(glossaryTermDelete, "string", "accountId", "", "Account ID", true)
	
	AddFlag(glossaryTermDelete, "string", "glossaryId", "", "Glossary ID", true)
	
	AddFlag(glossaryTermDelete, "string", "id", "", "ID", true)
	

	params.BindPFlags(glossaryTermDelete.Flags())
}

func initGlossaryTermShow() {
	params := viper.New()
	var glossaryTermShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GlossaryTermShow", strings.TrimSuffix("GlossaryTermsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GlossaryTermsApi", "Api"), "s"))),
		Short: "Get a single glossary term",
		Long:  `Get details on a single glossary term.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryTermShowOpts{}

			
			accountId := params.GetString("accountId")

			
			glossaryId := params.GetString("glossaryId")

			
			id := params.GetString("id")

			

			data, api_response, err := client.GlossaryTermsApi.GlossaryTermShow(auth, accountId, glossaryId, id, &localVarOptionals)

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

	glossaryTermsApiCmd.AddCommand(glossaryTermShow)

	
	AddFlag(glossaryTermShow, "string", "accountId", "", "Account ID", true)
	
	AddFlag(glossaryTermShow, "string", "glossaryId", "", "Glossary ID", true)
	
	AddFlag(glossaryTermShow, "string", "id", "", "ID", true)
	

	params.BindPFlags(glossaryTermShow.Flags())
}

func initGlossaryTermUpdate() {
	params := viper.New()
	var glossaryTermUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GlossaryTermUpdate", strings.TrimSuffix("GlossaryTermsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GlossaryTermsApi", "Api"), "s"))),
		Short: "Update a glossary term",
		Long:  `Update an existing glossary term.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryTermUpdateOpts{}

			
			accountId := params.GetString("accountId")

			
			glossaryId := params.GetString("glossaryId")

			
			id := params.GetString("id")

			

			glossaryTermUpdateParameters := api.GlossaryTermUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &glossaryTermUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", glossaryTermUpdateParameters)
			}
			

			data, api_response, err := client.GlossaryTermsApi.GlossaryTermUpdate(auth, accountId, glossaryId, id, glossaryTermUpdateParameters, &localVarOptionals)

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

	glossaryTermsApiCmd.AddCommand(glossaryTermUpdate)

	
	AddFlag(glossaryTermUpdate, "string", "accountId", "", "Account ID", true)
	
	AddFlag(glossaryTermUpdate, "string", "glossaryId", "", "Glossary ID", true)
	
	AddFlag(glossaryTermUpdate, "string", "id", "", "ID", true)
	
	AddFlag(glossaryTermUpdate, "string", "data", "d", "payload in JSON format", true)
	// glossaryTermUpdateParameters := api.GlossaryTermUpdateParameters{}
	

	params.BindPFlags(glossaryTermUpdate.Flags())
}

func initGlossaryTermsList() {
	params := viper.New()
	var glossaryTermsList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GlossaryTermsList", strings.TrimSuffix("GlossaryTermsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GlossaryTermsApi", "Api"), "s"))),
		Short: "List glossary terms",
		Long:  `List all glossary terms the current user has access to.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryTermsListOpts{}

			
			accountId := params.GetString("accountId")

			
			glossaryId := params.GetString("glossaryId")

			

			data, api_response, err := client.GlossaryTermsApi.GlossaryTermsList(auth, accountId, glossaryId, &localVarOptionals)

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

	glossaryTermsApiCmd.AddCommand(glossaryTermsList)

	
	AddFlag(glossaryTermsList, "string", "accountId", "", "Account ID", true)
	
	AddFlag(glossaryTermsList, "string", "glossaryId", "", "Glossary ID", true)
	

	params.BindPFlags(glossaryTermsList.Flags())
}

