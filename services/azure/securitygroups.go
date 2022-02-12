package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-03-01/network"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	. "github.com/mlabouardy/komiser/models/azure"
)

func getSGClient(subscriptionID string) network.SecurityGroupsClient {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	sgClient := network.NewSecurityGroupsClient(subscriptionID)
	sgClient.Authorizer = a
	return sgClient
}

func (azure Azure) GetSecurityGroups(subscriptionID string) ([]SecurityGroup, error) {
	securityGroups := make([]SecurityGroup, 0)
	sgClient := getSGClient(subscriptionID)
	ctx := context.Background()
	rGroups, err := getGroups(subscriptionID)
	if err != nil {
		return securityGroups, err
	}
	for _, rGroup := range rGroups {
		for sgItr, err := sgClient.ListComplete(ctx, rGroup); sgItr.NotDone(); sgItr.Next() {
			if err != nil {
				return securityGroups, err
			}
			securityGroups = append(securityGroups, SecurityGroup{
				Name: *sgItr.Value().Name,
				ID:   *sgItr.Value().ID,
			})
		}
	}
	return securityGroups, nil
}
