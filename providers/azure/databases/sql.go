package databases

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"

	log "github.com/sirupsen/logrus"
)

func Sql(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	svc, err := armsql.NewServersClient(client.AzureClient.SubscriptionId, client.AzureClient.Credentials, &arm.ClientOptions{})

	if err != nil {
		return resources, err
	}

	pager := svc.NewListPager(nil)

	for pager.More() {
		page, err := pager.NextPage(ctx)

		if err != nil {
			return resources, err
		}

		for _, db := range page.ServerListResult.Value {
			tags := make([]models.Tag, 0)

			for key, value := range db.Tags {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: *value,
				})
			}

			resources = append(resources, models.Resource{
				Provider:   "Azure",
				Account:    client.Name,
				Service:    "SQL Database Server",
				Region:     *db.Location,
				ResourceId: *db.ID,
				Cost:       0,
				Name:       *db.Name,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link:       fmt.Sprintf("https://portal.azure.com/#resource%s", *db.ID),
			})

		}
	}

	log.WithFields(log.Fields{
		"provider":  "Azure",
		"account":   client.Name,
		"service":   "PostgreSQL Database Servers",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
