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
	initRepliesList()
	initReplyCreate()
	initReplyDelete()
	initReplyMarkAsRead()
	initReplyMarkAsUnread()
	initReplyShow()

	rootCmd.AddCommand(CommentRepliesApiCmd)
}

var CommentRepliesApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("CommentReplies"),
	Short: "CommentReplies API",
}

func initRepliesList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("replies/list", "/")[1:], "_")
	var RepliesList = &cobra.Command{
		Use:   use,
		Short: "List replies",
		Long:  `List all replies for a comment.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.RepliesListOpts{}

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

			if params.IsSet(helpers.ToSnakeCase("query")) {
				localVarOptionals.Query = optional.NewString(params.GetString(helpers.ToSnakeCase("Query")))
			}

			if params.IsSet(helpers.ToSnakeCase("filters")) {

				var filters []string

				if err := json.Unmarshal([]byte(params.GetString(helpers.ToSnakeCase("filters"))), &filters); err != nil {
					HandleError(err)
				}
				localVarOptionals.Filters = filters
			}

			if params.IsSet(helpers.ToSnakeCase("order")) {
				localVarOptionals.Order = optional.NewString(params.GetString(helpers.ToSnakeCase("Order")))
			}

			data, api_response, err := client.CommentRepliesApi.RepliesList(auth, projectId, keyId, commentId, &localVarOptionals)

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

	CommentRepliesApiCmd.AddCommand(RepliesList)
	AddFlag(RepliesList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(RepliesList, "string", helpers.ToSnakeCase("KeyId"), "", "Translation Key ID", true)
	AddFlag(RepliesList, "string", helpers.ToSnakeCase("CommentId"), "", "Comment ID", true)
	AddFlag(RepliesList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(RepliesList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(RepliesList, "int32", helpers.ToSnakeCase("PerPage"), "", "Limit on the number of objects to be returned, between 1 and 100. 25 by default", false)
	AddFlag(RepliesList, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)
	AddFlag(RepliesList, "string", helpers.ToSnakeCase("Query"), "", "Search query for comment messages", false)
	AddFlag(RepliesList, "string", helpers.ToSnakeCase("Filters"), "", "payload in JSON format", false)
	AddFlag(RepliesList, "string", helpers.ToSnakeCase("Order"), "", "Order direction. Can be one of: asc, desc.", false)

	params.BindPFlags(RepliesList.Flags())
}
func initReplyCreate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("reply/create", "/")[1:], "_")
	var ReplyCreate = &cobra.Command{
		Use:   use,
		Short: "Create a reply",
		Long:  `Create a new reply for a comment.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.ReplyCreateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			keyId := params.GetString(helpers.ToSnakeCase("KeyId"))

			commentId := params.GetString(helpers.ToSnakeCase("CommentId"))

			var commentCreateParameters1 api.CommentCreateParameters1
			if err := json.Unmarshal([]byte(params.GetString("data")), &commentCreateParameters1); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", commentCreateParameters1)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.CommentRepliesApi.ReplyCreate(auth, projectId, keyId, commentId, commentCreateParameters1, &localVarOptionals)

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

	CommentRepliesApiCmd.AddCommand(ReplyCreate)
	AddFlag(ReplyCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(ReplyCreate, "string", helpers.ToSnakeCase("KeyId"), "", "Translation Key ID", true)
	AddFlag(ReplyCreate, "string", helpers.ToSnakeCase("CommentId"), "", "Comment ID", true)
	AddFlag(ReplyCreate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(ReplyCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(ReplyCreate.Flags())
}
func initReplyDelete() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("reply/delete", "/")[1:], "_")
	var ReplyDelete = &cobra.Command{
		Use:   use,
		Short: "Delete a reply",
		Long:  `Delete an existing reply.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.ReplyDeleteOpts{}

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

			data, api_response, err := client.CommentRepliesApi.ReplyDelete(auth, projectId, keyId, commentId, id, &localVarOptionals)

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

	CommentRepliesApiCmd.AddCommand(ReplyDelete)
	AddFlag(ReplyDelete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(ReplyDelete, "string", helpers.ToSnakeCase("KeyId"), "", "Translation Key ID", true)
	AddFlag(ReplyDelete, "string", helpers.ToSnakeCase("CommentId"), "", "Comment ID", true)
	AddFlag(ReplyDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(ReplyDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(ReplyDelete, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(ReplyDelete.Flags())
}
func initReplyMarkAsRead() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("reply/mark_as_read", "/")[1:], "_")
	var ReplyMarkAsRead = &cobra.Command{
		Use:   use,
		Short: "Mark a reply as read",
		Long:  `Mark a reply as read.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.ReplyMarkAsReadOpts{}

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

			data, api_response, err := client.CommentRepliesApi.ReplyMarkAsRead(auth, projectId, keyId, commentId, id, &localVarOptionals)

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

	CommentRepliesApiCmd.AddCommand(ReplyMarkAsRead)
	AddFlag(ReplyMarkAsRead, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(ReplyMarkAsRead, "string", helpers.ToSnakeCase("KeyId"), "", "Translation Key ID", true)
	AddFlag(ReplyMarkAsRead, "string", helpers.ToSnakeCase("CommentId"), "", "Comment ID", true)
	AddFlag(ReplyMarkAsRead, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(ReplyMarkAsRead, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(ReplyMarkAsRead, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(ReplyMarkAsRead.Flags())
}
func initReplyMarkAsUnread() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("reply/mark_as_unread", "/")[1:], "_")
	var ReplyMarkAsUnread = &cobra.Command{
		Use:   use,
		Short: "Mark a reply as unread",
		Long:  `Mark a reply as unread.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.ReplyMarkAsUnreadOpts{}

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

			data, api_response, err := client.CommentRepliesApi.ReplyMarkAsUnread(auth, projectId, keyId, commentId, id, &localVarOptionals)

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

	CommentRepliesApiCmd.AddCommand(ReplyMarkAsUnread)
	AddFlag(ReplyMarkAsUnread, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(ReplyMarkAsUnread, "string", helpers.ToSnakeCase("KeyId"), "", "Translation Key ID", true)
	AddFlag(ReplyMarkAsUnread, "string", helpers.ToSnakeCase("CommentId"), "", "Comment ID", true)
	AddFlag(ReplyMarkAsUnread, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(ReplyMarkAsUnread, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(ReplyMarkAsUnread, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(ReplyMarkAsUnread.Flags())
}
func initReplyShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("reply/show", "/")[1:], "_")
	var ReplyShow = &cobra.Command{
		Use:   use,
		Short: "Get a single reply",
		Long:  `Get details on a single reply.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.ReplyShowOpts{}

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

			data, api_response, err := client.CommentRepliesApi.ReplyShow(auth, projectId, keyId, commentId, id, &localVarOptionals)

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

	CommentRepliesApiCmd.AddCommand(ReplyShow)
	AddFlag(ReplyShow, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(ReplyShow, "string", helpers.ToSnakeCase("KeyId"), "", "Translation Key ID", true)
	AddFlag(ReplyShow, "string", helpers.ToSnakeCase("CommentId"), "", "Comment ID", true)
	AddFlag(ReplyShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(ReplyShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(ReplyShow, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(ReplyShow.Flags())
}
