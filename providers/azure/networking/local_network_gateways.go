package networking

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/providers/azure/resourcegroup"
	"time"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func LocalNetworkGateways(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	resourceGroups, err := resourcegroup.ResourceGroups(ctx, client)
	if err != nil {
		return resources, err
	}

	if len(resourceGroups) < 1 {
		return resources, nil
	}

	localNetworkGatewayClient, err := armnetwork.NewLocalNetworkGatewaysClient(
		client.AzureClient.SubscriptionId,
		client.AzureClient.Credentials,
		&arm.ClientOptions{},
	)
	if err != nil {
		return resources, err
	}

	// check resources on each resource group
	for _, name := range resourceGroups {
		pager := localNetworkGatewayClient.NewListPager(
			name, nil)
		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				return resources, err
			}

			for _, lng := range page.LocalNetworkGatewayListResult.Value {
				tags := make([]models.Tag, 0)

				for key, value := range lng.Tags {
					tags = append(tags, models.Tag{
						Key:   key,
						Value: *value,
					})
				}

				resources = append(resources, models.Resource{
					Provider:   "Azure",
					Account:    client.Name,
					Service:    "Local Network Gateway",
					Region:     *lng.Location,
					ResourceId: *lng.ID,
					Cost:       0,
					Name:       *lng.Name,
					FetchedAt:  time.Now(),
					Tags:       tags,
					Link:       fmt.Sprint("https://portal.azure.com/#resource%s", *lng.ID),
				})
			}
		}
	}

	log.WithFields(log.Fields{
		"provider":  "Azure",
		"account":   client.Name,
		"service":   "Local Network Gateway",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
