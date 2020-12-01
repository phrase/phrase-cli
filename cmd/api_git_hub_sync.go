package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/antihax/optional"
	helpers "github.com/phrase/phrase-cli/helpers"
	api "github.com/phrase/phrase-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initGithubSyncExport()
	initGithubSyncImport()

	rootCmd.AddCommand(GitHubSyncApiCmd)
}

var GitHubSyncApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("GitHubSync"),
	Short: "GitHubSync API",
}

func initGithubSyncExport() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("github_sync/export", "/")[1:], "_")
	var GithubSyncExport = &cobra.Command{
		Use:   use,
		Short: "Export from Phrase to GitHub",
		Long:  `Export translations from Phrase to GitHub according to the .phraseapp.yml file within the GitHub repository.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.GithubSyncExportOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			githubSyncExportParameters := api.GithubSyncExportParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &githubSyncExportParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", githubSyncExportParameters)
			}
			data, api_response, err := client.GitHubSyncApi.GithubSyncExport(auth, githubSyncExportParameters, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
				os.Stdout.Write(data)
			}
			if err != nil {
				switch castedError := err.(type) {
				case api.GenericOpenAPIError:
					fmt.Printf("\n%s\n\n", string(castedError.Body()))
					HandleError(castedError)

				default:
					HandleError(castedError)
				}
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	GitHubSyncApiCmd.AddCommand(GithubSyncExport)
	AddFlag(GithubSyncExport, "string", "data", "d", "payload in JSON format", true)
	AddFlag(GithubSyncExport, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(GithubSyncExport.Flags())
}
func initGithubSyncImport() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("github_sync/import", "/")[1:], "_")
	var GithubSyncImport = &cobra.Command{
		Use:   use,
		Short: "Import to Phrase from GitHub",
		Long:  `Import files to Phrase from your connected GitHub repository.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.GithubSyncImportOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			githubSyncImportParameters := api.GithubSyncImportParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &githubSyncImportParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", githubSyncImportParameters)
			}
			data, api_response, err := client.GitHubSyncApi.GithubSyncImport(auth, githubSyncImportParameters, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
				os.Stdout.Write(data)
			}
			if err != nil {
				switch castedError := err.(type) {
				case api.GenericOpenAPIError:
					fmt.Printf("\n%s\n\n", string(castedError.Body()))
					HandleError(castedError)

				default:
					HandleError(castedError)
				}
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	GitHubSyncApiCmd.AddCommand(GithubSyncImport)
	AddFlag(GithubSyncImport, "string", "data", "d", "payload in JSON format", true)
	AddFlag(GithubSyncImport, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(GithubSyncImport.Flags())
}
