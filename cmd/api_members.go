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
	initMemberDelete()
	initMemberShow()
	initMemberUpdate()
	initMembersList()

	rootCmd.AddCommand(membersApiCmd)
}

var membersApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("membersapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("MembersApi", "Api"), "API"}, " "),
}


func initMemberDelete() {
	params := viper.New()
	var memberDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("MemberDelete", strings.TrimSuffix("MembersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("MembersApi", "Api"), "s"))),
		Short: "Remove a user from the account",
		Long:  `Remove a user from the account. The user will be removed from the account but not deleted from Phrase. Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.AccessToken,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.MemberDeleteOpts{}

			
			accountId := params.GetString("accountId")
			
			id := params.GetString("id")
			

			api_response, err := client.MembersApi.MemberDelete(auth, accountId, id, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	membersApiCmd.AddCommand(memberDelete)

	
	AddFlag(memberDelete, "string", "accountId", "", "ID")
	
	AddFlag(memberDelete, "string", "id", "", "ID")
	

	params.BindPFlags(memberDelete.Flags())
}

func initMemberShow() {
	params := viper.New()
	var memberShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("MemberShow", strings.TrimSuffix("MembersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("MembersApi", "Api"), "s"))),
		Short: "Get single member",
		Long:  `Get details on a single user in the account. Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.AccessToken,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.MemberShowOpts{}

			
			accountId := params.GetString("accountId")
			
			id := params.GetString("id")
			

			data, api_response, err := client.MembersApi.MemberShow(auth, accountId, id, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	membersApiCmd.AddCommand(memberShow)

	
	AddFlag(memberShow, "string", "accountId", "", "ID")
	
	AddFlag(memberShow, "string", "id", "", "ID")
	

	params.BindPFlags(memberShow.Flags())
}

func initMemberUpdate() {
	params := viper.New()
	var memberUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("MemberUpdate", strings.TrimSuffix("MembersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("MembersApi", "Api"), "s"))),
		Short: "Update a member",
		Long:  `Update user permissions in the account. Developers and translators need &lt;code&gt;project_ids&lt;/code&gt; and &lt;code&gt;locale_ids&lt;/code&gt; assigned to access them. Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.AccessToken,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.MemberUpdateOpts{}

			
			accountId := params.GetString("accountId")
			
			id := params.GetString("id")
			
			memberUpdateParameters := api.MemberUpdateParameters{}
			

			data, api_response, err := client.MembersApi.MemberUpdate(auth, accountId, id, memberUpdateParameters, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	membersApiCmd.AddCommand(memberUpdate)

	
	AddFlag(memberUpdate, "string", "accountId", "", "ID")
	
	AddFlag(memberUpdate, "string", "id", "", "ID")
	
	// memberUpdateParameters := api.MemberUpdateParameters{}
	

	params.BindPFlags(memberUpdate.Flags())
}

func initMembersList() {
	params := viper.New()
	var membersList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("MembersList", strings.TrimSuffix("MembersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("MembersApi", "Api"), "s"))),
		Short: "List members",
		Long:  `Get all users active in the account. It also lists resources like projects and locales the member has access to. In case nothing is shown the default access from the role is used. Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.AccessToken,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.MembersListOpts{}

			
			accountId := params.GetString("accountId")
			

			data, api_response, err := client.MembersApi.MembersList(auth, accountId, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	membersApiCmd.AddCommand(membersList)

	
	AddFlag(membersList, "string", "accountId", "", "ID")
	

	params.BindPFlags(membersList.Flags())
}

