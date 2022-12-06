package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/antihax/optional"
	helpers "github.com/phrase/phrase-cli/helpers"
	api "github.com/phrase/phrase-go/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initSearchInAccount()

	rootCmd.AddCommand(SearchApiCmd)
}

var SearchApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("Search"),
	Short: "Search API",
}

func initSearchInAccount() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("search/in_account", "/")[1:], "_")
	var SearchInAccount = &cobra.Command{
		Use:   use,
		Short: "Search across projects",
		Long:  `Search for keys and translations in all account projects &lt;br&gt;&lt;br&gt;&lt;i&gt;Note: Search is limited to 10000 results and may not include recently updated data depending on the project sizes.&lt;/i&gt;`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.SearchInAccountOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			searchInAccountParameters := api.SearchInAccountParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &searchInAccountParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", searchInAccountParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.SearchApi.SearchInAccount(auth, accountId, searchInAccountParameters, &localVarOptionals)

			if err != nil {
				switch castedError := err.(type) {
				case api.GenericOpenAPIError:
					fmt.Printf("\n%s\n\n", string(castedError.Body()))
					HandleError(castedError)

				default:
					HandleError(castedError)
				}
			} else if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}
				fmt.Printf("%s\n", string(jsonBuf))

				if Config.Debug {
					fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
				}
			}
		},
	}

	SearchApiCmd.AddCommand(SearchInAccount)
	AddFlag(SearchInAccount, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(SearchInAccount, "string", "data", "d", "payload in JSON format", true)
	AddFlag(SearchInAccount, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(SearchInAccount.Flags())
}
