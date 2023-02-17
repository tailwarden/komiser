package storage

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Databases(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	paginatedDatabases, err := client.CivoClient.ListDatabases()
	if err != nil {
		return resources, err
	}
	// pagination is not yet implemented in client.CivoClient.ListDatabases() in Civo client library. Once its implemented, it would look like something below
	// func (c *Client) ListDatabases(page int, perPage int) (*PaginatedDatabaseList, error) {}

	for _, resource := range paginatedDatabases.Items {
		resources = append(resources, models.Resource{
			Provider:   "Civo",
			Account:    client.Name,
			Service:    "Database",
			Region:     client.CivoClient.Region,
			ResourceId: resource.ID,
			Name:       resource.Name,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://dashboard.civo.com/databases/%s", resource.ID),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Civo",
		"account":   client.Name,
		"service":   "Database",
		"region":    client.CivoClient.Region,
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
