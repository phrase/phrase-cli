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
	initReleaseCreate()
	initReleaseDelete()
	initReleasePublish()
	initReleaseShow()
	initReleaseUpdate()
	initReleasesList()

	rootCmd.AddCommand(releasesApiCmd)
}

var releasesApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("releasesapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("ReleasesApi", "Api"), "API"}, " "),
}


func initReleaseCreate() {
	params := viper.New()
	var releaseCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ReleaseCreate", strings.TrimSuffix("ReleasesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ReleasesApi", "Api"), "s"))),
		Short: "Create a release",
		Long:  `Create a new release.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ReleaseCreateOpts{}

			
			accountId := params.GetString("accountId")

			
			distributionId := params.GetString("distributionId")

			

			releaseCreateParameters := api.ReleaseCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &releaseCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", releaseCreateParameters)
			}
			

			data, api_response, err := client.ReleasesApi.ReleaseCreate(auth, accountId, distributionId, releaseCreateParameters, &localVarOptionals)

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

	releasesApiCmd.AddCommand(releaseCreate)

	
	AddFlag(releaseCreate, "string", "accountId", "", "Account ID", true)
	
	AddFlag(releaseCreate, "string", "distributionId", "", "Distribution ID", true)
	
	AddFlag(releaseCreate, "string", "data", "d", "payload in JSON format", true)
	// releaseCreateParameters := api.ReleaseCreateParameters{}
	

	params.BindPFlags(releaseCreate.Flags())
}

func initReleaseDelete() {
	params := viper.New()
	var releaseDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ReleaseDelete", strings.TrimSuffix("ReleasesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ReleasesApi", "Api"), "s"))),
		Short: "Delete a release",
		Long:  `Delete an existing release.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ReleaseDeleteOpts{}

			
			accountId := params.GetString("accountId")

			
			distributionId := params.GetString("distributionId")

			
			id := params.GetString("id")

			

			data, api_response, err := client.ReleasesApi.ReleaseDelete(auth, accountId, distributionId, id, &localVarOptionals)

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

	releasesApiCmd.AddCommand(releaseDelete)

	
	AddFlag(releaseDelete, "string", "accountId", "", "Account ID", true)
	
	AddFlag(releaseDelete, "string", "distributionId", "", "Distribution ID", true)
	
	AddFlag(releaseDelete, "string", "id", "", "ID", true)
	

	params.BindPFlags(releaseDelete.Flags())
}

func initReleasePublish() {
	params := viper.New()
	var releasePublish = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ReleasePublish", strings.TrimSuffix("ReleasesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ReleasesApi", "Api"), "s"))),
		Short: "Publish a release",
		Long:  `Publish a release for production.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ReleasePublishOpts{}

			
			accountId := params.GetString("accountId")

			
			distributionId := params.GetString("distributionId")

			
			id := params.GetString("id")

			

			data, api_response, err := client.ReleasesApi.ReleasePublish(auth, accountId, distributionId, id, &localVarOptionals)

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

	releasesApiCmd.AddCommand(releasePublish)

	
	AddFlag(releasePublish, "string", "accountId", "", "Account ID", true)
	
	AddFlag(releasePublish, "string", "distributionId", "", "Distribution ID", true)
	
	AddFlag(releasePublish, "string", "id", "", "ID", true)
	

	params.BindPFlags(releasePublish.Flags())
}

func initReleaseShow() {
	params := viper.New()
	var releaseShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ReleaseShow", strings.TrimSuffix("ReleasesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ReleasesApi", "Api"), "s"))),
		Short: "Get a single release",
		Long:  `Get details on a single release.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ReleaseShowOpts{}

			
			accountId := params.GetString("accountId")

			
			distributionId := params.GetString("distributionId")

			
			id := params.GetString("id")

			

			data, api_response, err := client.ReleasesApi.ReleaseShow(auth, accountId, distributionId, id, &localVarOptionals)

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

	releasesApiCmd.AddCommand(releaseShow)

	
	AddFlag(releaseShow, "string", "accountId", "", "Account ID", true)
	
	AddFlag(releaseShow, "string", "distributionId", "", "Distribution ID", true)
	
	AddFlag(releaseShow, "string", "id", "", "ID", true)
	

	params.BindPFlags(releaseShow.Flags())
}

func initReleaseUpdate() {
	params := viper.New()
	var releaseUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ReleaseUpdate", strings.TrimSuffix("ReleasesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ReleasesApi", "Api"), "s"))),
		Short: "Update a release",
		Long:  `Update an existing release.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ReleaseUpdateOpts{}

			
			accountId := params.GetString("accountId")

			
			distributionId := params.GetString("distributionId")

			
			id := params.GetString("id")

			

			releaseUpdateParameters := api.ReleaseUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &releaseUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", releaseUpdateParameters)
			}
			

			data, api_response, err := client.ReleasesApi.ReleaseUpdate(auth, accountId, distributionId, id, releaseUpdateParameters, &localVarOptionals)

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

	releasesApiCmd.AddCommand(releaseUpdate)

	
	AddFlag(releaseUpdate, "string", "accountId", "", "Account ID", true)
	
	AddFlag(releaseUpdate, "string", "distributionId", "", "Distribution ID", true)
	
	AddFlag(releaseUpdate, "string", "id", "", "ID", true)
	
	AddFlag(releaseUpdate, "string", "data", "d", "payload in JSON format", true)
	// releaseUpdateParameters := api.ReleaseUpdateParameters{}
	

	params.BindPFlags(releaseUpdate.Flags())
}

func initReleasesList() {
	params := viper.New()
	var releasesList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ReleasesList", strings.TrimSuffix("ReleasesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ReleasesApi", "Api"), "s"))),
		Short: "List releases",
		Long:  `List all releases for the given distribution.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ReleasesListOpts{}

			
			accountId := params.GetString("accountId")

			
			distributionId := params.GetString("distributionId")

			

			data, api_response, err := client.ReleasesApi.ReleasesList(auth, accountId, distributionId, &localVarOptionals)

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

	releasesApiCmd.AddCommand(releasesList)

	
	AddFlag(releasesList, "string", "accountId", "", "Account ID", true)
	
	AddFlag(releasesList, "string", "distributionId", "", "Distribution ID", true)
	

	params.BindPFlags(releasesList.Flags())
}

