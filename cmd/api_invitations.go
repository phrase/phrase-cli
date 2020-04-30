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
	initInvitationCreate()
	initInvitationDelete()
	initInvitationResend()
	initInvitationShow()
	initInvitationUpdate()
	initInvitationsList()

	rootCmd.AddCommand(invitationsApiCmd)
}

var invitationsApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("invitationsapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("InvitationsApi", "Api"), "API"}, " "),
}


func initInvitationCreate() {
	params := viper.New()
	var invitationCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("InvitationCreate", strings.TrimSuffix("InvitationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("InvitationsApi", "Api"), "s"))),
		Short: "Create a new invitation",
		Long:  `Invite a person to an account. Developers and translators need &lt;code&gt;project_ids&lt;/code&gt; and &lt;code&gt;locale_ids&lt;/code&gt; assigned to access them. Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.InvitationCreateOpts{}

			
			accountId := params.GetString("accountId")
			
			invitationCreateParameters := api.InvitationCreateParameters{}
			

			api_response, err := client.InvitationsApi.InvitationCreate(auth, accountId, invitationCreateParameters, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	invitationsApiCmd.AddCommand(invitationCreate)

	
	AddFlag(invitationCreate, "string", "accountId", "", "ID")
	
	// invitationCreateParameters := api.InvitationCreateParameters{}
	

	params.BindPFlags(invitationCreate.Flags())
}

func initInvitationDelete() {
	params := viper.New()
	var invitationDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("InvitationDelete", strings.TrimSuffix("InvitationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("InvitationsApi", "Api"), "s"))),
		Short: "Delete an invitation",
		Long:  `Delete an existing invitation (must not be accepted yet). Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.InvitationDeleteOpts{}

			
			accountId := params.GetString("accountId")
			
			id := params.GetString("id")
			

			api_response, err := client.InvitationsApi.InvitationDelete(auth, accountId, id, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	invitationsApiCmd.AddCommand(invitationDelete)

	
	AddFlag(invitationDelete, "string", "accountId", "", "ID")
	
	AddFlag(invitationDelete, "string", "id", "", "ID")
	

	params.BindPFlags(invitationDelete.Flags())
}

func initInvitationResend() {
	params := viper.New()
	var invitationResend = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("InvitationResend", strings.TrimSuffix("InvitationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("InvitationsApi", "Api"), "s"))),
		Short: "Resend an invitation",
		Long:  `Resend the invitation email (must not be accepted yet). Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.InvitationResendOpts{}

			
			accountId := params.GetString("accountId")
			
			id := params.GetString("id")
			

			data, api_response, err := client.InvitationsApi.InvitationResend(auth, accountId, id, &localVarOptionals)

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

	invitationsApiCmd.AddCommand(invitationResend)

	
	AddFlag(invitationResend, "string", "accountId", "", "ID")
	
	AddFlag(invitationResend, "string", "id", "", "ID")
	

	params.BindPFlags(invitationResend.Flags())
}

func initInvitationShow() {
	params := viper.New()
	var invitationShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("InvitationShow", strings.TrimSuffix("InvitationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("InvitationsApi", "Api"), "s"))),
		Short: "Get a single invitation",
		Long:  `Get details on a single invitation. Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.InvitationShowOpts{}

			
			accountId := params.GetString("accountId")
			
			id := params.GetString("id")
			

			data, api_response, err := client.InvitationsApi.InvitationShow(auth, accountId, id, &localVarOptionals)

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

	invitationsApiCmd.AddCommand(invitationShow)

	
	AddFlag(invitationShow, "string", "accountId", "", "ID")
	
	AddFlag(invitationShow, "string", "id", "", "ID")
	

	params.BindPFlags(invitationShow.Flags())
}

func initInvitationUpdate() {
	params := viper.New()
	var invitationUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("InvitationUpdate", strings.TrimSuffix("InvitationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("InvitationsApi", "Api"), "s"))),
		Short: "Update an invitation",
		Long:  `Update an existing invitation (must not be accepted yet). The &lt;code&gt;email&lt;/code&gt; cannot be updated. Developers and translators need &lt;code&gt;project_ids&lt;/code&gt; and &lt;code&gt;locale_ids&lt;/code&gt; assigned to access them. Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.InvitationUpdateOpts{}

			
			accountId := params.GetString("accountId")
			
			id := params.GetString("id")
			
			invitationUpdateParameters := api.InvitationUpdateParameters{}
			

			data, api_response, err := client.InvitationsApi.InvitationUpdate(auth, accountId, id, invitationUpdateParameters, &localVarOptionals)

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

	invitationsApiCmd.AddCommand(invitationUpdate)

	
	AddFlag(invitationUpdate, "string", "accountId", "", "ID")
	
	AddFlag(invitationUpdate, "string", "id", "", "ID")
	
	// invitationUpdateParameters := api.InvitationUpdateParameters{}
	

	params.BindPFlags(invitationUpdate.Flags())
}

func initInvitationsList() {
	params := viper.New()
	var invitationsList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("InvitationsList", strings.TrimSuffix("InvitationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("InvitationsApi", "Api"), "s"))),
		Short: "List invitations",
		Long:  `List invitations for an account. It will also list the accessible resources like projects and locales the invited user has access to. In case nothing is shown the default access from the role is used. Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.InvitationsListOpts{}

			
			accountId := params.GetString("accountId")
			

			data, api_response, err := client.InvitationsApi.InvitationsList(auth, accountId, &localVarOptionals)

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

	invitationsApiCmd.AddCommand(invitationsList)

	
	AddFlag(invitationsList, "string", "accountId", "", "ID")
	

	params.BindPFlags(invitationsList.Flags())
}

