package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/antihax/optional"
	helpers "github.com/phrase/phrase-cli/helpers"
	api "github.com/phrase/phrase-go/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initOrganizationJobTemplateLocaleDelete()
	initOrganizationJobTemplateLocaleShow()
	initOrganizationJobTemplateLocaleUpdate()
	initOrganizationJobTemplateLocalesCreate()
	initOrganizationJobTemplateLocalesList()

	rootCmd.AddCommand(OrganizationJobTemplateLocalesApiCmd)
}

var OrganizationJobTemplateLocalesApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("OrganizationJobTemplateLocales"),
	Short: "OrganizationJobTemplateLocales API",
}

func initOrganizationJobTemplateLocaleDelete() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("organization_job_template_locale/delete", "/")[1:], "_")
	var OrganizationJobTemplateLocaleDelete = &cobra.Command{
		Use:   use,
		Short: "Delete an organization job template locale",
		Long:  `Delete an existing organization job template locale.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.OrganizationJobTemplateLocaleDeleteOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			jobTemplateId := params.GetString(helpers.ToSnakeCase("JobTemplateId"))

			jobTemplateLocaleId := params.GetString(helpers.ToSnakeCase("JobTemplateLocaleId"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.OrganizationJobTemplateLocalesApi.OrganizationJobTemplateLocaleDelete(auth, accountId, jobTemplateId, jobTemplateLocaleId, &localVarOptionals)

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

	OrganizationJobTemplateLocalesApiCmd.AddCommand(OrganizationJobTemplateLocaleDelete)
	AddFlag(OrganizationJobTemplateLocaleDelete, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(OrganizationJobTemplateLocaleDelete, "string", helpers.ToSnakeCase("JobTemplateId"), "", "Job Template ID", true)
	AddFlag(OrganizationJobTemplateLocaleDelete, "string", helpers.ToSnakeCase("JobTemplateLocaleId"), "", "Job Template Locale ID", true)
	AddFlag(OrganizationJobTemplateLocaleDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(OrganizationJobTemplateLocaleDelete.Flags())
}
func initOrganizationJobTemplateLocaleShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("organization_job_template_locale/show", "/")[1:], "_")
	var OrganizationJobTemplateLocaleShow = &cobra.Command{
		Use:   use,
		Short: "Get a single organization job template locale",
		Long:  `Get a single job template locale for a given organization job template.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.OrganizationJobTemplateLocaleShowOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			jobTemplateId := params.GetString(helpers.ToSnakeCase("JobTemplateId"))

			jobTemplateLocaleId := params.GetString(helpers.ToSnakeCase("JobTemplateLocaleId"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.OrganizationJobTemplateLocalesApi.OrganizationJobTemplateLocaleShow(auth, accountId, jobTemplateId, jobTemplateLocaleId, &localVarOptionals)

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

	OrganizationJobTemplateLocalesApiCmd.AddCommand(OrganizationJobTemplateLocaleShow)
	AddFlag(OrganizationJobTemplateLocaleShow, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(OrganizationJobTemplateLocaleShow, "string", helpers.ToSnakeCase("JobTemplateId"), "", "Job Template ID", true)
	AddFlag(OrganizationJobTemplateLocaleShow, "string", helpers.ToSnakeCase("JobTemplateLocaleId"), "", "Job Template Locale ID", true)
	AddFlag(OrganizationJobTemplateLocaleShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(OrganizationJobTemplateLocaleShow.Flags())
}
func initOrganizationJobTemplateLocaleUpdate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("organization_job_template_locale/update", "/")[1:], "_")
	var OrganizationJobTemplateLocaleUpdate = &cobra.Command{
		Use:   use,
		Short: "Update an organization job template locale",
		Long:  `Update an existing organization job template locale.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.OrganizationJobTemplateLocaleUpdateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			jobTemplateId := params.GetString(helpers.ToSnakeCase("JobTemplateId"))

			jobTemplateLocaleId := params.GetString(helpers.ToSnakeCase("JobTemplateLocaleId"))

			var organizationJobTemplateLocaleUpdateParameters api.OrganizationJobTemplateLocaleUpdateParameters
			if err := json.Unmarshal([]byte(params.GetString("data")), &organizationJobTemplateLocaleUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", organizationJobTemplateLocaleUpdateParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.OrganizationJobTemplateLocalesApi.OrganizationJobTemplateLocaleUpdate(auth, accountId, jobTemplateId, jobTemplateLocaleId, organizationJobTemplateLocaleUpdateParameters, &localVarOptionals)

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

	OrganizationJobTemplateLocalesApiCmd.AddCommand(OrganizationJobTemplateLocaleUpdate)
	AddFlag(OrganizationJobTemplateLocaleUpdate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(OrganizationJobTemplateLocaleUpdate, "string", helpers.ToSnakeCase("JobTemplateId"), "", "Job Template ID", true)
	AddFlag(OrganizationJobTemplateLocaleUpdate, "string", helpers.ToSnakeCase("JobTemplateLocaleId"), "", "Job Template Locale ID", true)
	AddFlag(OrganizationJobTemplateLocaleUpdate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(OrganizationJobTemplateLocaleUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(OrganizationJobTemplateLocaleUpdate.Flags())
}
func initOrganizationJobTemplateLocalesCreate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("organization_job_template_locales/create", "/")[1:], "_")
	var OrganizationJobTemplateLocalesCreate = &cobra.Command{
		Use:   use,
		Short: "Create an organization job template locale",
		Long:  `Create a new organization job template locale.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.OrganizationJobTemplateLocalesCreateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			jobTemplateId := params.GetString(helpers.ToSnakeCase("JobTemplateId"))

			var organizationJobTemplateLocalesCreateParameters api.OrganizationJobTemplateLocalesCreateParameters
			if err := json.Unmarshal([]byte(params.GetString("data")), &organizationJobTemplateLocalesCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", organizationJobTemplateLocalesCreateParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.OrganizationJobTemplateLocalesApi.OrganizationJobTemplateLocalesCreate(auth, accountId, jobTemplateId, organizationJobTemplateLocalesCreateParameters, &localVarOptionals)

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

	OrganizationJobTemplateLocalesApiCmd.AddCommand(OrganizationJobTemplateLocalesCreate)
	AddFlag(OrganizationJobTemplateLocalesCreate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(OrganizationJobTemplateLocalesCreate, "string", helpers.ToSnakeCase("JobTemplateId"), "", "Job Template ID", true)
	AddFlag(OrganizationJobTemplateLocalesCreate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(OrganizationJobTemplateLocalesCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(OrganizationJobTemplateLocalesCreate.Flags())
}
func initOrganizationJobTemplateLocalesList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("organization_job_template_locales/list", "/")[1:], "_")
	var OrganizationJobTemplateLocalesList = &cobra.Command{
		Use:   use,
		Short: "List organization job template locales",
		Long:  `List all job template locales for a given organization job template.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.OrganizationJobTemplateLocalesListOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			jobTemplateId := params.GetString(helpers.ToSnakeCase("JobTemplateId"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("page")) {
				localVarOptionals.Page = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("Page")))
			}

			if params.IsSet(helpers.ToSnakeCase("perPage")) {
				localVarOptionals.PerPage = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("PerPage")))
			}

			data, api_response, err := client.OrganizationJobTemplateLocalesApi.OrganizationJobTemplateLocalesList(auth, accountId, jobTemplateId, &localVarOptionals)

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

	OrganizationJobTemplateLocalesApiCmd.AddCommand(OrganizationJobTemplateLocalesList)
	AddFlag(OrganizationJobTemplateLocalesList, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(OrganizationJobTemplateLocalesList, "string", helpers.ToSnakeCase("JobTemplateId"), "", "Job Template ID", true)
	AddFlag(OrganizationJobTemplateLocalesList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(OrganizationJobTemplateLocalesList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(OrganizationJobTemplateLocalesList, "int32", helpers.ToSnakeCase("PerPage"), "", "Limit on the number of objects to be returned, between 1 and 100. 25 by default", false)

	params.BindPFlags(OrganizationJobTemplateLocalesList.Flags())
}
