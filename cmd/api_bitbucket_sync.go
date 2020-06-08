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
	initBitbucketSyncExport()
	initBitbucketSyncImport()
	initBitbucketSyncsList()

	rootCmd.AddCommand(BitbucketSyncApiCmd)
}

var BitbucketSyncApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("BitbucketSync"),
	Short: "BitbucketSync API",
}

func initBitbucketSyncExport() {
	params := viper.New()
	var BitbucketSyncExport = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("BitbucketSyncExport", strings.TrimSuffix("BitbucketSyncApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("BitbucketSyncApi", "Api"), "s"))),
		Short: "Export from Phrase to Bitbucket",
		Long:  `Export translations from Phrase to Bitbucket according to the .phraseapp.yml file within the Bitbucket Repository.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.BitbucketSyncExportOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			id := params.GetString(helpers.ToSnakeCase("Id"))

			bitbucketSyncExportParameters := api.BitbucketSyncExportParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &bitbucketSyncExportParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", bitbucketSyncExportParameters)
			}
			data, api_response, err := client.BitbucketSyncApi.BitbucketSyncExport(auth, id, bitbucketSyncExportParameters, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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

	BitbucketSyncApiCmd.AddCommand(BitbucketSyncExport)

	AddFlag(BitbucketSyncExport, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(BitbucketSyncExport, "string", "data", "d", "payload in JSON format", true)

	AddFlag(BitbucketSyncExport, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(BitbucketSyncExport.Flags())
}
func initBitbucketSyncImport() {
	params := viper.New()
	var BitbucketSyncImport = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("BitbucketSyncImport", strings.TrimSuffix("BitbucketSyncApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("BitbucketSyncApi", "Api"), "s"))),
		Short: "Import to Phrase from Bitbucket",
		Long:  `Import translations from Bitbucket to Phrase according to the .phraseapp.yml file within the Bitbucket repository.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.BitbucketSyncImportOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			id := params.GetString(helpers.ToSnakeCase("Id"))

			bitbucketSyncImportParameters := api.BitbucketSyncImportParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &bitbucketSyncImportParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", bitbucketSyncImportParameters)
			}
			data, api_response, err := client.BitbucketSyncApi.BitbucketSyncImport(auth, id, bitbucketSyncImportParameters, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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

	BitbucketSyncApiCmd.AddCommand(BitbucketSyncImport)

	AddFlag(BitbucketSyncImport, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(BitbucketSyncImport, "string", "data", "d", "payload in JSON format", true)

	AddFlag(BitbucketSyncImport, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(BitbucketSyncImport.Flags())
}
func initBitbucketSyncsList() {
	params := viper.New()
	var BitbucketSyncsList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("BitbucketSyncsList", strings.TrimSuffix("BitbucketSyncApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("BitbucketSyncApi", "Api"), "s"))),
		Short: "List Bitbucket syncs",
		Long:  `List all Bitbucket repositories for which synchronisation with Phrase is activated.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.BitbucketSyncsListOpts{}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("accountId")) {
				localVarOptionals.AccountId = optional.NewString(params.GetString(helpers.ToSnakeCase("AccountId")))
			}

			data, api_response, err := client.BitbucketSyncApi.BitbucketSyncsList(auth, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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

	BitbucketSyncApiCmd.AddCommand(BitbucketSyncsList)

	AddFlag(BitbucketSyncsList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(BitbucketSyncsList, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID to specify the actual account the project should be created in. Required if the requesting user is a member of multiple accounts.", false)
	params.BindPFlags(BitbucketSyncsList.Flags())
}
