package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	helpers "github.com/phrase/phrase-cli/helpers"
	api "github.com/phrase/phrase-go/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initLanguagesList()

	rootCmd.AddCommand(SupportedLanguagesApiCmd)
}

var SupportedLanguagesApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("SupportedLanguages"),
	Short: "SupportedLanguages API",
}

func initLanguagesList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("languages/list", "/")[1:], "_")
	var LanguagesList = &cobra.Command{
		Use:   use,
		Short: "List supported languages",
		Long:  `Returns all languages/locale codes that Phrase Strings recognizes, mapped to their human-readable display name. This endpoint does not require authentication. `,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)

			data, api_response, err := client.SupportedLanguagesApi.LanguagesList(auth)

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

	SupportedLanguagesApiCmd.AddCommand(LanguagesList)

	params.BindPFlags(LanguagesList.Flags())
}
