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

func Firewalls(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	svc, err := armnetwork.NewAzureFirewallsClient(client.AzureClient.SubscriptionId, client.AzureClient.Credentials, &arm.ClientOptions{})
	if err != nil {
		return resources, err
	}

	pager := svc.NewListAllPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return resources, err
		}

		for _, firewall := range page.AzureFirewallListResult.Value {
			tags := make([]models.Tag, 0)

			for key, value := range firewall.Tags {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: *value,
				})
			}

			resources = append(resources, models.Resource{
				Provider:   "Azure",
				Account:    client.Name,
				Service:    "Firewall",
				Region:     *firewall.Location,
				ResourceId: *firewall.ID,
				Cost:       0,
				Name:       *firewall.Name,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link:       fmt.Sprintf("https://portal.azure.com/#resource%s", *firewall.ID),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "Azure",
		"account":   client.Name,
		"service":   "Firewall",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
