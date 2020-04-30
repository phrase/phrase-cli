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
	initAuthorizationCreate()
	initAuthorizationDelete()
	initAuthorizationShow()
	initAuthorizationUpdate()
	initAuthorizationsList()

	rootCmd.AddCommand(authorizationsApiCmd)
}

var authorizationsApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("authorizationsapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("AuthorizationsApi", "Api"), "API"}, " "),
}


func initAuthorizationCreate() {
	params := viper.New()
	var authorizationCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("AuthorizationCreate", strings.TrimSuffix("AuthorizationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("AuthorizationsApi", "Api"), "s"))),
		Short: "Create an authorization",
		Long:  `Create a new authorization.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.AuthorizationCreateOpts{}

			
			authorizationCreateParameters := api.AuthorizationCreateParameters{}
			

			data, api_response, err := client.AuthorizationsApi.AuthorizationCreate(auth, authorizationCreateParameters, &localVarOptionals)

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

	authorizationsApiCmd.AddCommand(authorizationCreate)

	
	// authorizationCreateParameters := api.AuthorizationCreateParameters{}
	

	params.BindPFlags(authorizationCreate.Flags())
}

func initAuthorizationDelete() {
	params := viper.New()
	var authorizationDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("AuthorizationDelete", strings.TrimSuffix("AuthorizationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("AuthorizationsApi", "Api"), "s"))),
		Short: "Delete an authorization",
		Long:  `Delete an existing authorization. API calls using that token will stop working.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.AuthorizationDeleteOpts{}

			
			id := params.GetString("id")
			

			data, api_response, err := client.AuthorizationsApi.AuthorizationDelete(auth, id, &localVarOptionals)

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

	authorizationsApiCmd.AddCommand(authorizationDelete)

	
	AddFlag(authorizationDelete, "string", "id", "", "ID")
	

	params.BindPFlags(authorizationDelete.Flags())
}

func initAuthorizationShow() {
	params := viper.New()
	var authorizationShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("AuthorizationShow", strings.TrimSuffix("AuthorizationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("AuthorizationsApi", "Api"), "s"))),
		Short: "Get a single authorization",
		Long:  `Get details on a single authorization.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.AuthorizationShowOpts{}

			
			id := params.GetString("id")
			

			data, api_response, err := client.AuthorizationsApi.AuthorizationShow(auth, id, &localVarOptionals)

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

	authorizationsApiCmd.AddCommand(authorizationShow)

	
	AddFlag(authorizationShow, "string", "id", "", "ID")
	

	params.BindPFlags(authorizationShow.Flags())
}

func initAuthorizationUpdate() {
	params := viper.New()
	var authorizationUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("AuthorizationUpdate", strings.TrimSuffix("AuthorizationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("AuthorizationsApi", "Api"), "s"))),
		Short: "Update an authorization",
		Long:  `Update an existing authorization.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.AuthorizationUpdateOpts{}

			
			id := params.GetString("id")
			
			authorizationUpdateParameters := api.AuthorizationUpdateParameters{}
			

			data, api_response, err := client.AuthorizationsApi.AuthorizationUpdate(auth, id, authorizationUpdateParameters, &localVarOptionals)

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

	authorizationsApiCmd.AddCommand(authorizationUpdate)

	
	AddFlag(authorizationUpdate, "string", "id", "", "ID")
	
	// authorizationUpdateParameters := api.AuthorizationUpdateParameters{}
	

	params.BindPFlags(authorizationUpdate.Flags())
}

func initAuthorizationsList() {
	params := viper.New()
	var authorizationsList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("AuthorizationsList", strings.TrimSuffix("AuthorizationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("AuthorizationsApi", "Api"), "s"))),
		Short: "List authorizations",
		Long:  `List all your authorizations.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.AuthorizationsListOpts{}

			

			data, api_response, err := client.AuthorizationsApi.AuthorizationsList(auth, &localVarOptionals)

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

	authorizationsApiCmd.AddCommand(authorizationsList)

	

	params.BindPFlags(authorizationsList.Flags())
}

