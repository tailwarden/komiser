package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-03-01/network"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func getSubnetsClient(subscriptionID string) network.SubnetsClient {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	subnetsClient := network.NewSubnetsClient(subscriptionID)
	subnetsClient.Authorizer = a
	return subnetsClient
}

func (azure Azure) GetSubnetsCount(subscriptionID string) (int, error) {
	var subnetsCount int
	subnetsClient := getSubnetsClient(subscriptionID)
	ctx := context.Background()
	rGoups, err := getGroups(subscriptionID)
	if err != nil {
		return subnetsCount, err
	}
	vns, err := azure.GetVirtualNetworks(subscriptionID)
	if err != nil {
		return subnetsCount, err
	}
	for _, rGroup := range rGoups {
		for _, vn := range vns {
			for subnetItr, err := subnetsClient.ListComplete(ctx, rGroup, vn.Name); subnetItr.NotDone(); subnetItr.Next() {
				if err != nil {
					return subnetsCount, err
				}
				subnetsCount++
			}
		}
	}
	return subnetsCount, nil
}
