package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-03-01/network"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func getRouteTablesClient(subscriptionID string) network.RouteTablesClient {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	rtClient := network.NewRouteTablesClient(subscriptionID)
	rtClient.Authorizer = a
	return rtClient
}

func (azure Azure) GetRouteTablesCount(subscriptionID string) (int, error) {
	var rtCount int
	rtClient := getRouteTablesClient(subscriptionID)
	ctx := context.Background()
	rGroups, err := getGroups(subscriptionID)
	if err != nil {
		return rtCount, err
	}
	for _, rGroup := range rGroups {
		for rtItr, err := rtClient.ListComplete(ctx, rGroup); rtItr.NotDone(); rtItr.Next() {
			if err != nil {
				return rtCount, err
			}
			rtCount++
		}
	}
	return rtCount, nil
}
