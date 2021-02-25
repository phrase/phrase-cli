package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/antihax/optional"
	helpers "github.com/phrase/phrase-cli/helpers"
	api "github.com/phrase/phrase-go/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initGitlabSyncDelete()
	initGitlabSyncExport()
	initGitlabSyncHistory()
	initGitlabSyncImport()
	initGitlabSyncList()
	initGitlabSyncShow()
	initGitlabSyncUpdate()

	rootCmd.AddCommand(GitLabSyncApiCmd)
}

var GitLabSyncApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("GitLabSync"),
	Short: "GitLabSync API",
}

func initGitlabSyncDelete() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("gitlab_sync/delete", "/")[1:], "_")
	var GitlabSyncDelete = &cobra.Command{
		Use:   use,
		Short: "Delete single Sync Setting",
		Long:  `Deletes a single GitLab Sync Setting.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.GitlabSyncDeleteOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("accountId")) {
				localVarOptionals.AccountId = optional.NewString(params.GetString(helpers.ToSnakeCase("AccountId")))
			}

			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.GitLabSyncApi.GitlabSyncDelete(auth, id, &localVarOptionals)

			if err != nil {
				switch castedError := err.(type) {
				case api.GenericOpenAPIError:
					fmt.Printf("\n%s\n\n", string(castedError.Body()))
					HandleError(castedError)

				default:
					HandleError(castedError)
				}
			} else if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
				os.Stdout.Write(data)

				if Config.Debug {
					fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
				}
			}
		},
	}

	GitLabSyncApiCmd.AddCommand(GitlabSyncDelete)
	AddFlag(GitlabSyncDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(GitlabSyncDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(GitlabSyncDelete, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID to specify the actual account the GitLab Sync should be created in. Required if the requesting user is a member of multiple accounts.", false)

	params.BindPFlags(GitlabSyncDelete.Flags())
}
func initGitlabSyncExport() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("gitlab_sync/export", "/")[1:], "_")
	var GitlabSyncExport = &cobra.Command{
		Use:   use,
		Short: "Export from Phrase to GitLab",
		Long:  `Export translations from Phrase to GitLab according to the .phraseapp.yml file within the GitLab repository.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.GitlabSyncExportOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			gitlabSyncId := params.GetString(helpers.ToSnakeCase("GitlabSyncId"))

			gitlabSyncExportParameters := api.GitlabSyncExportParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &gitlabSyncExportParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", gitlabSyncExportParameters)
			}
			data, api_response, err := client.GitLabSyncApi.GitlabSyncExport(auth, gitlabSyncId, gitlabSyncExportParameters, &localVarOptionals)

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

	GitLabSyncApiCmd.AddCommand(GitlabSyncExport)
	AddFlag(GitlabSyncExport, "string", helpers.ToSnakeCase("GitlabSyncId"), "", "Gitlab Sync ID", true)
	AddFlag(GitlabSyncExport, "string", "data", "d", "payload in JSON format", true)
	AddFlag(GitlabSyncExport, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(GitlabSyncExport.Flags())
}
func initGitlabSyncHistory() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("gitlab_sync/history", "/")[1:], "_")
	var GitlabSyncHistory = &cobra.Command{
		Use:   use,
		Short: "History of single Sync Setting",
		Long:  `List history for a single Sync Setting.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.GitlabSyncHistoryOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("page")) {
				localVarOptionals.Page = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("Page")))
			}
			if params.IsSet(helpers.ToSnakeCase("perPage")) {
				localVarOptionals.PerPage = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("PerPage")))
			}
			if params.IsSet(helpers.ToSnakeCase("accountId")) {
				localVarOptionals.AccountId = optional.NewString(params.GetString(helpers.ToSnakeCase("AccountId")))
			}

			gitlabSyncId := params.GetString(helpers.ToSnakeCase("GitlabSyncId"))

			data, api_response, err := client.GitLabSyncApi.GitlabSyncHistory(auth, gitlabSyncId, &localVarOptionals)

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

	GitLabSyncApiCmd.AddCommand(GitlabSyncHistory)
	AddFlag(GitlabSyncHistory, "string", helpers.ToSnakeCase("GitlabSyncId"), "", "Gitlab Sync ID", true)
	AddFlag(GitlabSyncHistory, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(GitlabSyncHistory, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(GitlabSyncHistory, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 25 by default", false)
	AddFlag(GitlabSyncHistory, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID to specify the actual account the GitLab Sync should be created in. Required if the requesting user is a member of multiple accounts.", false)

	params.BindPFlags(GitlabSyncHistory.Flags())
}
func initGitlabSyncImport() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("gitlab_sync/import", "/")[1:], "_")
	var GitlabSyncImport = &cobra.Command{
		Use:   use,
		Short: "Import from GitLab to Phrase",
		Long:  `Import translations from GitLab to Phrase according to the .phraseapp.yml file within the GitLab repository.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.GitlabSyncImportOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			gitlabSyncId := params.GetString(helpers.ToSnakeCase("GitlabSyncId"))

			gitlabSyncImportParameters := api.GitlabSyncImportParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &gitlabSyncImportParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", gitlabSyncImportParameters)
			}
			data, api_response, err := client.GitLabSyncApi.GitlabSyncImport(auth, gitlabSyncId, gitlabSyncImportParameters, &localVarOptionals)

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

	GitLabSyncApiCmd.AddCommand(GitlabSyncImport)
	AddFlag(GitlabSyncImport, "string", helpers.ToSnakeCase("GitlabSyncId"), "", "Gitlab Sync ID", true)
	AddFlag(GitlabSyncImport, "string", "data", "d", "payload in JSON format", true)
	AddFlag(GitlabSyncImport, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(GitlabSyncImport.Flags())
}
func initGitlabSyncList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("gitlab_sync/list", "/")[1:], "_")
	var GitlabSyncList = &cobra.Command{
		Use:   use,
		Short: "List GitLab syncs",
		Long:  `List all GitLab Sync Settings for which synchronisation with Phrase and GitLab is activated.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.GitlabSyncListOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("accountId")) {
				localVarOptionals.AccountId = optional.NewString(params.GetString(helpers.ToSnakeCase("AccountId")))
			}

			data, api_response, err := client.GitLabSyncApi.GitlabSyncList(auth, &localVarOptionals)

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

	GitLabSyncApiCmd.AddCommand(GitlabSyncList)
	AddFlag(GitlabSyncList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(GitlabSyncList, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID to specify the actual account the GitLab Sync should be created in. Required if the requesting user is a member of multiple accounts.", false)

	params.BindPFlags(GitlabSyncList.Flags())
}
func initGitlabSyncShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("gitlab_sync/show", "/")[1:], "_")
	var GitlabSyncShow = &cobra.Command{
		Use:   use,
		Short: "Get single Sync Setting",
		Long:  `Shows a single GitLab Sync Setting.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.GitlabSyncShowOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("accountId")) {
				localVarOptionals.AccountId = optional.NewString(params.GetString(helpers.ToSnakeCase("AccountId")))
			}

			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.GitLabSyncApi.GitlabSyncShow(auth, id, &localVarOptionals)

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

	GitLabSyncApiCmd.AddCommand(GitlabSyncShow)
	AddFlag(GitlabSyncShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(GitlabSyncShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(GitlabSyncShow, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID to specify the actual account the GitLab Sync should be created in. Required if the requesting user is a member of multiple accounts.", false)

	params.BindPFlags(GitlabSyncShow.Flags())
}
func initGitlabSyncUpdate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("gitlab_sync/update", "/")[1:], "_")
	var GitlabSyncUpdate = &cobra.Command{
		Use:   use,
		Short: "Update single Sync Setting",
		Long:  `Updates a single GitLab Sync Setting.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.GitlabSyncUpdateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("accountId")) {
				localVarOptionals.AccountId = optional.NewString(params.GetString(helpers.ToSnakeCase("AccountId")))
			}
			if params.IsSet(helpers.ToSnakeCase("phraseProjectCode")) {
				localVarOptionals.PhraseProjectCode = optional.NewString(params.GetString(helpers.ToSnakeCase("PhraseProjectCode")))
			}
			if params.IsSet(helpers.ToSnakeCase("gitlabProjectId")) {
				localVarOptionals.GitlabProjectId = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("GitlabProjectId")))
			}
			if params.IsSet(helpers.ToSnakeCase("gitlabBranchName")) {
				localVarOptionals.GitlabBranchName = optional.NewString(params.GetString(helpers.ToSnakeCase("GitlabBranchName")))
			}

			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.GitLabSyncApi.GitlabSyncUpdate(auth, id, &localVarOptionals)

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

	GitLabSyncApiCmd.AddCommand(GitlabSyncUpdate)
	AddFlag(GitlabSyncUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(GitlabSyncUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(GitlabSyncUpdate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID to specify the actual account the GitLab Sync should be created in. Required if the requesting user is a member of multiple accounts.", false)
	AddFlag(GitlabSyncUpdate, "string", helpers.ToSnakeCase("PhraseProjectCode"), "", "Code of the related Phrase Project.", false)
	AddFlag(GitlabSyncUpdate, "int32", helpers.ToSnakeCase("GitlabProjectId"), "", "ID of the related GitLab Project.", false)
	AddFlag(GitlabSyncUpdate, "string", helpers.ToSnakeCase("GitlabBranchName"), "", "Name of the GitLab Branch.", false)

	params.BindPFlags(GitlabSyncUpdate.Flags())
}
