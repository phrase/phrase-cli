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
	initSpaceCreate()
	initSpaceDelete()
	initSpaceShow()
	initSpaceUpdate()
	initSpacesList()
	initSpacesProjectsCreate()
	initSpacesProjectsDelete()
	initSpacesProjectsList()

	rootCmd.AddCommand(SpacesApiCmd)
}

var SpacesApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("Spaces"),
	Short: "Spaces API",
}

func initSpaceCreate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("space/create", "/")[1:], "_")
	var SpaceCreate = &cobra.Command{
		Use:   use,
		Short: "Create a Space",
		Long:  `Create a new Space.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.SpaceCreateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			spaceCreateParameters := api.SpaceCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &spaceCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", spaceCreateParameters)
			}
			data, api_response, err := client.SpacesApi.SpaceCreate(auth, accountId, spaceCreateParameters, &localVarOptionals)

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

	SpacesApiCmd.AddCommand(SpaceCreate)
	AddFlag(SpaceCreate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(SpaceCreate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(SpaceCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(SpaceCreate.Flags())
}
func initSpaceDelete() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("space/delete", "/")[1:], "_")
	var SpaceDelete = &cobra.Command{
		Use:   use,
		Short: "Delete Space",
		Long:  `Delete the specified Space.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.SpaceDeleteOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.SpacesApi.SpaceDelete(auth, accountId, id, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
				os.Stdout.Write(data)
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

	SpacesApiCmd.AddCommand(SpaceDelete)
	AddFlag(SpaceDelete, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(SpaceDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(SpaceDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(SpaceDelete.Flags())
}
func initSpaceShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("space/show", "/")[1:], "_")
	var SpaceShow = &cobra.Command{
		Use:   use,
		Short: "Get Space",
		Long:  `Show the specified Space.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.SpaceShowOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.SpacesApi.SpaceShow(auth, accountId, id, &localVarOptionals)

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

	SpacesApiCmd.AddCommand(SpaceShow)
	AddFlag(SpaceShow, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(SpaceShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(SpaceShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(SpaceShow.Flags())
}
func initSpaceUpdate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("space/update", "/")[1:], "_")
	var SpaceUpdate = &cobra.Command{
		Use:   use,
		Short: "Update Space",
		Long:  `Update the specified Space.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.SpaceUpdateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			spaceUpdateParameters := api.SpaceUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &spaceUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", spaceUpdateParameters)
			}
			data, api_response, err := client.SpacesApi.SpaceUpdate(auth, accountId, id, spaceUpdateParameters, &localVarOptionals)

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

	SpacesApiCmd.AddCommand(SpaceUpdate)
	AddFlag(SpaceUpdate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(SpaceUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(SpaceUpdate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(SpaceUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(SpaceUpdate.Flags())
}
func initSpacesList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("spaces/list", "/")[1:], "_")
	var SpacesList = &cobra.Command{
		Use:   use,
		Short: "List Spaces",
		Long:  `List all Spaces for the given account.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.SpacesListOpts{}

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

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			data, api_response, err := client.SpacesApi.SpacesList(auth, accountId, &localVarOptionals)

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

	SpacesApiCmd.AddCommand(SpacesList)
	AddFlag(SpacesList, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(SpacesList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(SpacesList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(SpacesList, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 25 by default", false)

	params.BindPFlags(SpacesList.Flags())
}
func initSpacesProjectsCreate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("spaces/projects/create", "/")[1:], "_")
	var SpacesProjectsCreate = &cobra.Command{
		Use:   use,
		Short: "Add Project",
		Long:  `Adds an existing project to the space.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.SpacesProjectsCreateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			spaceId := params.GetString(helpers.ToSnakeCase("SpaceId"))

			spacesProjectsCreateParameters := api.SpacesProjectsCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &spacesProjectsCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", spacesProjectsCreateParameters)
			}
			data, api_response, err := client.SpacesApi.SpacesProjectsCreate(auth, accountId, spaceId, spacesProjectsCreateParameters, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
				os.Stdout.Write(data)
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

	SpacesApiCmd.AddCommand(SpacesProjectsCreate)
	AddFlag(SpacesProjectsCreate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(SpacesProjectsCreate, "string", helpers.ToSnakeCase("SpaceId"), "", "Space ID", true)
	AddFlag(SpacesProjectsCreate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(SpacesProjectsCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(SpacesProjectsCreate.Flags())
}
func initSpacesProjectsDelete() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("spaces/projects/delete", "/")[1:], "_")
	var SpacesProjectsDelete = &cobra.Command{
		Use:   use,
		Short: "Remove Project",
		Long:  `Removes a specified project from the specified space.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.SpacesProjectsDeleteOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			spaceId := params.GetString(helpers.ToSnakeCase("SpaceId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.SpacesApi.SpacesProjectsDelete(auth, accountId, spaceId, id, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
				os.Stdout.Write(data)
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

	SpacesApiCmd.AddCommand(SpacesProjectsDelete)
	AddFlag(SpacesProjectsDelete, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(SpacesProjectsDelete, "string", helpers.ToSnakeCase("SpaceId"), "", "Space ID", true)
	AddFlag(SpacesProjectsDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(SpacesProjectsDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(SpacesProjectsDelete.Flags())
}
func initSpacesProjectsList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("spaces/projects/list", "/")[1:], "_")
	var SpacesProjectsList = &cobra.Command{
		Use:   use,
		Short: "List Projects",
		Long:  `List all projects for the specified Space.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.SpacesProjectsListOpts{}

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

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			spaceId := params.GetString(helpers.ToSnakeCase("SpaceId"))

			data, api_response, err := client.SpacesApi.SpacesProjectsList(auth, accountId, spaceId, &localVarOptionals)

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

	SpacesApiCmd.AddCommand(SpacesProjectsList)
	AddFlag(SpacesProjectsList, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(SpacesProjectsList, "string", helpers.ToSnakeCase("SpaceId"), "", "Space ID", true)
	AddFlag(SpacesProjectsList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(SpacesProjectsList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(SpacesProjectsList, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 25 by default", false)

	params.BindPFlags(SpacesProjectsList.Flags())
}
