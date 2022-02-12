package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-03-01/network"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func getProfileClient(subscriptionID string) network.ProfilesClient {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	profilesClient := network.NewProfilesClient(subscriptionID)
	profilesClient.Authorizer = a
	return profilesClient
}

func (azure Azure) GetProfileCount(subscriptionID string) (int, error) {
	var profilesCount int
	profilesClient := getProfileClient(subscriptionID)
	ctx := context.Background()
	rGroups, err := getGroups(subscriptionID)
	if err != nil {
		return profilesCount, err
	}
	for _, rGroup := range rGroups {
		for profileItr, err := profilesClient.ListComplete(ctx, rGroup); profileItr.NotDone(); profileItr.Next() {
			if err != nil {
				return profilesCount, err
			}
			profilesCount++
		}
	}
	return profilesCount, nil
}
