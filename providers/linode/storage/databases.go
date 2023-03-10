package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/linode/linodego"
	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Databases(ctx context.Context, client providers.ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)

	databases, err := client.LinodeClient.ListDatabases(ctx, &linodego.ListOptions{})
	if err != nil {
		return resources, err
	}

	for _, database := range databases {
		resources = append(resources, models.Resource{
			Provider:   "Linode",
			Account:    client.Name,
			Service:    "Database",
			Region:     database.Region,
			ResourceId: fmt.Sprintf("%d", database.ID),
			Cost:       0,
			Name:       database.Label,
			FetchedAt:  time.Now(),
			CreatedAt:  *database.Created,
			Link:       fmt.Sprintf("https://cloud.linode.com/databases/%d", database.ID),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Linode",
		"account":   client.Name,
		"service":   "Database",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
