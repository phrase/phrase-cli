package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/antihax/optional"
	helpers "github.com/phrase/phrase-cli/helpers"
	api "github.com/phrase/phrase-go/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initRepoSyncEventList()
	initRepoSyncEventShow()

	rootCmd.AddCommand(RepoSyncEventsApiCmd)
}

var RepoSyncEventsApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("RepoSyncEvents"),
	Short: "RepoSyncEvents API",
}

func initRepoSyncEventList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("repo_sync_event/list", "/")[1:], "_")
	var RepoSyncEventList = &cobra.Command{
		Use:   use,
		Short: "Repository Syncs History",
		Long:  `Get the history of a single Repo Sync. The history includes all imports and exports performed by the Repo Sync.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.RepoSyncEventListOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			id := params.GetString(helpers.ToSnakeCase("Id"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.RepoSyncEventsApi.RepoSyncEventList(auth, accountId, id, &localVarOptionals)

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

	RepoSyncEventsApiCmd.AddCommand(RepoSyncEventList)
	AddFlag(RepoSyncEventList, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(RepoSyncEventList, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(RepoSyncEventList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(RepoSyncEventList.Flags())
}
func initRepoSyncEventShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("repo_sync_event/show", "/")[1:], "_")
	var RepoSyncEventShow = &cobra.Command{
		Use:   use,
		Short: "Get a single Repo Sync Event",
		Long:  `Shows a single Repo Sync event.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.RepoSyncEventShowOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			repoSyncId := params.GetString(helpers.ToSnakeCase("RepoSyncId"))

			id := params.GetString(helpers.ToSnakeCase("Id"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.RepoSyncEventsApi.RepoSyncEventShow(auth, accountId, repoSyncId, id, &localVarOptionals)

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

	RepoSyncEventsApiCmd.AddCommand(RepoSyncEventShow)
	AddFlag(RepoSyncEventShow, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(RepoSyncEventShow, "string", helpers.ToSnakeCase("RepoSyncId"), "", "Repo Sync ID", true)
	AddFlag(RepoSyncEventShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(RepoSyncEventShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(RepoSyncEventShow.Flags())
}
