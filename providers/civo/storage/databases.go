package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/civo/civogo"
	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Databases(_ context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	// pagination is not yet implemented in client.CivoClient.ListDatabases() in Civo client library. Once its implemented, it would look like something below
	// func (c *Client) ListDatabases(page int, perPage int) (*PaginatedDatabaseList, error) {}
	paginatedDatabases, err := client.CivoClient.ListDatabases()
	if err != nil {
		return resources, err
	}

	sizes, err := client.CivoClient.ListInstanceSizes()
	if err != nil {
		return resources, err
	}
	sizeMap := getSizeMap(sizes)

	for _, resource := range paginatedDatabases.Items {

		resourceInGB := sizeMap[resource.Size]

		monthlyCost := float64((resourceInGB / 20) * (20 + (resource.Nodes-1)*15))

		relations := getDatabaseRelation(resource)
		resources = append(resources, models.Resource{
			Provider:   "Civo",
			Account:    client.Name,
			Service:    "Database",
			Region:     client.CivoClient.Region,
			ResourceId: resource.ID,
			Name:       resource.Name,
			Cost:       monthlyCost,
			Relations:  relations,
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

func getDatabaseRelation(db civogo.Database) []models.Link {

	var rel []models.Link

	if len(db.NetworkID) > 0 {
		rel = append(rel, models.Link{
			ResourceID: db.NetworkID,
			Type:       "Network",
			Name:       db.NetworkID,
			Relation:   "USES",
		})
	}

	if len(db.FirewallID) > 0 {
		rel = append(rel, models.Link{
			ResourceID: db.FirewallID,
			Type:       "Firewall",
			Name:       db.FirewallID,
			Relation:   "USES",
		})
	}

	return rel
}

func getSizeMap(sizes []civogo.InstanceSize) map[string]int {
	sm := make(map[string]int)
	for _, size := range sizes {
		sm[size.Name] = size.DiskGigabytes
	}
	return sm
}
