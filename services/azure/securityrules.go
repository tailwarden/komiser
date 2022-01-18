package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-03-01/network"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func getSecurityRulesClient(subscriptionID string) network.SecurityRulesClient {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	securityRulesClient := network.NewSecurityRulesClient(subscriptionID)
	securityRulesClient.Authorizer = a
	return securityRulesClient
}

func (azure Azure) GetSecurityRulesCount(subscriptionID string) (int, error) {
	var srCount int
	srClient := getSecurityRulesClient(subscriptionID)
	ctx := context.Background()
	rGroups, err := getGroups(subscriptionID)
	securityGroups, err := azure.GetSecurityGroups(subscriptionID)
	if err != nil {
		return srCount, err
	}
	for _, rGroup := range rGroups {
		for _, securityGroup := range securityGroups {
			for srItr, err := srClient.ListComplete(ctx, rGroup, securityGroup.Name); srItr.NotDone(); srItr.Next() {
				if err != nil {
					return srCount, err
				}
				srCount++
			}
		}

	}
	return srCount, nil
}
