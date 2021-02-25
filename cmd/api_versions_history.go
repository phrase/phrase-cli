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
	initVersionShow()
	initVersionsList()

	rootCmd.AddCommand(VersionsHistoryApiCmd)
}

var VersionsHistoryApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("VersionsHistory"),
	Short: "VersionsHistory API",
}

func initVersionShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("version/show", "/")[1:], "_")
	var VersionShow = &cobra.Command{
		Use:   use,
		Short: "Get a single version",
		Long:  `Get details on a single version.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.VersionShowOpts{}

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
			translationId := params.GetString(helpers.ToSnakeCase("TranslationId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.VersionsHistoryApi.VersionShow(auth, projectId, translationId, id, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}
				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				switch castedError := err.(type) {
				case api.GenericOpenAPIError:
					fmt.Printf("\n%s\n\n", string(castedError.Body()))
					HandleError(castedError)

				default:
					HandleError(castedError)
				}
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	VersionsHistoryApiCmd.AddCommand(VersionShow)
	AddFlag(VersionShow, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(VersionShow, "string", helpers.ToSnakeCase("TranslationId"), "", "Translation ID", true)
	AddFlag(VersionShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(VersionShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(VersionShow, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(VersionShow.Flags())
}
func initVersionsList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("versions/list", "/")[1:], "_")
	var VersionsList = &cobra.Command{
		Use:   use,
		Short: "List all versions",
		Long:  `List all versions for the given translation.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.VersionsListOpts{}

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
			translationId := params.GetString(helpers.ToSnakeCase("TranslationId"))

			data, api_response, err := client.VersionsHistoryApi.VersionsList(auth, projectId, translationId, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}
				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				switch castedError := err.(type) {
				case api.GenericOpenAPIError:
					fmt.Printf("\n%s\n\n", string(castedError.Body()))
					HandleError(castedError)

				default:
					HandleError(castedError)
				}
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	VersionsHistoryApiCmd.AddCommand(VersionsList)
	AddFlag(VersionsList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(VersionsList, "string", helpers.ToSnakeCase("TranslationId"), "", "Translation ID", true)
	AddFlag(VersionsList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(VersionsList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(VersionsList, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 25 by default", false)
	AddFlag(VersionsList, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(VersionsList.Flags())
}
