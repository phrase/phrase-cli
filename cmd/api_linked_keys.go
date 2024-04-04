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
	initKeyLinksBatchDestroy()
	initKeyLinksCreate()
	initKeyLinksDestroy()
	initKeyLinksIndex()

	rootCmd.AddCommand(LinkedKeysApiCmd)
}

var LinkedKeysApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("LinkedKeys"),
	Short: "LinkedKeys API",
}

func initKeyLinksBatchDestroy() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("key_links/batch_destroy", "/")[1:], "_")
	var KeyLinksBatchDestroy = &cobra.Command{
		Use:   use,
		Short: "Batch unlink child keys from a parent key",
		Long:  `Unlinks multiple child keys from a given parent key in a single operation.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.KeyLinksBatchDestroyOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			var keyLinksBatchDestroyParameters api.KeyLinksBatchDestroyParameters
			if err := json.Unmarshal([]byte(params.GetString("data")), &keyLinksBatchDestroyParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", keyLinksBatchDestroyParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("unlinkParent")) {
				localVarOptionals.UnlinkParent = optional.NewBool(params.GetBool(helpers.ToSnakeCase("UnlinkParent")))
			}

			data, api_response, err := client.LinkedKeysApi.KeyLinksBatchDestroy(auth, projectId, id, keyLinksBatchDestroyParameters, &localVarOptionals)

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

	LinkedKeysApiCmd.AddCommand(KeyLinksBatchDestroy)
	AddFlag(KeyLinksBatchDestroy, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(KeyLinksBatchDestroy, "string", helpers.ToSnakeCase("Id"), "", "Parent Translation Key ID", true)
	AddFlag(KeyLinksBatchDestroy, "string", "data", "d", "payload in JSON format", true)
	AddFlag(KeyLinksBatchDestroy, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(KeyLinksBatchDestroy, "bool", helpers.ToSnakeCase("UnlinkParent"), "", "Whether to unlink the parent key as well and unmark it as linked-key.", false)

	params.BindPFlags(KeyLinksBatchDestroy.Flags())
}
func initKeyLinksCreate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("key_links/create", "/")[1:], "_")
	var KeyLinksCreate = &cobra.Command{
		Use:   use,
		Short: "Link child keys to a parent key",
		Long:  `Creates links between a given parent key and one or more child keys.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.KeyLinksCreateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			var keyLinksCreateParameters api.KeyLinksCreateParameters
			if err := json.Unmarshal([]byte(params.GetString("data")), &keyLinksCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", keyLinksCreateParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.LinkedKeysApi.KeyLinksCreate(auth, projectId, id, keyLinksCreateParameters, &localVarOptionals)

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

	LinkedKeysApiCmd.AddCommand(KeyLinksCreate)
	AddFlag(KeyLinksCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(KeyLinksCreate, "string", helpers.ToSnakeCase("Id"), "", "Parent Translation Key ID", true)
	AddFlag(KeyLinksCreate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(KeyLinksCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(KeyLinksCreate.Flags())
}
func initKeyLinksDestroy() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("key_links/destroy", "/")[1:], "_")
	var KeyLinksDestroy = &cobra.Command{
		Use:   use,
		Short: "Unlink a child key from a parent key",
		Long:  `Unlinks a single child key from a given parent key.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.KeyLinksDestroyOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			childKeyId := params.GetString(helpers.ToSnakeCase("ChildKeyId"))
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.LinkedKeysApi.KeyLinksDestroy(auth, projectId, id, childKeyId, &localVarOptionals)

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

	LinkedKeysApiCmd.AddCommand(KeyLinksDestroy)
	AddFlag(KeyLinksDestroy, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(KeyLinksDestroy, "string", helpers.ToSnakeCase("Id"), "", "Parent Translation Key ID", true)
	AddFlag(KeyLinksDestroy, "string", helpers.ToSnakeCase("ChildKeyId"), "", "The ID of the child key to unlink.", true)
	AddFlag(KeyLinksDestroy, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(KeyLinksDestroy.Flags())
}
func initKeyLinksIndex() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("key_links/index", "/")[1:], "_")
	var KeyLinksIndex = &cobra.Command{
		Use:   use,
		Short: "Retrieve all child keys linked to a specific parent key",
		Long:  `Returns detailed information about a parent key, including its linked child keys.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.KeyLinksIndexOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.LinkedKeysApi.KeyLinksIndex(auth, projectId, id, &localVarOptionals)

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

	LinkedKeysApiCmd.AddCommand(KeyLinksIndex)
	AddFlag(KeyLinksIndex, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(KeyLinksIndex, "string", helpers.ToSnakeCase("Id"), "", "Parent Translation Key ID", true)
	AddFlag(KeyLinksIndex, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(KeyLinksIndex.Flags())
}
