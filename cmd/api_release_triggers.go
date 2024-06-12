package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/antihax/optional"
	helpers "github.com/phrase/phrase-cli/helpers"
	api "github.com/phrase/phrase-go/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initReleaseTriggersCreate()
	initReleaseTriggersDestroy()
	initReleaseTriggersList()
	initReleaseTriggersShow()
	initReleaseTriggersUpdate()

	rootCmd.AddCommand(ReleaseTriggersApiCmd)
}

var ReleaseTriggersApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("ReleaseTriggers"),
	Short: "ReleaseTriggers API",
}

func initReleaseTriggersCreate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("release_triggers/create", "/")[1:], "_")
	var ReleaseTriggersCreate = &cobra.Command{
		Use:   use,
		Short: "Create a release trigger",
		Long:  `Create a new recurring release. New releases will be published automatically, based on the cron schedule provided. Currently, only one release trigger can exist per distribution.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.ReleaseTriggersCreateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			distributionId := params.GetString(helpers.ToSnakeCase("DistributionId"))

			var releaseCreateParameters1 api.ReleaseCreateParameters1
			if err := json.Unmarshal([]byte(params.GetString("data")), &releaseCreateParameters1); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", releaseCreateParameters1)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.ReleaseTriggersApi.ReleaseTriggersCreate(auth, accountId, distributionId, releaseCreateParameters1, &localVarOptionals)

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

	ReleaseTriggersApiCmd.AddCommand(ReleaseTriggersCreate)
	AddFlag(ReleaseTriggersCreate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(ReleaseTriggersCreate, "string", helpers.ToSnakeCase("DistributionId"), "", "Distribution ID", true)
	AddFlag(ReleaseTriggersCreate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(ReleaseTriggersCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(ReleaseTriggersCreate.Flags())
}
func initReleaseTriggersDestroy() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("release_triggers/destroy", "/")[1:], "_")
	var ReleaseTriggersDestroy = &cobra.Command{
		Use:   use,
		Short: "Delete a single release trigger",
		Long:  `Delete a single release trigger.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.ReleaseTriggersDestroyOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			distributionId := params.GetString(helpers.ToSnakeCase("DistributionId"))

			id := params.GetString(helpers.ToSnakeCase("Id"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.ReleaseTriggersApi.ReleaseTriggersDestroy(auth, accountId, distributionId, id, &localVarOptionals)

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

	ReleaseTriggersApiCmd.AddCommand(ReleaseTriggersDestroy)
	AddFlag(ReleaseTriggersDestroy, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(ReleaseTriggersDestroy, "string", helpers.ToSnakeCase("DistributionId"), "", "Distribution ID", true)
	AddFlag(ReleaseTriggersDestroy, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(ReleaseTriggersDestroy, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(ReleaseTriggersDestroy.Flags())
}
func initReleaseTriggersList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("release_triggers/list", "/")[1:], "_")
	var ReleaseTriggersList = &cobra.Command{
		Use:   use,
		Short: "List release triggers",
		Long:  `List all release triggers for the given distribution.&lt;br&gt; Note: Currently only one release trigger can exist per distribution. `,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.ReleaseTriggersListOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			distributionId := params.GetString(helpers.ToSnakeCase("DistributionId"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.ReleaseTriggersApi.ReleaseTriggersList(auth, accountId, distributionId, &localVarOptionals)

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

	ReleaseTriggersApiCmd.AddCommand(ReleaseTriggersList)
	AddFlag(ReleaseTriggersList, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(ReleaseTriggersList, "string", helpers.ToSnakeCase("DistributionId"), "", "Distribution ID", true)
	AddFlag(ReleaseTriggersList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(ReleaseTriggersList.Flags())
}
func initReleaseTriggersShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("release_triggers/show", "/")[1:], "_")
	var ReleaseTriggersShow = &cobra.Command{
		Use:   use,
		Short: "Get a single release trigger",
		Long:  `Get details of a single release trigger.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.ReleaseTriggersShowOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			distributionId := params.GetString(helpers.ToSnakeCase("DistributionId"))

			id := params.GetString(helpers.ToSnakeCase("Id"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.ReleaseTriggersApi.ReleaseTriggersShow(auth, accountId, distributionId, id, &localVarOptionals)

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

	ReleaseTriggersApiCmd.AddCommand(ReleaseTriggersShow)
	AddFlag(ReleaseTriggersShow, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(ReleaseTriggersShow, "string", helpers.ToSnakeCase("DistributionId"), "", "Distribution ID", true)
	AddFlag(ReleaseTriggersShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(ReleaseTriggersShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(ReleaseTriggersShow.Flags())
}
func initReleaseTriggersUpdate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("release_triggers/update", "/")[1:], "_")
	var ReleaseTriggersUpdate = &cobra.Command{
		Use:   use,
		Short: "Update a release trigger",
		Long:  `Update a recurring release.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.ReleaseTriggersUpdateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			distributionId := params.GetString(helpers.ToSnakeCase("DistributionId"))

			id := params.GetString(helpers.ToSnakeCase("Id"))

			var releaseUpdateParameters1 api.ReleaseUpdateParameters1
			if err := json.Unmarshal([]byte(params.GetString("data")), &releaseUpdateParameters1); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", releaseUpdateParameters1)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.ReleaseTriggersApi.ReleaseTriggersUpdate(auth, accountId, distributionId, id, releaseUpdateParameters1, &localVarOptionals)

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

	ReleaseTriggersApiCmd.AddCommand(ReleaseTriggersUpdate)
	AddFlag(ReleaseTriggersUpdate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(ReleaseTriggersUpdate, "string", helpers.ToSnakeCase("DistributionId"), "", "Distribution ID", true)
	AddFlag(ReleaseTriggersUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(ReleaseTriggersUpdate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(ReleaseTriggersUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(ReleaseTriggersUpdate.Flags())
}
