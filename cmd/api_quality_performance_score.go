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
	initQualityPerformanceScoreList()

	rootCmd.AddCommand(QualityPerformanceScoreApiCmd)
}

var QualityPerformanceScoreApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("QualityPerformanceScore"),
	Short: "QualityPerformanceScore API",
}

func initQualityPerformanceScoreList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("quality_performance_score/list", "/")[1:], "_")
	var QualityPerformanceScoreList = &cobra.Command{
		Use:   use,
		Short: "Get Translation Quality",
		Long:  `Retrieves the quality scores for your Strings translations. Returns a score, measured by Phrase QPS`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.QualityPerformanceScoreListOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			var qualityPerformanceScoreListRequest api.QualityPerformanceScoreListRequest
			if err := json.Unmarshal([]byte(params.GetString("data")), &qualityPerformanceScoreListRequest); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", qualityPerformanceScoreListRequest)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.QualityPerformanceScoreApi.QualityPerformanceScoreList(auth, projectId, qualityPerformanceScoreListRequest, &localVarOptionals)

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

	QualityPerformanceScoreApiCmd.AddCommand(QualityPerformanceScoreList)
	AddFlag(QualityPerformanceScoreList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(QualityPerformanceScoreList, "string", "data", "d", "payload in JSON format", true)
	AddFlag(QualityPerformanceScoreList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(QualityPerformanceScoreList.Flags())
}
