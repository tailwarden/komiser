package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-03-01/network"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func getPublicIPClient(subscriptionID string) network.PublicIPAddressesClient {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	publicIPClient := network.NewPublicIPAddressesClient(subscriptionID)
	publicIPClient.Authorizer = a
	return publicIPClient
}

func (azure Azure) GetPublicIPsCount(subscriptionID string) (int, error) {
	var ipCount int
	publicIPClient := getPublicIPClient(subscriptionID)
	ctx := context.Background()
	rGroups, err := getGroups(subscriptionID)
	if err != nil {
		return ipCount, err
	}
	for _, rGroup := range rGroups {
		for ipItr, err := publicIPClient.ListComplete(ctx, rGroup); ipItr.NotDone(); ipItr.Next() {
			if err != nil {
				return ipCount, err
			}
			ipCount++
		}
	}
	return ipCount, nil
}
