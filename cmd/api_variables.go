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
	initVariableCreate()
	initVariableDelete()
	initVariableShow()
	initVariableUpdate()

	rootCmd.AddCommand(VariablesApiCmd)
}

var VariablesApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("Variables"),
	Short: "Variables API",
}

func initVariableCreate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("variable/create", "/")[1:], "_")
	var VariableCreate = &cobra.Command{
		Use:   use,
		Short: "Create a variable",
		Long:  `Create a new variable.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.VariableCreateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			variableCreateParameters := api.VariableCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &variableCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", variableCreateParameters)
			}
			data, api_response, err := client.VariablesApi.VariableCreate(auth, projectId, variableCreateParameters, &localVarOptionals)

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

	VariablesApiCmd.AddCommand(VariableCreate)
	AddFlag(VariableCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(VariableCreate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(VariableCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(VariableCreate.Flags())
}
func initVariableDelete() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("variable/delete", "/")[1:], "_")
	var VariableDelete = &cobra.Command{
		Use:   use,
		Short: "Delete a variable",
		Long:  `Delete an existing variable.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.VariableDeleteOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			name := params.GetString(helpers.ToSnakeCase("Name"))

			data, api_response, err := client.VariablesApi.VariableDelete(auth, projectId, name, &localVarOptionals)

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

	VariablesApiCmd.AddCommand(VariableDelete)
	AddFlag(VariableDelete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(VariableDelete, "string", helpers.ToSnakeCase("Name"), "", "name", true)
	AddFlag(VariableDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(VariableDelete.Flags())
}
func initVariableShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("variable/show", "/")[1:], "_")
	var VariableShow = &cobra.Command{
		Use:   use,
		Short: "Get a single variable",
		Long:  `Get details on a single variable for a given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.VariableShowOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			name := params.GetString(helpers.ToSnakeCase("Name"))

			data, api_response, err := client.VariablesApi.VariableShow(auth, projectId, name, &localVarOptionals)

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

	VariablesApiCmd.AddCommand(VariableShow)
	AddFlag(VariableShow, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(VariableShow, "string", helpers.ToSnakeCase("Name"), "", "name", true)
	AddFlag(VariableShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(VariableShow.Flags())
}
func initVariableUpdate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("variable/update", "/")[1:], "_")
	var VariableUpdate = &cobra.Command{
		Use:   use,
		Short: "Update a variable",
		Long:  `Update an existing variable.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			client := api.NewAPIClient(cfg)
			localVarOptionals := api.VariableUpdateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			name := params.GetString(helpers.ToSnakeCase("Name"))

			branchUpdateParameters := api.BranchUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &branchUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", branchUpdateParameters)
			}
			data, api_response, err := client.VariablesApi.VariableUpdate(auth, projectId, name, branchUpdateParameters, &localVarOptionals)

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

	VariablesApiCmd.AddCommand(VariableUpdate)
	AddFlag(VariableUpdate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(VariableUpdate, "string", helpers.ToSnakeCase("Name"), "", "name", true)
	AddFlag(VariableUpdate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(VariableUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(VariableUpdate.Flags())
}
