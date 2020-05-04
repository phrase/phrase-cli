package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

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

	rootCmd.AddCommand(jobLocalesApiCmd)
}

var jobLocalesApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("joblocalesapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("JobLocalesApi", "Api"), "API"}, " "),
}


func initJobLocaleComplete() {
	params := viper.New()
	var jobLocaleComplete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobLocaleComplete", strings.TrimSuffix("JobLocalesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobLocalesApi", "Api"), "s"))),
		Short: "Complete a job locale",
		Long:  `Mark a job locale as completed.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobLocaleCompleteOpts{}

			
			projectId := params.GetString("projectId")

			
			jobId := params.GetString("jobId")

			
			id := params.GetString("id")

			

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

	jobLocalesApiCmd.AddCommand(jobLocaleComplete)

	
	AddFlag(jobLocaleComplete, "string", "projectId", "", "Project ID", true)
	
	AddFlag(jobLocaleComplete, "string", "jobId", "", "Job ID", true)
	
	AddFlag(jobLocaleComplete, "string", "id", "", "ID", true)
	
	AddFlag(jobLocaleComplete, "string", "data", "d", "payload in JSON format", true)
	// jobLocaleCompleteParameters := api.JobLocaleCompleteParameters{}
	

	params.BindPFlags(jobLocaleComplete.Flags())
}

func initJobLocaleDelete() {
	params := viper.New()
	var jobLocaleDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobLocaleDelete", strings.TrimSuffix("JobLocalesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobLocalesApi", "Api"), "s"))),
		Short: "Delete a job locale",
		Long:  `Delete an existing job locale.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobLocaleDeleteOpts{}

			
			projectId := params.GetString("projectId")

			
			jobId := params.GetString("jobId")

			
			id := params.GetString("id")

			

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

	jobLocalesApiCmd.AddCommand(jobLocaleDelete)

	
	AddFlag(jobLocaleDelete, "string", "projectId", "", "Project ID", true)
	
	AddFlag(jobLocaleDelete, "string", "jobId", "", "Job ID", true)
	
	AddFlag(jobLocaleDelete, "string", "id", "", "ID", true)
	

	params.BindPFlags(jobLocaleDelete.Flags())
}

func initJobLocaleReopen() {
	params := viper.New()
	var jobLocaleReopen = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobLocaleReopen", strings.TrimSuffix("JobLocalesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobLocalesApi", "Api"), "s"))),
		Short: "Reopen a job locale",
		Long:  `Mark a job locale as uncompleted.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobLocaleReopenOpts{}

			
			projectId := params.GetString("projectId")

			
			jobId := params.GetString("jobId")

			
			id := params.GetString("id")

			

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

	jobLocalesApiCmd.AddCommand(jobLocaleReopen)

	
	AddFlag(jobLocaleReopen, "string", "projectId", "", "Project ID", true)
	
	AddFlag(jobLocaleReopen, "string", "jobId", "", "Job ID", true)
	
	AddFlag(jobLocaleReopen, "string", "id", "", "ID", true)
	
	AddFlag(jobLocaleReopen, "string", "data", "d", "payload in JSON format", true)
	// jobLocaleReopenParameters := api.JobLocaleReopenParameters{}
	

	params.BindPFlags(jobLocaleReopen.Flags())
}

func initJobLocaleShow() {
	params := viper.New()
	var jobLocaleShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobLocaleShow", strings.TrimSuffix("JobLocalesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobLocalesApi", "Api"), "s"))),
		Short: "Get a single job locale",
		Long:  `Get a single job locale for a given job.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobLocaleShowOpts{}

			
			projectId := params.GetString("projectId")

			
			jobId := params.GetString("jobId")

			
			id := params.GetString("id")

			

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

	jobLocalesApiCmd.AddCommand(jobLocaleShow)

	
	AddFlag(jobLocaleShow, "string", "projectId", "", "Project ID", true)
	
	AddFlag(jobLocaleShow, "string", "jobId", "", "Job ID", true)
	
	AddFlag(jobLocaleShow, "string", "id", "", "ID", true)
	

	params.BindPFlags(jobLocaleShow.Flags())
}

func initJobLocaleUpdate() {
	params := viper.New()
	var jobLocaleUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobLocaleUpdate", strings.TrimSuffix("JobLocalesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobLocalesApi", "Api"), "s"))),
		Short: "Update a job locale",
		Long:  `Update an existing job locale.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobLocaleUpdateOpts{}

			
			projectId := params.GetString("projectId")

			
			jobId := params.GetString("jobId")

			
			id := params.GetString("id")

			

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

	jobLocalesApiCmd.AddCommand(jobLocaleUpdate)

	
	AddFlag(jobLocaleUpdate, "string", "projectId", "", "Project ID", true)
	
	AddFlag(jobLocaleUpdate, "string", "jobId", "", "Job ID", true)
	
	AddFlag(jobLocaleUpdate, "string", "id", "", "ID", true)
	
	AddFlag(jobLocaleUpdate, "string", "data", "d", "payload in JSON format", true)
	// jobLocaleUpdateParameters := api.JobLocaleUpdateParameters{}
	

	params.BindPFlags(jobLocaleUpdate.Flags())
}

func initJobLocalesCreate() {
	params := viper.New()
	var jobLocalesCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobLocalesCreate", strings.TrimSuffix("JobLocalesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobLocalesApi", "Api"), "s"))),
		Short: "Create a job locale",
		Long:  `Create a new job locale.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobLocalesCreateOpts{}

			
			projectId := params.GetString("projectId")

			
			jobId := params.GetString("jobId")

			

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

	jobLocalesApiCmd.AddCommand(jobLocalesCreate)

	
	AddFlag(jobLocalesCreate, "string", "projectId", "", "Project ID", true)
	
	AddFlag(jobLocalesCreate, "string", "jobId", "", "Job ID", true)
	
	AddFlag(jobLocalesCreate, "string", "data", "d", "payload in JSON format", true)
	// jobLocalesCreateParameters := api.JobLocalesCreateParameters{}
	

	params.BindPFlags(jobLocalesCreate.Flags())
}

func initJobLocalesList() {
	params := viper.New()
	var jobLocalesList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobLocalesList", strings.TrimSuffix("JobLocalesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobLocalesApi", "Api"), "s"))),
		Short: "List job locales",
		Long:  `List all job locales for a given job.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobLocalesListOpts{}

			
			projectId := params.GetString("projectId")

			
			jobId := params.GetString("jobId")

			

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

	jobLocalesApiCmd.AddCommand(jobLocalesList)

	
	AddFlag(jobLocalesList, "string", "projectId", "", "Project ID", true)
	
	AddFlag(jobLocalesList, "string", "jobId", "", "Job ID", true)
	

	params.BindPFlags(jobLocalesList.Flags())
}

