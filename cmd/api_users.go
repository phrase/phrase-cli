package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/antihax/optional"
	helpers "github.com/phrase/phrase-cli/helpers"
	api "github.com/phrase/phrase-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initShowUser()

	rootCmd.AddCommand(UsersApiCmd)
}

var UsersApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("Users"),
	Short: "Users API",
}

func initShowUser() {
	params := viper.New()
	var ShowUser = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ShowUser", strings.TrimSuffix("UsersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("UsersApi", "Api"), "s"))),
		Short: "Show current User",
		Long:  `Show details for current User.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ShowUserOpts{}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.UsersApi.ShowUser(auth, &localVarOptionals)

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

	UsersApiCmd.AddCommand(ShowUser)

	AddFlag(ShowUser, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(ShowUser.Flags())
}
