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
	initMemberDelete()
	initMemberShow()
	initMemberUpdate()
	initMembersList()

	rootCmd.AddCommand(MembersApiCmd)
}

var MembersApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("Members"),
	Short: "Members API",
}

func initMemberDelete() {
	params := viper.New()
	var MemberDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("MemberDelete", strings.TrimSuffix("MembersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("MembersApi", "Api"), "s"))),
		Short: "Remove a user from the account",
		Long:  `Remove a user from the account. The user will be removed from the account but not deleted from Phrase. Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.MemberDeleteOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.MembersApi.MemberDelete(auth, accountId, id, &localVarOptionals)

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

	MembersApiCmd.AddCommand(MemberDelete)

	AddFlag(MemberDelete, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(MemberDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(MemberDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(MemberDelete.Flags())
}
func initMemberShow() {
	params := viper.New()
	var MemberShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("MemberShow", strings.TrimSuffix("MembersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("MembersApi", "Api"), "s"))),
		Short: "Get single member",
		Long:  `Get details on a single user in the account. Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.MemberShowOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.MembersApi.MemberShow(auth, accountId, id, &localVarOptionals)

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

	MembersApiCmd.AddCommand(MemberShow)

	AddFlag(MemberShow, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(MemberShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(MemberShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(MemberShow.Flags())
}
func initMemberUpdate() {
	params := viper.New()
	var MemberUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("MemberUpdate", strings.TrimSuffix("MembersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("MembersApi", "Api"), "s"))),
		Short: "Update a member",
		Long:  `Update user permissions in the account. Developers and translators need &lt;code&gt;project_ids&lt;/code&gt; and &lt;code&gt;locale_ids&lt;/code&gt; assigned to access them. Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.MemberUpdateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			memberUpdateParameters := api.MemberUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &memberUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", memberUpdateParameters)
			}
			data, api_response, err := client.MembersApi.MemberUpdate(auth, accountId, id, memberUpdateParameters, &localVarOptionals)

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

	MembersApiCmd.AddCommand(MemberUpdate)

	AddFlag(MemberUpdate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(MemberUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(MemberUpdate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(MemberUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(MemberUpdate.Flags())
}
func initMembersList() {
	params := viper.New()
	var MembersList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("MembersList", strings.TrimSuffix("MembersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("MembersApi", "Api"), "s"))),
		Short: "List members",
		Long:  `Get all users active in the account. It also lists resources like projects and locales the member has access to. In case nothing is shown the default access from the role is used. Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.MembersListOpts{}

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

			data, api_response, err := client.MembersApi.MembersList(auth, accountId, &localVarOptionals)

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

	MembersApiCmd.AddCommand(MembersList)

	AddFlag(MembersList, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(MembersList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(MembersList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(MembersList, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 10 by default", false)
	params.BindPFlags(MembersList.Flags())
}
