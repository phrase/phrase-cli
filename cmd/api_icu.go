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
	initIcuSkeleton()

	rootCmd.AddCommand(ICUApiCmd)
}

var ICUApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("ICU"),
	Short: "ICU API",
}

func initIcuSkeleton() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("icu/skeleton", "/")[1:], "_")
	var IcuSkeleton = &cobra.Command{
		Use:   use,
		Short: "Build ICU skeletons",
		Long:  `Generates ICU (International Components for Unicode) message format skeletons for a given source string across one or more locales. An ICU skeleton strips the literal text from a pluralized or select message while preserving its structural rules — argument names, plural categories, select cases, and ordinal forms — adjusted to the pluralization rules of each requested locale.  Use this endpoint to normalize translation templates before importing them into locale files, or to validate that a source string carries the plural forms required by a target language.  Either &#x60;content&#x60; or &#x60;id&#x60; must be provided — supplying both or neither returns 400. When &#x60;id&#x60; is used and the referenced translation does not exist, the endpoint returns 404. When the source string is not valid ICU message format syntax, the endpoint returns 422 with an &#x60;error&#x60; field describing the parse failure. `,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.IcuSkeletonOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			var icuSkeletonParameters api.IcuSkeletonParameters
			if err := json.Unmarshal([]byte(params.GetString("data")), &icuSkeletonParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", icuSkeletonParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.ICUApi.IcuSkeleton(auth, icuSkeletonParameters, &localVarOptionals)

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

	ICUApiCmd.AddCommand(IcuSkeleton)
	AddFlag(IcuSkeleton, "string", "data", "d", "payload in JSON format", true)
	AddFlag(IcuSkeleton, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(IcuSkeleton.Flags())
}
