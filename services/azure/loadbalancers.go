package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-03-01/network"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func getLoadBalancersClient(subscriptionID string) network.LoadBalancersClient {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	lbClient := network.NewLoadBalancersClient(subscriptionID)
	lbClient.Authorizer = a
	return lbClient
}

func (azure Azure) GetLoadBalancersCount(subscriptionID string) (int, error) {
	var lbCount int
	lbClient := getLoadBalancersClient(subscriptionID)
	ctx := context.Background()
	rGroups, err := getGroups(subscriptionID)
	if err != nil {
		return lbCount, err
	}
	for _, rGroup := range rGroups {
		for lbItr, err := lbClient.ListComplete(ctx, rGroup); lbItr.NotDone(); lbItr.Next() {
			if err != nil {
				return lbCount, err
			}
			lbCount++
		}
	}
	return lbCount, nil
}
