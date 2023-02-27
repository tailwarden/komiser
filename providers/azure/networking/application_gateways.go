package networking

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func ApplicationGateways(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	appGatewayClient, err := armnetwork.NewApplicationGatewaysClient(
		client.AzureClient.SubscriptionId,
		client.AzureClient.Credentials,
		&arm.ClientOptions{},
	)
	if err != nil {
		return resources, err
	}

	pager := appGatewayClient.NewListAllPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return resources, err
		}

		for _, gateway := range page.ApplicationGatewayListResult.Value {
			tags := make([]models.Tag, 0)

			for key, value := range gateway.Tags {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: *value,
				})
			}

			resources = append(resources, models.Resource{
				Provider:   "Azure",
				Account:    client.Name,
				Service:    "Application Gateway",
				Region:     *gateway.Location,
				ResourceId: *gateway.ID,
				Cost:       0,
				Name:       *gateway.Name,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link:       fmt.Sprintf("https://portal.azure.com/#resource%s", *gateway.ID),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "Azure",
		"account":   client.Name,
		"service":   "Application Gateway",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
