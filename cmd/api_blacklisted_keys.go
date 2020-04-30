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
	initBlacklistedKeyCreate()
	initBlacklistedKeyDelete()
	initBlacklistedKeyShow()
	initBlacklistedKeyUpdate()
	initBlacklistedKeysList()

	rootCmd.AddCommand(blacklistedKeysApiCmd)
}

var blacklistedKeysApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("blacklistedkeysapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("BlacklistedKeysApi", "Api"), "API"}, " "),
}


func initBlacklistedKeyCreate() {
	params := viper.New()
	var blacklistedKeyCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("BlacklistedKeyCreate", strings.TrimSuffix("BlacklistedKeysApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("BlacklistedKeysApi", "Api"), "s"))),
		Short: "Create a blacklisted key",
		Long:  `Create a new rule for blacklisting keys.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.BlacklistedKeyCreateOpts{}

			
			projectId := params.GetString("projectId")
			
			blacklistedKeyCreateParameters := api.BlacklistedKeyCreateParameters{}
			

			api_response, err := client.BlacklistedKeysApi.BlacklistedKeyCreate(auth, projectId, blacklistedKeyCreateParameters, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	blacklistedKeysApiCmd.AddCommand(blacklistedKeyCreate)

	
	AddFlag(blacklistedKeyCreate, "string", "projectId", "", "ID")
	
	// blacklistedKeyCreateParameters := api.BlacklistedKeyCreateParameters{}
	

	params.BindPFlags(blacklistedKeyCreate.Flags())
}

func initBlacklistedKeyDelete() {
	params := viper.New()
	var blacklistedKeyDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("BlacklistedKeyDelete", strings.TrimSuffix("BlacklistedKeysApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("BlacklistedKeysApi", "Api"), "s"))),
		Short: "Delete a blacklisted key",
		Long:  `Delete an existing rule for blacklisting keys.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.BlacklistedKeyDeleteOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			

			api_response, err := client.BlacklistedKeysApi.BlacklistedKeyDelete(auth, projectId, id, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	blacklistedKeysApiCmd.AddCommand(blacklistedKeyDelete)

	
	AddFlag(blacklistedKeyDelete, "string", "projectId", "", "ID")
	
	AddFlag(blacklistedKeyDelete, "string", "id", "", "ID")
	

	params.BindPFlags(blacklistedKeyDelete.Flags())
}

func initBlacklistedKeyShow() {
	params := viper.New()
	var blacklistedKeyShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("BlacklistedKeyShow", strings.TrimSuffix("BlacklistedKeysApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("BlacklistedKeysApi", "Api"), "s"))),
		Short: "Get a single blacklisted key",
		Long:  `Get details on a single rule for blacklisting keys for a given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.BlacklistedKeyShowOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			

			data, api_response, err := client.BlacklistedKeysApi.BlacklistedKeyShow(auth, projectId, id, &localVarOptionals)

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

	blacklistedKeysApiCmd.AddCommand(blacklistedKeyShow)

	
	AddFlag(blacklistedKeyShow, "string", "projectId", "", "ID")
	
	AddFlag(blacklistedKeyShow, "string", "id", "", "ID")
	

	params.BindPFlags(blacklistedKeyShow.Flags())
}

func initBlacklistedKeyUpdate() {
	params := viper.New()
	var blacklistedKeyUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("BlacklistedKeyUpdate", strings.TrimSuffix("BlacklistedKeysApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("BlacklistedKeysApi", "Api"), "s"))),
		Short: "Update a blacklisted key",
		Long:  `Update an existing rule for blacklisting keys.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.BlacklistedKeyUpdateOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			
			blacklistedKeyUpdateParameters := api.BlacklistedKeyUpdateParameters{}
			

			data, api_response, err := client.BlacklistedKeysApi.BlacklistedKeyUpdate(auth, projectId, id, blacklistedKeyUpdateParameters, &localVarOptionals)

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

	blacklistedKeysApiCmd.AddCommand(blacklistedKeyUpdate)

	
	AddFlag(blacklistedKeyUpdate, "string", "projectId", "", "ID")
	
	AddFlag(blacklistedKeyUpdate, "string", "id", "", "ID")
	
	// blacklistedKeyUpdateParameters := api.BlacklistedKeyUpdateParameters{}
	

	params.BindPFlags(blacklistedKeyUpdate.Flags())
}

func initBlacklistedKeysList() {
	params := viper.New()
	var blacklistedKeysList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("BlacklistedKeysList", strings.TrimSuffix("BlacklistedKeysApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("BlacklistedKeysApi", "Api"), "s"))),
		Short: "List blacklisted keys",
		Long:  `List all rules for blacklisting keys for the given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.BlacklistedKeysListOpts{}

			
			projectId := params.GetString("projectId")
			

			data, api_response, err := client.BlacklistedKeysApi.BlacklistedKeysList(auth, projectId, &localVarOptionals)

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

	blacklistedKeysApiCmd.AddCommand(blacklistedKeysList)

	
	AddFlag(blacklistedKeysList, "string", "projectId", "", "ID")
	

	params.BindPFlags(blacklistedKeysList.Flags())
}

