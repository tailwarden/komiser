package resourcegroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/providers"
)

func ResourceGroups(ctx context.Context, client providers.ProviderClient) ([]string, error) {
	resourceGroupNames := make([]string, 0)

	resourceGroupClient, err := armresources.NewResourceGroupsClient(
		client.AzureClient.SubscriptionId,
		client.AzureClient.Credentials,
		&arm.ClientOptions{},
	)

	if err != nil {
		return resourceGroupNames, err
	}

	pager := resourceGroupClient.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return resourceGroupNames, err
		}

		for _, resourceGroup := range page.ResourceGroupListResult.Value {

			resourceGroupNames = append(resourceGroupNames, *resourceGroup.Name)
		}
	}

	log.WithFields(log.Fields{
		"provider":  "Azure",
		"account":   client.Name,
		"service":   "Resource Group",
		"resources": len(resourceGroupNames),
	}).Info("Fetched resource groups")

	return resourceGroupNames, nil
}
