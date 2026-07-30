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
	initAccountAutomationEventsList()
	initAutomationEventsList()

	rootCmd.AddCommand(AutomationEventsApiCmd)
}

var AutomationEventsApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("AutomationEvents"),
	Short: "AutomationEvents API",
}

func initAccountAutomationEventsList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("account_automation_events/list", "/")[1:], "_")
	var AccountAutomationEventsList = &cobra.Command{
		Use:   use,
		Short: "List automation events for an account",
		Long:  `Returns the run history across all automations in the account, newest-first.  Use &#x60;automation_id&#x60; to narrow results to a single automation. Use &#x60;project_id&#x60; or &#x60;project_ids&#x60; to narrow by project.  For feature availability, see [Jobs (Strings)](https://support.phrase.com/hc/en-us/articles/5784100517788-Jobs-Strings). `,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.AccountAutomationEventsListOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("page")) {
				localVarOptionals.Page = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("Page")))
			}

			if params.IsSet(helpers.ToSnakeCase("perPage")) {
				localVarOptionals.PerPage = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("PerPage")))
			}

			if params.IsSet(helpers.ToSnakeCase("automationId")) {
				localVarOptionals.AutomationId = optional.NewString(params.GetString(helpers.ToSnakeCase("AutomationId")))
			}

			if params.IsSet(helpers.ToSnakeCase("state")) {
				localVarOptionals.State = optional.NewString(params.GetString(helpers.ToSnakeCase("State")))
			}

			if params.IsSet(helpers.ToSnakeCase("triggeredBy")) {
				localVarOptionals.TriggeredBy = optional.NewString(params.GetString(helpers.ToSnakeCase("TriggeredBy")))
			}

			if params.IsSet(helpers.ToSnakeCase("projectId")) {
				localVarOptionals.ProjectId = optional.NewString(params.GetString(helpers.ToSnakeCase("ProjectId")))
			}

			if params.IsSet(helpers.ToSnakeCase("projectIds")) {

				var projectIds []string

				if err := json.Unmarshal([]byte(params.GetString(helpers.ToSnakeCase("projectIds"))), &projectIds); err != nil {
					HandleError(err)
				}
				localVarOptionals.ProjectIds = projectIds
			}

			if params.IsSet(helpers.ToSnakeCase("createdAfter")) {
				localVarOptionals.CreatedAfter = optional.NewString(params.GetString(helpers.ToSnakeCase("CreatedAfter")))
			}

			if params.IsSet(helpers.ToSnakeCase("createdBefore")) {
				localVarOptionals.CreatedBefore = optional.NewString(params.GetString(helpers.ToSnakeCase("CreatedBefore")))
			}

			data, api_response, err := client.AutomationEventsApi.AccountAutomationEventsList(auth, accountId, &localVarOptionals)

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

	AutomationEventsApiCmd.AddCommand(AccountAutomationEventsList)
	AddFlag(AccountAutomationEventsList, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(AccountAutomationEventsList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(AccountAutomationEventsList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(AccountAutomationEventsList, "int32", helpers.ToSnakeCase("PerPage"), "", "Limit on the number of objects to be returned, between 1 and 100. 25 by default", false)
	AddFlag(AccountAutomationEventsList, "string", helpers.ToSnakeCase("AutomationId"), "", "Filter events to a single automation by its ID.", false)
	AddFlag(AccountAutomationEventsList, "string", helpers.ToSnakeCase("State"), "", "Filter events by outcome state. Unrecognized values are ignored.", false)
	AddFlag(AccountAutomationEventsList, "string", helpers.ToSnakeCase("TriggeredBy"), "", "Filter events by what triggered the automation run. Unrecognized values are ignored.", false)
	AddFlag(AccountAutomationEventsList, "string", helpers.ToSnakeCase("ProjectId"), "", "Filter events by project ID. Accepts a single ID or a comma-separated list of IDs.", false)
	AddFlag(AccountAutomationEventsList, "string", helpers.ToSnakeCase("ProjectIds"), "", "payload in JSON format", false)
	AddFlag(AccountAutomationEventsList, "string", helpers.ToSnakeCase("CreatedAfter"), "", "Return only events created after this ISO 8601 timestamp. Returns 400 if the value is not a valid date-time.", false)
	AddFlag(AccountAutomationEventsList, "string", helpers.ToSnakeCase("CreatedBefore"), "", "Return only events created before this ISO 8601 timestamp. Returns 400 if the value is not a valid date-time.", false)

	params.BindPFlags(AccountAutomationEventsList.Flags())
}
func initAutomationEventsList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("automation_events/list", "/")[1:], "_")
	var AutomationEventsList = &cobra.Command{
		Use:   use,
		Short: "List events for an automation",
		Long:  `Returns the run history for a specific automation, newest-first.  For feature availability, see [Jobs (Strings)](https://support.phrase.com/hc/en-us/articles/5784100517788-Jobs-Strings). `,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.AutomationEventsListOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			id := params.GetString(helpers.ToSnakeCase("Id"))

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

			if params.IsSet(helpers.ToSnakeCase("triggeredBy")) {
				localVarOptionals.TriggeredBy = optional.NewString(params.GetString(helpers.ToSnakeCase("TriggeredBy")))
			}

			if params.IsSet(helpers.ToSnakeCase("projectId")) {
				localVarOptionals.ProjectId = optional.NewString(params.GetString(helpers.ToSnakeCase("ProjectId")))
			}

			if params.IsSet(helpers.ToSnakeCase("projectIds")) {

				var projectIds []string

				if err := json.Unmarshal([]byte(params.GetString(helpers.ToSnakeCase("projectIds"))), &projectIds); err != nil {
					HandleError(err)
				}
				localVarOptionals.ProjectIds = projectIds
			}

			if params.IsSet(helpers.ToSnakeCase("createdAfter")) {
				localVarOptionals.CreatedAfter = optional.NewString(params.GetString(helpers.ToSnakeCase("CreatedAfter")))
			}

			if params.IsSet(helpers.ToSnakeCase("createdBefore")) {
				localVarOptionals.CreatedBefore = optional.NewString(params.GetString(helpers.ToSnakeCase("CreatedBefore")))
			}

			data, api_response, err := client.AutomationEventsApi.AutomationEventsList(auth, accountId, id, &localVarOptionals)

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

	AutomationEventsApiCmd.AddCommand(AutomationEventsList)
	AddFlag(AutomationEventsList, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(AutomationEventsList, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(AutomationEventsList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(AutomationEventsList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(AutomationEventsList, "int32", helpers.ToSnakeCase("PerPage"), "", "Limit on the number of objects to be returned, between 1 and 100. 25 by default", false)
	AddFlag(AutomationEventsList, "string", helpers.ToSnakeCase("State"), "", "Filter events by outcome state. Unrecognized values are ignored.", false)
	AddFlag(AutomationEventsList, "string", helpers.ToSnakeCase("TriggeredBy"), "", "Filter events by what triggered the automation run. Unrecognized values are ignored.", false)
	AddFlag(AutomationEventsList, "string", helpers.ToSnakeCase("ProjectId"), "", "Filter events by project ID. Accepts a single ID or a comma-separated list of IDs.", false)
	AddFlag(AutomationEventsList, "string", helpers.ToSnakeCase("ProjectIds"), "", "payload in JSON format", false)
	AddFlag(AutomationEventsList, "string", helpers.ToSnakeCase("CreatedAfter"), "", "Return only events created after this ISO 8601 timestamp. Returns 400 if the value is not a valid date-time.", false)
	AddFlag(AutomationEventsList, "string", helpers.ToSnakeCase("CreatedBefore"), "", "Return only events created before this ISO 8601 timestamp. Returns 400 if the value is not a valid date-time.", false)

	params.BindPFlags(AutomationEventsList.Flags())
}
