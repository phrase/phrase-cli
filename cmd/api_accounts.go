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
	initAccountShow()
	initAccountsList()

	rootCmd.AddCommand(accountsApiCmd)
}

var accountsApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("accountsapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("AccountsApi", "Api"), "API"}, " "),
}


func initAccountShow() {
	params := viper.New()
	var accountShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("AccountShow", strings.TrimSuffix("AccountsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("AccountsApi", "Api"), "s"))),
		Short: "Get a single account",
		Long:  `Get details on a single account.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    viper.GetString("access_token"),
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.AccountShowOpts{}

			
			id := params.GetString("id")
			

			data, api_response, err := client.AccountsApi.AccountShow(auth, id, &localVarOptionals)

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

	accountsApiCmd.AddCommand(accountShow)

	
	AddFlag(accountShow, "string", "id", "", "ID")
	

	params.BindPFlags(accountShow.Flags())
}

func initAccountsList() {
	params := viper.New()
	var accountsList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("AccountsList", strings.TrimSuffix("AccountsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("AccountsApi", "Api"), "s"))),
		Short: "List accounts",
		Long:  `List all accounts the current user has access to.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    viper.GetString("access_token"),
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.AccountsListOpts{}

			

			data, api_response, err := client.AccountsApi.AccountsList(auth, &localVarOptionals)

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

	accountsApiCmd.AddCommand(accountsList)

	

	params.BindPFlags(accountsList.Flags())
}

