package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/antihax/optional"
	helpers "github.com/phrase/phrase-cli/helpers"
	api "github.com/phrase/phrase-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initLocaleCreate()
	initLocaleDelete()
	initLocaleDownload()
	initLocaleShow()
	initLocaleUpdate()
	initLocalesList()

	rootCmd.AddCommand(LocalesApiCmd)
}

var LocalesApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("Locales"),
	Short: "Locales API",
}

func initLocaleCreate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("locale/create", "/")[1:], "_")
	var LocaleCreate = &cobra.Command{
		Use:   use,
		Short: "Create a locale",
		Long:  `Create a new locale.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.LocaleCreateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			localeCreateParameters := api.LocaleCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &localeCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", localeCreateParameters)
			}
			data, api_response, err := client.LocalesApi.LocaleCreate(auth, projectId, localeCreateParameters, &localVarOptionals)

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

	LocalesApiCmd.AddCommand(LocaleCreate)
	AddFlag(LocaleCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(LocaleCreate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(LocaleCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(LocaleCreate.Flags())
}
func initLocaleDelete() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("locale/delete", "/")[1:], "_")
	var LocaleDelete = &cobra.Command{
		Use:   use,
		Short: "Delete a locale",
		Long:  `Delete an existing locale.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.LocaleDeleteOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.LocalesApi.LocaleDelete(auth, projectId, id, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
				os.Stdout.Write(data)
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	LocalesApiCmd.AddCommand(LocaleDelete)
	AddFlag(LocaleDelete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(LocaleDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(LocaleDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(LocaleDelete, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(LocaleDelete.Flags())
}
func initLocaleDownload() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("locale/download", "/")[1:], "_")
	var LocaleDownload = &cobra.Command{
		Use:   use,
		Short: "Download a locale",
		Long:  `Download a locale in a specific file format.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.LocaleDownloadOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}
			if params.IsSet(helpers.ToSnakeCase("fileFormat")) {
				localVarOptionals.FileFormat = optional.NewString(params.GetString(helpers.ToSnakeCase("FileFormat")))
			}
			if params.IsSet(helpers.ToSnakeCase("tags")) {
				localVarOptionals.Tags = optional.NewString(params.GetString(helpers.ToSnakeCase("Tags")))
			}
			if params.IsSet(helpers.ToSnakeCase("tag")) {
				localVarOptionals.Tag = optional.NewString(params.GetString(helpers.ToSnakeCase("Tag")))
			}
			if params.IsSet(helpers.ToSnakeCase("includeEmptyTranslations")) {
				localVarOptionals.IncludeEmptyTranslations = optional.NewBool(params.GetBool(helpers.ToSnakeCase("IncludeEmptyTranslations")))
			}
			if params.IsSet(helpers.ToSnakeCase("includeTranslatedKeys")) {
				localVarOptionals.IncludeTranslatedKeys = optional.NewBool(params.GetBool(helpers.ToSnakeCase("IncludeTranslatedKeys")))
			}
			if params.IsSet(helpers.ToSnakeCase("keepNotranslateTags")) {
				localVarOptionals.KeepNotranslateTags = optional.NewBool(params.GetBool(helpers.ToSnakeCase("KeepNotranslateTags")))
			}
			if params.IsSet(helpers.ToSnakeCase("convertEmoji")) {
				localVarOptionals.ConvertEmoji = optional.NewBool(params.GetBool(helpers.ToSnakeCase("ConvertEmoji")))
			}
			if params.IsSet(helpers.ToSnakeCase("formatOptions")) {
				var formatOptions map[string]interface{}
				if err := json.Unmarshal([]byte(params.GetString(helpers.ToSnakeCase("FormatOptions"))), &formatOptions); err != nil {
					HandleError(err)
				}
				localVarOptionals.FormatOptions = optional.NewInterface(formatOptions)
			}
			if params.IsSet(helpers.ToSnakeCase("encoding")) {
				localVarOptionals.Encoding = optional.NewString(params.GetString(helpers.ToSnakeCase("Encoding")))
			}
			if params.IsSet(helpers.ToSnakeCase("skipUnverifiedTranslations")) {
				localVarOptionals.SkipUnverifiedTranslations = optional.NewBool(params.GetBool(helpers.ToSnakeCase("SkipUnverifiedTranslations")))
			}
			if params.IsSet(helpers.ToSnakeCase("includeUnverifiedTranslations")) {
				localVarOptionals.IncludeUnverifiedTranslations = optional.NewBool(params.GetBool(helpers.ToSnakeCase("IncludeUnverifiedTranslations")))
			}
			if params.IsSet(helpers.ToSnakeCase("useLastReviewedVersion")) {
				localVarOptionals.UseLastReviewedVersion = optional.NewBool(params.GetBool(helpers.ToSnakeCase("UseLastReviewedVersion")))
			}
			if params.IsSet(helpers.ToSnakeCase("fallbackLocaleId")) {
				localVarOptionals.FallbackLocaleId = optional.NewString(params.GetString(helpers.ToSnakeCase("FallbackLocaleId")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.LocalesApi.LocaleDownload(auth, projectId, id, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
				os.Stdout.Write(data)
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	LocalesApiCmd.AddCommand(LocaleDownload)
	AddFlag(LocaleDownload, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(LocaleDownload, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(LocaleDownload, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(LocaleDownload, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)
	AddFlag(LocaleDownload, "string", helpers.ToSnakeCase("FileFormat"), "", "File format name. See the format guide for all supported file formats.", false)
	AddFlag(LocaleDownload, "string", helpers.ToSnakeCase("Tags"), "", "Limit results to keys tagged with a list of comma separated tag names.", false)
	AddFlag(LocaleDownload, "string", helpers.ToSnakeCase("Tag"), "", "Limit download to tagged keys. This parameter is deprecated. Please use the \"tags\" parameter instead", false)
	AddFlag(LocaleDownload, "bool", helpers.ToSnakeCase("IncludeEmptyTranslations"), "", "Indicates whether keys without translations should be included in the output as well.", false)
	AddFlag(LocaleDownload, "bool", helpers.ToSnakeCase("IncludeTranslatedKeys"), "", "Include translated keys in the locale file. Use in combination with include_empty_translations to obtain only untranslated keys.", false)
	AddFlag(LocaleDownload, "bool", helpers.ToSnakeCase("KeepNotranslateTags"), "", "Indicates whether [NOTRANSLATE] tags should be kept.", false)
	AddFlag(LocaleDownload, "bool", helpers.ToSnakeCase("ConvertEmoji"), "", "This option is obsolete. Projects that were created on or after Nov 29th 2019 or that did not contain emoji by then will not require this flag any longer since emoji are now supported natively.", false)
	AddFlag(LocaleDownload, "string", helpers.ToSnakeCase("FormatOptions"), "", "payload in JSON format", false)
	AddFlag(LocaleDownload, "string", helpers.ToSnakeCase("Encoding"), "", "Enforces a specific encoding on the file contents. Valid options are \"UTF-8\", \"UTF-16\" and \"ISO-8859-1\".", false)
	AddFlag(LocaleDownload, "bool", helpers.ToSnakeCase("SkipUnverifiedTranslations"), "", "Indicates whether the locale file should skip all unverified translations. This parameter is deprecated and should be replaced with <code>include_unverified_translations</code>.", false)
	AddFlag(LocaleDownload, "bool", helpers.ToSnakeCase("IncludeUnverifiedTranslations"), "", "if set to false unverified translations are excluded", false)
	AddFlag(LocaleDownload, "bool", helpers.ToSnakeCase("UseLastReviewedVersion"), "", "If set to true the last reviewed version of a translation is used. This is only available if the review workflow (currently in beta) is enabled for the project.", false)
	AddFlag(LocaleDownload, "string", helpers.ToSnakeCase("FallbackLocaleId"), "", "If a key has no translation in the locale being downloaded the translation in the fallback locale will be used. Provide the public ID of the locale that should be used as the fallback. Requires include_empty_translations to be set to <code>true</code>.", false)

	params.BindPFlags(LocaleDownload.Flags())
}
func initLocaleShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("locale/show", "/")[1:], "_")
	var LocaleShow = &cobra.Command{
		Use:   use,
		Short: "Get a single locale",
		Long:  `Get details on a single locale for a given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.LocaleShowOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.LocalesApi.LocaleShow(auth, projectId, id, &localVarOptionals)

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

	LocalesApiCmd.AddCommand(LocaleShow)
	AddFlag(LocaleShow, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(LocaleShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(LocaleShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(LocaleShow, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(LocaleShow.Flags())
}
func initLocaleUpdate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("locale/update", "/")[1:], "_")
	var LocaleUpdate = &cobra.Command{
		Use:   use,
		Short: "Update a locale",
		Long:  `Update an existing locale.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.LocaleUpdateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			localeUpdateParameters := api.LocaleUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &localeUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", localeUpdateParameters)
			}
			data, api_response, err := client.LocalesApi.LocaleUpdate(auth, projectId, id, localeUpdateParameters, &localVarOptionals)

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

	LocalesApiCmd.AddCommand(LocaleUpdate)
	AddFlag(LocaleUpdate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(LocaleUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(LocaleUpdate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(LocaleUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(LocaleUpdate.Flags())
}
func initLocalesList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("locales/list", "/")[1:], "_")
	var LocalesList = &cobra.Command{
		Use:   use,
		Short: "List locales",
		Long:  `List all locales for the given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.LocalesListOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
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
			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			data, api_response, err := client.LocalesApi.LocalesList(auth, projectId, &localVarOptionals)

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

	LocalesApiCmd.AddCommand(LocalesList)
	AddFlag(LocalesList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(LocalesList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(LocalesList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(LocalesList, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 10 by default", false)
	AddFlag(LocalesList, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(LocalesList.Flags())
}
