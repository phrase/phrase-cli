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
	initTranslationCreate()
	initTranslationExclude()
	initTranslationInclude()
	initTranslationReview()
	initTranslationShow()
	initTranslationUnverify()
	initTranslationUpdate()
	initTranslationVerify()
	initTranslationsByKey()
	initTranslationsByLocale()
	initTranslationsExclude()
	initTranslationsInclude()
	initTranslationsList()
	initTranslationsReview()
	initTranslationsSearch()
	initTranslationsUnverify()
	initTranslationsVerify()

	rootCmd.AddCommand(translationsApiCmd)
}

var translationsApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("translationsapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("TranslationsApi", "Api"), "API"}, " "),
}


func initTranslationCreate() {
	params := viper.New()
	var translationCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("TranslationCreate", strings.TrimSuffix("TranslationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("TranslationsApi", "Api"), "s"))),
		Short: "Create a translation",
		Long:  `Create a translation.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TranslationCreateOpts{}

			
			projectId := params.GetString("projectId")
			
			translationCreateParameters := api.TranslationCreateParameters{}
			

			api_response, err := client.TranslationsApi.TranslationCreate(auth, projectId, translationCreateParameters, &localVarOptionals)

			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	translationsApiCmd.AddCommand(translationCreate)

	
	AddFlag(translationCreate, "string", "projectId", "", "ID")
	
	// translationCreateParameters := api.TranslationCreateParameters{}
	

	params.BindPFlags(translationCreate.Flags())
}

func initTranslationExclude() {
	params := viper.New()
	var translationExclude = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("TranslationExclude", strings.TrimSuffix("TranslationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("TranslationsApi", "Api"), "s"))),
		Short: "Exclude a translation from export",
		Long:  `Set exclude from export flag on an existing translation.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TranslationExcludeOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			
			translationExcludeParameters := api.TranslationExcludeParameters{}
			

			data, api_response, err := client.TranslationsApi.TranslationExclude(auth, projectId, id, translationExcludeParameters, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	translationsApiCmd.AddCommand(translationExclude)

	
	AddFlag(translationExclude, "string", "projectId", "", "ID")
	
	AddFlag(translationExclude, "string", "id", "", "ID")
	
	// translationExcludeParameters := api.TranslationExcludeParameters{}
	

	params.BindPFlags(translationExclude.Flags())
}

func initTranslationInclude() {
	params := viper.New()
	var translationInclude = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("TranslationInclude", strings.TrimSuffix("TranslationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("TranslationsApi", "Api"), "s"))),
		Short: "Revoke exclusion of a translation in export",
		Long:  `Remove exclude from export flag from an existing translation.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TranslationIncludeOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			
			translationIncludeParameters := api.TranslationIncludeParameters{}
			

			data, api_response, err := client.TranslationsApi.TranslationInclude(auth, projectId, id, translationIncludeParameters, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	translationsApiCmd.AddCommand(translationInclude)

	
	AddFlag(translationInclude, "string", "projectId", "", "ID")
	
	AddFlag(translationInclude, "string", "id", "", "ID")
	
	// translationIncludeParameters := api.TranslationIncludeParameters{}
	

	params.BindPFlags(translationInclude.Flags())
}

func initTranslationReview() {
	params := viper.New()
	var translationReview = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("TranslationReview", strings.TrimSuffix("TranslationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("TranslationsApi", "Api"), "s"))),
		Short: "Review a translation",
		Long:  `Mark an existing translation as reviewed.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TranslationReviewOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			
			translationReviewParameters := api.TranslationReviewParameters{}
			

			data, api_response, err := client.TranslationsApi.TranslationReview(auth, projectId, id, translationReviewParameters, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	translationsApiCmd.AddCommand(translationReview)

	
	AddFlag(translationReview, "string", "projectId", "", "ID")
	
	AddFlag(translationReview, "string", "id", "", "ID")
	
	// translationReviewParameters := api.TranslationReviewParameters{}
	

	params.BindPFlags(translationReview.Flags())
}

func initTranslationShow() {
	params := viper.New()
	var translationShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("TranslationShow", strings.TrimSuffix("TranslationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("TranslationsApi", "Api"), "s"))),
		Short: "Get a single translation",
		Long:  `Get details on a single translation.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TranslationShowOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			

			data, api_response, err := client.TranslationsApi.TranslationShow(auth, projectId, id, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	translationsApiCmd.AddCommand(translationShow)

	
	AddFlag(translationShow, "string", "projectId", "", "ID")
	
	AddFlag(translationShow, "string", "id", "", "ID")
	

	params.BindPFlags(translationShow.Flags())
}

func initTranslationUnverify() {
	params := viper.New()
	var translationUnverify = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("TranslationUnverify", strings.TrimSuffix("TranslationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("TranslationsApi", "Api"), "s"))),
		Short: "Mark a translation as unverified",
		Long:  `Mark an existing translation as unverified.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TranslationUnverifyOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			
			translationUnverifyParameters := api.TranslationUnverifyParameters{}
			

			data, api_response, err := client.TranslationsApi.TranslationUnverify(auth, projectId, id, translationUnverifyParameters, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	translationsApiCmd.AddCommand(translationUnverify)

	
	AddFlag(translationUnverify, "string", "projectId", "", "ID")
	
	AddFlag(translationUnverify, "string", "id", "", "ID")
	
	// translationUnverifyParameters := api.TranslationUnverifyParameters{}
	

	params.BindPFlags(translationUnverify.Flags())
}

func initTranslationUpdate() {
	params := viper.New()
	var translationUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("TranslationUpdate", strings.TrimSuffix("TranslationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("TranslationsApi", "Api"), "s"))),
		Short: "Update a translation",
		Long:  `Update an existing translation.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TranslationUpdateOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			
			translationUpdateParameters := api.TranslationUpdateParameters{}
			

			data, api_response, err := client.TranslationsApi.TranslationUpdate(auth, projectId, id, translationUpdateParameters, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	translationsApiCmd.AddCommand(translationUpdate)

	
	AddFlag(translationUpdate, "string", "projectId", "", "ID")
	
	AddFlag(translationUpdate, "string", "id", "", "ID")
	
	// translationUpdateParameters := api.TranslationUpdateParameters{}
	

	params.BindPFlags(translationUpdate.Flags())
}

func initTranslationVerify() {
	params := viper.New()
	var translationVerify = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("TranslationVerify", strings.TrimSuffix("TranslationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("TranslationsApi", "Api"), "s"))),
		Short: "Verify a translation",
		Long:  `Verify an existing translation.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TranslationVerifyOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			
			translationVerifyParameters := api.TranslationVerifyParameters{}
			

			data, api_response, err := client.TranslationsApi.TranslationVerify(auth, projectId, id, translationVerifyParameters, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	translationsApiCmd.AddCommand(translationVerify)

	
	AddFlag(translationVerify, "string", "projectId", "", "ID")
	
	AddFlag(translationVerify, "string", "id", "", "ID")
	
	// translationVerifyParameters := api.TranslationVerifyParameters{}
	

	params.BindPFlags(translationVerify.Flags())
}

func initTranslationsByKey() {
	params := viper.New()
	var translationsByKey = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("TranslationsByKey", strings.TrimSuffix("TranslationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("TranslationsApi", "Api"), "s"))),
		Short: "List translations by key",
		Long:  `List translations for a specific key.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TranslationsByKeyOpts{}

			
			projectId := params.GetString("projectId")
			
			keyId := params.GetString("keyId")
			

			data, api_response, err := client.TranslationsApi.TranslationsByKey(auth, projectId, keyId, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	translationsApiCmd.AddCommand(translationsByKey)

	
	AddFlag(translationsByKey, "string", "projectId", "", "ID")
	
	AddFlag(translationsByKey, "string", "keyId", "", "ID")
	

	params.BindPFlags(translationsByKey.Flags())
}

func initTranslationsByLocale() {
	params := viper.New()
	var translationsByLocale = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("TranslationsByLocale", strings.TrimSuffix("TranslationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("TranslationsApi", "Api"), "s"))),
		Short: "List translations by locale",
		Long:  `List translations for a specific locale. If you want to download all translations for one locale we recommend to use the &lt;code&gt;locales#download&lt;/code&gt; endpoint.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TranslationsByLocaleOpts{}

			
			projectId := params.GetString("projectId")
			
			localeId := params.GetString("localeId")
			

			data, api_response, err := client.TranslationsApi.TranslationsByLocale(auth, projectId, localeId, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	translationsApiCmd.AddCommand(translationsByLocale)

	
	AddFlag(translationsByLocale, "string", "projectId", "", "ID")
	
	AddFlag(translationsByLocale, "string", "localeId", "", "ID")
	

	params.BindPFlags(translationsByLocale.Flags())
}

func initTranslationsExclude() {
	params := viper.New()
	var translationsExclude = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("TranslationsExclude", strings.TrimSuffix("TranslationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("TranslationsApi", "Api"), "s"))),
		Short: "Set exclude from export flag on translations selected by query",
		Long:  `Exclude translations matching query from locale export.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TranslationsExcludeOpts{}

			
			projectId := params.GetString("projectId")
			
			translationsExcludeParameters := api.TranslationsExcludeParameters{}
			

			data, api_response, err := client.TranslationsApi.TranslationsExclude(auth, projectId, translationsExcludeParameters, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	translationsApiCmd.AddCommand(translationsExclude)

	
	AddFlag(translationsExclude, "string", "projectId", "", "ID")
	
	// translationsExcludeParameters := api.TranslationsExcludeParameters{}
	

	params.BindPFlags(translationsExclude.Flags())
}

func initTranslationsInclude() {
	params := viper.New()
	var translationsInclude = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("TranslationsInclude", strings.TrimSuffix("TranslationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("TranslationsApi", "Api"), "s"))),
		Short: "Remove exlude from import flag from translations selected by query",
		Long:  `Include translations matching query in locale export.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TranslationsIncludeOpts{}

			
			projectId := params.GetString("projectId")
			
			translationsIncludeParameters := api.TranslationsIncludeParameters{}
			

			data, api_response, err := client.TranslationsApi.TranslationsInclude(auth, projectId, translationsIncludeParameters, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	translationsApiCmd.AddCommand(translationsInclude)

	
	AddFlag(translationsInclude, "string", "projectId", "", "ID")
	
	// translationsIncludeParameters := api.TranslationsIncludeParameters{}
	

	params.BindPFlags(translationsInclude.Flags())
}

func initTranslationsList() {
	params := viper.New()
	var translationsList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("TranslationsList", strings.TrimSuffix("TranslationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("TranslationsApi", "Api"), "s"))),
		Short: "List all translations",
		Long:  `List translations for the given project. If you want to download all translations for one locale we recommend to use the &lt;code&gt;locales#download&lt;/code&gt; endpoint.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TranslationsListOpts{}

			
			projectId := params.GetString("projectId")
			

			data, api_response, err := client.TranslationsApi.TranslationsList(auth, projectId, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	translationsApiCmd.AddCommand(translationsList)

	
	AddFlag(translationsList, "string", "projectId", "", "ID")
	

	params.BindPFlags(translationsList.Flags())
}

func initTranslationsReview() {
	params := viper.New()
	var translationsReview = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("TranslationsReview", strings.TrimSuffix("TranslationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("TranslationsApi", "Api"), "s"))),
		Short: "Review translations selected by query",
		Long:  `Review translations matching query.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TranslationsReviewOpts{}

			
			projectId := params.GetString("projectId")
			
			translationsReviewParameters := api.TranslationsReviewParameters{}
			

			data, api_response, err := client.TranslationsApi.TranslationsReview(auth, projectId, translationsReviewParameters, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	translationsApiCmd.AddCommand(translationsReview)

	
	AddFlag(translationsReview, "string", "projectId", "", "ID")
	
	// translationsReviewParameters := api.TranslationsReviewParameters{}
	

	params.BindPFlags(translationsReview.Flags())
}

func initTranslationsSearch() {
	params := viper.New()
	var translationsSearch = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("TranslationsSearch", strings.TrimSuffix("TranslationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("TranslationsApi", "Api"), "s"))),
		Short: "Search translations",
		Long:  `Search translations for the given project. Provides the same search interface as &lt;code&gt;translations#index&lt;/code&gt; but allows POST requests to avoid limitations imposed by GET requests. If you want to download all translations for one locale we recommend to use the &lt;code&gt;locales#download&lt;/code&gt; endpoint.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TranslationsSearchOpts{}

			
			projectId := params.GetString("projectId")
			
			translationsSearchParameters := api.TranslationsSearchParameters{}
			

			data, api_response, err := client.TranslationsApi.TranslationsSearch(auth, projectId, translationsSearchParameters, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	translationsApiCmd.AddCommand(translationsSearch)

	
	AddFlag(translationsSearch, "string", "projectId", "", "ID")
	
	// translationsSearchParameters := api.TranslationsSearchParameters{}
	

	params.BindPFlags(translationsSearch.Flags())
}

func initTranslationsUnverify() {
	params := viper.New()
	var translationsUnverify = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("TranslationsUnverify", strings.TrimSuffix("TranslationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("TranslationsApi", "Api"), "s"))),
		Short: "Mark translations selected by query as unverified",
		Long:  `Mark translations matching query as unverified.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TranslationsUnverifyOpts{}

			
			projectId := params.GetString("projectId")
			
			translationsUnverifyParameters := api.TranslationsUnverifyParameters{}
			

			data, api_response, err := client.TranslationsApi.TranslationsUnverify(auth, projectId, translationsUnverifyParameters, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	translationsApiCmd.AddCommand(translationsUnverify)

	
	AddFlag(translationsUnverify, "string", "projectId", "", "ID")
	
	// translationsUnverifyParameters := api.TranslationsUnverifyParameters{}
	

	params.BindPFlags(translationsUnverify.Flags())
}

func initTranslationsVerify() {
	params := viper.New()
	var translationsVerify = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("TranslationsVerify", strings.TrimSuffix("TranslationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("TranslationsApi", "Api"), "s"))),
		Short: "Verify translations selected by query",
		Long:  `Verify translations matching query.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.TranslationsVerifyOpts{}

			
			projectId := params.GetString("projectId")
			
			translationsVerifyParameters := api.TranslationsVerifyParameters{}
			

			data, api_response, err := client.TranslationsApi.TranslationsVerify(auth, projectId, translationsVerifyParameters, &localVarOptionals)

			jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
			if jsonErr != nil {
				fmt.Printf("%v\n", data)
				HandleError(err)
			}

			fmt.Printf("%s\n", string(jsonBuf))
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	translationsApiCmd.AddCommand(translationsVerify)

	
	AddFlag(translationsVerify, "string", "projectId", "", "ID")
	
	// translationsVerifyParameters := api.TranslationsVerifyParameters{}
	

	params.BindPFlags(translationsVerify.Flags())
}

