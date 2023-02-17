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

func Disks(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	svc, err := armcompute.NewDisksClient(client.AzureClient.SubscriptionId, client.AzureClient.Credentials, &arm.ClientOptions{})
	if err != nil {
		return resources, err
	}

	pager := svc.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return resources, err
		}

		for _, disk := range page.DiskList.Value {
			tags := make([]models.Tag, 0)

			for key, value := range disk.Tags {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: *value,
				})
			}

			resources = append(resources, models.Resource{
				Provider:   "Azure",
				Account:    client.Name,
				Service:    "Disk",
				Region:     *disk.Location,
				ResourceId: *disk.ID,
				Cost:       0,
				Name:       *disk.Name,
				FetchedAt:  time.Now(),
				Tags:       tags,
				CreatedAt:  *disk.Properties.TimeCreated,
				Link:       fmt.Sprintf("https://portal.azure.com/#resource%s", *disk.ID),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "Azure",
		"account":   client.Name,
		"service":   "Disk",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
