package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/antihax/optional"
	helpers "github.com/phrase/phrase-cli/helpers"
	api "github.com/phrase/phrase-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initScreenshotMarkerCreate()
	initScreenshotMarkerDelete()
	initScreenshotMarkerShow()
	initScreenshotMarkerUpdate()
	initScreenshotMarkersList()

	rootCmd.AddCommand(ScreenshotMarkersApiCmd)
}

var ScreenshotMarkersApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("ScreenshotMarkers"),
	Short: "ScreenshotMarkers API",
}

func initScreenshotMarkerCreate() {
	params := viper.New()
	var ScreenshotMarkerCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ScreenshotMarkerCreate", strings.TrimSuffix("ScreenshotMarkersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ScreenshotMarkersApi", "Api"), "s"))),
		Short: "Create a screenshot marker",
		Long:  `Create a new screenshot marker.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ScreenshotMarkerCreateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			screenshotId := params.GetString(helpers.ToSnakeCase("ScreenshotId"))

			screenshotMarkerCreateParameters := api.ScreenshotMarkerCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &screenshotMarkerCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", screenshotMarkerCreateParameters)
			}
			data, api_response, err := client.ScreenshotMarkersApi.ScreenshotMarkerCreate(auth, projectId, screenshotId, screenshotMarkerCreateParameters, &localVarOptionals)

			if api_response.StatusCode == 200 {
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

	ScreenshotMarkersApiCmd.AddCommand(ScreenshotMarkerCreate)

	AddFlag(ScreenshotMarkerCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(ScreenshotMarkerCreate, "string", helpers.ToSnakeCase("ScreenshotId"), "", "Screenshot ID", true)
	AddFlag(ScreenshotMarkerCreate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(ScreenshotMarkerCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(ScreenshotMarkerCreate.Flags())
}
func initScreenshotMarkerDelete() {
	params := viper.New()
	var ScreenshotMarkerDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ScreenshotMarkerDelete", strings.TrimSuffix("ScreenshotMarkersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ScreenshotMarkersApi", "Api"), "s"))),
		Short: "Delete a screenshot marker",
		Long:  `Delete an existing screenshot marker.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ScreenshotMarkerDeleteOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			screenshotId := params.GetString(helpers.ToSnakeCase("ScreenshotId"))

			data, api_response, err := client.ScreenshotMarkersApi.ScreenshotMarkerDelete(auth, projectId, screenshotId, &localVarOptionals)

			if api_response.StatusCode == 200 {
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

	ScreenshotMarkersApiCmd.AddCommand(ScreenshotMarkerDelete)

	AddFlag(ScreenshotMarkerDelete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(ScreenshotMarkerDelete, "string", helpers.ToSnakeCase("ScreenshotId"), "", "Screenshot ID", true)
	AddFlag(ScreenshotMarkerDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(ScreenshotMarkerDelete.Flags())
}
func initScreenshotMarkerShow() {
	params := viper.New()
	var ScreenshotMarkerShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ScreenshotMarkerShow", strings.TrimSuffix("ScreenshotMarkersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ScreenshotMarkersApi", "Api"), "s"))),
		Short: "Get a single screenshot marker",
		Long:  `Get details on a single screenshot marker for a given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ScreenshotMarkerShowOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			screenshotId := params.GetString(helpers.ToSnakeCase("ScreenshotId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.ScreenshotMarkersApi.ScreenshotMarkerShow(auth, projectId, screenshotId, id, &localVarOptionals)

			if api_response.StatusCode == 200 {
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

	ScreenshotMarkersApiCmd.AddCommand(ScreenshotMarkerShow)

	AddFlag(ScreenshotMarkerShow, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(ScreenshotMarkerShow, "string", helpers.ToSnakeCase("ScreenshotId"), "", "Screenshot ID", true)
	AddFlag(ScreenshotMarkerShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(ScreenshotMarkerShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(ScreenshotMarkerShow.Flags())
}
func initScreenshotMarkerUpdate() {
	params := viper.New()
	var ScreenshotMarkerUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ScreenshotMarkerUpdate", strings.TrimSuffix("ScreenshotMarkersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ScreenshotMarkersApi", "Api"), "s"))),
		Short: "Update a screenshot marker",
		Long:  `Update an existing screenshot marker.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ScreenshotMarkerUpdateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			screenshotId := params.GetString(helpers.ToSnakeCase("ScreenshotId"))

			screenshotMarkerUpdateParameters := api.ScreenshotMarkerUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &screenshotMarkerUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", screenshotMarkerUpdateParameters)
			}
			data, api_response, err := client.ScreenshotMarkersApi.ScreenshotMarkerUpdate(auth, projectId, screenshotId, screenshotMarkerUpdateParameters, &localVarOptionals)

			if api_response.StatusCode == 200 {
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

	ScreenshotMarkersApiCmd.AddCommand(ScreenshotMarkerUpdate)

	AddFlag(ScreenshotMarkerUpdate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(ScreenshotMarkerUpdate, "string", helpers.ToSnakeCase("ScreenshotId"), "", "Screenshot ID", true)
	AddFlag(ScreenshotMarkerUpdate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(ScreenshotMarkerUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(ScreenshotMarkerUpdate.Flags())
}
func initScreenshotMarkersList() {
	params := viper.New()
	var ScreenshotMarkersList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ScreenshotMarkersList", strings.TrimSuffix("ScreenshotMarkersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ScreenshotMarkersApi", "Api"), "s"))),
		Short: "List screenshot markers",
		Long:  `List all screenshot markers for the given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ScreenshotMarkersListOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("page")) {
				localVarOptionals.Page = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("Page")))
			}
			if params.IsSet(helpers.ToSnakeCase("perPage")) {
				localVarOptionals.PerPage = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("PerPage")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.ScreenshotMarkersApi.ScreenshotMarkersList(auth, projectId, id, &localVarOptionals)

			if api_response.StatusCode == 200 {
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

	ScreenshotMarkersApiCmd.AddCommand(ScreenshotMarkersList)

	AddFlag(ScreenshotMarkersList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(ScreenshotMarkersList, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(ScreenshotMarkersList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(ScreenshotMarkersList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(ScreenshotMarkersList, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 10 by default", false)
	params.BindPFlags(ScreenshotMarkersList.Flags())
}
