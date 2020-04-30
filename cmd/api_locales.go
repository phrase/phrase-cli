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
	initLocaleCreate()
	initLocaleDelete()
	initLocaleDownload()
	initLocaleShow()
	initLocaleUpdate()
	initLocalesList()

	rootCmd.AddCommand(localesApiCmd)
}

var localesApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("localesapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("LocalesApi", "Api"), "API"}, " "),
}


func initLocaleCreate() {
	params := viper.New()
	var localeCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("LocaleCreate", strings.TrimSuffix("LocalesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("LocalesApi", "Api"), "s"))),
		Short: "Create a locale",
		Long:  `Create a new locale.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.LocaleCreateOpts{}

			
			projectId := params.GetString("projectId")
			
			localeCreateParameters := api.LocaleCreateParameters{}
			

			api_response, err := client.LocalesApi.LocaleCreate(auth, projectId, localeCreateParameters, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	localesApiCmd.AddCommand(localeCreate)

	
	AddFlag(localeCreate, "string", "projectId", "", "ID")
	
	// localeCreateParameters := api.LocaleCreateParameters{}
	

	params.BindPFlags(localeCreate.Flags())
}

func initLocaleDelete() {
	params := viper.New()
	var localeDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("LocaleDelete", strings.TrimSuffix("LocalesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("LocalesApi", "Api"), "s"))),
		Short: "Delete a locale",
		Long:  `Delete an existing locale.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.LocaleDeleteOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			
			localeDeleteParameters := api.LocaleDeleteParameters{}
			

			api_response, err := client.LocalesApi.LocaleDelete(auth, projectId, id, localeDeleteParameters, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	localesApiCmd.AddCommand(localeDelete)

	
	AddFlag(localeDelete, "string", "projectId", "", "ID")
	
	AddFlag(localeDelete, "string", "id", "", "ID")
	
	// localeDeleteParameters := api.LocaleDeleteParameters{}
	

	params.BindPFlags(localeDelete.Flags())
}

func initLocaleDownload() {
	params := viper.New()
	var localeDownload = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("LocaleDownload", strings.TrimSuffix("LocalesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("LocalesApi", "Api"), "s"))),
		Short: "Download a locale",
		Long:  `Download a locale in a specific file format.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.LocaleDownloadOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			
			localeDownloadParameters := api.LocaleDownloadParameters{}
			

			api_response, err := client.LocalesApi.LocaleDownload(auth, projectId, id, localeDownloadParameters, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	localesApiCmd.AddCommand(localeDownload)

	
	AddFlag(localeDownload, "string", "projectId", "", "ID")
	
	AddFlag(localeDownload, "string", "id", "", "ID")
	
	// localeDownloadParameters := api.LocaleDownloadParameters{}
	

	params.BindPFlags(localeDownload.Flags())
}

func initLocaleShow() {
	params := viper.New()
	var localeShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("LocaleShow", strings.TrimSuffix("LocalesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("LocalesApi", "Api"), "s"))),
		Short: "Get a single locale",
		Long:  `Get details on a single locale for a given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.LocaleShowOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			
			localeShowParameters := api.LocaleShowParameters{}
			

			data, api_response, err := client.LocalesApi.LocaleShow(auth, projectId, id, localeShowParameters, &localVarOptionals)

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

	localesApiCmd.AddCommand(localeShow)

	
	AddFlag(localeShow, "string", "projectId", "", "ID")
	
	AddFlag(localeShow, "string", "id", "", "ID")
	
	// localeShowParameters := api.LocaleShowParameters{}
	

	params.BindPFlags(localeShow.Flags())
}

func initLocaleUpdate() {
	params := viper.New()
	var localeUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("LocaleUpdate", strings.TrimSuffix("LocalesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("LocalesApi", "Api"), "s"))),
		Short: "Update a locale",
		Long:  `Update an existing locale.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.LocaleUpdateOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			
			localeUpdateParameters := api.LocaleUpdateParameters{}
			

			data, api_response, err := client.LocalesApi.LocaleUpdate(auth, projectId, id, localeUpdateParameters, &localVarOptionals)

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

	localesApiCmd.AddCommand(localeUpdate)

	
	AddFlag(localeUpdate, "string", "projectId", "", "ID")
	
	AddFlag(localeUpdate, "string", "id", "", "ID")
	
	// localeUpdateParameters := api.LocaleUpdateParameters{}
	

	params.BindPFlags(localeUpdate.Flags())
}

func initLocalesList() {
	params := viper.New()
	var localesList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("LocalesList", strings.TrimSuffix("LocalesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("LocalesApi", "Api"), "s"))),
		Short: "List locales",
		Long:  `List all locales for the given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.LocalesListOpts{}

			
			projectId := params.GetString("projectId")
			
			localesListParameters := api.LocalesListParameters{}
			

			data, api_response, err := client.LocalesApi.LocalesList(auth, projectId, localesListParameters, &localVarOptionals)

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

	localesApiCmd.AddCommand(localesList)

	
	AddFlag(localesList, "string", "projectId", "", "ID")
	
	// localesListParameters := api.LocalesListParameters{}
	

	params.BindPFlags(localesList.Flags())
}

