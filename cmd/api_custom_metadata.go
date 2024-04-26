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
	initCustomMetadataPropertiesDelete()
	initCustomMetadataPropertiesList()
	initCustomMetadataPropertyCreate()
	initCustomMetadataPropertyShow()
	initCustomMetadataPropertyUpdate()

	rootCmd.AddCommand(CustomMetadataApiCmd)
}

var CustomMetadataApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("CustomMetadata"),
	Short: "CustomMetadata API",
}

func initCustomMetadataPropertiesDelete() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("custom_metadata_properties/delete", "/")[1:], "_")
	var CustomMetadataPropertiesDelete = &cobra.Command{
		Use:   use,
		Short: "Destroy property",
		Long:  `Destroy a custom metadata property of an account.  This endpoint is only available to accounts with advanced plans or above. `,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.CustomMetadataPropertiesDeleteOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			id := params.GetString(helpers.ToSnakeCase("Id"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.CustomMetadataApi.CustomMetadataPropertiesDelete(auth, accountId, id, &localVarOptionals)

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

	CustomMetadataApiCmd.AddCommand(CustomMetadataPropertiesDelete)
	AddFlag(CustomMetadataPropertiesDelete, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(CustomMetadataPropertiesDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(CustomMetadataPropertiesDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(CustomMetadataPropertiesDelete.Flags())
}
func initCustomMetadataPropertiesList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("custom_metadata_properties/list", "/")[1:], "_")
	var CustomMetadataPropertiesList = &cobra.Command{
		Use:   use,
		Short: "List properties",
		Long:  `List all custom metadata properties for an account.  This endpoint is only available to accounts with advanced plans or above. `,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.CustomMetadataPropertiesListOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("dataType")) {
				var dataType map[string]interface{}
				if err := json.Unmarshal([]byte(params.GetString(helpers.ToSnakeCase("DataType"))), &dataType); err != nil {
					HandleError(err)
				}
				localVarOptionals.DataType = optional.NewInterface(dataType)
			}

			if params.IsSet(helpers.ToSnakeCase("projectId")) {
				localVarOptionals.ProjectId = optional.NewString(params.GetString(helpers.ToSnakeCase("ProjectId")))
			}

			if params.IsSet(helpers.ToSnakeCase("page")) {
				localVarOptionals.Page = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("Page")))
			}

			if params.IsSet(helpers.ToSnakeCase("perPage")) {
				localVarOptionals.PerPage = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("PerPage")))
			}

			if params.IsSet(helpers.ToSnakeCase("q")) {
				localVarOptionals.Q = optional.NewString(params.GetString(helpers.ToSnakeCase("Q")))
			}

			if params.IsSet(helpers.ToSnakeCase("sort")) {
				localVarOptionals.Sort = optional.NewString(params.GetString(helpers.ToSnakeCase("Sort")))
			}

			if params.IsSet(helpers.ToSnakeCase("order")) {
				localVarOptionals.Order = optional.NewString(params.GetString(helpers.ToSnakeCase("Order")))
			}

			data, api_response, err := client.CustomMetadataApi.CustomMetadataPropertiesList(auth, accountId, &localVarOptionals)

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

	CustomMetadataApiCmd.AddCommand(CustomMetadataPropertiesList)
	AddFlag(CustomMetadataPropertiesList, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(CustomMetadataPropertiesList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(CustomMetadataPropertiesList, "string", helpers.ToSnakeCase("DataType"), "", "payload in JSON format", false)
	AddFlag(CustomMetadataPropertiesList, "string", helpers.ToSnakeCase("ProjectId"), "", "id of project that the properties belong to", false)
	AddFlag(CustomMetadataPropertiesList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(CustomMetadataPropertiesList, "int32", helpers.ToSnakeCase("PerPage"), "", "Limit on the number of objects to be returned, between 1 and 100. 25 by default", false)
	AddFlag(CustomMetadataPropertiesList, "string", helpers.ToSnakeCase("Q"), "", "query to find a property by name", false)
	AddFlag(CustomMetadataPropertiesList, "string", helpers.ToSnakeCase("Sort"), "", "Sort criteria. Can be one of: name, data_type, created_at.", false)
	AddFlag(CustomMetadataPropertiesList, "string", helpers.ToSnakeCase("Order"), "", "Order direction. Can be one of: asc, desc.", false)

	params.BindPFlags(CustomMetadataPropertiesList.Flags())
}
func initCustomMetadataPropertyCreate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("custom_metadata_property/create", "/")[1:], "_")
	var CustomMetadataPropertyCreate = &cobra.Command{
		Use:   use,
		Short: "Create a property",
		Long:  `Create a new custom metadata property.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.CustomMetadataPropertyCreateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			var customMetadataPropertiesCreateParameters api.CustomMetadataPropertiesCreateParameters
			if err := json.Unmarshal([]byte(params.GetString("data")), &customMetadataPropertiesCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", customMetadataPropertiesCreateParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.CustomMetadataApi.CustomMetadataPropertyCreate(auth, accountId, customMetadataPropertiesCreateParameters, &localVarOptionals)

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

	CustomMetadataApiCmd.AddCommand(CustomMetadataPropertyCreate)
	AddFlag(CustomMetadataPropertyCreate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(CustomMetadataPropertyCreate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(CustomMetadataPropertyCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(CustomMetadataPropertyCreate.Flags())
}
func initCustomMetadataPropertyShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("custom_metadata_property/show", "/")[1:], "_")
	var CustomMetadataPropertyShow = &cobra.Command{
		Use:   use,
		Short: "Get a single property",
		Long:  `Get details of a single custom property.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.CustomMetadataPropertyShowOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			id := params.GetString(helpers.ToSnakeCase("Id"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.CustomMetadataApi.CustomMetadataPropertyShow(auth, accountId, id, &localVarOptionals)

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

	CustomMetadataApiCmd.AddCommand(CustomMetadataPropertyShow)
	AddFlag(CustomMetadataPropertyShow, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(CustomMetadataPropertyShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(CustomMetadataPropertyShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(CustomMetadataPropertyShow.Flags())
}
func initCustomMetadataPropertyUpdate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("custom_metadata_property/update", "/")[1:], "_")
	var CustomMetadataPropertyUpdate = &cobra.Command{
		Use:   use,
		Short: "Update a property",
		Long:  `Update an existing custom metadata property.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.CustomMetadataPropertyUpdateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			id := params.GetString(helpers.ToSnakeCase("Id"))

			var customMetadataPropertiesUpdateParameters api.CustomMetadataPropertiesUpdateParameters
			if err := json.Unmarshal([]byte(params.GetString("data")), &customMetadataPropertiesUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", customMetadataPropertiesUpdateParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.CustomMetadataApi.CustomMetadataPropertyUpdate(auth, accountId, id, customMetadataPropertiesUpdateParameters, &localVarOptionals)

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

	CustomMetadataApiCmd.AddCommand(CustomMetadataPropertyUpdate)
	AddFlag(CustomMetadataPropertyUpdate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(CustomMetadataPropertyUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(CustomMetadataPropertyUpdate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(CustomMetadataPropertyUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(CustomMetadataPropertyUpdate.Flags())
}
