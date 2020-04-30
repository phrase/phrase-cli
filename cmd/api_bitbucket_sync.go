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
	initBitbucketSyncExport()
	initBitbucketSyncImport()
	initBitbucketSyncsList()

	rootCmd.AddCommand(bitbucketSyncApiCmd)
}

var bitbucketSyncApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("bitbucketsyncapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("BitbucketSyncApi", "Api"), "API"}, " "),
}


func initBitbucketSyncExport() {
	params := viper.New()
	var bitbucketSyncExport = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("BitbucketSyncExport", strings.TrimSuffix("BitbucketSyncApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("BitbucketSyncApi", "Api"), "s"))),
		Short: "Export from Phrase to Bitbucket",
		Long:  `Export translations from Phrase to Bitbucket according to the .phraseapp.yml file within the Bitbucket Repository.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.BitbucketSyncExportOpts{}

			
			id := params.GetString("id")
			
			bitbucketSyncExportParameters := api.BitbucketSyncExportParameters{}
			

			data, api_response, err := client.BitbucketSyncApi.BitbucketSyncExport(auth, id, bitbucketSyncExportParameters, &localVarOptionals)

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

	bitbucketSyncApiCmd.AddCommand(bitbucketSyncExport)

	
	AddFlag(bitbucketSyncExport, "string", "id", "", "ID")
	
	// bitbucketSyncExportParameters := api.BitbucketSyncExportParameters{}
	

	params.BindPFlags(bitbucketSyncExport.Flags())
}

func initBitbucketSyncImport() {
	params := viper.New()
	var bitbucketSyncImport = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("BitbucketSyncImport", strings.TrimSuffix("BitbucketSyncApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("BitbucketSyncApi", "Api"), "s"))),
		Short: "Import to Phrase from Bitbucket",
		Long:  `Import translations from Bitbucket to Phrase according to the .phraseapp.yml file within the Bitbucket repository.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.BitbucketSyncImportOpts{}

			
			id := params.GetString("id")
			
			bitbucketSyncImportParameters := api.BitbucketSyncImportParameters{}
			

			api_response, err := client.BitbucketSyncApi.BitbucketSyncImport(auth, id, bitbucketSyncImportParameters, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	bitbucketSyncApiCmd.AddCommand(bitbucketSyncImport)

	
	AddFlag(bitbucketSyncImport, "string", "id", "", "ID")
	
	// bitbucketSyncImportParameters := api.BitbucketSyncImportParameters{}
	

	params.BindPFlags(bitbucketSyncImport.Flags())
}

func initBitbucketSyncsList() {
	params := viper.New()
	var bitbucketSyncsList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("BitbucketSyncsList", strings.TrimSuffix("BitbucketSyncApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("BitbucketSyncApi", "Api"), "s"))),
		Short: "List Bitbucket syncs",
		Long:  `List all Bitbucket repositories for which synchronisation with Phrase is activated.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.BitbucketSyncsListOpts{}

			
			bitbucketSyncsListParameters := api.BitbucketSyncsListParameters{}
			

			data, api_response, err := client.BitbucketSyncApi.BitbucketSyncsList(auth, bitbucketSyncsListParameters, &localVarOptionals)

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

	bitbucketSyncApiCmd.AddCommand(bitbucketSyncsList)

	
	// bitbucketSyncsListParameters := api.BitbucketSyncsListParameters{}
	

	params.BindPFlags(bitbucketSyncsList.Flags())
}

