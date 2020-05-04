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
	initDistributionCreate()
	initDistributionDelete()
	initDistributionShow()
	initDistributionUpdate()
	initDistributionsList()

	rootCmd.AddCommand(distributionsApiCmd)
}

var distributionsApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("distributionsapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("DistributionsApi", "Api"), "API"}, " "),
}


func initDistributionCreate() {
	params := viper.New()
	var distributionCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("DistributionCreate", strings.TrimSuffix("DistributionsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("DistributionsApi", "Api"), "s"))),
		Short: "Create a distribution",
		Long:  `Create a new distribution.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.DistributionCreateOpts{}

			
			accountId := params.GetString("accountId")

			

			distributionCreateParameters := api.DistributionCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &distributionCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", distributionCreateParameters)
			}
			

			data, api_response, err := client.DistributionsApi.DistributionCreate(auth, accountId, distributionCreateParameters, &localVarOptionals)

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

	distributionsApiCmd.AddCommand(distributionCreate)

	
	AddFlag(distributionCreate, "string", "accountId", "", "Account ID", true)
	
	AddFlag(distributionCreate, "string", "data", "d", "payload in JSON format", true)
	// distributionCreateParameters := api.DistributionCreateParameters{}
	

	params.BindPFlags(distributionCreate.Flags())
}

func initDistributionDelete() {
	params := viper.New()
	var distributionDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("DistributionDelete", strings.TrimSuffix("DistributionsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("DistributionsApi", "Api"), "s"))),
		Short: "Delete a distribution",
		Long:  `Delete an existing distribution.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.DistributionDeleteOpts{}

			
			accountId := params.GetString("accountId")

			
			id := params.GetString("id")

			

			data, api_response, err := client.DistributionsApi.DistributionDelete(auth, accountId, id, &localVarOptionals)

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

	distributionsApiCmd.AddCommand(distributionDelete)

	
	AddFlag(distributionDelete, "string", "accountId", "", "Account ID", true)
	
	AddFlag(distributionDelete, "string", "id", "", "ID", true)
	

	params.BindPFlags(distributionDelete.Flags())
}

func initDistributionShow() {
	params := viper.New()
	var distributionShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("DistributionShow", strings.TrimSuffix("DistributionsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("DistributionsApi", "Api"), "s"))),
		Short: "Get a single distribution",
		Long:  `Get details on a single distribution.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.DistributionShowOpts{}

			
			accountId := params.GetString("accountId")

			
			id := params.GetString("id")

			

			data, api_response, err := client.DistributionsApi.DistributionShow(auth, accountId, id, &localVarOptionals)

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

	distributionsApiCmd.AddCommand(distributionShow)

	
	AddFlag(distributionShow, "string", "accountId", "", "Account ID", true)
	
	AddFlag(distributionShow, "string", "id", "", "ID", true)
	

	params.BindPFlags(distributionShow.Flags())
}

func initDistributionUpdate() {
	params := viper.New()
	var distributionUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("DistributionUpdate", strings.TrimSuffix("DistributionsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("DistributionsApi", "Api"), "s"))),
		Short: "Update a distribution",
		Long:  `Update an existing distribution.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.DistributionUpdateOpts{}

			
			accountId := params.GetString("accountId")

			
			id := params.GetString("id")

			

			distributionUpdateParameters := api.DistributionUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &distributionUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", distributionUpdateParameters)
			}
			

			data, api_response, err := client.DistributionsApi.DistributionUpdate(auth, accountId, id, distributionUpdateParameters, &localVarOptionals)

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

	distributionsApiCmd.AddCommand(distributionUpdate)

	
	AddFlag(distributionUpdate, "string", "accountId", "", "Account ID", true)
	
	AddFlag(distributionUpdate, "string", "id", "", "ID", true)
	
	AddFlag(distributionUpdate, "string", "data", "d", "payload in JSON format", true)
	// distributionUpdateParameters := api.DistributionUpdateParameters{}
	

	params.BindPFlags(distributionUpdate.Flags())
}

func initDistributionsList() {
	params := viper.New()
	var distributionsList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("DistributionsList", strings.TrimSuffix("DistributionsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("DistributionsApi", "Api"), "s"))),
		Short: "List distributions",
		Long:  `List all distributions for the given account.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.DistributionsListOpts{}

			
			accountId := params.GetString("accountId")

			

			data, api_response, err := client.DistributionsApi.DistributionsList(auth, accountId, &localVarOptionals)

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

	distributionsApiCmd.AddCommand(distributionsList)

	
	AddFlag(distributionsList, "string", "accountId", "", "Account ID", true)
	

	params.BindPFlags(distributionsList.Flags())
}

