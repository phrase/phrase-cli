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
	initReleaseCreate()
	initReleaseDelete()
	initReleasePublish()
	initReleaseShow()
	initReleaseUpdate()
	initReleasesList()

	rootCmd.AddCommand(ReleasesApiCmd)
}

var ReleasesApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("Releases"),
	Short: "Releases API",
}

func initReleaseCreate() {
	params := viper.New()
	var ReleaseCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ReleaseCreate", strings.TrimSuffix("ReleasesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ReleasesApi", "Api"), "s"))),
		Short: "Create a release",
		Long:  `Create a new release.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ReleaseCreateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			distributionId := params.GetString(helpers.ToSnakeCase("DistributionId"))

			releaseCreateParameters := api.ReleaseCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &releaseCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", releaseCreateParameters)
			}
			data, api_response, err := client.ReleasesApi.ReleaseCreate(auth, accountId, distributionId, releaseCreateParameters, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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

	ReleasesApiCmd.AddCommand(ReleaseCreate)

	AddFlag(ReleaseCreate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(ReleaseCreate, "string", helpers.ToSnakeCase("DistributionId"), "", "Distribution ID", true)
	AddFlag(ReleaseCreate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(ReleaseCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(ReleaseCreate.Flags())
}
func initReleaseDelete() {
	params := viper.New()
	var ReleaseDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ReleaseDelete", strings.TrimSuffix("ReleasesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ReleasesApi", "Api"), "s"))),
		Short: "Delete a release",
		Long:  `Delete an existing release.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ReleaseDeleteOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			distributionId := params.GetString(helpers.ToSnakeCase("DistributionId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.ReleasesApi.ReleaseDelete(auth, accountId, distributionId, id, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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

	ReleasesApiCmd.AddCommand(ReleaseDelete)

	AddFlag(ReleaseDelete, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(ReleaseDelete, "string", helpers.ToSnakeCase("DistributionId"), "", "Distribution ID", true)
	AddFlag(ReleaseDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(ReleaseDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(ReleaseDelete.Flags())
}
func initReleasePublish() {
	params := viper.New()
	var ReleasePublish = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ReleasePublish", strings.TrimSuffix("ReleasesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ReleasesApi", "Api"), "s"))),
		Short: "Publish a release",
		Long:  `Publish a release for production.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ReleasePublishOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			distributionId := params.GetString(helpers.ToSnakeCase("DistributionId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.ReleasesApi.ReleasePublish(auth, accountId, distributionId, id, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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

	ReleasesApiCmd.AddCommand(ReleasePublish)

	AddFlag(ReleasePublish, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(ReleasePublish, "string", helpers.ToSnakeCase("DistributionId"), "", "Distribution ID", true)
	AddFlag(ReleasePublish, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(ReleasePublish, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(ReleasePublish.Flags())
}
func initReleaseShow() {
	params := viper.New()
	var ReleaseShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ReleaseShow", strings.TrimSuffix("ReleasesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ReleasesApi", "Api"), "s"))),
		Short: "Get a single release",
		Long:  `Get details on a single release.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ReleaseShowOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			distributionId := params.GetString(helpers.ToSnakeCase("DistributionId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.ReleasesApi.ReleaseShow(auth, accountId, distributionId, id, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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

	ReleasesApiCmd.AddCommand(ReleaseShow)

	AddFlag(ReleaseShow, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(ReleaseShow, "string", helpers.ToSnakeCase("DistributionId"), "", "Distribution ID", true)
	AddFlag(ReleaseShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(ReleaseShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(ReleaseShow.Flags())
}
func initReleaseUpdate() {
	params := viper.New()
	var ReleaseUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ReleaseUpdate", strings.TrimSuffix("ReleasesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ReleasesApi", "Api"), "s"))),
		Short: "Update a release",
		Long:  `Update an existing release.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ReleaseUpdateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			distributionId := params.GetString(helpers.ToSnakeCase("DistributionId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			releaseUpdateParameters := api.ReleaseUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &releaseUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", releaseUpdateParameters)
			}
			data, api_response, err := client.ReleasesApi.ReleaseUpdate(auth, accountId, distributionId, id, releaseUpdateParameters, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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

	ReleasesApiCmd.AddCommand(ReleaseUpdate)

	AddFlag(ReleaseUpdate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(ReleaseUpdate, "string", helpers.ToSnakeCase("DistributionId"), "", "Distribution ID", true)
	AddFlag(ReleaseUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(ReleaseUpdate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(ReleaseUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(ReleaseUpdate.Flags())
}
func initReleasesList() {
	params := viper.New()
	var ReleasesList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ReleasesList", strings.TrimSuffix("ReleasesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ReleasesApi", "Api"), "s"))),
		Short: "List releases",
		Long:  `List all releases for the given distribution.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ReleasesListOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("page")) {
				localVarOptionals.Page = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("Page")))
			}
			if params.IsSet(helpers.ToSnakeCase("perPage")) {
				localVarOptionals.PerPage = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("PerPage")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			distributionId := params.GetString(helpers.ToSnakeCase("DistributionId"))

			data, api_response, err := client.ReleasesApi.ReleasesList(auth, accountId, distributionId, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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

	ReleasesApiCmd.AddCommand(ReleasesList)

	AddFlag(ReleasesList, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(ReleasesList, "string", helpers.ToSnakeCase("DistributionId"), "", "Distribution ID", true)
	AddFlag(ReleasesList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(ReleasesList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(ReleasesList, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 10 by default", false)
	params.BindPFlags(ReleasesList.Flags())
}
