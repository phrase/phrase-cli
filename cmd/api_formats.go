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
	initFormatsList()

	rootCmd.AddCommand(formatsApiCmd)
}

var formatsApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("formatsapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("FormatsApi", "Api"), "API"}, " "),
}


func initFormatsList() {
	params := viper.New()
	var formatsList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("FormatsList", strings.TrimSuffix("FormatsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("FormatsApi", "Api"), "s"))),
		Short: "List formats",
		Long:  `Get a handy list of all localization file formats supported in Phrase.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.AccessToken,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.FormatsListOpts{}

			

			data, api_response, err := client.FormatsApi.FormatsList(auth, &localVarOptionals)

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

	formatsApiCmd.AddCommand(formatsList)

	

	params.BindPFlags(formatsList.Flags())
}

