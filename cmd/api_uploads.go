package cmd

import (
	"context"
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
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.UploadCreateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}
			if params.IsSet(helpers.ToSnakeCase("file")) {
				file, err := os.Open(params.GetString(helpers.ToSnakeCase("file")))
				localVarOptionals.File = optional.NewInterface(file)

				if err != nil {
					HandleError(err)
				}
			}
			if params.IsSet(helpers.ToSnakeCase("fileFormat")) {
				localVarOptionals.FileFormat = optional.NewString(params.GetString(helpers.ToSnakeCase("FileFormat")))
			}
			if params.IsSet(helpers.ToSnakeCase("localeId")) {
				localVarOptionals.LocaleId = optional.NewString(params.GetString(helpers.ToSnakeCase("LocaleId")))
			}
			if params.IsSet(helpers.ToSnakeCase("tags")) {
				localVarOptionals.Tags = optional.NewString(params.GetString(helpers.ToSnakeCase("Tags")))
			}
			if params.IsSet(helpers.ToSnakeCase("updateTranslations")) {
				localVarOptionals.UpdateTranslations = optional.NewBool(params.GetBool(helpers.ToSnakeCase("UpdateTranslations")))
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
				file, err := os.Open(params.GetString(helpers.ToSnakeCase("file")))
				localVarOptionals.File = optional.NewInterface(file)

				if err != nil {
					HandleError(err)
				}
			}
			if params.IsSet(helpers.ToSnakeCase("autotranslate")) {
				localVarOptionals.Autotranslate = optional.NewBool(params.GetBool(helpers.ToSnakeCase("Autotranslate")))
			}
			if params.IsSet(helpers.ToSnakeCase("markReviewed")) {
				localVarOptionals.MarkReviewed = optional.NewBool(params.GetBool(helpers.ToSnakeCase("MarkReviewed")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			data, api_response, err := client.UploadsApi.UploadCreate(auth, projectId, &localVarOptionals)

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

	UploadsApiCmd.AddCommand(UploadCreate)
	AddFlag(UploadCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(UploadCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(UploadCreate, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)
	AddFlag(UploadCreate, "*os.File", helpers.ToSnakeCase("File"), "", "File to be imported", false)
	AddFlag(UploadCreate, "string", helpers.ToSnakeCase("FileFormat"), "", "File format. Auto-detected when possible and not specified.", false)
	AddFlag(UploadCreate, "string", helpers.ToSnakeCase("LocaleId"), "", "Locale of the file's content. Can be the name or public id of the locale. Preferred is the public id.", false)
	AddFlag(UploadCreate, "string", helpers.ToSnakeCase("Tags"), "", "List of tags separated by comma to be associated with the new keys contained in the upload.", false)
	AddFlag(UploadCreate, "bool", helpers.ToSnakeCase("UpdateTranslations"), "", "Indicates whether existing translations should be updated with the file content.", false)
	AddFlag(UploadCreate, "bool", helpers.ToSnakeCase("UpdateDescriptions"), "", "Existing key descriptions will be updated with the file content. Empty descriptions overwrite existing descriptions.", false)
	AddFlag(UploadCreate, "bool", helpers.ToSnakeCase("ConvertEmoji"), "", "This option is obsolete. Providing the option will cause a bad request error.", false)
	AddFlag(UploadCreate, "bool", helpers.ToSnakeCase("SkipUploadTags"), "", "Indicates whether the upload should not create upload tags.", false)
	AddFlag(UploadCreate, "bool", helpers.ToSnakeCase("SkipUnverification"), "", "Indicates whether the upload should unverify updated translations.", false)
	AddFlag(UploadCreate, "string", helpers.ToSnakeCase("FileEncoding"), "", "Enforces a specific encoding on the file contents. Valid options are \\\"UTF-8\\\", \\\"UTF-16\\\" and \\\"ISO-8859-1\\\".", false)
	AddFlag(UploadCreate, "bool", helpers.ToSnakeCase("Autotranslate"), "", "If set, translations for the uploaded language will be fetched automatically.", false)
	AddFlag(UploadCreate, "bool", helpers.ToSnakeCase("MarkReviewed"), "", "Indicated whether the imported translations should be marked as reviewed. This setting is available if the review workflow (currently beta) is enabled for the project.", false)

	params.BindPFlags(UploadCreate.Flags())
}
func initUploadShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("upload/show", "/")[1:], "_")
	var UploadShow = &cobra.Command{
		Use:   use,
		Short: "View upload details",
		Long:  `View details and summary for a single upload.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.UploadShowOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.UploadsApi.UploadShow(auth, projectId, id, &localVarOptionals)

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
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.UploadsListOpts{}

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

			data, api_response, err := client.UploadsApi.UploadsList(auth, projectId, &localVarOptionals)

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

	UploadsApiCmd.AddCommand(UploadsList)
	AddFlag(UploadsList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(UploadsList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(UploadsList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(UploadsList, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 10 by default", false)
	AddFlag(UploadsList, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(UploadsList.Flags())
}
