package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/antihax/optional"
	helpers "github.com/phrase/phrase-cli/helpers"
	api "github.com/phrase/phrase-go/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initReactionCreate()
	initReactionDelete()
	initReactionShow()
	initReactionsList()

	rootCmd.AddCommand(CommentReactionsApiCmd)
}

var CommentReactionsApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("CommentReactions"),
	Short: "CommentReactions API",
}

func initReactionCreate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("reaction/create", "/")[1:], "_")
	var ReactionCreate = &cobra.Command{
		Use:   use,
		Short: "Create a reaction",
		Long:  `Create a new reaction for a comment.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.ReactionCreateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			keyId := params.GetString(helpers.ToSnakeCase("KeyId"))
			commentId := params.GetString(helpers.ToSnakeCase("CommentId"))
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			if params.IsSet(helpers.ToSnakeCase("emoji")) {
				localVarOptionals.Emoji = optional.NewString(params.GetString(helpers.ToSnakeCase("Emoji")))
			}

			data, api_response, err := client.CommentReactionsApi.ReactionCreate(auth, projectId, keyId, commentId, &localVarOptionals)

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

	CommentReactionsApiCmd.AddCommand(ReactionCreate)
	AddFlag(ReactionCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(ReactionCreate, "string", helpers.ToSnakeCase("KeyId"), "", "Translation Key ID", true)
	AddFlag(ReactionCreate, "string", helpers.ToSnakeCase("CommentId"), "", "Comment ID", true)
	AddFlag(ReactionCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(ReactionCreate, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)
	AddFlag(ReactionCreate, "string", helpers.ToSnakeCase("Emoji"), "", "specify the emoji for the reaction", false)

	params.BindPFlags(ReactionCreate.Flags())
}
func initReactionDelete() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("reaction/delete", "/")[1:], "_")
	var ReactionDelete = &cobra.Command{
		Use:   use,
		Short: "Delete a reaction",
		Long:  `Delete an existing reaction.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.ReactionDeleteOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			keyId := params.GetString(helpers.ToSnakeCase("KeyId"))
			commentId := params.GetString(helpers.ToSnakeCase("CommentId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			data, api_response, err := client.CommentReactionsApi.ReactionDelete(auth, projectId, keyId, commentId, id, &localVarOptionals)

			if err != nil {
				switch castedError := err.(type) {
				case api.GenericOpenAPIError:
					fmt.Printf("\n%s\n\n", string(castedError.Body()))
					HandleError(castedError)

				default:
					HandleError(castedError)
				}
			} else if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
				os.Stdout.Write(data)

				if Config.Debug {
					fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
				}
			}
		},
	}

	CommentReactionsApiCmd.AddCommand(ReactionDelete)
	AddFlag(ReactionDelete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(ReactionDelete, "string", helpers.ToSnakeCase("KeyId"), "", "Translation Key ID", true)
	AddFlag(ReactionDelete, "string", helpers.ToSnakeCase("CommentId"), "", "Comment ID", true)
	AddFlag(ReactionDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(ReactionDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(ReactionDelete, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(ReactionDelete.Flags())
}
func initReactionShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("reaction/show", "/")[1:], "_")
	var ReactionShow = &cobra.Command{
		Use:   use,
		Short: "Get a single reaction",
		Long:  `Get details on a single reaction.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.ReactionShowOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			keyId := params.GetString(helpers.ToSnakeCase("KeyId"))
			commentId := params.GetString(helpers.ToSnakeCase("CommentId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			data, api_response, err := client.CommentReactionsApi.ReactionShow(auth, projectId, keyId, commentId, id, &localVarOptionals)

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

	CommentReactionsApiCmd.AddCommand(ReactionShow)
	AddFlag(ReactionShow, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(ReactionShow, "string", helpers.ToSnakeCase("KeyId"), "", "Translation Key ID", true)
	AddFlag(ReactionShow, "string", helpers.ToSnakeCase("CommentId"), "", "Comment ID", true)
	AddFlag(ReactionShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(ReactionShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(ReactionShow, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(ReactionShow.Flags())
}
func initReactionsList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("reactions/list", "/")[1:], "_")
	var ReactionsList = &cobra.Command{
		Use:   use,
		Short: "List reactions",
		Long:  `List all reactions for a comment.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.ReactionsListOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			keyId := params.GetString(helpers.ToSnakeCase("KeyId"))
			commentId := params.GetString(helpers.ToSnakeCase("CommentId"))
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

			data, api_response, err := client.CommentReactionsApi.ReactionsList(auth, projectId, keyId, commentId, &localVarOptionals)

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

	CommentReactionsApiCmd.AddCommand(ReactionsList)
	AddFlag(ReactionsList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(ReactionsList, "string", helpers.ToSnakeCase("KeyId"), "", "Translation Key ID", true)
	AddFlag(ReactionsList, "string", helpers.ToSnakeCase("CommentId"), "", "Comment ID", true)
	AddFlag(ReactionsList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(ReactionsList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(ReactionsList, "int32", helpers.ToSnakeCase("PerPage"), "", "Limit on the number of objects to be returned, between 1 and 100. 25 by default", false)
	AddFlag(ReactionsList, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(ReactionsList.Flags())
}
