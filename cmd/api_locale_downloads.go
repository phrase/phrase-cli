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
	initLocaleDownloadCreate()
	initLocaleDownloadShow()

	rootCmd.AddCommand(LocaleDownloadsApiCmd)
}

var LocaleDownloadsApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("LocaleDownloads"),
	Short: "LocaleDownloads API",
}

func initLocaleDownloadCreate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("locale_download/create", "/")[1:], "_")
	var LocaleDownloadCreate = &cobra.Command{
		Use:   use,
		Short: "Initiate async download of a locale",
		Long:  `Prepare a locale for download in a specific file format.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.LocaleDownloadCreateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			localeId := params.GetString(helpers.ToSnakeCase("LocaleId"))

			var localeDownloadCreateParameters api.LocaleDownloadCreateParameters
			if err := json.Unmarshal([]byte(params.GetString("data")), &localeDownloadCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", localeDownloadCreateParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("ifModifiedSince")) {
				localVarOptionals.IfModifiedSince = optional.NewString(params.GetString(helpers.ToSnakeCase("IfModifiedSince")))
			}

			if params.IsSet(helpers.ToSnakeCase("ifNoneMatch")) {
				localVarOptionals.IfNoneMatch = optional.NewString(params.GetString(helpers.ToSnakeCase("IfNoneMatch")))
			}

			data, api_response, err := client.LocaleDownloadsApi.LocaleDownloadCreate(auth, projectId, localeId, localeDownloadCreateParameters, &localVarOptionals)

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

	LocaleDownloadsApiCmd.AddCommand(LocaleDownloadCreate)
	AddFlag(LocaleDownloadCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(LocaleDownloadCreate, "string", helpers.ToSnakeCase("LocaleId"), "", "Locale ID", true)
	AddFlag(LocaleDownloadCreate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(LocaleDownloadCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(LocaleDownloadCreate, "string", helpers.ToSnakeCase("IfModifiedSince"), "", "Last modified condition, see [Conditional GET requests / HTTP Caching](/en/api/strings/pagination#conditional-get-requests-%2F-http-caching) (optional)", false)
	AddFlag(LocaleDownloadCreate, "string", helpers.ToSnakeCase("IfNoneMatch"), "", "ETag condition, see [Conditional GET requests / HTTP Caching](/en/api/strings/pagination#conditional-get-requests-%2F-http-caching) (optional)", false)

	params.BindPFlags(LocaleDownloadCreate.Flags())
}
func initLocaleDownloadShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("locale_download/show", "/")[1:], "_")
	var LocaleDownloadShow = &cobra.Command{
		Use:   use,
		Short: "Show status of an async locale download",
		Long:  `Show status of already started async locale download. If the download is finished, the download link will be returned.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.LocaleDownloadShowOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			localeId := params.GetString(helpers.ToSnakeCase("LocaleId"))

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

			data, api_response, err := client.LocaleDownloadsApi.LocaleDownloadShow(auth, projectId, localeId, id, &localVarOptionals)

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

	LocaleDownloadsApiCmd.AddCommand(LocaleDownloadShow)
	AddFlag(LocaleDownloadShow, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(LocaleDownloadShow, "string", helpers.ToSnakeCase("LocaleId"), "", "Locale ID", true)
	AddFlag(LocaleDownloadShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(LocaleDownloadShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(LocaleDownloadShow, "string", helpers.ToSnakeCase("IfModifiedSince"), "", "Last modified condition, see [Conditional GET requests / HTTP Caching](/en/api/strings/pagination#conditional-get-requests-%2F-http-caching) (optional)", false)
	AddFlag(LocaleDownloadShow, "string", helpers.ToSnakeCase("IfNoneMatch"), "", "ETag condition, see [Conditional GET requests / HTTP Caching](/en/api/strings/pagination#conditional-get-requests-%2F-http-caching) (optional)", false)

	params.BindPFlags(LocaleDownloadShow.Flags())
}
