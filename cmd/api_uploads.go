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
	initUploadCreate()
	initUploadShow()
	initUploadsList()

	rootCmd.AddCommand(UploadsApiCmd)
}

var UploadsApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("Uploads"),
	Short: "Uploads API",
}

func initUploadCreate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("upload/create", "/")[1:], "_")
	var UploadCreate = &cobra.Command{
		Use:   use,
		Short: "Upload a new file",
		Long:  `Upload a new language file. Creates necessary resources in your project.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.UploadCreateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			file, err := os.Open(params.GetString(helpers.ToSnakeCase("file")))
			if err != nil {
				HandleError(err)
			}

			fileFormat := params.GetString(helpers.ToSnakeCase("FileFormat"))

			localeId := params.GetString(helpers.ToSnakeCase("LocaleId"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			if params.IsSet(helpers.ToSnakeCase("tags")) {
				localVarOptionals.Tags = optional.NewString(params.GetString(helpers.ToSnakeCase("Tags")))
			}

			if params.IsSet(helpers.ToSnakeCase("updateTranslations")) {
				localVarOptionals.UpdateTranslations = optional.NewBool(params.GetBool(helpers.ToSnakeCase("UpdateTranslations")))
			}

			if params.IsSet(helpers.ToSnakeCase("updateCustomMetadata")) {
				localVarOptionals.UpdateCustomMetadata = optional.NewBool(params.GetBool(helpers.ToSnakeCase("UpdateCustomMetadata")))
			}

			if params.IsSet(helpers.ToSnakeCase("updateTranslationKeys")) {
				localVarOptionals.UpdateTranslationKeys = optional.NewBool(params.GetBool(helpers.ToSnakeCase("UpdateTranslationKeys")))
			}

			if params.IsSet(helpers.ToSnakeCase("updateTranslationsOnSourceMatch")) {
				localVarOptionals.UpdateTranslationsOnSourceMatch = optional.NewBool(params.GetBool(helpers.ToSnakeCase("UpdateTranslationsOnSourceMatch")))
			}

			if params.IsSet(helpers.ToSnakeCase("sourceLocaleId")) {
				localVarOptionals.SourceLocaleId = optional.NewString(params.GetString(helpers.ToSnakeCase("SourceLocaleId")))
			}

			if params.IsSet(helpers.ToSnakeCase("updateDescriptions")) {
				localVarOptionals.UpdateDescriptions = optional.NewBool(params.GetBool(helpers.ToSnakeCase("UpdateDescriptions")))
			}

			if params.IsSet(helpers.ToSnakeCase("convertEmoji")) {
				localVarOptionals.ConvertEmoji = optional.NewBool(params.GetBool(helpers.ToSnakeCase("ConvertEmoji")))
			}

			if params.IsSet(helpers.ToSnakeCase("skipUploadTags")) {
				localVarOptionals.SkipUploadTags = optional.NewBool(params.GetBool(helpers.ToSnakeCase("SkipUploadTags")))
			}

			if params.IsSet(helpers.ToSnakeCase("skipUnverification")) {
				localVarOptionals.SkipUnverification = optional.NewBool(params.GetBool(helpers.ToSnakeCase("SkipUnverification")))
			}

			if params.IsSet(helpers.ToSnakeCase("fileEncoding")) {
				localVarOptionals.FileEncoding = optional.NewString(params.GetString(helpers.ToSnakeCase("FileEncoding")))
			}

			if params.IsSet(helpers.ToSnakeCase("localeMapping")) {

				var localeMapping map[string]interface{}

				if err := json.Unmarshal([]byte(params.GetString(helpers.ToSnakeCase("localeMapping"))), &localeMapping); err != nil {
					HandleError(err)
				}
				localVarOptionals.LocaleMapping = optional.NewInterface(localeMapping)
			}

			if params.IsSet(helpers.ToSnakeCase("formatOptions")) {

				var formatOptions map[string]interface{}

				if err := json.Unmarshal([]byte(params.GetString(helpers.ToSnakeCase("formatOptions"))), &formatOptions); err != nil {
					HandleError(err)
				}
				localVarOptionals.FormatOptions = optional.NewInterface(formatOptions)
			}

			if params.IsSet(helpers.ToSnakeCase("autotranslate")) {
				localVarOptionals.Autotranslate = optional.NewBool(params.GetBool(helpers.ToSnakeCase("Autotranslate")))
			}

			if params.IsSet(helpers.ToSnakeCase("verifyMentionedTranslations")) {
				localVarOptionals.VerifyMentionedTranslations = optional.NewBool(params.GetBool(helpers.ToSnakeCase("VerifyMentionedTranslations")))
			}

			if params.IsSet(helpers.ToSnakeCase("markReviewed")) {
				localVarOptionals.MarkReviewed = optional.NewBool(params.GetBool(helpers.ToSnakeCase("MarkReviewed")))
			}

			if params.IsSet(helpers.ToSnakeCase("tagOnlyAffectedKeys")) {
				localVarOptionals.TagOnlyAffectedKeys = optional.NewBool(params.GetBool(helpers.ToSnakeCase("TagOnlyAffectedKeys")))
			}

			if params.IsSet(helpers.ToSnakeCase("translationKeyPrefix")) {
				localVarOptionals.TranslationKeyPrefix = optional.NewString(params.GetString(helpers.ToSnakeCase("TranslationKeyPrefix")))
			}

			data, api_response, err := client.UploadsApi.UploadCreate(auth, projectId, file, fileFormat, localeId, &localVarOptionals)

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

	UploadsApiCmd.AddCommand(UploadCreate)
	AddFlag(UploadCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(UploadCreate, "*os.File", helpers.ToSnakeCase("File"), "", "File to be imported", true)
	AddFlag(UploadCreate, "string", helpers.ToSnakeCase("FileFormat"), "", "File format. Auto-detected when possible and not specified.", true)
	AddFlag(UploadCreate, "string", helpers.ToSnakeCase("LocaleId"), "", "Locale of the file's content. Can be the name or id of the locale. Preferred is id.", true)
	AddFlag(UploadCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(UploadCreate, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)
	AddFlag(UploadCreate, "string", helpers.ToSnakeCase("Tags"), "", "List of tags separated by comma to be associated with the new keys contained in the upload.", false)
	AddFlag(UploadCreate, "bool", helpers.ToSnakeCase("UpdateTranslations"), "", "Indicates whether existing translations should be updated with the file content.", false)
	AddFlag(UploadCreate, "bool", helpers.ToSnakeCase("UpdateCustomMetadata"), "", "Determines whether to update custom metadata values when uploading a file. If set to true, existing metadata can be changed or removed. Passing an empty value deletes the corresponding metadata property.", false)
	AddFlag(UploadCreate, "bool", helpers.ToSnakeCase("UpdateTranslationKeys"), "", "Pass `false` here to prevent new keys from being created and existing keys updated.", false)
	AddFlag(UploadCreate, "bool", helpers.ToSnakeCase("UpdateTranslationsOnSourceMatch"), "", "Update target translations only if the source translations of the uploaded multilingual file match the stored translations.", false)
	AddFlag(UploadCreate, "string", helpers.ToSnakeCase("SourceLocaleId"), "", "Specifies the source locale for multilingual files. Can be the name or id of the locale. Preferred is id.", false)
	AddFlag(UploadCreate, "bool", helpers.ToSnakeCase("UpdateDescriptions"), "", "Existing key descriptions will be updated with the file content. Empty descriptions overwrite existing descriptions.", false)
	AddFlag(UploadCreate, "bool", helpers.ToSnakeCase("ConvertEmoji"), "", "This option is obsolete. Providing the option will cause a bad request error.", false)
	AddFlag(UploadCreate, "bool", helpers.ToSnakeCase("SkipUploadTags"), "", "Indicates whether the upload should not create upload tags.", false)
	AddFlag(UploadCreate, "bool", helpers.ToSnakeCase("SkipUnverification"), "", "Indicates whether the upload should unverify updated translations.", false)
	AddFlag(UploadCreate, "string", helpers.ToSnakeCase("FileEncoding"), "", "Enforces a specific encoding on the file contents. Valid options are \\\"UTF-8\\\", \\\"UTF-16\\\" and \\\"ISO-8859-1\\\".", false)
	AddFlag(UploadCreate, "string", helpers.ToSnakeCase("LocaleMapping"), "", "payload in JSON format", false)
	AddFlag(UploadCreate, "string", helpers.ToSnakeCase("FormatOptions"), "", "payload in JSON format", false)
	AddFlag(UploadCreate, "bool", helpers.ToSnakeCase("Autotranslate"), "", "If set, translations for the uploaded language will be fetched automatically.", false)
	AddFlag(UploadCreate, "bool", helpers.ToSnakeCase("VerifyMentionedTranslations"), "", "Indicates whether all translations mentioned in the upload should be verified.", false)
	AddFlag(UploadCreate, "bool", helpers.ToSnakeCase("MarkReviewed"), "", "Indicated whether the imported translations should be marked as reviewed. This setting is available if the review workflow is enabled for the project.", false)
	AddFlag(UploadCreate, "bool", helpers.ToSnakeCase("TagOnlyAffectedKeys"), "", "Indicates whether only keys affected (created or updated) by the upload should be tagged. The default is `false`", false)
	AddFlag(UploadCreate, "string", helpers.ToSnakeCase("TranslationKeyPrefix"), "", "This prefix will be added to all uploaded translation key names to prevent collisions. Use a meaningful prefix related to your project or file to keep key names organized.", false)

	params.BindPFlags(UploadCreate.Flags())
}
func initUploadShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("upload/show", "/")[1:], "_")
	var UploadShow = &cobra.Command{
		Use:   use,
		Short: "Get a single upload",
		Long:  `View details and summary for a single upload.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.UploadShowOpts{}

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

			data, api_response, err := client.UploadsApi.UploadShow(auth, projectId, id, &localVarOptionals)

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

	UploadsApiCmd.AddCommand(UploadShow)
	AddFlag(UploadShow, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(UploadShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(UploadShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(UploadShow, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(UploadShow.Flags())
}
func initUploadsList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("uploads/list", "/")[1:], "_")
	var UploadsList = &cobra.Command{
		Use:   use,
		Short: "List uploads",
		Long:  `List all uploads for the given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.UploadsListOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

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

			data, api_response, err := client.UploadsApi.UploadsList(auth, projectId, &localVarOptionals)

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

	UploadsApiCmd.AddCommand(UploadsList)
	AddFlag(UploadsList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(UploadsList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(UploadsList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(UploadsList, "int32", helpers.ToSnakeCase("PerPage"), "", "Limit on the number of objects to be returned, between 1 and 100. 25 by default", false)
	AddFlag(UploadsList, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(UploadsList.Flags())
}
