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
	initJobTemplateLocaleDelete()
	initJobTemplateLocaleShow()
	initJobTemplateLocaleUpdate()
	initJobTemplateLocalesCreate()
	initJobTemplateLocalesList()

	rootCmd.AddCommand(JobTemplateLocalesApiCmd)
}

var JobTemplateLocalesApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("JobTemplateLocales"),
	Short: "JobTemplateLocales API",
}

func initJobTemplateLocaleDelete() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job_template_locale/delete", "/")[1:], "_")
	var JobTemplateLocaleDelete = &cobra.Command{
		Use:   use,
		Short: "Delete a job template locale",
		Long:  `Delete an existing job template locale.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobTemplateLocaleDeleteOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			jobTemplateId := params.GetString(helpers.ToSnakeCase("JobTemplateId"))
			jobTemplateLocaleId := params.GetString(helpers.ToSnakeCase("JobTemplateLocaleId"))
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			data, api_response, err := client.JobTemplateLocalesApi.JobTemplateLocaleDelete(auth, projectId, jobTemplateId, jobTemplateLocaleId, &localVarOptionals)

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

	JobTemplateLocalesApiCmd.AddCommand(JobTemplateLocaleDelete)
	AddFlag(JobTemplateLocaleDelete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobTemplateLocaleDelete, "string", helpers.ToSnakeCase("JobTemplateId"), "", "Job Template ID", true)
	AddFlag(JobTemplateLocaleDelete, "string", helpers.ToSnakeCase("JobTemplateLocaleId"), "", "Job Template Locale ID", true)
	AddFlag(JobTemplateLocaleDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobTemplateLocaleDelete, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(JobTemplateLocaleDelete.Flags())
}
func initJobTemplateLocaleShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job_template_locale/show", "/")[1:], "_")
	var JobTemplateLocaleShow = &cobra.Command{
		Use:   use,
		Short: "Get a single job template locale",
		Long:  `Get a single job template locale for a given job template.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobTemplateLocaleShowOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			jobTemplateId := params.GetString(helpers.ToSnakeCase("JobTemplateId"))
			jobTemplateLocaleId := params.GetString(helpers.ToSnakeCase("JobTemplateLocaleId"))
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			data, api_response, err := client.JobTemplateLocalesApi.JobTemplateLocaleShow(auth, projectId, jobTemplateId, jobTemplateLocaleId, &localVarOptionals)

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

	JobTemplateLocalesApiCmd.AddCommand(JobTemplateLocaleShow)
	AddFlag(JobTemplateLocaleShow, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobTemplateLocaleShow, "string", helpers.ToSnakeCase("JobTemplateId"), "", "Job Template ID", true)
	AddFlag(JobTemplateLocaleShow, "string", helpers.ToSnakeCase("JobTemplateLocaleId"), "", "Job Template Locale ID", true)
	AddFlag(JobTemplateLocaleShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobTemplateLocaleShow, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(JobTemplateLocaleShow.Flags())
}
func initJobTemplateLocaleUpdate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job_template_locale/update", "/")[1:], "_")
	var JobTemplateLocaleUpdate = &cobra.Command{
		Use:   use,
		Short: "Update a job template locale",
		Long:  `Update an existing job template locale.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobTemplateLocaleUpdateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			jobTemplateId := params.GetString(helpers.ToSnakeCase("JobTemplateId"))
			jobTemplateLocaleId := params.GetString(helpers.ToSnakeCase("JobTemplateLocaleId"))
			var jobTemplateLocaleUpdateParameters api.JobTemplateLocaleUpdateParameters
			if err := json.Unmarshal([]byte(params.GetString("data")), &jobTemplateLocaleUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", jobTemplateLocaleUpdateParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.JobTemplateLocalesApi.JobTemplateLocaleUpdate(auth, projectId, jobTemplateId, jobTemplateLocaleId, jobTemplateLocaleUpdateParameters, &localVarOptionals)

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

	JobTemplateLocalesApiCmd.AddCommand(JobTemplateLocaleUpdate)
	AddFlag(JobTemplateLocaleUpdate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobTemplateLocaleUpdate, "string", helpers.ToSnakeCase("JobTemplateId"), "", "Job Template ID", true)
	AddFlag(JobTemplateLocaleUpdate, "string", helpers.ToSnakeCase("JobTemplateLocaleId"), "", "Job Template Locale ID", true)
	AddFlag(JobTemplateLocaleUpdate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(JobTemplateLocaleUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(JobTemplateLocaleUpdate.Flags())
}
func initJobTemplateLocalesCreate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job_template_locales/create", "/")[1:], "_")
	var JobTemplateLocalesCreate = &cobra.Command{
		Use:   use,
		Short: "Create a job template locale",
		Long:  `Create a new job template locale.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobTemplateLocalesCreateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			jobTemplateId := params.GetString(helpers.ToSnakeCase("JobTemplateId"))
			var jobTemplateLocalesCreateParameters api.JobTemplateLocalesCreateParameters
			if err := json.Unmarshal([]byte(params.GetString("data")), &jobTemplateLocalesCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", jobTemplateLocalesCreateParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.JobTemplateLocalesApi.JobTemplateLocalesCreate(auth, projectId, jobTemplateId, jobTemplateLocalesCreateParameters, &localVarOptionals)

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

	JobTemplateLocalesApiCmd.AddCommand(JobTemplateLocalesCreate)
	AddFlag(JobTemplateLocalesCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobTemplateLocalesCreate, "string", helpers.ToSnakeCase("JobTemplateId"), "", "Job Template ID", true)
	AddFlag(JobTemplateLocalesCreate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(JobTemplateLocalesCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(JobTemplateLocalesCreate.Flags())
}
func initJobTemplateLocalesList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job_template_locales/list", "/")[1:], "_")
	var JobTemplateLocalesList = &cobra.Command{
		Use:   use,
		Short: "List job template locales",
		Long:  `List all job template locales for a given job template.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobTemplateLocalesListOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
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

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			data, api_response, err := client.JobTemplateLocalesApi.JobTemplateLocalesList(auth, projectId, jobTemplateId, &localVarOptionals)

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

	JobTemplateLocalesApiCmd.AddCommand(JobTemplateLocalesList)
	AddFlag(JobTemplateLocalesList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobTemplateLocalesList, "string", helpers.ToSnakeCase("JobTemplateId"), "", "Job Template ID", true)
	AddFlag(JobTemplateLocalesList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobTemplateLocalesList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(JobTemplateLocalesList, "int32", helpers.ToSnakeCase("PerPage"), "", "Limit on the number of objects to be returned, between 1 and 100. 25 by default", false)
	AddFlag(JobTemplateLocalesList, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(JobTemplateLocalesList.Flags())
}
