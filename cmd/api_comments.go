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
	initCommentCreate()
	initCommentDelete()
	initCommentMarkCheck()
	initCommentMarkRead()
	initCommentMarkUnread()
	initCommentShow()
	initCommentUpdate()
	initCommentsList()

	rootCmd.AddCommand(CommentsApiCmd)
}

var CommentsApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("Comments"),
	Short: "Comments API",
}

func initCommentCreate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("comment/create", "/")[1:], "_")
	var CommentCreate = &cobra.Command{
		Use:   use,
		Short: "Create a comment",
		Long:  `Create a new comment for a key.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.CommentCreateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			keyId := params.GetString(helpers.ToSnakeCase("KeyId"))

			var commentCreateParameters api.CommentCreateParameters
			if err := json.Unmarshal([]byte(params.GetString("data")), &commentCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", commentCreateParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.CommentsApi.CommentCreate(auth, projectId, keyId, commentCreateParameters, &localVarOptionals)

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

	CommentsApiCmd.AddCommand(CommentCreate)
	AddFlag(CommentCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(CommentCreate, "string", helpers.ToSnakeCase("KeyId"), "", "Translation Key ID", true)
	AddFlag(CommentCreate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(CommentCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(CommentCreate.Flags())
}
func initCommentDelete() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("comment/delete", "/")[1:], "_")
	var CommentDelete = &cobra.Command{
		Use:   use,
		Short: "Delete a comment",
		Long:  `Delete an existing comment.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.CommentDeleteOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			keyId := params.GetString(helpers.ToSnakeCase("KeyId"))

			id := params.GetString(helpers.ToSnakeCase("Id"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			data, api_response, err := client.CommentsApi.CommentDelete(auth, projectId, keyId, id, &localVarOptionals)

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

	CommentsApiCmd.AddCommand(CommentDelete)
	AddFlag(CommentDelete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(CommentDelete, "string", helpers.ToSnakeCase("KeyId"), "", "Translation Key ID", true)
	AddFlag(CommentDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(CommentDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(CommentDelete, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(CommentDelete.Flags())
}
func initCommentMarkCheck() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("comment/mark/check", "/")[1:], "_")
	var CommentMarkCheck = &cobra.Command{
		Use:   use,
		Short: "Check if comment is read",
		Long:  `Check if comment was marked as read. Returns 204 if read, 404 if unread.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.CommentMarkCheckOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			keyId := params.GetString(helpers.ToSnakeCase("KeyId"))

			id := params.GetString(helpers.ToSnakeCase("Id"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			data, api_response, err := client.CommentsApi.CommentMarkCheck(auth, projectId, keyId, id, &localVarOptionals)

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

	CommentsApiCmd.AddCommand(CommentMarkCheck)
	AddFlag(CommentMarkCheck, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(CommentMarkCheck, "string", helpers.ToSnakeCase("KeyId"), "", "Translation Key ID", true)
	AddFlag(CommentMarkCheck, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(CommentMarkCheck, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(CommentMarkCheck, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(CommentMarkCheck.Flags())
}
func initCommentMarkRead() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("comment/mark/read", "/")[1:], "_")
	var CommentMarkRead = &cobra.Command{
		Use:   use,
		Short: "Mark a comment as read",
		Long:  `Mark a comment as read.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.CommentMarkReadOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			keyId := params.GetString(helpers.ToSnakeCase("KeyId"))

			id := params.GetString(helpers.ToSnakeCase("Id"))

			var commentMarkReadParameters api.CommentMarkReadParameters
			if err := json.Unmarshal([]byte(params.GetString("data")), &commentMarkReadParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", commentMarkReadParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.CommentsApi.CommentMarkRead(auth, projectId, keyId, id, commentMarkReadParameters, &localVarOptionals)

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

	CommentsApiCmd.AddCommand(CommentMarkRead)
	AddFlag(CommentMarkRead, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(CommentMarkRead, "string", helpers.ToSnakeCase("KeyId"), "", "Translation Key ID", true)
	AddFlag(CommentMarkRead, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(CommentMarkRead, "string", "data", "d", "payload in JSON format", true)
	AddFlag(CommentMarkRead, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(CommentMarkRead.Flags())
}
func initCommentMarkUnread() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("comment/mark/unread", "/")[1:], "_")
	var CommentMarkUnread = &cobra.Command{
		Use:   use,
		Short: "Mark a comment as unread",
		Long:  `Mark a comment as unread.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.CommentMarkUnreadOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			keyId := params.GetString(helpers.ToSnakeCase("KeyId"))

			id := params.GetString(helpers.ToSnakeCase("Id"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			data, api_response, err := client.CommentsApi.CommentMarkUnread(auth, projectId, keyId, id, &localVarOptionals)

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

	CommentsApiCmd.AddCommand(CommentMarkUnread)
	AddFlag(CommentMarkUnread, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(CommentMarkUnread, "string", helpers.ToSnakeCase("KeyId"), "", "Translation Key ID", true)
	AddFlag(CommentMarkUnread, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(CommentMarkUnread, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(CommentMarkUnread, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(CommentMarkUnread.Flags())
}
func initCommentShow() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("comment/show", "/")[1:], "_")
	var CommentShow = &cobra.Command{
		Use:   use,
		Short: "Get a single comment",
		Long:  `Get details on a single comment.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.CommentShowOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			keyId := params.GetString(helpers.ToSnakeCase("KeyId"))

			id := params.GetString(helpers.ToSnakeCase("Id"))

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			data, api_response, err := client.CommentsApi.CommentShow(auth, projectId, keyId, id, &localVarOptionals)

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

	CommentsApiCmd.AddCommand(CommentShow)
	AddFlag(CommentShow, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(CommentShow, "string", helpers.ToSnakeCase("KeyId"), "", "Translation Key ID", true)
	AddFlag(CommentShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(CommentShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(CommentShow, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)

	params.BindPFlags(CommentShow.Flags())
}
func initCommentUpdate() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("comment/update", "/")[1:], "_")
	var CommentUpdate = &cobra.Command{
		Use:   use,
		Short: "Update a comment",
		Long:  `Update an existing comment.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.CommentUpdateOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			keyId := params.GetString(helpers.ToSnakeCase("KeyId"))

			id := params.GetString(helpers.ToSnakeCase("Id"))

			var commentUpdateParameters api.CommentUpdateParameters
			if err := json.Unmarshal([]byte(params.GetString("data")), &commentUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", commentUpdateParameters)
			}
			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			data, api_response, err := client.CommentsApi.CommentUpdate(auth, projectId, keyId, id, commentUpdateParameters, &localVarOptionals)

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

	CommentsApiCmd.AddCommand(CommentUpdate)
	AddFlag(CommentUpdate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(CommentUpdate, "string", helpers.ToSnakeCase("KeyId"), "", "Translation Key ID", true)
	AddFlag(CommentUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(CommentUpdate, "string", "data", "d", "payload in JSON format", true)
	AddFlag(CommentUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)

	params.BindPFlags(CommentUpdate.Flags())
}
func initCommentsList() {
	params := viper.New()
	var use string
	// this weird approach is due to mustache template limitations
	use = strings.Join(strings.Split("comments/list", "/")[1:], "_")
	var CommentsList = &cobra.Command{
		Use:   use,
		Short: "List comments",
		Long:  `List all comments for a key.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := Auth()

			cfg := api.NewConfiguration()
			cfg.SetUserAgent(Config.UserAgent)
			if Config.Credentials.Host != "" {
				cfg.BasePath = Config.Credentials.Host
			}

			client := api.NewAPIClient(cfg)
			localVarOptionals := api.CommentsListOpts{}

			if Config.Credentials.TFA && Config.Credentials.TFAToken != "" {
				localVarOptionals.XPhraseAppOTP = optional.NewString(Config.Credentials.TFAToken)
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			keyId := params.GetString(helpers.ToSnakeCase("KeyId"))

			var commentsListParameters api.CommentsListParameters
			if err := json.Unmarshal([]byte(params.GetString("data")), &commentsListParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", commentsListParameters)
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

			if params.IsSet(helpers.ToSnakeCase("query")) {
				localVarOptionals.Query = optional.NewString(params.GetString(helpers.ToSnakeCase("Query")))
			}

			if params.IsSet(helpers.ToSnakeCase("localeIds")) {

				var localeIds []string

				if err := json.Unmarshal([]byte(params.GetString(helpers.ToSnakeCase("localeIds"))), &localeIds); err != nil {
					HandleError(err)
				}
				localVarOptionals.LocaleIds = localeIds
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

			data, api_response, err := client.CommentsApi.CommentsList(auth, projectId, keyId, commentsListParameters, &localVarOptionals)

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

	CommentsApiCmd.AddCommand(CommentsList)
	AddFlag(CommentsList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(CommentsList, "string", helpers.ToSnakeCase("KeyId"), "", "Translation Key ID", true)
	AddFlag(CommentsList, "string", "data", "d", "payload in JSON format", true)
	AddFlag(CommentsList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(CommentsList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(CommentsList, "int32", helpers.ToSnakeCase("PerPage"), "", "Limit on the number of objects to be returned, between 1 and 100. 25 by default", false)
	AddFlag(CommentsList, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)
	AddFlag(CommentsList, "string", helpers.ToSnakeCase("Query"), "", "Search query for comment messages", false)
	AddFlag(CommentsList, "string", helpers.ToSnakeCase("LocaleIds"), "", "payload in JSON format", false)
	AddFlag(CommentsList, "string", helpers.ToSnakeCase("Filters"), "", "payload in JSON format", false)
	AddFlag(CommentsList, "string", helpers.ToSnakeCase("Order"), "", "Order direction. Can be one of: asc, desc.", false)

	params.BindPFlags(CommentsList.Flags())
}
