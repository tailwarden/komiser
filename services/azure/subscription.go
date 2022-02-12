package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/subscription/mgmt/2020-09-01/subscription"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func getSubscriptionClient() subscription.SubscriptionsClient {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	subsClient := subscription.NewSubscriptionsClient()
	subsClient.Authorizer = a
	return subsClient
}

func (azure Azure) DescribeRegions(subscriptionID string) ([]string, error) {
	subscriptionClient := getSubscriptionClient()
	ctx := context.Background()
	result, err := subscriptionClient.ListLocations(ctx, subscriptionID)
	regions := make([]string, 10)
	if err != nil {
		return nil, err
	} else {
		locations := *result.Value
		for _, location := range locations {
			if *location.Name != "" {
				regions = append(regions, *location.Name)
			}
		}
	}
	return regions, nil
}
