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
	initBranchCompare()
	initBranchCreate()
	initBranchDelete()
	initBranchMerge()
	initBranchShow()
	initBranchUpdate()
	initBranchesList()

	rootCmd.AddCommand(branchesApiCmd)
}

var branchesApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("branchesapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("BranchesApi", "Api"), "API"}, " "),
}


func initBranchCompare() {
	params := viper.New()
	var branchCompare = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("BranchCompare", strings.TrimSuffix("BranchesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("BranchesApi", "Api"), "s"))),
		Short: "Compare branches",
		Long:  `Compare branch with main branch.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.BranchCompareOpts{}

			
			projectId := params.GetString("projectId")
			
			name := params.GetString("name")
			
			branchCompareParameters := api.BranchCompareParameters{}
			

			api_response, err := client.BranchesApi.BranchCompare(auth, projectId, name, branchCompareParameters, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	branchesApiCmd.AddCommand(branchCompare)

	
	AddFlag(branchCompare, "string", "projectId", "", "ID")
	
	AddFlag(branchCompare, "string", "name", "", "ID")
	
	// branchCompareParameters := api.BranchCompareParameters{}
	

	params.BindPFlags(branchCompare.Flags())
}

func initBranchCreate() {
	params := viper.New()
	var branchCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("BranchCreate", strings.TrimSuffix("BranchesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("BranchesApi", "Api"), "s"))),
		Short: "Create a branch",
		Long:  `Create a new branch.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.BranchCreateOpts{}

			
			projectId := params.GetString("projectId")
			
			branchCreateParameters := api.BranchCreateParameters{}
			

			api_response, err := client.BranchesApi.BranchCreate(auth, projectId, branchCreateParameters, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	branchesApiCmd.AddCommand(branchCreate)

	
	AddFlag(branchCreate, "string", "projectId", "", "ID")
	
	// branchCreateParameters := api.BranchCreateParameters{}
	

	params.BindPFlags(branchCreate.Flags())
}

func initBranchDelete() {
	params := viper.New()
	var branchDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("BranchDelete", strings.TrimSuffix("BranchesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("BranchesApi", "Api"), "s"))),
		Short: "Delete a branch",
		Long:  `Delete an existing branch.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.BranchDeleteOpts{}

			
			projectId := params.GetString("projectId")
			
			name := params.GetString("name")
			

			api_response, err := client.BranchesApi.BranchDelete(auth, projectId, name, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	branchesApiCmd.AddCommand(branchDelete)

	
	AddFlag(branchDelete, "string", "projectId", "", "ID")
	
	AddFlag(branchDelete, "string", "name", "", "ID")
	

	params.BindPFlags(branchDelete.Flags())
}

func initBranchMerge() {
	params := viper.New()
	var branchMerge = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("BranchMerge", strings.TrimSuffix("BranchesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("BranchesApi", "Api"), "s"))),
		Short: "Merge a branch",
		Long:  `Merge an existing branch.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.BranchMergeOpts{}

			
			projectId := params.GetString("projectId")
			
			name := params.GetString("name")
			
			branchMergeParameters := api.BranchMergeParameters{}
			

			api_response, err := client.BranchesApi.BranchMerge(auth, projectId, name, branchMergeParameters, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	branchesApiCmd.AddCommand(branchMerge)

	
	AddFlag(branchMerge, "string", "projectId", "", "ID")
	
	AddFlag(branchMerge, "string", "name", "", "ID")
	
	// branchMergeParameters := api.BranchMergeParameters{}
	

	params.BindPFlags(branchMerge.Flags())
}

func initBranchShow() {
	params := viper.New()
	var branchShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("BranchShow", strings.TrimSuffix("BranchesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("BranchesApi", "Api"), "s"))),
		Short: "Get a single branch",
		Long:  `Get details on a single branch for a given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.BranchShowOpts{}

			
			projectId := params.GetString("projectId")
			
			name := params.GetString("name")
			

			data, api_response, err := client.BranchesApi.BranchShow(auth, projectId, name, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	branchesApiCmd.AddCommand(branchShow)

	
	AddFlag(branchShow, "string", "projectId", "", "ID")
	
	AddFlag(branchShow, "string", "name", "", "ID")
	

	params.BindPFlags(branchShow.Flags())
}

func initBranchUpdate() {
	params := viper.New()
	var branchUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("BranchUpdate", strings.TrimSuffix("BranchesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("BranchesApi", "Api"), "s"))),
		Short: "Update a branch",
		Long:  `Update an existing branch.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.BranchUpdateOpts{}

			
			projectId := params.GetString("projectId")
			
			name := params.GetString("name")
			
			branchUpdateParameters := api.BranchUpdateParameters{}
			

			data, api_response, err := client.BranchesApi.BranchUpdate(auth, projectId, name, branchUpdateParameters, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	branchesApiCmd.AddCommand(branchUpdate)

	
	AddFlag(branchUpdate, "string", "projectId", "", "ID")
	
	AddFlag(branchUpdate, "string", "name", "", "ID")
	
	// branchUpdateParameters := api.BranchUpdateParameters{}
	

	params.BindPFlags(branchUpdate.Flags())
}

func initBranchesList() {
	params := viper.New()
	var branchesList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("BranchesList", strings.TrimSuffix("BranchesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("BranchesApi", "Api"), "s"))),
		Short: "List branches",
		Long:  `List all branches the of the current project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.BranchesListOpts{}

			
			projectId := params.GetString("projectId")
			

			data, api_response, err := client.BranchesApi.BranchesList(auth, projectId, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	branchesApiCmd.AddCommand(branchesList)

	
	AddFlag(branchesList, "string", "projectId", "", "ID")
	

	params.BindPFlags(branchesList.Flags())
}

