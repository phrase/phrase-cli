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
	var CommentCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("CommentCreate", strings.TrimSuffix("CommentsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("CommentsApi", "Api"), "s"))),
		Short: "Create a comment",
		Long:  `Create a new comment for a key.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.CommentCreateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			keyId := params.GetString(helpers.ToSnakeCase("KeyId"))

			commentCreateParameters := api.CommentCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &commentCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", commentCreateParameters)
			}
			data, api_response, err := client.CommentsApi.CommentCreate(auth, projectId, keyId, commentCreateParameters, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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

	CommentsApiCmd.AddCommand(CommentCreate)

	AddFlag(CommentCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(CommentCreate, "string", helpers.ToSnakeCase("KeyId"), "", "Translation Key ID", true)
	AddFlag(CommentCreate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(CommentCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(CommentCreate.Flags())
}
func initCommentDelete() {
	params := viper.New()
	var CommentDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("CommentDelete", strings.TrimSuffix("CommentsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("CommentsApi", "Api"), "s"))),
		Short: "Delete a comment",
		Long:  `Delete an existing comment.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.CommentDeleteOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			keyId := params.GetString(helpers.ToSnakeCase("KeyId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.CommentsApi.CommentDelete(auth, projectId, keyId, id, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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
	var CommentMarkCheck = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("CommentMarkCheck", strings.TrimSuffix("CommentsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("CommentsApi", "Api"), "s"))),
		Short: "Check if comment is read",
		Long:  `Check if comment was marked as read. Returns 204 if read, 404 if unread.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.CommentMarkCheckOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			keyId := params.GetString(helpers.ToSnakeCase("KeyId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.CommentsApi.CommentMarkCheck(auth, projectId, keyId, id, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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
	var CommentMarkRead = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("CommentMarkRead", strings.TrimSuffix("CommentsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("CommentsApi", "Api"), "s"))),
		Short: "Mark a comment as read",
		Long:  `Mark a comment as read.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.CommentMarkReadOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			keyId := params.GetString(helpers.ToSnakeCase("KeyId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			commentMarkReadParameters := api.CommentMarkReadParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &commentMarkReadParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", commentMarkReadParameters)
			}
			data, api_response, err := client.CommentsApi.CommentMarkRead(auth, projectId, keyId, id, commentMarkReadParameters, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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
	var CommentMarkUnread = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("CommentMarkUnread", strings.TrimSuffix("CommentsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("CommentsApi", "Api"), "s"))),
		Short: "Mark a comment as unread",
		Long:  `Mark a comment as unread.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.CommentMarkUnreadOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			keyId := params.GetString(helpers.ToSnakeCase("KeyId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.CommentsApi.CommentMarkUnread(auth, projectId, keyId, id, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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
	var CommentShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("CommentShow", strings.TrimSuffix("CommentsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("CommentsApi", "Api"), "s"))),
		Short: "Get a single comment",
		Long:  `Get details on a single comment.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.CommentShowOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			keyId := params.GetString(helpers.ToSnakeCase("KeyId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.CommentsApi.CommentShow(auth, projectId, keyId, id, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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
	var CommentUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("CommentUpdate", strings.TrimSuffix("CommentsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("CommentsApi", "Api"), "s"))),
		Short: "Update a comment",
		Long:  `Update an existing comment.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.CommentUpdateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			keyId := params.GetString(helpers.ToSnakeCase("KeyId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			commentUpdateParameters := api.CommentUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &commentUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", commentUpdateParameters)
			}
			data, api_response, err := client.CommentsApi.CommentUpdate(auth, projectId, keyId, id, commentUpdateParameters, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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
	var CommentsList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("CommentsList", strings.TrimSuffix("CommentsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("CommentsApi", "Api"), "s"))),
		Short: "List comments",
		Long:  `List all comments for a key.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.CommentsListOpts{}

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
			keyId := params.GetString(helpers.ToSnakeCase("KeyId"))

			data, api_response, err := client.CommentsApi.CommentsList(auth, projectId, keyId, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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

	CommentsApiCmd.AddCommand(CommentsList)

	AddFlag(CommentsList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(CommentsList, "string", helpers.ToSnakeCase("KeyId"), "", "Translation Key ID", true)
	AddFlag(CommentsList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(CommentsList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(CommentsList, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 10 by default", false)
	AddFlag(CommentsList, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)
	params.BindPFlags(CommentsList.Flags())
}
