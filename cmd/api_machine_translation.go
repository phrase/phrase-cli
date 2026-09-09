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
	initMachineTranslationLocaleProviderMappingsCreate()
	initMachineTranslationLocaleProviderMappingsDestroy()
	initMachineTranslationSettingsShow()
	initMachineTranslationSettingsUpdate()

	rootCmd.AddCommand(MachineTranslationApiCmd)
}

var MachineTranslationApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("MachineTranslation"),
	Short: "MachineTranslation API",
}

func initMachineTranslationLocaleProviderMappingsCreate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("machine_translation_locale_provider_mappings/create", "/")[1:], "_")
	var MachineTranslationLocaleProviderMappingsCreate = &cobra.Command{
		Use:   use,
		Short: "Create a locale provider mapping",
		Long:  `Creates a locale-pair-specific machine translation provider override for the account. When a mapping exists for a given source/target locale pair, that provider is used instead of the account default. Only one mapping may exist per source/target locale pair; attempting to create a duplicate returns a validation error. The source and target locale codes must differ. `,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.MachineTranslationLocaleProviderMappingsCreateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			var machineTranslationLocaleProviderMappingsCreateParameters api.MachineTranslationLocaleProviderMappingsCreateParameters
			if err := json.Unmarshal([]byte(params.GetString("data")), &machineTranslationLocaleProviderMappingsCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", machineTranslationLocaleProviderMappingsCreateParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.MachineTranslationApi.MachineTranslationLocaleProviderMappingsCreate(auth, accountId, machineTranslationLocaleProviderMappingsCreateParameters, &localVarOptionals)

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

	MachineTranslationApiCmd.AddCommand(MachineTranslationLocaleProviderMappingsCreate)
	AddFlag(MachineTranslationLocaleProviderMappingsCreate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(MachineTranslationLocaleProviderMappingsCreate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(MachineTranslationLocaleProviderMappingsCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(MachineTranslationLocaleProviderMappingsCreate.Flags())
}
func initMachineTranslationLocaleProviderMappingsDestroy() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("machine_translation_locale_provider_mappings/destroy", "/")[1:], "_")
	var MachineTranslationLocaleProviderMappingsDestroy = &cobra.Command{
		Use:   use,
		Short: "Delete a locale provider mapping",
		Long:  `Removes the machine translation provider override for the specified source and target locale pair. The mapping is identified by locale codes supplied as query parameters rather than a path ID. `,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.MachineTranslationLocaleProviderMappingsDestroyOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			sourceLocaleCode := params.GetString(helpers.ToSnakeCase("SourceLocaleCode"))

			targetLocaleCode := params.GetString(helpers.ToSnakeCase("TargetLocaleCode"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.MachineTranslationApi.MachineTranslationLocaleProviderMappingsDestroy(auth, accountId, sourceLocaleCode, targetLocaleCode, &localVarOptionals)

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

	MachineTranslationApiCmd.AddCommand(MachineTranslationLocaleProviderMappingsDestroy)
	AddFlag(MachineTranslationLocaleProviderMappingsDestroy, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(MachineTranslationLocaleProviderMappingsDestroy, "string", helpers.ToSnakeCase("SourceLocaleCode"), "", "The locale code of the source language of the mapping to delete.", true)
	AddFlag(MachineTranslationLocaleProviderMappingsDestroy, "string", helpers.ToSnakeCase("TargetLocaleCode"), "", "The locale code of the target language of the mapping to delete.", true)
	AddFlag(MachineTranslationLocaleProviderMappingsDestroy, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(MachineTranslationLocaleProviderMappingsDestroy.Flags())
}
func initMachineTranslationSettingsShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("machine_translation_settings/show", "/")[1:], "_")
	var MachineTranslationSettingsShow = &cobra.Command{
		Use:   use,
		Short: "Get machine translation settings",
		Long:  `Returns the machine translation configuration for the account, including the default translation service, current machine translation unit usage, and any locale-pair-specific provider mappings. `,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.MachineTranslationSettingsShowOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.MachineTranslationApi.MachineTranslationSettingsShow(auth, accountId, &localVarOptionals)

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

	MachineTranslationApiCmd.AddCommand(MachineTranslationSettingsShow)
	AddFlag(MachineTranslationSettingsShow, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(MachineTranslationSettingsShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(MachineTranslationSettingsShow.Flags())
}
func initMachineTranslationSettingsUpdate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("machine_translation_settings/update", "/")[1:], "_")
	var MachineTranslationSettingsUpdate = &cobra.Command{
		Use:   use,
		Short: "Update machine translation settings",
		Long:  `Sets the default machine translation service for the account. Requires write access to the account&#x27;s machine translation settings. Passing an empty or absent value for &#x60;default_service&#x60; resets the account to its plan default (Microsoft Translate). `,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.MachineTranslationSettingsUpdateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			var machineTranslationSettingsUpdateParameters api.MachineTranslationSettingsUpdateParameters
			if err := json.Unmarshal([]byte(params.GetString("data")), &machineTranslationSettingsUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", machineTranslationSettingsUpdateParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.MachineTranslationApi.MachineTranslationSettingsUpdate(auth, accountId, machineTranslationSettingsUpdateParameters, &localVarOptionals)

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

	MachineTranslationApiCmd.AddCommand(MachineTranslationSettingsUpdate)
	AddFlag(MachineTranslationSettingsUpdate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(MachineTranslationSettingsUpdate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(MachineTranslationSettingsUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(MachineTranslationSettingsUpdate.Flags())
}
