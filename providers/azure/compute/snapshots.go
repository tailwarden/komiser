package compute

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Snapshots(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	svc, err := armcompute.NewSnapshotsClient(client.AzureClient.SubscriptionId, client.AzureClient.Credentials, &arm.ClientOptions{})
	if err != nil {
		return resources, err
	}

	pager := svc.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return resources, err
		}

		for _, snapshot := range page.SnapshotList.Value {
			tags := make([]models.Tag, 0)

			for key, value := range snapshot.Tags {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: *value,
				})
			}

			resources = append(resources, models.Resource{
				Provider:   "Azure",
				Account:    client.Name,
				Service:    "Snapshot",
				Region:     *snapshot.Location,
				ResourceId: *snapshot.ID,
				Cost:       0,
				Name:       *snapshot.Name,
				FetchedAt:  time.Now(),
				Tags:       tags,
				CreatedAt:  *snapshot.Properties.TimeCreated,
				Link:       fmt.Sprintf("https://portal.azure.com/#resource%s", *snapshot.ID),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "Azure",
		"account":   client.Name,
		"service":   "Snapshot",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
