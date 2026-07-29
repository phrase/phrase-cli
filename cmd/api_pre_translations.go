package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/antihax/optional"
	helpers "github.com/phrase/phrase-cli/helpers"
	api "github.com/phrase/phrase-go/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initPreTranslationCreate()
	initPreTranslationShow()
	initPreTranslationsList()

	rootCmd.AddCommand(PreTranslationsApiCmd)
}

var PreTranslationsApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("PreTranslations"),
	Short: "PreTranslations API",
}

func initPreTranslationCreate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("pre_translation/create", "/")[1:], "_")
	var PreTranslationCreate = &cobra.Command{
		Use:   use,
		Short: "Create a pre-translation job",
		Long:  `Triggers a pre-translation job for a resource within a project, addressed by &#x60;translatable_type&#x60; (&#x60;locale&#x60;, &#x60;job&#x60;, &#x60;translation_key&#x60;, or &#x60;upload&#x60;) and &#x60;translatable_id&#x60; (its ID).  Enqueues machine translation using the project&#x27;s configured MT engine. `,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.PreTranslationCreateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			var preTranslationCreateParameters api.PreTranslationCreateParameters
			if err := json.Unmarshal([]byte(params.GetString("data")), &preTranslationCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", preTranslationCreateParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.PreTranslationsApi.PreTranslationCreate(auth, projectId, preTranslationCreateParameters, &localVarOptionals)

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

	PreTranslationsApiCmd.AddCommand(PreTranslationCreate)
	AddFlag(PreTranslationCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(PreTranslationCreate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(PreTranslationCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(PreTranslationCreate.Flags())
}
func initPreTranslationShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("pre_translation/show", "/")[1:], "_")
	var PreTranslationShow = &cobra.Command{
		Use:   use,
		Short: "Get a single pre-translation job",
		Long:  `Returns a single pre-translation job identified by its ID. `,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.PreTranslationShowOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			id := params.GetString(helpers.ToSnakeCase("Id"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("ifModifiedSince")) {
				localVarOptionals.IfModifiedSince = optional.NewString(params.GetString(helpers.ToSnakeCase("IfModifiedSince")))
			}

			if params.IsSet(helpers.ToSnakeCase("ifNoneMatch")) {
				localVarOptionals.IfNoneMatch = optional.NewString(params.GetString(helpers.ToSnakeCase("IfNoneMatch")))
			}

			data, api_response, err := client.PreTranslationsApi.PreTranslationShow(auth, projectId, id, &localVarOptionals)

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

	PreTranslationsApiCmd.AddCommand(PreTranslationShow)
	AddFlag(PreTranslationShow, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(PreTranslationShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(PreTranslationShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(PreTranslationShow, "string", helpers.ToSnakeCase("IfModifiedSince"), "", "Last modified condition, see [Conditional GET requests / HTTP Caching](/en/api/strings/pagination#conditional-get-requests-%2F-http-caching) (optional)", false)
	AddFlag(PreTranslationShow, "string", helpers.ToSnakeCase("IfNoneMatch"), "", "ETag condition, see [Conditional GET requests / HTTP Caching](/en/api/strings/pagination#conditional-get-requests-%2F-http-caching) (optional)", false)

	params.BindPFlags(PreTranslationShow.Flags())
}
func initPreTranslationsList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("pre_translations/list", "/")[1:], "_")
	var PreTranslationsList = &cobra.Command{
		Use:   use,
		Short: "List pre-translation jobs",
		Long:  `Returns all pre-translation jobs scoped to a project, ordered by creation date descending. `,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.PreTranslationsListOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("ifModifiedSince")) {
				localVarOptionals.IfModifiedSince = optional.NewString(params.GetString(helpers.ToSnakeCase("IfModifiedSince")))
			}

			if params.IsSet(helpers.ToSnakeCase("ifNoneMatch")) {
				localVarOptionals.IfNoneMatch = optional.NewString(params.GetString(helpers.ToSnakeCase("IfNoneMatch")))
			}

			if params.IsSet(helpers.ToSnakeCase("page")) {
				localVarOptionals.Page = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("Page")))
			}

			if params.IsSet(helpers.ToSnakeCase("perPage")) {
				localVarOptionals.PerPage = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("PerPage")))
			}

			data, api_response, err := client.PreTranslationsApi.PreTranslationsList(auth, projectId, &localVarOptionals)

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

	PreTranslationsApiCmd.AddCommand(PreTranslationsList)
	AddFlag(PreTranslationsList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(PreTranslationsList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(PreTranslationsList, "string", helpers.ToSnakeCase("IfModifiedSince"), "", "Last modified condition, see [Conditional GET requests / HTTP Caching](/en/api/strings/pagination#conditional-get-requests-%2F-http-caching) (optional)", false)
	AddFlag(PreTranslationsList, "string", helpers.ToSnakeCase("IfNoneMatch"), "", "ETag condition, see [Conditional GET requests / HTTP Caching](/en/api/strings/pagination#conditional-get-requests-%2F-http-caching) (optional)", false)
	AddFlag(PreTranslationsList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(PreTranslationsList, "int32", helpers.ToSnakeCase("PerPage"), "", "Limit on the number of objects to be returned, between 1 and 100. 25 by default", false)

	params.BindPFlags(PreTranslationsList.Flags())
}
