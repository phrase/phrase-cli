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
	initVersionShow()
	initVersionsList()

	rootCmd.AddCommand(versionsHistoryApiCmd)
}

var versionsHistoryApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("versionshistoryapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("VersionsHistoryApi", "Api"), "API"}, " "),
}


func initVersionShow() {
	params := viper.New()
	var versionShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("VersionShow", strings.TrimSuffix("VersionsHistoryApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("VersionsHistoryApi", "Api"), "s"))),
		Short: "Get a single version",
		Long:  `Get details on a single version.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.AccessToken,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.VersionShowOpts{}

			
			projectId := params.GetString("projectId")
			
			translationId := params.GetString("translationId")
			
			id := params.GetString("id")
			
			versionShowParameters := api.VersionShowParameters{}
			

			data, api_response, err := client.VersionsHistoryApi.VersionShow(auth, projectId, translationId, id, versionShowParameters, &localVarOptionals)

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

	versionsHistoryApiCmd.AddCommand(versionShow)

	
	AddFlag(versionShow, "string", "projectId", "", "ID")
	
	AddFlag(versionShow, "string", "translationId", "", "ID")
	
	AddFlag(versionShow, "string", "id", "", "ID")
	
	// versionShowParameters := api.VersionShowParameters{}
	

	params.BindPFlags(versionShow.Flags())
}

func initVersionsList() {
	params := viper.New()
	var versionsList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("VersionsList", strings.TrimSuffix("VersionsHistoryApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("VersionsHistoryApi", "Api"), "s"))),
		Short: "List all versions",
		Long:  `List all versions for the given translation.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.AccessToken,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.VersionsListOpts{}

			
			projectId := params.GetString("projectId")
			
			translationId := params.GetString("translationId")
			
			versionsListParameters := api.VersionsListParameters{}
			

			data, api_response, err := client.VersionsHistoryApi.VersionsList(auth, projectId, translationId, versionsListParameters, &localVarOptionals)

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

	versionsHistoryApiCmd.AddCommand(versionsList)

	
	AddFlag(versionsList, "string", "projectId", "", "ID")
	
	AddFlag(versionsList, "string", "translationId", "", "ID")
	
	// versionsListParameters := api.VersionsListParameters{}
	

	params.BindPFlags(versionsList.Flags())
}

