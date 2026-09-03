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
	initCheckIssueDismiss()
	initCheckIssuesList()

	rootCmd.AddCommand(ChecksApiCmd)
}

var ChecksApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("Checks"),
	Short: "Checks API",
}

func initCheckIssueDismiss() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("check_issue/dismiss", "/")[1:], "_")
	var CheckIssueDismiss = &cobra.Command{
		Use:   use,
		Short: "Dismiss a check issue",
		Long:  `**Note:** The Checks API is still in development and might change in subsequent releases.  Mark a check issue as dismissed so it no longer appears on the list of active check issues.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.CheckIssueDismissOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			id := params.GetString(helpers.ToSnakeCase("Id"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.ChecksApi.CheckIssueDismiss(auth, projectId, id, &localVarOptionals)

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

	ChecksApiCmd.AddCommand(CheckIssueDismiss)
	AddFlag(CheckIssueDismiss, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(CheckIssueDismiss, "string", helpers.ToSnakeCase("Id"), "", "Check Issue ID", true)
	AddFlag(CheckIssueDismiss, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(CheckIssueDismiss.Flags())
}
func initCheckIssuesList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("check_issues/list", "/")[1:], "_")
	var CheckIssuesList = &cobra.Command{
		Use:   use,
		Short: "List check issues",
		Long:  `**Note:** The Checks API is still in development and might change in subsequent releases.  List check issues for the given project. Results can be filtered by locale, check name, and state.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.CheckIssuesListOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("page")) {
				localVarOptionals.Page = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("Page")))
			}

			if params.IsSet(helpers.ToSnakeCase("perPage")) {
				localVarOptionals.PerPage = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("PerPage")))
			}

			if params.IsSet(helpers.ToSnakeCase("state")) {
				localVarOptionals.State = optional.NewString(params.GetString(helpers.ToSnakeCase("State")))
			}

			if params.IsSet(helpers.ToSnakeCase("localeIds")) {

				var localeIds []string

				if err := json.Unmarshal([]byte(params.GetString(helpers.ToSnakeCase("localeIds"))), &localeIds); err != nil {
					HandleError(err)
				}
				localVarOptionals.LocaleIds = localeIds
			}

			if params.IsSet(helpers.ToSnakeCase("checkNames")) {

				var checkNames []string

				if err := json.Unmarshal([]byte(params.GetString(helpers.ToSnakeCase("checkNames"))), &checkNames); err != nil {
					HandleError(err)
				}
				localVarOptionals.CheckNames = checkNames
			}

			data, api_response, err := client.ChecksApi.CheckIssuesList(auth, projectId, &localVarOptionals)

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

	ChecksApiCmd.AddCommand(CheckIssuesList)
	AddFlag(CheckIssuesList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(CheckIssuesList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(CheckIssuesList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(CheckIssuesList, "int32", helpers.ToSnakeCase("PerPage"), "", "Limit on the number of objects to be returned, between 1 and 100. 25 by default", false)
	AddFlag(CheckIssuesList, "string", helpers.ToSnakeCase("State"), "", "Filter by state of the check issue. Can be one of: `active`, `solved`, `dismissed`, `all`. Defaults to `active`.", false)
	AddFlag(CheckIssuesList, "string", helpers.ToSnakeCase("LocaleIds"), "", "payload in JSON format", false)
	AddFlag(CheckIssuesList, "string", helpers.ToSnakeCase("CheckNames"), "", "payload in JSON format", false)

	params.BindPFlags(CheckIssuesList.Flags())
}
