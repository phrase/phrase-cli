package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/antihax/optional"
	helpers "github.com/phrase/phrase-cli/helpers"
	api "github.com/phrase/phrase-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initJobLocaleComplete()
	initJobLocaleDelete()
	initJobLocaleReopen()
	initJobLocaleShow()
	initJobLocaleUpdate()
	initJobLocalesCreate()
	initJobLocalesList()

	rootCmd.AddCommand(JobLocalesApiCmd)
}

var JobLocalesApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("JobLocales"),
	Short: "JobLocales API",
}

func initJobLocaleComplete() {
	params := viper.New()
	var JobLocaleComplete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobLocaleComplete", strings.TrimSuffix("JobLocalesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobLocalesApi", "Api"), "s"))),
		Short: "Complete a job locale",
		Long:  `Mark a job locale as completed.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobLocaleCompleteOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			jobId := params.GetString(helpers.ToSnakeCase("JobId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			jobLocaleCompleteParameters := api.JobLocaleCompleteParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &jobLocaleCompleteParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", jobLocaleCompleteParameters)
			}
			data, api_response, err := client.JobLocalesApi.JobLocaleComplete(auth, projectId, jobId, id, jobLocaleCompleteParameters, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	JobLocalesApiCmd.AddCommand(JobLocaleComplete)

	AddFlag(JobLocaleComplete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobLocaleComplete, "string", helpers.ToSnakeCase("JobId"), "", "Job ID", true)
	AddFlag(JobLocaleComplete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobLocaleComplete, "string", "data", "d", "payload in JSON format", true)

	AddFlag(JobLocaleComplete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(JobLocaleComplete.Flags())
}
func initJobLocaleDelete() {
	params := viper.New()
	var JobLocaleDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobLocaleDelete", strings.TrimSuffix("JobLocalesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobLocalesApi", "Api"), "s"))),
		Short: "Delete a job locale",
		Long:  `Delete an existing job locale.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobLocaleDeleteOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			jobId := params.GetString(helpers.ToSnakeCase("JobId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.JobLocalesApi.JobLocaleDelete(auth, projectId, jobId, id, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	JobLocalesApiCmd.AddCommand(JobLocaleDelete)

	AddFlag(JobLocaleDelete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobLocaleDelete, "string", helpers.ToSnakeCase("JobId"), "", "Job ID", true)
	AddFlag(JobLocaleDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobLocaleDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobLocaleDelete, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)
	params.BindPFlags(JobLocaleDelete.Flags())
}
func initJobLocaleReopen() {
	params := viper.New()
	var JobLocaleReopen = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobLocaleReopen", strings.TrimSuffix("JobLocalesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobLocalesApi", "Api"), "s"))),
		Short: "Reopen a job locale",
		Long:  `Mark a job locale as uncompleted.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobLocaleReopenOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			jobId := params.GetString(helpers.ToSnakeCase("JobId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			jobLocaleReopenParameters := api.JobLocaleReopenParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &jobLocaleReopenParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", jobLocaleReopenParameters)
			}
			data, api_response, err := client.JobLocalesApi.JobLocaleReopen(auth, projectId, jobId, id, jobLocaleReopenParameters, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	JobLocalesApiCmd.AddCommand(JobLocaleReopen)

	AddFlag(JobLocaleReopen, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobLocaleReopen, "string", helpers.ToSnakeCase("JobId"), "", "Job ID", true)
	AddFlag(JobLocaleReopen, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobLocaleReopen, "string", "data", "d", "payload in JSON format", true)

	AddFlag(JobLocaleReopen, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(JobLocaleReopen.Flags())
}
func initJobLocaleShow() {
	params := viper.New()
	var JobLocaleShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobLocaleShow", strings.TrimSuffix("JobLocalesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobLocalesApi", "Api"), "s"))),
		Short: "Get a single job locale",
		Long:  `Get a single job locale for a given job.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobLocaleShowOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			jobId := params.GetString(helpers.ToSnakeCase("JobId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.JobLocalesApi.JobLocaleShow(auth, projectId, jobId, id, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	JobLocalesApiCmd.AddCommand(JobLocaleShow)

	AddFlag(JobLocaleShow, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobLocaleShow, "string", helpers.ToSnakeCase("JobId"), "", "Job ID", true)
	AddFlag(JobLocaleShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobLocaleShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobLocaleShow, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)
	params.BindPFlags(JobLocaleShow.Flags())
}
func initJobLocaleUpdate() {
	params := viper.New()
	var JobLocaleUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobLocaleUpdate", strings.TrimSuffix("JobLocalesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobLocalesApi", "Api"), "s"))),
		Short: "Update a job locale",
		Long:  `Update an existing job locale.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobLocaleUpdateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			jobId := params.GetString(helpers.ToSnakeCase("JobId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			jobLocaleUpdateParameters := api.JobLocaleUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &jobLocaleUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", jobLocaleUpdateParameters)
			}
			data, api_response, err := client.JobLocalesApi.JobLocaleUpdate(auth, projectId, jobId, id, jobLocaleUpdateParameters, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	JobLocalesApiCmd.AddCommand(JobLocaleUpdate)

	AddFlag(JobLocaleUpdate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobLocaleUpdate, "string", helpers.ToSnakeCase("JobId"), "", "Job ID", true)
	AddFlag(JobLocaleUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobLocaleUpdate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(JobLocaleUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(JobLocaleUpdate.Flags())
}
func initJobLocalesCreate() {
	params := viper.New()
	var JobLocalesCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobLocalesCreate", strings.TrimSuffix("JobLocalesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobLocalesApi", "Api"), "s"))),
		Short: "Create a job locale",
		Long:  `Create a new job locale.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobLocalesCreateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			jobId := params.GetString(helpers.ToSnakeCase("JobId"))

			jobLocalesCreateParameters := api.JobLocalesCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &jobLocalesCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", jobLocalesCreateParameters)
			}
			data, api_response, err := client.JobLocalesApi.JobLocalesCreate(auth, projectId, jobId, jobLocalesCreateParameters, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	JobLocalesApiCmd.AddCommand(JobLocalesCreate)

	AddFlag(JobLocalesCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobLocalesCreate, "string", helpers.ToSnakeCase("JobId"), "", "Job ID", true)
	AddFlag(JobLocalesCreate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(JobLocalesCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(JobLocalesCreate.Flags())
}
func initJobLocalesList() {
	params := viper.New()
	var JobLocalesList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobLocalesList", strings.TrimSuffix("JobLocalesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobLocalesApi", "Api"), "s"))),
		Short: "List job locales",
		Long:  `List all job locales for a given job.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobLocalesListOpts{}

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

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			jobId := params.GetString(helpers.ToSnakeCase("JobId"))

			data, api_response, err := client.JobLocalesApi.JobLocalesList(auth, projectId, jobId, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	JobLocalesApiCmd.AddCommand(JobLocalesList)

	AddFlag(JobLocalesList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobLocalesList, "string", helpers.ToSnakeCase("JobId"), "", "Job ID", true)
	AddFlag(JobLocalesList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobLocalesList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(JobLocalesList, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 10 by default", false)
	AddFlag(JobLocalesList, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)
	params.BindPFlags(JobLocalesList.Flags())
}
