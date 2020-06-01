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
	initScreenshotCreate()
	initScreenshotDelete()
	initScreenshotShow()
	initScreenshotUpdate()
	initScreenshotsList()

	rootCmd.AddCommand(ScreenshotsApiCmd)
}

var ScreenshotsApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("Screenshots"),
	Short: "Screenshots API",
}

func initScreenshotCreate() {
	params := viper.New()
	var ScreenshotCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ScreenshotCreate", strings.TrimSuffix("ScreenshotsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ScreenshotsApi", "Api"), "s"))),
		Short: "Create a screenshot",
		Long:  `Create a new screenshot.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ScreenshotCreateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			screenshotCreateParameters := api.ScreenshotCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &screenshotCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", screenshotCreateParameters)
			}
			data, api_response, err := client.ScreenshotsApi.ScreenshotCreate(auth, projectId, screenshotCreateParameters, &localVarOptionals)

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

	ScreenshotsApiCmd.AddCommand(ScreenshotCreate)

	AddFlag(ScreenshotCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(ScreenshotCreate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(ScreenshotCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(ScreenshotCreate.Flags())
}
func initScreenshotDelete() {
	params := viper.New()
	var ScreenshotDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ScreenshotDelete", strings.TrimSuffix("ScreenshotsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ScreenshotsApi", "Api"), "s"))),
		Short: "Delete a screenshot",
		Long:  `Delete an existing screenshot.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ScreenshotDeleteOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.ScreenshotsApi.ScreenshotDelete(auth, projectId, id, &localVarOptionals)

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

	ScreenshotsApiCmd.AddCommand(ScreenshotDelete)

	AddFlag(ScreenshotDelete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(ScreenshotDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(ScreenshotDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(ScreenshotDelete.Flags())
}
func initScreenshotShow() {
	params := viper.New()
	var ScreenshotShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ScreenshotShow", strings.TrimSuffix("ScreenshotsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ScreenshotsApi", "Api"), "s"))),
		Short: "Get a single screenshot",
		Long:  `Get details on a single screenshot for a given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ScreenshotShowOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.ScreenshotsApi.ScreenshotShow(auth, projectId, id, &localVarOptionals)

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

	ScreenshotsApiCmd.AddCommand(ScreenshotShow)

	AddFlag(ScreenshotShow, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(ScreenshotShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(ScreenshotShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(ScreenshotShow.Flags())
}
func initScreenshotUpdate() {
	params := viper.New()
	var ScreenshotUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ScreenshotUpdate", strings.TrimSuffix("ScreenshotsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ScreenshotsApi", "Api"), "s"))),
		Short: "Update a screenshot",
		Long:  `Update an existing screenshot.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ScreenshotUpdateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			screenshotUpdateParameters := api.ScreenshotUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &screenshotUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", screenshotUpdateParameters)
			}
			data, api_response, err := client.ScreenshotsApi.ScreenshotUpdate(auth, projectId, id, screenshotUpdateParameters, &localVarOptionals)

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

	ScreenshotsApiCmd.AddCommand(ScreenshotUpdate)

	AddFlag(ScreenshotUpdate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(ScreenshotUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(ScreenshotUpdate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(ScreenshotUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(ScreenshotUpdate.Flags())
}
func initScreenshotsList() {
	params := viper.New()
	var ScreenshotsList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("ScreenshotsList", strings.TrimSuffix("ScreenshotsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("ScreenshotsApi", "Api"), "s"))),
		Short: "List screenshots",
		Long:  `List all screenshots for the given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.ScreenshotsListOpts{}

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

			data, api_response, err := client.ScreenshotsApi.ScreenshotsList(auth, projectId, &localVarOptionals)

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

	ScreenshotsApiCmd.AddCommand(ScreenshotsList)

	AddFlag(ScreenshotsList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(ScreenshotsList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(ScreenshotsList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(ScreenshotsList, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 10 by default", false)
	params.BindPFlags(ScreenshotsList.Flags())
}
