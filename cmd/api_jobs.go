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
	initJobComplete()
	initJobCreate()
	initJobDelete()
	initJobKeysCreate()
	initJobKeysDelete()
	initJobReopen()
	initJobShow()
	initJobStart()
	initJobUpdate()
	initJobsList()

	rootCmd.AddCommand(jobsApiCmd)
}

var jobsApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("jobsapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("JobsApi", "Api"), "API"}, " "),
}


func initJobComplete() {
	params := viper.New()
	var jobComplete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobComplete", strings.TrimSuffix("JobsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobsApi", "Api"), "s"))),
		Short: "Complete a job",
		Long:  `Mark a job as completed.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobCompleteOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			
			jobCompleteParameters := api.JobCompleteParameters{}
			

			data, api_response, err := client.JobsApi.JobComplete(auth, projectId, id, jobCompleteParameters, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	jobsApiCmd.AddCommand(jobComplete)

	
	AddFlag(jobComplete, "string", "projectId", "", "ID")
	
	AddFlag(jobComplete, "string", "id", "", "ID")
	
	// jobCompleteParameters := api.JobCompleteParameters{}
	

	params.BindPFlags(jobComplete.Flags())
}

func initJobCreate() {
	params := viper.New()
	var jobCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobCreate", strings.TrimSuffix("JobsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobsApi", "Api"), "s"))),
		Short: "Create a job",
		Long:  `Create a new job.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobCreateOpts{}

			
			projectId := params.GetString("projectId")
			
			jobCreateParameters := api.JobCreateParameters{}
			

			data, api_response, err := client.JobsApi.JobCreate(auth, projectId, jobCreateParameters, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	jobsApiCmd.AddCommand(jobCreate)

	
	AddFlag(jobCreate, "string", "projectId", "", "ID")
	
	// jobCreateParameters := api.JobCreateParameters{}
	

	params.BindPFlags(jobCreate.Flags())
}

func initJobDelete() {
	params := viper.New()
	var jobDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobDelete", strings.TrimSuffix("JobsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobsApi", "Api"), "s"))),
		Short: "Delete a job",
		Long:  `Delete an existing job.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobDeleteOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			

			data, api_response, err := client.JobsApi.JobDelete(auth, projectId, id, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	jobsApiCmd.AddCommand(jobDelete)

	
	AddFlag(jobDelete, "string", "projectId", "", "ID")
	
	AddFlag(jobDelete, "string", "id", "", "ID")
	

	params.BindPFlags(jobDelete.Flags())
}

func initJobKeysCreate() {
	params := viper.New()
	var jobKeysCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobKeysCreate", strings.TrimSuffix("JobsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobsApi", "Api"), "s"))),
		Short: "Add keys to job",
		Long:  `Add multiple keys to a existing job.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobKeysCreateOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			
			jobKeysCreateParameters := api.JobKeysCreateParameters{}
			

			data, api_response, err := client.JobsApi.JobKeysCreate(auth, projectId, id, jobKeysCreateParameters, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	jobsApiCmd.AddCommand(jobKeysCreate)

	
	AddFlag(jobKeysCreate, "string", "projectId", "", "ID")
	
	AddFlag(jobKeysCreate, "string", "id", "", "ID")
	
	// jobKeysCreateParameters := api.JobKeysCreateParameters{}
	

	params.BindPFlags(jobKeysCreate.Flags())
}

func initJobKeysDelete() {
	params := viper.New()
	var jobKeysDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobKeysDelete", strings.TrimSuffix("JobsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobsApi", "Api"), "s"))),
		Short: "Remove keys from job",
		Long:  `Remove multiple keys from existing job.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobKeysDeleteOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			

			data, api_response, err := client.JobsApi.JobKeysDelete(auth, projectId, id, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	jobsApiCmd.AddCommand(jobKeysDelete)

	
	AddFlag(jobKeysDelete, "string", "projectId", "", "ID")
	
	AddFlag(jobKeysDelete, "string", "id", "", "ID")
	

	params.BindPFlags(jobKeysDelete.Flags())
}

func initJobReopen() {
	params := viper.New()
	var jobReopen = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobReopen", strings.TrimSuffix("JobsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobsApi", "Api"), "s"))),
		Short: "Reopen a job",
		Long:  `Mark a job as uncompleted.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobReopenOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			
			jobReopenParameters := api.JobReopenParameters{}
			

			data, api_response, err := client.JobsApi.JobReopen(auth, projectId, id, jobReopenParameters, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	jobsApiCmd.AddCommand(jobReopen)

	
	AddFlag(jobReopen, "string", "projectId", "", "ID")
	
	AddFlag(jobReopen, "string", "id", "", "ID")
	
	// jobReopenParameters := api.JobReopenParameters{}
	

	params.BindPFlags(jobReopen.Flags())
}

func initJobShow() {
	params := viper.New()
	var jobShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobShow", strings.TrimSuffix("JobsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobsApi", "Api"), "s"))),
		Short: "Get a single job",
		Long:  `Get details on a single job for a given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobShowOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			

			data, api_response, err := client.JobsApi.JobShow(auth, projectId, id, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	jobsApiCmd.AddCommand(jobShow)

	
	AddFlag(jobShow, "string", "projectId", "", "ID")
	
	AddFlag(jobShow, "string", "id", "", "ID")
	

	params.BindPFlags(jobShow.Flags())
}

func initJobStart() {
	params := viper.New()
	var jobStart = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobStart", strings.TrimSuffix("JobsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobsApi", "Api"), "s"))),
		Short: "Start a job",
		Long:  `Starts an existing job in state draft.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobStartOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			
			jobStartParameters := api.JobStartParameters{}
			

			data, api_response, err := client.JobsApi.JobStart(auth, projectId, id, jobStartParameters, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	jobsApiCmd.AddCommand(jobStart)

	
	AddFlag(jobStart, "string", "projectId", "", "ID")
	
	AddFlag(jobStart, "string", "id", "", "ID")
	
	// jobStartParameters := api.JobStartParameters{}
	

	params.BindPFlags(jobStart.Flags())
}

func initJobUpdate() {
	params := viper.New()
	var jobUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobUpdate", strings.TrimSuffix("JobsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobsApi", "Api"), "s"))),
		Short: "Update a job",
		Long:  `Update an existing job.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobUpdateOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			
			jobUpdateParameters := api.JobUpdateParameters{}
			

			data, api_response, err := client.JobsApi.JobUpdate(auth, projectId, id, jobUpdateParameters, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	jobsApiCmd.AddCommand(jobUpdate)

	
	AddFlag(jobUpdate, "string", "projectId", "", "ID")
	
	AddFlag(jobUpdate, "string", "id", "", "ID")
	
	// jobUpdateParameters := api.JobUpdateParameters{}
	

	params.BindPFlags(jobUpdate.Flags())
}

func initJobsList() {
	params := viper.New()
	var jobsList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobsList", strings.TrimSuffix("JobsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobsApi", "Api"), "s"))),
		Short: "List jobs",
		Long:  `List all jobs for the given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobsListOpts{}

			
			projectId := params.GetString("projectId")
			

			data, api_response, err := client.JobsApi.JobsList(auth, projectId, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	jobsApiCmd.AddCommand(jobsList)

	
	AddFlag(jobsList, "string", "projectId", "", "ID")
	

	params.BindPFlags(jobsList.Flags())
}

