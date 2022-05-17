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
	initJobComplete()
	initJobCreate()
	initJobDelete()
	initJobKeysCreate()
	initJobKeysDelete()
	initJobLock()
	initJobReopen()
	initJobShow()
	initJobStart()
	initJobUnlock()
	initJobUpdate()
	initJobsByAccount()
	initJobsList()

	rootCmd.AddCommand(JobsApiCmd)
}

var JobsApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("Jobs"),
	Short: "Jobs API",
}

func initJobComplete() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job/complete", "/")[1:], "_")
	var JobComplete = &cobra.Command{
		Use:   use,
		Short: "Complete a job",
		Long:  `Mark a job as completed.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobCompleteOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			jobCompleteParameters := api.JobCompleteParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &jobCompleteParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", jobCompleteParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.JobsApi.JobComplete(auth, projectId, id, jobCompleteParameters, &localVarOptionals)

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

	JobsApiCmd.AddCommand(JobComplete)
	AddFlag(JobComplete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobComplete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobComplete, "string", "data", "d", "payload in JSON format", true)
	AddFlag(JobComplete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(JobComplete.Flags())
}
func initJobCreate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job/create", "/")[1:], "_")
	var JobCreate = &cobra.Command{
		Use:   use,
		Short: "Create a job",
		Long:  `Create a new job.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobCreateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			jobCreateParameters := api.JobCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &jobCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", jobCreateParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.JobsApi.JobCreate(auth, projectId, jobCreateParameters, &localVarOptionals)

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

	JobsApiCmd.AddCommand(JobCreate)
	AddFlag(JobCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobCreate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(JobCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(JobCreate.Flags())
}
func initJobDelete() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job/delete", "/")[1:], "_")
	var JobDelete = &cobra.Command{
		Use:   use,
		Short: "Delete a job",
		Long:  `Delete an existing job.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobDeleteOpts{}

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

			data, api_response, err := client.JobsApi.JobDelete(auth, projectId, id, &localVarOptionals)

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

	JobsApiCmd.AddCommand(JobDelete)
	AddFlag(JobDelete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobDelete, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(JobDelete.Flags())
}
func initJobKeysCreate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job/keys/create", "/")[1:], "_")
	var JobKeysCreate = &cobra.Command{
		Use:   use,
		Short: "Add keys to job",
		Long:  `Add multiple keys to a existing job.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobKeysCreateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			jobKeysCreateParameters := api.JobKeysCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &jobKeysCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", jobKeysCreateParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.JobsApi.JobKeysCreate(auth, projectId, id, jobKeysCreateParameters, &localVarOptionals)

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

	JobsApiCmd.AddCommand(JobKeysCreate)
	AddFlag(JobKeysCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobKeysCreate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobKeysCreate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(JobKeysCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(JobKeysCreate.Flags())
}
func initJobKeysDelete() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job/keys/delete", "/")[1:], "_")
	var JobKeysDelete = &cobra.Command{
		Use:   use,
		Short: "Remove keys from job",
		Long:  `Remove multiple keys from existing job.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobKeysDeleteOpts{}

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

			if params.IsSet(helpers.ToSnakeCase("translationKeyIds")) {
				var translationKeyIds map[string]interface{}
				if err := json.Unmarshal([]byte(params.GetString(helpers.ToSnakeCase("TranslationKeyIds"))), &translationKeyIds); err != nil {
					HandleError(err)
				}
				localVarOptionals.TranslationKeyIds = optional.NewInterface(translationKeyIds)
			}

			data, api_response, err := client.JobsApi.JobKeysDelete(auth, projectId, id, &localVarOptionals)

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

	JobsApiCmd.AddCommand(JobKeysDelete)
	AddFlag(JobKeysDelete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobKeysDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobKeysDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobKeysDelete, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)
	AddFlag(JobKeysDelete, "string", helpers.ToSnakeCase("TranslationKeyIds"), "", "payload in JSON format", false)

	params.BindPFlags(JobKeysDelete.Flags())
}
func initJobLock() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job/lock", "/")[1:], "_")
	var JobLock = &cobra.Command{
		Use:   use,
		Short: "Lock a job",
		Long:  `If you are the job owner, you may lock a job using this API request.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobLockOpts{}

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

			data, api_response, err := client.JobsApi.JobLock(auth, projectId, id, &localVarOptionals)

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

	JobsApiCmd.AddCommand(JobLock)
	AddFlag(JobLock, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobLock, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobLock, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobLock, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(JobLock.Flags())
}
func initJobReopen() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job/reopen", "/")[1:], "_")
	var JobReopen = &cobra.Command{
		Use:   use,
		Short: "Reopen a job",
		Long:  `Mark a job as uncompleted.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobReopenOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			jobReopenParameters := api.JobReopenParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &jobReopenParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", jobReopenParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.JobsApi.JobReopen(auth, projectId, id, jobReopenParameters, &localVarOptionals)

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

	JobsApiCmd.AddCommand(JobReopen)
	AddFlag(JobReopen, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobReopen, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobReopen, "string", "data", "d", "payload in JSON format", true)
	AddFlag(JobReopen, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(JobReopen.Flags())
}
func initJobShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job/show", "/")[1:], "_")
	var JobShow = &cobra.Command{
		Use:   use,
		Short: "Get a single job",
		Long:  `Get details on a single job for a given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobShowOpts{}

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

			data, api_response, err := client.JobsApi.JobShow(auth, projectId, id, &localVarOptionals)

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

	JobsApiCmd.AddCommand(JobShow)
	AddFlag(JobShow, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobShow, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(JobShow.Flags())
}
func initJobStart() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job/start", "/")[1:], "_")
	var JobStart = &cobra.Command{
		Use:   use,
		Short: "Start a job",
		Long:  `Starts an existing job in state draft.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobStartOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			jobStartParameters := api.JobStartParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &jobStartParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", jobStartParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.JobsApi.JobStart(auth, projectId, id, jobStartParameters, &localVarOptionals)

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

	JobsApiCmd.AddCommand(JobStart)
	AddFlag(JobStart, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobStart, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobStart, "string", "data", "d", "payload in JSON format", true)
	AddFlag(JobStart, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(JobStart.Flags())
}
func initJobUnlock() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job/unlock", "/")[1:], "_")
	var JobUnlock = &cobra.Command{
		Use:   use,
		Short: "Unlock a job",
		Long:  `If you are the job owner, you may unlock a locked job using this API request.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobUnlockOpts{}

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

			data, api_response, err := client.JobsApi.JobUnlock(auth, projectId, id, &localVarOptionals)

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

	JobsApiCmd.AddCommand(JobUnlock)
	AddFlag(JobUnlock, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobUnlock, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobUnlock, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobUnlock, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(JobUnlock.Flags())
}
func initJobUpdate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("job/update", "/")[1:], "_")
	var JobUpdate = &cobra.Command{
		Use:   use,
		Short: "Update a job",
		Long:  `Update an existing job.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobUpdateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			jobUpdateParameters := api.JobUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &jobUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", jobUpdateParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.JobsApi.JobUpdate(auth, projectId, id, jobUpdateParameters, &localVarOptionals)

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

	JobsApiCmd.AddCommand(JobUpdate)
	AddFlag(JobUpdate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobUpdate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(JobUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(JobUpdate.Flags())
}
func initJobsByAccount() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("jobs/by_account", "/")[1:], "_")
	var JobsByAccount = &cobra.Command{
		Use:   use,
		Short: "List account jobs",
		Long:  `List all jobs for the given account.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobsByAccountOpts{}

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

			if params.IsSet(helpers.ToSnakeCase("ownedBy")) {
				localVarOptionals.OwnedBy = optional.NewString(params.GetString(helpers.ToSnakeCase("OwnedBy")))
			}

			if params.IsSet(helpers.ToSnakeCase("assignedTo")) {
				localVarOptionals.AssignedTo = optional.NewString(params.GetString(helpers.ToSnakeCase("AssignedTo")))
			}

			if params.IsSet(helpers.ToSnakeCase("state")) {
				localVarOptionals.State = optional.NewString(params.GetString(helpers.ToSnakeCase("State")))
			}

			data, api_response, err := client.JobsApi.JobsByAccount(auth, accountId, &localVarOptionals)

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

	JobsApiCmd.AddCommand(JobsByAccount)
	AddFlag(JobsByAccount, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(JobsByAccount, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobsByAccount, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(JobsByAccount, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 25 by default", false)
	AddFlag(JobsByAccount, "string", helpers.ToSnakeCase("OwnedBy"), "", "filter by user owning job", false)
	AddFlag(JobsByAccount, "string", helpers.ToSnakeCase("AssignedTo"), "", "filter by user assigned to job", false)
	AddFlag(JobsByAccount, "string", helpers.ToSnakeCase("State"), "", "filter by state of job Valid states are <code>draft</code>, <code>in_progress</code>, <code>completed</code>", false)

	params.BindPFlags(JobsByAccount.Flags())
}
func initJobsList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("jobs/list", "/")[1:], "_")
	var JobsList = &cobra.Command{
		Use:   use,
		Short: "List jobs",
		Long:  `List all jobs for the given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.JobsListOpts{}

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

			if params.IsSet(helpers.ToSnakeCase("ownedBy")) {
				localVarOptionals.OwnedBy = optional.NewString(params.GetString(helpers.ToSnakeCase("OwnedBy")))
			}

			if params.IsSet(helpers.ToSnakeCase("assignedTo")) {
				localVarOptionals.AssignedTo = optional.NewString(params.GetString(helpers.ToSnakeCase("AssignedTo")))
			}

			if params.IsSet(helpers.ToSnakeCase("state")) {
				localVarOptionals.State = optional.NewString(params.GetString(helpers.ToSnakeCase("State")))
			}

			data, api_response, err := client.JobsApi.JobsList(auth, projectId, &localVarOptionals)

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

	JobsApiCmd.AddCommand(JobsList)
	AddFlag(JobsList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobsList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobsList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(JobsList, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 25 by default", false)
	AddFlag(JobsList, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)
	AddFlag(JobsList, "string", helpers.ToSnakeCase("OwnedBy"), "", "filter by user owning job", false)
	AddFlag(JobsList, "string", helpers.ToSnakeCase("AssignedTo"), "", "filter by user assigned to job", false)
	AddFlag(JobsList, "string", helpers.ToSnakeCase("State"), "", "filter by state of job Valid states are <code>draft</code>, <code>in_progress</code>, <code>completed</code>", false)

	params.BindPFlags(JobsList.Flags())
}
