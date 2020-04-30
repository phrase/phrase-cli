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
	initKeyCreate()
	initKeyDelete()
	initKeyShow()
	initKeyUpdate()
	initKeysDelete()
	initKeysList()
	initKeysSearch()
	initKeysTag()
	initKeysUntag()

	rootCmd.AddCommand(keysApiCmd)
}

var keysApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("keysapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("KeysApi", "Api"), "API"}, " "),
}


func initKeyCreate() {
	params := viper.New()
	var keyCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("KeyCreate", strings.TrimSuffix("KeysApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("KeysApi", "Api"), "s"))),
		Short: "Create a key",
		Long:  `Create a new key.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.KeyCreateOpts{}

			
			projectId := params.GetString("projectId")
			
			keyCreateParameters := api.KeyCreateParameters{}
			

			api_response, err := client.KeysApi.KeyCreate(auth, projectId, keyCreateParameters, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	keysApiCmd.AddCommand(keyCreate)

	
	AddFlag(keyCreate, "string", "projectId", "", "ID")
	
	// keyCreateParameters := api.KeyCreateParameters{}
	

	params.BindPFlags(keyCreate.Flags())
}

func initKeyDelete() {
	params := viper.New()
	var keyDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("KeyDelete", strings.TrimSuffix("KeysApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("KeysApi", "Api"), "s"))),
		Short: "Delete a key",
		Long:  `Delete an existing key.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.KeyDeleteOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			

			api_response, err := client.KeysApi.KeyDelete(auth, projectId, id, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	keysApiCmd.AddCommand(keyDelete)

	
	AddFlag(keyDelete, "string", "projectId", "", "ID")
	
	AddFlag(keyDelete, "string", "id", "", "ID")
	

	params.BindPFlags(keyDelete.Flags())
}

func initKeyShow() {
	params := viper.New()
	var keyShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("KeyShow", strings.TrimSuffix("KeysApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("KeysApi", "Api"), "s"))),
		Short: "Get a single key",
		Long:  `Get details on a single key for a given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.KeyShowOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			

			data, api_response, err := client.KeysApi.KeyShow(auth, projectId, id, &localVarOptionals)

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

	keysApiCmd.AddCommand(keyShow)

	
	AddFlag(keyShow, "string", "projectId", "", "ID")
	
	AddFlag(keyShow, "string", "id", "", "ID")
	

	params.BindPFlags(keyShow.Flags())
}

func initKeyUpdate() {
	params := viper.New()
	var keyUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("KeyUpdate", strings.TrimSuffix("KeysApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("KeysApi", "Api"), "s"))),
		Short: "Update a key",
		Long:  `Update an existing key.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.KeyUpdateOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			
			keyUpdateParameters := api.KeyUpdateParameters{}
			

			data, api_response, err := client.KeysApi.KeyUpdate(auth, projectId, id, keyUpdateParameters, &localVarOptionals)

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

	keysApiCmd.AddCommand(keyUpdate)

	
	AddFlag(keyUpdate, "string", "projectId", "", "ID")
	
	AddFlag(keyUpdate, "string", "id", "", "ID")
	
	// keyUpdateParameters := api.KeyUpdateParameters{}
	

	params.BindPFlags(keyUpdate.Flags())
}

func initKeysDelete() {
	params := viper.New()
	var keysDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("KeysDelete", strings.TrimSuffix("KeysApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("KeysApi", "Api"), "s"))),
		Short: "Delete collection of keys",
		Long:  `Delete all keys matching query. Same constraints as list. Please limit the number of affected keys to about 1,000 as you might experience timeouts otherwise.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.KeysDeleteOpts{}

			
			projectId := params.GetString("projectId")
			

			data, api_response, err := client.KeysApi.KeysDelete(auth, projectId, &localVarOptionals)

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

	keysApiCmd.AddCommand(keysDelete)

	
	AddFlag(keysDelete, "string", "projectId", "", "ID")
	

	params.BindPFlags(keysDelete.Flags())
}

func initKeysList() {
	params := viper.New()
	var keysList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("KeysList", strings.TrimSuffix("KeysApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("KeysApi", "Api"), "s"))),
		Short: "List keys",
		Long:  `List all keys for the given project. Alternatively you can POST requests to /search.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.KeysListOpts{}

			
			projectId := params.GetString("projectId")
			

			data, api_response, err := client.KeysApi.KeysList(auth, projectId, &localVarOptionals)

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

	keysApiCmd.AddCommand(keysList)

	
	AddFlag(keysList, "string", "projectId", "", "ID")
	

	params.BindPFlags(keysList.Flags())
}

func initKeysSearch() {
	params := viper.New()
	var keysSearch = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("KeysSearch", strings.TrimSuffix("KeysApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("KeysApi", "Api"), "s"))),
		Short: "Search keys",
		Long:  `Search keys for the given project matching query.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.KeysSearchOpts{}

			
			projectId := params.GetString("projectId")
			
			keysSearchParameters := api.KeysSearchParameters{}
			

			data, api_response, err := client.KeysApi.KeysSearch(auth, projectId, keysSearchParameters, &localVarOptionals)

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

	keysApiCmd.AddCommand(keysSearch)

	
	AddFlag(keysSearch, "string", "projectId", "", "ID")
	
	// keysSearchParameters := api.KeysSearchParameters{}
	

	params.BindPFlags(keysSearch.Flags())
}

func initKeysTag() {
	params := viper.New()
	var keysTag = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("KeysTag", strings.TrimSuffix("KeysApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("KeysApi", "Api"), "s"))),
		Short: "Add tags to collection of keys",
		Long:  `Tags all keys matching query. Same constraints as list.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.KeysTagOpts{}

			
			projectId := params.GetString("projectId")
			
			keysTagParameters := api.KeysTagParameters{}
			

			data, api_response, err := client.KeysApi.KeysTag(auth, projectId, keysTagParameters, &localVarOptionals)

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

	keysApiCmd.AddCommand(keysTag)

	
	AddFlag(keysTag, "string", "projectId", "", "ID")
	
	// keysTagParameters := api.KeysTagParameters{}
	

	params.BindPFlags(keysTag.Flags())
}

func initKeysUntag() {
	params := viper.New()
	var keysUntag = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("KeysUntag", strings.TrimSuffix("KeysApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("KeysApi", "Api"), "s"))),
		Short: "Remove tags from collection of keys",
		Long:  `Removes specified tags from keys matching query.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.KeysUntagOpts{}

			
			projectId := params.GetString("projectId")
			
			keysUntagParameters := api.KeysUntagParameters{}
			

			data, api_response, err := client.KeysApi.KeysUntag(auth, projectId, keysUntagParameters, &localVarOptionals)

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

	keysApiCmd.AddCommand(keysUntag)

	
	AddFlag(keysUntag, "string", "projectId", "", "ID")
	
	// keysUntagParameters := api.KeysUntagParameters{}
	

	params.BindPFlags(keysUntag.Flags())
}

