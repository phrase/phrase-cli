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
	initGlossaryTermTranslationCreate()
	initGlossaryTermTranslationDelete()
	initGlossaryTermTranslationUpdate()

	rootCmd.AddCommand(glossaryTermTranslationsApiCmd)
}

var glossaryTermTranslationsApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("glossarytermtranslationsapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("GlossaryTermTranslationsApi", "Api"), "API"}, " "),
}


func initGlossaryTermTranslationCreate() {
	params := viper.New()
	var glossaryTermTranslationCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GlossaryTermTranslationCreate", strings.TrimSuffix("GlossaryTermTranslationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GlossaryTermTranslationsApi", "Api"), "s"))),
		Short: "Create a glossary term translation",
		Long:  `Create a new glossary term translation.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryTermTranslationCreateOpts{}

			
			accountId := params.GetString("accountId")

			
			glossaryId := params.GetString("glossaryId")

			
			termId := params.GetString("termId")

			

			glossaryTermTranslationCreateParameters := api.GlossaryTermTranslationCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &glossaryTermTranslationCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", glossaryTermTranslationCreateParameters)
			}
			

			data, api_response, err := client.GlossaryTermTranslationsApi.GlossaryTermTranslationCreate(auth, accountId, glossaryId, termId, glossaryTermTranslationCreateParameters, &localVarOptionals)

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

	glossaryTermTranslationsApiCmd.AddCommand(glossaryTermTranslationCreate)

	
	AddFlag(glossaryTermTranslationCreate, "string", "accountId", "", "Account ID", true)
	
	AddFlag(glossaryTermTranslationCreate, "string", "glossaryId", "", "Glossary ID", true)
	
	AddFlag(glossaryTermTranslationCreate, "string", "termId", "", "Term ID", true)
	
	AddFlag(glossaryTermTranslationCreate, "string", "data", "d", "payload in JSON format", true)
	// glossaryTermTranslationCreateParameters := api.GlossaryTermTranslationCreateParameters{}
	

	params.BindPFlags(glossaryTermTranslationCreate.Flags())
}

func initGlossaryTermTranslationDelete() {
	params := viper.New()
	var glossaryTermTranslationDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GlossaryTermTranslationDelete", strings.TrimSuffix("GlossaryTermTranslationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GlossaryTermTranslationsApi", "Api"), "s"))),
		Short: "Delete a glossary term translation",
		Long:  `Delete an existing glossary term translation.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryTermTranslationDeleteOpts{}

			
			accountId := params.GetString("accountId")

			
			glossaryId := params.GetString("glossaryId")

			
			termId := params.GetString("termId")

			
			id := params.GetString("id")

			

			data, api_response, err := client.GlossaryTermTranslationsApi.GlossaryTermTranslationDelete(auth, accountId, glossaryId, termId, id, &localVarOptionals)

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

	glossaryTermTranslationsApiCmd.AddCommand(glossaryTermTranslationDelete)

	
	AddFlag(glossaryTermTranslationDelete, "string", "accountId", "", "Account ID", true)
	
	AddFlag(glossaryTermTranslationDelete, "string", "glossaryId", "", "Glossary ID", true)
	
	AddFlag(glossaryTermTranslationDelete, "string", "termId", "", "Term ID", true)
	
	AddFlag(glossaryTermTranslationDelete, "string", "id", "", "ID", true)
	

	params.BindPFlags(glossaryTermTranslationDelete.Flags())
}

func initGlossaryTermTranslationUpdate() {
	params := viper.New()
	var glossaryTermTranslationUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GlossaryTermTranslationUpdate", strings.TrimSuffix("GlossaryTermTranslationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GlossaryTermTranslationsApi", "Api"), "s"))),
		Short: "Update a glossary term translation",
		Long:  `Update an existing glossary term translation.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GlossaryTermTranslationUpdateOpts{}

			
			accountId := params.GetString("accountId")

			
			glossaryId := params.GetString("glossaryId")

			
			termId := params.GetString("termId")

			
			id := params.GetString("id")

			

			glossaryTermTranslationUpdateParameters := api.GlossaryTermTranslationUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &glossaryTermTranslationUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", glossaryTermTranslationUpdateParameters)
			}
			

			data, api_response, err := client.GlossaryTermTranslationsApi.GlossaryTermTranslationUpdate(auth, accountId, glossaryId, termId, id, glossaryTermTranslationUpdateParameters, &localVarOptionals)

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

	glossaryTermTranslationsApiCmd.AddCommand(glossaryTermTranslationUpdate)

	
	AddFlag(glossaryTermTranslationUpdate, "string", "accountId", "", "Account ID", true)
	
	AddFlag(glossaryTermTranslationUpdate, "string", "glossaryId", "", "Glossary ID", true)
	
	AddFlag(glossaryTermTranslationUpdate, "string", "termId", "", "Term ID", true)
	
	AddFlag(glossaryTermTranslationUpdate, "string", "id", "", "ID", true)
	
	AddFlag(glossaryTermTranslationUpdate, "string", "data", "d", "payload in JSON format", true)
	// glossaryTermTranslationUpdateParameters := api.GlossaryTermTranslationUpdateParameters{}
	

	params.BindPFlags(glossaryTermTranslationUpdate.Flags())
}

