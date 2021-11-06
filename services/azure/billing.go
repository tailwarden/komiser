package azure

import (
	"context"
	"fmt"

	cost "github.com/Azure/azure-sdk-for-go/services/costmanagement/mgmt/2019-11-01/costmanagement"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func getCostQueryClient(subscriptionID string) cost.QueryClient {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	queryClient := cost.NewQueryClient(subscriptionID)
	queryClient.Authorizer = a
	return queryClient
}

func (azure Azure) Usage(subscriptionID string) (cost.QueryResult, error) {
	costQueryClient := getCostQueryClient(subscriptionID)
	ctx := context.Background()
	scope := fmt.Sprintf("/subscriptions/%s", subscriptionID)
	parameters := &cost.QueryDefinition{
		Type: "ExportTypeUsage",
	}
	result, err := costQueryClient.Usage(ctx, scope, *parameters)
	return result, err
}
