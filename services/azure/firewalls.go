package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-03-01/network"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func getAzureFirewallsClient(subscriptionID string) network.AzureFirewallsClient {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	azureFirewallsClient := network.NewAzureFirewallsClient(subscriptionID)
	azureFirewallsClient.Authorizer = a
	return azureFirewallsClient
}

func (azure Azure) GetFirewallsCount(subscriptionID string) (int, error) {
	var fwCount int
	azureFirewallsClient := getAzureFirewallsClient(subscriptionID)
	ctx := context.Background()
	rGroups, err := getGroups(subscriptionID)
	if err != nil {
		return fwCount, err
	}
	for _, rGroup := range rGroups {
		for fwItr, err := azureFirewallsClient.ListComplete(ctx, rGroup); fwItr.NotDone(); fwItr.Next() {
			if err != nil {
				return fwCount, err
			}
			fwCount++
		}
	}
	return fwCount, nil
}
