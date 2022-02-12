package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-03-01/network"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	. "github.com/mlabouardy/komiser/models/azure"
)

func getVirtualNetworksClient(subscriptionID string) network.VirtualNetworksClient {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	vnClient := network.NewVirtualNetworksClient(subscriptionID)
	vnClient.Authorizer = a
	return vnClient
}

func (azure Azure) GetVirtualNetworks(subscriptionID string) ([]VirtualNetwork, error) {
	virtualNetworks := make([]VirtualNetwork, 0)
	vnClient := getVirtualNetworksClient(subscriptionID)
	ctx := context.Background()
	rGroups, err := getGroups(subscriptionID)
	if err != nil {
		return virtualNetworks, err
	}
	for _, rGroup := range rGroups {
		for vnItr, err := vnClient.ListComplete(ctx, rGroup); vnItr.NotDone(); vnItr.Next() {
			if err != nil {
				return virtualNetworks, err
			}
			virtualNetworks = append(virtualNetworks, VirtualNetwork{
				Name: *vnItr.Value().Name,
				ID:   *vnItr.Value().ID,
			})
		}
	}
	return virtualNetworks, nil
}
