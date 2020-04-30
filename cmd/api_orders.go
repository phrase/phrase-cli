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
	initOrderConfirm()
	initOrderCreate()
	initOrderDelete()
	initOrderShow()
	initOrdersList()

	rootCmd.AddCommand(ordersApiCmd)
}

var ordersApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("ordersapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("OrdersApi", "Api"), "API"}, " "),
}


func initOrderConfirm() {
	params := viper.New()
	var orderConfirm = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("OrderConfirm", strings.TrimSuffix("OrdersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("OrdersApi", "Api"), "s"))),
		Short: "Confirm an order",
		Long:  `Confirm an existing order and send it to the provider for translation. Same constraints as for create.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.OrderConfirmOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			
			orderConfirmParameters := api.OrderConfirmParameters{}
			

			data, api_response, err := client.OrdersApi.OrderConfirm(auth, projectId, id, orderConfirmParameters, &localVarOptionals)

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

	ordersApiCmd.AddCommand(orderConfirm)

	
	AddFlag(orderConfirm, "string", "projectId", "", "ID")
	
	AddFlag(orderConfirm, "string", "id", "", "ID")
	
	// orderConfirmParameters := api.OrderConfirmParameters{}
	

	params.BindPFlags(orderConfirm.Flags())
}

func initOrderCreate() {
	params := viper.New()
	var orderCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("OrderCreate", strings.TrimSuffix("OrdersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("OrdersApi", "Api"), "s"))),
		Short: "Create a new order",
		Long:  `Create a new order. Access token scope must include &lt;code&gt;orders.create&lt;/code&gt;.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.OrderCreateOpts{}

			
			projectId := params.GetString("projectId")
			
			orderCreateParameters := api.OrderCreateParameters{}
			

			data, api_response, err := client.OrdersApi.OrderCreate(auth, projectId, orderCreateParameters, &localVarOptionals)

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

	ordersApiCmd.AddCommand(orderCreate)

	
	AddFlag(orderCreate, "string", "projectId", "", "ID")
	
	// orderCreateParameters := api.OrderCreateParameters{}
	

	params.BindPFlags(orderCreate.Flags())
}

func initOrderDelete() {
	params := viper.New()
	var orderDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("OrderDelete", strings.TrimSuffix("OrdersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("OrdersApi", "Api"), "s"))),
		Short: "Cancel an order",
		Long:  `Cancel an existing order. Must not yet be confirmed.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.OrderDeleteOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			

			data, api_response, err := client.OrdersApi.OrderDelete(auth, projectId, id, &localVarOptionals)

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

	ordersApiCmd.AddCommand(orderDelete)

	
	AddFlag(orderDelete, "string", "projectId", "", "ID")
	
	AddFlag(orderDelete, "string", "id", "", "ID")
	

	params.BindPFlags(orderDelete.Flags())
}

func initOrderShow() {
	params := viper.New()
	var orderShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("OrderShow", strings.TrimSuffix("OrdersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("OrdersApi", "Api"), "s"))),
		Short: "Get a single order",
		Long:  `Get details on a single order.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.OrderShowOpts{}

			
			projectId := params.GetString("projectId")
			
			id := params.GetString("id")
			

			data, api_response, err := client.OrdersApi.OrderShow(auth, projectId, id, &localVarOptionals)

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

	ordersApiCmd.AddCommand(orderShow)

	
	AddFlag(orderShow, "string", "projectId", "", "ID")
	
	AddFlag(orderShow, "string", "id", "", "ID")
	

	params.BindPFlags(orderShow.Flags())
}

func initOrdersList() {
	params := viper.New()
	var ordersList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("OrdersList", strings.TrimSuffix("OrdersApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("OrdersApi", "Api"), "s"))),
		Short: "List orders",
		Long:  `List all orders for the given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.OrdersListOpts{}

			
			projectId := params.GetString("projectId")
			

			data, api_response, err := client.OrdersApi.OrdersList(auth, projectId, &localVarOptionals)

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

	ordersApiCmd.AddCommand(ordersList)

	
	AddFlag(ordersList, "string", "projectId", "", "ID")
	

	params.BindPFlags(ordersList.Flags())
}

