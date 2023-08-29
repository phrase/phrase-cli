package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/antihax/optional"
	helpers "github.com/phrase/phrase-cli/helpers"
	api "github.com/phrase/phrase-go/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initTranslationCreate()
	initTranslationExclude()
	initTranslationInclude()
	initTranslationReview()
	initTranslationShow()
	initTranslationUnverify()
	initTranslationUpdate()
	initTranslationVerify()
	initTranslationsByKey()
	initTranslationsByLocale()
	initTranslationsExcludeCollection()
	initTranslationsIncludeCollection()
	initTranslationsList()
	initTranslationsReviewCollection()
	initTranslationsSearch()
	initTranslationsUnverifyCollection()
	initTranslationsVerifyCollection()

	rootCmd.AddCommand(TranslationsApiCmd)
}

var TranslationsApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("Translations"),
	Short: "Translations API",
}

func initTranslationCreate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("translation/create", "/")[1:], "_")
	var TranslationCreate = &cobra.Command{
		Use:   use,
		Short: "Create a translation",
		Long:  `Create a translation.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.TranslationCreateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			translationCreateParameters := api.TranslationCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &translationCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", translationCreateParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.TranslationsApi.TranslationCreate(auth, projectId, translationCreateParameters, &localVarOptionals)

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

	TranslationsApiCmd.AddCommand(TranslationCreate)
	AddFlag(TranslationCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(TranslationCreate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(TranslationCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(TranslationCreate.Flags())
}
func initTranslationExclude() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("translation/exclude", "/")[1:], "_")
	var TranslationExclude = &cobra.Command{
		Use:   use,
		Short: "Exclude a translation from export",
		Long:  `Set exclude from export flag on an existing translation.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.TranslationExcludeOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			translationExcludeParameters := api.TranslationExcludeParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &translationExcludeParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", translationExcludeParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.TranslationsApi.TranslationExclude(auth, projectId, id, translationExcludeParameters, &localVarOptionals)

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

	TranslationsApiCmd.AddCommand(TranslationExclude)
	AddFlag(TranslationExclude, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(TranslationExclude, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(TranslationExclude, "string", "data", "d", "payload in JSON format", true)
	AddFlag(TranslationExclude, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(TranslationExclude.Flags())
}
func initTranslationInclude() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("translation/include", "/")[1:], "_")
	var TranslationInclude = &cobra.Command{
		Use:   use,
		Short: "Include a translation",
		Long:  `Remove exclude from export flag from an existing translation.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.TranslationIncludeOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			translationIncludeParameters := api.TranslationIncludeParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &translationIncludeParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", translationIncludeParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.TranslationsApi.TranslationInclude(auth, projectId, id, translationIncludeParameters, &localVarOptionals)

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

	TranslationsApiCmd.AddCommand(TranslationInclude)
	AddFlag(TranslationInclude, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(TranslationInclude, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(TranslationInclude, "string", "data", "d", "payload in JSON format", true)
	AddFlag(TranslationInclude, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(TranslationInclude.Flags())
}
func initTranslationReview() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("translation/review", "/")[1:], "_")
	var TranslationReview = &cobra.Command{
		Use:   use,
		Short: "Review a translation",
		Long:  `Mark an existing translation as reviewed.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.TranslationReviewOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			translationReviewParameters := api.TranslationReviewParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &translationReviewParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", translationReviewParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.TranslationsApi.TranslationReview(auth, projectId, id, translationReviewParameters, &localVarOptionals)

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

	TranslationsApiCmd.AddCommand(TranslationReview)
	AddFlag(TranslationReview, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(TranslationReview, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(TranslationReview, "string", "data", "d", "payload in JSON format", true)
	AddFlag(TranslationReview, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(TranslationReview.Flags())
}
func initTranslationShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("translation/show", "/")[1:], "_")
	var TranslationShow = &cobra.Command{
		Use:   use,
		Short: "Get a single translation",
		Long:  `Get details on a single translation.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.TranslationShowOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			data, api_response, err := client.TranslationsApi.TranslationShow(auth, projectId, id, &localVarOptionals)

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

	TranslationsApiCmd.AddCommand(TranslationShow)
	AddFlag(TranslationShow, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(TranslationShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(TranslationShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(TranslationShow, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(TranslationShow.Flags())
}
func initTranslationUnverify() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("translation/unverify", "/")[1:], "_")
	var TranslationUnverify = &cobra.Command{
		Use:   use,
		Short: "Mark a translation as unverified",
		Long:  `Mark an existing translation as unverified.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.TranslationUnverifyOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			translationUnverifyParameters := api.TranslationUnverifyParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &translationUnverifyParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", translationUnverifyParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.TranslationsApi.TranslationUnverify(auth, projectId, id, translationUnverifyParameters, &localVarOptionals)

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

	TranslationsApiCmd.AddCommand(TranslationUnverify)
	AddFlag(TranslationUnverify, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(TranslationUnverify, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(TranslationUnverify, "string", "data", "d", "payload in JSON format", true)
	AddFlag(TranslationUnverify, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(TranslationUnverify.Flags())
}
func initTranslationUpdate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("translation/update", "/")[1:], "_")
	var TranslationUpdate = &cobra.Command{
		Use:   use,
		Short: "Update a translation",
		Long:  `Update an existing translation.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.TranslationUpdateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			translationUpdateParameters := api.TranslationUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &translationUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", translationUpdateParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.TranslationsApi.TranslationUpdate(auth, projectId, id, translationUpdateParameters, &localVarOptionals)

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

	TranslationsApiCmd.AddCommand(TranslationUpdate)
	AddFlag(TranslationUpdate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(TranslationUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(TranslationUpdate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(TranslationUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(TranslationUpdate.Flags())
}
func initTranslationVerify() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("translation/verify", "/")[1:], "_")
	var TranslationVerify = &cobra.Command{
		Use:   use,
		Short: "Verify a translation",
		Long:  `Verify an existing translation.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.TranslationVerifyOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			translationVerifyParameters := api.TranslationVerifyParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &translationVerifyParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", translationVerifyParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.TranslationsApi.TranslationVerify(auth, projectId, id, translationVerifyParameters, &localVarOptionals)

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

	TranslationsApiCmd.AddCommand(TranslationVerify)
	AddFlag(TranslationVerify, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(TranslationVerify, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(TranslationVerify, "string", "data", "d", "payload in JSON format", true)
	AddFlag(TranslationVerify, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(TranslationVerify.Flags())
}
func initTranslationsByKey() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("translations/by_key", "/")[1:], "_")
	var TranslationsByKey = &cobra.Command{
		Use:   use,
		Short: "List translations by key",
		Long:  `List translations for a specific key.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.TranslationsByKeyOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			keyId := params.GetString(helpers.ToSnakeCase("KeyId"))
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("page")) {
				localVarOptionals.Page = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("Page")))
			}

			if params.IsSet(helpers.ToSnakeCase("perPage")) {
				localVarOptionals.PerPage = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("PerPage")))
			}

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			if params.IsSet(helpers.ToSnakeCase("sort")) {
				localVarOptionals.Sort = optional.NewString(params.GetString(helpers.ToSnakeCase("Sort")))
			}

			if params.IsSet(helpers.ToSnakeCase("order")) {
				localVarOptionals.Order = optional.NewString(params.GetString(helpers.ToSnakeCase("Order")))
			}

			if params.IsSet(helpers.ToSnakeCase("q")) {
				localVarOptionals.Q = optional.NewString(params.GetString(helpers.ToSnakeCase("Q")))
			}

			data, api_response, err := client.TranslationsApi.TranslationsByKey(auth, projectId, keyId, &localVarOptionals)

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

	TranslationsApiCmd.AddCommand(TranslationsByKey)
	AddFlag(TranslationsByKey, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(TranslationsByKey, "string", helpers.ToSnakeCase("KeyId"), "", "Translation Key ID", true)
	AddFlag(TranslationsByKey, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(TranslationsByKey, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(TranslationsByKey, "int32", helpers.ToSnakeCase("PerPage"), "", "Limit on the number of objects to be returned, between 1 and 100. 25 by default", false)
	AddFlag(TranslationsByKey, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)
	AddFlag(TranslationsByKey, "string", helpers.ToSnakeCase("Sort"), "", "Sort criteria. Can be one of: key_name, created_at, updated_at.", false)
	AddFlag(TranslationsByKey, "string", helpers.ToSnakeCase("Order"), "", "Order direction. Can be one of: asc, desc.", false)
	AddFlag(TranslationsByKey, "string", helpers.ToSnakeCase("Q"), "", "Specify a query to find translations by content (including wildcards).<br><br> The following qualifiers are supported in the query:<br> <ul>   <li><code>id:translation_id,...</code> for queries on a comma-separated list of ids</li>   <li><code>unverified:{true|false}</code> for verification status</li>   <li><code>tags:XYZ</code> for tags on the translation</li>   <li><code>excluded:{true|false}</code> for exclusion status</li>   <li><code>updated_at:{>=|<=}2013-02-21T00:00:00Z</code> for date range queries</li> </ul> Find more examples <a href=\"#overview--usage-examples\">here</a>. ", false)

	params.BindPFlags(TranslationsByKey.Flags())
}
func initTranslationsByLocale() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("translations/by_locale", "/")[1:], "_")
	var TranslationsByLocale = &cobra.Command{
		Use:   use,
		Short: "List translations by locale",
		Long:  `List translations for a specific locale. If you want to download all translations for one locale we recommend to use the &lt;code&gt;locales#download&lt;/code&gt; endpoint.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.TranslationsByLocaleOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			localeId := params.GetString(helpers.ToSnakeCase("LocaleId"))
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("page")) {
				localVarOptionals.Page = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("Page")))
			}

			if params.IsSet(helpers.ToSnakeCase("perPage")) {
				localVarOptionals.PerPage = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("PerPage")))
			}

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			if params.IsSet(helpers.ToSnakeCase("sort")) {
				localVarOptionals.Sort = optional.NewString(params.GetString(helpers.ToSnakeCase("Sort")))
			}

			if params.IsSet(helpers.ToSnakeCase("order")) {
				localVarOptionals.Order = optional.NewString(params.GetString(helpers.ToSnakeCase("Order")))
			}

			if params.IsSet(helpers.ToSnakeCase("q")) {
				localVarOptionals.Q = optional.NewString(params.GetString(helpers.ToSnakeCase("Q")))
			}

			data, api_response, err := client.TranslationsApi.TranslationsByLocale(auth, projectId, localeId, &localVarOptionals)

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

	TranslationsApiCmd.AddCommand(TranslationsByLocale)
	AddFlag(TranslationsByLocale, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(TranslationsByLocale, "string", helpers.ToSnakeCase("LocaleId"), "", "Locale ID", true)
	AddFlag(TranslationsByLocale, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(TranslationsByLocale, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(TranslationsByLocale, "int32", helpers.ToSnakeCase("PerPage"), "", "Limit on the number of objects to be returned, between 1 and 100. 25 by default", false)
	AddFlag(TranslationsByLocale, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)
	AddFlag(TranslationsByLocale, "string", helpers.ToSnakeCase("Sort"), "", "Sort criteria. Can be one of: key_name, created_at, updated_at.", false)
	AddFlag(TranslationsByLocale, "string", helpers.ToSnakeCase("Order"), "", "Order direction. Can be one of: asc, desc.", false)
	AddFlag(TranslationsByLocale, "string", helpers.ToSnakeCase("Q"), "", "Specify a query to find translations by content (including wildcards).<br><br> <i>Note: Search is limited to 10000 results and may not include recently updated data (depending on the project size).</i><br> The following qualifiers are supported in the query:<br> <ul>   <li><code>id:translation_id,...</code> for queries on a comma-separated list of ids</li>   <li><code>unverified:{true|false}</code> for verification status</li>   <li><code>tags:XYZ</code> for tags on the translation</li>   <li><code>excluded:{true|false}</code> for exclusion status</li>   <li><code>updated_at:{>=|<=}2013-02-21T00:00:00Z</code> for date range queries</li> </ul> Find more examples <a href=\"#overview--usage-examples\">here</a>. ", false)

	params.BindPFlags(TranslationsByLocale.Flags())
}
func initTranslationsExcludeCollection() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("translations/exclude-collection", "/")[1:], "_")
	var TranslationsExcludeCollection = &cobra.Command{
		Use:   use,
		Short: "Exclude translations by query",
		Long:  `Exclude translations matching query from locale export.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.TranslationsExcludeCollectionOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			translationsExcludeParameters := api.TranslationsExcludeParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &translationsExcludeParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", translationsExcludeParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.TranslationsApi.TranslationsExcludeCollection(auth, projectId, translationsExcludeParameters, &localVarOptionals)

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

	TranslationsApiCmd.AddCommand(TranslationsExcludeCollection)
	AddFlag(TranslationsExcludeCollection, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(TranslationsExcludeCollection, "string", "data", "d", "payload in JSON format", true)
	AddFlag(TranslationsExcludeCollection, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(TranslationsExcludeCollection.Flags())
}
func initTranslationsIncludeCollection() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("translations/include-collection", "/")[1:], "_")
	var TranslationsIncludeCollection = &cobra.Command{
		Use:   use,
		Short: "Include translations by query",
		Long:  `Include translations matching query in locale export.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.TranslationsIncludeCollectionOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			translationsIncludeParameters := api.TranslationsIncludeParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &translationsIncludeParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", translationsIncludeParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.TranslationsApi.TranslationsIncludeCollection(auth, projectId, translationsIncludeParameters, &localVarOptionals)

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

	TranslationsApiCmd.AddCommand(TranslationsIncludeCollection)
	AddFlag(TranslationsIncludeCollection, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(TranslationsIncludeCollection, "string", "data", "d", "payload in JSON format", true)
	AddFlag(TranslationsIncludeCollection, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(TranslationsIncludeCollection.Flags())
}
func initTranslationsList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("translations/list", "/")[1:], "_")
	var TranslationsList = &cobra.Command{
		Use:   use,
		Short: "List all translations",
		Long:  `List translations for the given project. If you want to download all translations for one locale we recommend to use the &lt;code&gt;locales#download&lt;/code&gt; endpoint.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.TranslationsListOpts{}

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

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			if params.IsSet(helpers.ToSnakeCase("sort")) {
				localVarOptionals.Sort = optional.NewString(params.GetString(helpers.ToSnakeCase("Sort")))
			}

			if params.IsSet(helpers.ToSnakeCase("order")) {
				localVarOptionals.Order = optional.NewString(params.GetString(helpers.ToSnakeCase("Order")))
			}

			if params.IsSet(helpers.ToSnakeCase("q")) {
				localVarOptionals.Q = optional.NewString(params.GetString(helpers.ToSnakeCase("Q")))
			}

			data, api_response, err := client.TranslationsApi.TranslationsList(auth, projectId, &localVarOptionals)

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

	TranslationsApiCmd.AddCommand(TranslationsList)
	AddFlag(TranslationsList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(TranslationsList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(TranslationsList, "string", helpers.ToSnakeCase("IfModifiedSince"), "", "Last modified condition, see <a href=\"#overview--conditional-get-requests--http-caching\">Conditional GET requests / HTTP Caching</a> (optional)", false)
	AddFlag(TranslationsList, "string", helpers.ToSnakeCase("IfNoneMatch"), "", "ETag condition, see <a href=\"#overview--conditional-get-requests--http-caching\">Conditional GET requests / HTTP Caching</a> (optional)", false)
	AddFlag(TranslationsList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(TranslationsList, "int32", helpers.ToSnakeCase("PerPage"), "", "Limit on the number of objects to be returned, between 1 and 100. 25 by default", false)
	AddFlag(TranslationsList, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)
	AddFlag(TranslationsList, "string", helpers.ToSnakeCase("Sort"), "", "Sort criteria. Can be one of: key_name, created_at, updated_at.", false)
	AddFlag(TranslationsList, "string", helpers.ToSnakeCase("Order"), "", "Order direction. Can be one of: asc, desc.", false)
	AddFlag(TranslationsList, "string", helpers.ToSnakeCase("Q"), "", "Specify a query to find translations by content (including wildcards).<br><br> <i>Note: Search is limited to 10000 results and may not include recently updated data (depending on the project size).</i><br> The following qualifiers are supported in the query:<br> <ul>   <li><code>id:translation_id,...</code> for queries on a comma-separated list of ids</li>   <li><code>tags:XYZ</code> for tags on the translation</li>   <li><code>unverified:{true|false}</code> for verification status</li>   <li><code>excluded:{true|false}</code> for exclusion status</li>   <li><code>updated_at:{>=|<=}2013-02-21T00:00:00Z</code> for date range queries</li>   <li><code>reviewed_after:2013-02-21T00:00:00Z</code> for fetching translations that were reviewed after the given timestamp</li> </ul> Find more examples <a href=\"#overview--usage-examples\">here</a>. ", false)

	params.BindPFlags(TranslationsList.Flags())
}
func initTranslationsReviewCollection() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("translations/review-collection", "/")[1:], "_")
	var TranslationsReviewCollection = &cobra.Command{
		Use:   use,
		Short: "Review translations selected by query",
		Long:  `Review translations matching query.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.TranslationsReviewCollectionOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			translationsReviewParameters := api.TranslationsReviewParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &translationsReviewParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", translationsReviewParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.TranslationsApi.TranslationsReviewCollection(auth, projectId, translationsReviewParameters, &localVarOptionals)

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

	TranslationsApiCmd.AddCommand(TranslationsReviewCollection)
	AddFlag(TranslationsReviewCollection, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(TranslationsReviewCollection, "string", "data", "d", "payload in JSON format", true)
	AddFlag(TranslationsReviewCollection, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(TranslationsReviewCollection.Flags())
}
func initTranslationsSearch() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("translations/search", "/")[1:], "_")
	var TranslationsSearch = &cobra.Command{
		Use:   use,
		Short: "Search translations",
		Long:  `Search translations for the given project. Provides the same search interface as &lt;code&gt;translations#index&lt;/code&gt; but allows POST requests to avoid limitations imposed by GET requests. If you want to download all translations for one locale we recommend to use the &lt;code&gt;locales#download&lt;/code&gt; endpoint.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.TranslationsSearchOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			translationsSearchParameters := api.TranslationsSearchParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &translationsSearchParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", translationsSearchParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("page")) {
				localVarOptionals.Page = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("Page")))
			}

			if params.IsSet(helpers.ToSnakeCase("perPage")) {
				localVarOptionals.PerPage = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("PerPage")))
			}

			data, api_response, err := client.TranslationsApi.TranslationsSearch(auth, projectId, translationsSearchParameters, &localVarOptionals)

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

	TranslationsApiCmd.AddCommand(TranslationsSearch)
	AddFlag(TranslationsSearch, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(TranslationsSearch, "string", "data", "d", "payload in JSON format", true)
	AddFlag(TranslationsSearch, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(TranslationsSearch, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(TranslationsSearch, "int32", helpers.ToSnakeCase("PerPage"), "", "Limit on the number of objects to be returned, between 1 and 100. 25 by default", false)

	params.BindPFlags(TranslationsSearch.Flags())
}
func initTranslationsUnverifyCollection() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("translations/unverify-collection", "/")[1:], "_")
	var TranslationsUnverifyCollection = &cobra.Command{
		Use:   use,
		Short: "Unverify translations by query",
		Long:  `Mark translations matching query as unverified.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.TranslationsUnverifyCollectionOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			translationsUnverifyParameters := api.TranslationsUnverifyParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &translationsUnverifyParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", translationsUnverifyParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.TranslationsApi.TranslationsUnverifyCollection(auth, projectId, translationsUnverifyParameters, &localVarOptionals)

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

	TranslationsApiCmd.AddCommand(TranslationsUnverifyCollection)
	AddFlag(TranslationsUnverifyCollection, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(TranslationsUnverifyCollection, "string", "data", "d", "payload in JSON format", true)
	AddFlag(TranslationsUnverifyCollection, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(TranslationsUnverifyCollection.Flags())
}
func initTranslationsVerifyCollection() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("translations/verify-collection", "/")[1:], "_")
	var TranslationsVerifyCollection = &cobra.Command{
		Use:   use,
		Short: "Verify translations by query",
		Long:  `Verify translations matching query.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.TranslationsVerifyCollectionOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			translationsVerifyParameters := api.TranslationsVerifyParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &translationsVerifyParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", translationsVerifyParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.TranslationsApi.TranslationsVerifyCollection(auth, projectId, translationsVerifyParameters, &localVarOptionals)

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

	TranslationsApiCmd.AddCommand(TranslationsVerifyCollection)
	AddFlag(TranslationsVerifyCollection, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(TranslationsVerifyCollection, "string", "data", "d", "payload in JSON format", true)
	AddFlag(TranslationsVerifyCollection, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(TranslationsVerifyCollection.Flags())
}
