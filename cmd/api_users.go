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
	initShowUser()

	rootCmd.AddCommand(usersApiCmd)
}

var usersApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("usersapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("UsersApi", "Api"), "API"}, " "),
}


func initShowUser() {
	params := viper.New()
	var showUser = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ShowUser", strings.TrimSuffix("UsersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("UsersApi", "Api"), "s"))),
		Short: "Show current User",
		Long:  `Show details for current User.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ShowUserOpts{}

			

			data, api_response, err := client.UsersApi.ShowUser(auth, &localVarOptionals)

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

	usersApiCmd.AddCommand(showUser)

	

	params.BindPFlags(showUser.Flags())
}

