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
	initJobAnnotationDelete()
	initJobAnnotationUpdate()
	initJobAnnotationsList()
	initJobLocaleAnnotationDelete()
	initJobLocaleAnnotationUpdate()
	initJobLocaleAnnotationsList()

	rootCmd.AddCommand(JobAnnotationsApiCmd)
}

var JobAnnotationsApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("JobAnnotations"),
	Short: "JobAnnotations API",
}

func initJobAnnotationDelete() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job_annotation/delete", "/")[1:], "_")
	var JobAnnotationDelete = &cobra.Command{
		Use:   use,
		Short: "Delete a job annotation",
		Long:  `Delete an annotation for a job.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobAnnotationDeleteOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			jobId := params.GetString(helpers.ToSnakeCase("JobId"))

			id := params.GetString(helpers.ToSnakeCase("Id"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			data, api_response, err := client.JobAnnotationsApi.JobAnnotationDelete(auth, projectId, jobId, id, &localVarOptionals)

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

	JobAnnotationsApiCmd.AddCommand(JobAnnotationDelete)
	AddFlag(JobAnnotationDelete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobAnnotationDelete, "string", helpers.ToSnakeCase("JobId"), "", "Job ID", true)
	AddFlag(JobAnnotationDelete, "string", helpers.ToSnakeCase("Id"), "", "Name of the annotation to delete.", true)
	AddFlag(JobAnnotationDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobAnnotationDelete, "string", helpers.ToSnakeCase("Branch"), "", "Branch to use", false)

	params.BindPFlags(JobAnnotationDelete.Flags())
}
func initJobAnnotationUpdate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job_annotation/update", "/")[1:], "_")
	var JobAnnotationUpdate = &cobra.Command{
		Use:   use,
		Short: "Create/Update a job annotation",
		Long:  `Create or update an annotation for a job. If the annotation already exists, it will be updated; otherwise, a new annotation will be created.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobAnnotationUpdateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			jobId := params.GetString(helpers.ToSnakeCase("JobId"))

			id := params.GetString(helpers.ToSnakeCase("Id"))

			var jobAnnotationUpdateParameters api.JobAnnotationUpdateParameters
			if err := json.Unmarshal([]byte(params.GetString("data")), &jobAnnotationUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", jobAnnotationUpdateParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.JobAnnotationsApi.JobAnnotationUpdate(auth, projectId, jobId, id, jobAnnotationUpdateParameters, &localVarOptionals)

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

	JobAnnotationsApiCmd.AddCommand(JobAnnotationUpdate)
	AddFlag(JobAnnotationUpdate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobAnnotationUpdate, "string", helpers.ToSnakeCase("JobId"), "", "Job ID", true)
	AddFlag(JobAnnotationUpdate, "string", helpers.ToSnakeCase("Id"), "", "Name of the annotation to set or update.", true)
	AddFlag(JobAnnotationUpdate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(JobAnnotationUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(JobAnnotationUpdate.Flags())
}
func initJobAnnotationsList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job_annotations/list", "/")[1:], "_")
	var JobAnnotationsList = &cobra.Command{
		Use:   use,
		Short: "List job annotations",
		Long:  `Retrieve a list of annotations for a job.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobAnnotationsListOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			jobId := params.GetString(helpers.ToSnakeCase("JobId"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			data, api_response, err := client.JobAnnotationsApi.JobAnnotationsList(auth, projectId, jobId, &localVarOptionals)

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

	JobAnnotationsApiCmd.AddCommand(JobAnnotationsList)
	AddFlag(JobAnnotationsList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobAnnotationsList, "string", helpers.ToSnakeCase("JobId"), "", "Job ID", true)
	AddFlag(JobAnnotationsList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobAnnotationsList, "string", helpers.ToSnakeCase("Branch"), "", "Branch to use", false)

	params.BindPFlags(JobAnnotationsList.Flags())
}
func initJobLocaleAnnotationDelete() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job_locale_annotation/delete", "/")[1:], "_")
	var JobLocaleAnnotationDelete = &cobra.Command{
		Use:   use,
		Short: "Delete a job locale annotation",
		Long:  `Delete an annotation for a job locale.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobLocaleAnnotationDeleteOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			jobId := params.GetString(helpers.ToSnakeCase("JobId"))

			jobLocaleId := params.GetString(helpers.ToSnakeCase("JobLocaleId"))

			id := params.GetString(helpers.ToSnakeCase("Id"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			data, api_response, err := client.JobAnnotationsApi.JobLocaleAnnotationDelete(auth, projectId, jobId, jobLocaleId, id, &localVarOptionals)

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

	JobAnnotationsApiCmd.AddCommand(JobLocaleAnnotationDelete)
	AddFlag(JobLocaleAnnotationDelete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobLocaleAnnotationDelete, "string", helpers.ToSnakeCase("JobId"), "", "Job ID", true)
	AddFlag(JobLocaleAnnotationDelete, "string", helpers.ToSnakeCase("JobLocaleId"), "", "Job Locale ID", true)
	AddFlag(JobLocaleAnnotationDelete, "string", helpers.ToSnakeCase("Id"), "", "Name of the annotation to delete.", true)
	AddFlag(JobLocaleAnnotationDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobLocaleAnnotationDelete, "string", helpers.ToSnakeCase("Branch"), "", "Branch to use", false)

	params.BindPFlags(JobLocaleAnnotationDelete.Flags())
}
func initJobLocaleAnnotationUpdate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job_locale_annotation/update", "/")[1:], "_")
	var JobLocaleAnnotationUpdate = &cobra.Command{
		Use:   use,
		Short: "Create/Update a job locale annotation",
		Long:  `Create or update an annotation for a job locale. If the annotation already exists, it will be updated; otherwise, a new annotation will be created.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobLocaleAnnotationUpdateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			jobId := params.GetString(helpers.ToSnakeCase("JobId"))

			jobLocaleId := params.GetString(helpers.ToSnakeCase("JobLocaleId"))

			id := params.GetString(helpers.ToSnakeCase("Id"))

			var jobAnnotationUpdateParameters api.JobAnnotationUpdateParameters
			if err := json.Unmarshal([]byte(params.GetString("data")), &jobAnnotationUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", jobAnnotationUpdateParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.JobAnnotationsApi.JobLocaleAnnotationUpdate(auth, projectId, jobId, jobLocaleId, id, jobAnnotationUpdateParameters, &localVarOptionals)

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

	JobAnnotationsApiCmd.AddCommand(JobLocaleAnnotationUpdate)
	AddFlag(JobLocaleAnnotationUpdate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobLocaleAnnotationUpdate, "string", helpers.ToSnakeCase("JobId"), "", "Job ID", true)
	AddFlag(JobLocaleAnnotationUpdate, "string", helpers.ToSnakeCase("JobLocaleId"), "", "Job Locale ID", true)
	AddFlag(JobLocaleAnnotationUpdate, "string", helpers.ToSnakeCase("Id"), "", "Name of the annotation to set or update.", true)
	AddFlag(JobLocaleAnnotationUpdate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(JobLocaleAnnotationUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(JobLocaleAnnotationUpdate.Flags())
}
func initJobLocaleAnnotationsList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job_locale_annotations/list", "/")[1:], "_")
	var JobLocaleAnnotationsList = &cobra.Command{
		Use:   use,
		Short: "List job locale annotations",
		Long:  `Retrieve a list of annotations for a job locale.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobLocaleAnnotationsListOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			jobId := params.GetString(helpers.ToSnakeCase("JobId"))

			jobLocaleId := params.GetString(helpers.ToSnakeCase("JobLocaleId"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			data, api_response, err := client.JobAnnotationsApi.JobLocaleAnnotationsList(auth, projectId, jobId, jobLocaleId, &localVarOptionals)

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

	JobAnnotationsApiCmd.AddCommand(JobLocaleAnnotationsList)
	AddFlag(JobLocaleAnnotationsList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobLocaleAnnotationsList, "string", helpers.ToSnakeCase("JobId"), "", "Job ID", true)
	AddFlag(JobLocaleAnnotationsList, "string", helpers.ToSnakeCase("JobLocaleId"), "", "Job Locale ID", true)
	AddFlag(JobLocaleAnnotationsList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobLocaleAnnotationsList, "string", helpers.ToSnakeCase("Branch"), "", "Branch to use", false)

	params.BindPFlags(JobLocaleAnnotationsList.Flags())
}
