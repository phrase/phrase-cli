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
	initJobTemplateCreate()
	initJobTemplateDelete()
	initJobTemplateUpdate()
	initJobTemplatesList()
	initJobTemplatesShow()

	rootCmd.AddCommand(JobTemplatesApiCmd)
}

var JobTemplatesApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("JobTemplates"),
	Short: "JobTemplates API",
}

func initJobTemplateCreate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job_template/create", "/")[1:], "_")
	var JobTemplateCreate = &cobra.Command{
		Use:   use,
		Short: "Create a job template",
		Long:  `Create a new job template.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobTemplateCreateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			jobTemplateCreateParameters := api.JobTemplateCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &jobTemplateCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", jobTemplateCreateParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.JobTemplatesApi.JobTemplateCreate(auth, projectId, jobTemplateCreateParameters, &localVarOptionals)

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

	JobTemplatesApiCmd.AddCommand(JobTemplateCreate)
	AddFlag(JobTemplateCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobTemplateCreate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(JobTemplateCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(JobTemplateCreate.Flags())
}
func initJobTemplateDelete() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job_template/delete", "/")[1:], "_")
	var JobTemplateDelete = &cobra.Command{
		Use:   use,
		Short: "Delete a job template",
		Long:  `Delete an existing job template.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobTemplateDeleteOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			data, api_response, err := client.JobTemplatesApi.JobTemplateDelete(auth, projectId, id, &localVarOptionals)

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

	JobTemplatesApiCmd.AddCommand(JobTemplateDelete)
	AddFlag(JobTemplateDelete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobTemplateDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobTemplateDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobTemplateDelete, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(JobTemplateDelete.Flags())
}
func initJobTemplateUpdate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job_template/update", "/")[1:], "_")
	var JobTemplateUpdate = &cobra.Command{
		Use:   use,
		Short: "Update a job template",
		Long:  `Update an existing job template.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobTemplateUpdateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			jobTemplateUpdateParameters := api.JobTemplateUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &jobTemplateUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", jobTemplateUpdateParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.JobTemplatesApi.JobTemplateUpdate(auth, projectId, id, jobTemplateUpdateParameters, &localVarOptionals)

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

	JobTemplatesApiCmd.AddCommand(JobTemplateUpdate)
	AddFlag(JobTemplateUpdate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobTemplateUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobTemplateUpdate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(JobTemplateUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(JobTemplateUpdate.Flags())
}
func initJobTemplatesList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job_templates/list", "/")[1:], "_")
	var JobTemplatesList = &cobra.Command{
		Use:   use,
		Short: "List job templates",
		Long:  `List all job templates for the given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobTemplatesListOpts{}

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

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			data, api_response, err := client.JobTemplatesApi.JobTemplatesList(auth, projectId, &localVarOptionals)

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

	JobTemplatesApiCmd.AddCommand(JobTemplatesList)
	AddFlag(JobTemplatesList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobTemplatesList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobTemplatesList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(JobTemplatesList, "int32", helpers.ToSnakeCase("PerPage"), "", "Limit on the number of objects to be returned, between 1 and 100. 25 by default", false)
	AddFlag(JobTemplatesList, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(JobTemplatesList.Flags())
}
func initJobTemplatesShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job_templates/show", "/")[1:], "_")
	var JobTemplatesShow = &cobra.Command{
		Use:   use,
		Short: "Get a single job template",
		Long:  `Get details on a single job template for a given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobTemplatesShowOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			data, api_response, err := client.JobTemplatesApi.JobTemplatesShow(auth, projectId, id, &localVarOptionals)

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

	JobTemplatesApiCmd.AddCommand(JobTemplatesShow)
	AddFlag(JobTemplatesShow, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobTemplatesShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobTemplatesShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobTemplatesShow, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(JobTemplatesShow.Flags())
}
