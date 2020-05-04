package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

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

	rootCmd.AddCommand(commentsApiCmd)
}

var commentsApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("commentsapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("CommentsApi", "Api"), "API"}, " "),
}


func initCommentCreate() {
	params := viper.New()
	var commentCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("CommentCreate", strings.TrimSuffix("CommentsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("CommentsApi", "Api"), "s"))),
		Short: "Create a comment",
		Long:  `Create a new comment for a key.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.CommentCreateOpts{}

			
			projectId := params.GetString("projectId")

			
			keyId := params.GetString("keyId")

			

			commentCreateParameters := api.CommentCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &commentCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", commentCreateParameters)
			}
			

			data, api_response, err := client.CommentsApi.CommentCreate(auth, projectId, keyId, commentCreateParameters, &localVarOptionals)

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

	commentsApiCmd.AddCommand(commentCreate)

	
	AddFlag(commentCreate, "string", "projectId", "", "Project ID", true)
	
	AddFlag(commentCreate, "string", "keyId", "", "Translation Key ID", true)
	
	AddFlag(commentCreate, "string", "data", "d", "payload in JSON format", true)
	// commentCreateParameters := api.CommentCreateParameters{}
	

	params.BindPFlags(commentCreate.Flags())
}

func initCommentDelete() {
	params := viper.New()
	var commentDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("CommentDelete", strings.TrimSuffix("CommentsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("CommentsApi", "Api"), "s"))),
		Short: "Delete a comment",
		Long:  `Delete an existing comment.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.CommentDeleteOpts{}

			
			projectId := params.GetString("projectId")

			
			keyId := params.GetString("keyId")

			
			id := params.GetString("id")

			

			data, api_response, err := client.CommentsApi.CommentDelete(auth, projectId, keyId, id, &localVarOptionals)

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

	commentsApiCmd.AddCommand(commentDelete)

	
	AddFlag(commentDelete, "string", "projectId", "", "Project ID", true)
	
	AddFlag(commentDelete, "string", "keyId", "", "Translation Key ID", true)
	
	AddFlag(commentDelete, "string", "id", "", "ID", true)
	

	params.BindPFlags(commentDelete.Flags())
}

func initCommentMarkCheck() {
	params := viper.New()
	var commentMarkCheck = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("CommentMarkCheck", strings.TrimSuffix("CommentsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("CommentsApi", "Api"), "s"))),
		Short: "Check if comment is read",
		Long:  `Check if comment was marked as read. Returns 204 if read, 404 if unread.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.CommentMarkCheckOpts{}

			
			projectId := params.GetString("projectId")

			
			keyId := params.GetString("keyId")

			
			id := params.GetString("id")

			

			data, api_response, err := client.CommentsApi.CommentMarkCheck(auth, projectId, keyId, id, &localVarOptionals)

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

	commentsApiCmd.AddCommand(commentMarkCheck)

	
	AddFlag(commentMarkCheck, "string", "projectId", "", "Project ID", true)
	
	AddFlag(commentMarkCheck, "string", "keyId", "", "Translation Key ID", true)
	
	AddFlag(commentMarkCheck, "string", "id", "", "ID", true)
	

	params.BindPFlags(commentMarkCheck.Flags())
}

func initCommentMarkRead() {
	params := viper.New()
	var commentMarkRead = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("CommentMarkRead", strings.TrimSuffix("CommentsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("CommentsApi", "Api"), "s"))),
		Short: "Mark a comment as read",
		Long:  `Mark a comment as read.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.CommentMarkReadOpts{}

			
			projectId := params.GetString("projectId")

			
			keyId := params.GetString("keyId")

			
			id := params.GetString("id")

			

			commentMarkReadParameters := api.CommentMarkReadParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &commentMarkReadParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", commentMarkReadParameters)
			}
			

			data, api_response, err := client.CommentsApi.CommentMarkRead(auth, projectId, keyId, id, commentMarkReadParameters, &localVarOptionals)

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

	commentsApiCmd.AddCommand(commentMarkRead)

	
	AddFlag(commentMarkRead, "string", "projectId", "", "Project ID", true)
	
	AddFlag(commentMarkRead, "string", "keyId", "", "Translation Key ID", true)
	
	AddFlag(commentMarkRead, "string", "id", "", "ID", true)
	
	AddFlag(commentMarkRead, "string", "data", "d", "payload in JSON format", true)
	// commentMarkReadParameters := api.CommentMarkReadParameters{}
	

	params.BindPFlags(commentMarkRead.Flags())
}

func initCommentMarkUnread() {
	params := viper.New()
	var commentMarkUnread = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("CommentMarkUnread", strings.TrimSuffix("CommentsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("CommentsApi", "Api"), "s"))),
		Short: "Mark a comment as unread",
		Long:  `Mark a comment as unread.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.CommentMarkUnreadOpts{}

			
			projectId := params.GetString("projectId")

			
			keyId := params.GetString("keyId")

			
			id := params.GetString("id")

			

			data, api_response, err := client.CommentsApi.CommentMarkUnread(auth, projectId, keyId, id, &localVarOptionals)

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

	commentsApiCmd.AddCommand(commentMarkUnread)

	
	AddFlag(commentMarkUnread, "string", "projectId", "", "Project ID", true)
	
	AddFlag(commentMarkUnread, "string", "keyId", "", "Translation Key ID", true)
	
	AddFlag(commentMarkUnread, "string", "id", "", "ID", true)
	

	params.BindPFlags(commentMarkUnread.Flags())
}

func initCommentShow() {
	params := viper.New()
	var commentShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("CommentShow", strings.TrimSuffix("CommentsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("CommentsApi", "Api"), "s"))),
		Short: "Get a single comment",
		Long:  `Get details on a single comment.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.CommentShowOpts{}

			
			projectId := params.GetString("projectId")

			
			keyId := params.GetString("keyId")

			
			id := params.GetString("id")

			

			data, api_response, err := client.CommentsApi.CommentShow(auth, projectId, keyId, id, &localVarOptionals)

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

	commentsApiCmd.AddCommand(commentShow)

	
	AddFlag(commentShow, "string", "projectId", "", "Project ID", true)
	
	AddFlag(commentShow, "string", "keyId", "", "Translation Key ID", true)
	
	AddFlag(commentShow, "string", "id", "", "ID", true)
	

	params.BindPFlags(commentShow.Flags())
}

func initCommentUpdate() {
	params := viper.New()
	var commentUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("CommentUpdate", strings.TrimSuffix("CommentsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("CommentsApi", "Api"), "s"))),
		Short: "Update a comment",
		Long:  `Update an existing comment.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.CommentUpdateOpts{}

			
			projectId := params.GetString("projectId")

			
			keyId := params.GetString("keyId")

			
			id := params.GetString("id")

			

			commentUpdateParameters := api.CommentUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &commentUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", commentUpdateParameters)
			}
			

			data, api_response, err := client.CommentsApi.CommentUpdate(auth, projectId, keyId, id, commentUpdateParameters, &localVarOptionals)

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

	commentsApiCmd.AddCommand(commentUpdate)

	
	AddFlag(commentUpdate, "string", "projectId", "", "Project ID", true)
	
	AddFlag(commentUpdate, "string", "keyId", "", "Translation Key ID", true)
	
	AddFlag(commentUpdate, "string", "id", "", "ID", true)
	
	AddFlag(commentUpdate, "string", "data", "d", "payload in JSON format", true)
	// commentUpdateParameters := api.CommentUpdateParameters{}
	

	params.BindPFlags(commentUpdate.Flags())
}

func initCommentsList() {
	params := viper.New()
	var commentsList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("CommentsList", strings.TrimSuffix("CommentsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("CommentsApi", "Api"), "s"))),
		Short: "List comments",
		Long:  `List all comments for a key.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.CommentsListOpts{}

			
			projectId := params.GetString("projectId")

			
			keyId := params.GetString("keyId")

			

			data, api_response, err := client.CommentsApi.CommentsList(auth, projectId, keyId, &localVarOptionals)

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

	commentsApiCmd.AddCommand(commentsList)

	
	AddFlag(commentsList, "string", "projectId", "", "Project ID", true)
	
	AddFlag(commentsList, "string", "keyId", "", "Translation Key ID", true)
	

	params.BindPFlags(commentsList.Flags())
}

