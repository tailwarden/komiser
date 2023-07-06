package sql

import (
	"context"
	"fmt"
	"strings"
	"time"

	// "github.com/linode/linodego"
	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

// Cost data for Dedicated CPU instances
var dedicatedCPUCosts = map[string]float64{
	"Dedicated 4GB":   65.00,
	"Dedicated 8GB":   130.00,
	"Dedicated 16GB":  260.00,
	"Dedicated 32GB":  520.00,
	"Dedicated 64GB":  1040.00,
	"Dedicated 96GB":  1560.00,
	"Dedicated 128GB": 2080.00,
	"Dedicated 256GB": 4160.00,
	"Dedicated 512GB": 8320.00,
}

// Cost data for Shared CPU instances
var sharedCPUCosts = map[string]float64{
	"Shared 1GB":   15.00,
	"Shared 2GB":   30.00,
	"Shared 4GB":   60.00,
	"Shared 8GB":   120.00,
	"Shared 16GB":  240.00,
	"Shared 32GB":  480.00,
	"Shared 64GB":  960.00,
	"Shared 96GB":  1440.00,
	"Shared 128GB": 1920.00,
	"Shared 192GB": 2880.00,
	"Shared 256GB": 3840.00,
}

// Instances fetches MySQL instances from the provider and returns them as resources.
func Instances(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	// Fetch MySQL databases from the Linode provider
	databases, err := client.LinodeClient.ListMySQLDatabases(ctx, nil)
	if err != nil {
		return resources, err
	}

	for _, database := range databases {
		// Get the cluster size for the database
		clusterSize, err := GetClusterSize(ctx, client, database.ID)
		if err != nil {
			log.Warnf("Failed to get cluster size for MySQL database: %d, Error: %s", database.ID, err.Error())
			// Skip this database and continue with the next one
			continue
		}

		// Calculate the cost based on the database type and cluster size
		cost, ok := InstancesCost(database.Type, clusterSize)
		if !ok {
			log.Warnf("Failed to calculate cost for MySQL database: %d, Type: %s", database.ID, database.Type)
			// Skip this database and continue with the next one
			continue
		}

		resources = append(resources, models.Resource{
			Provider:   "Linode",
			Account:    client.Name,
			Service:    "MySQL",
			Region:     database.Region,
			ResourceId: fmt.Sprintf("%d", database.ID),
			Cost:       cost,
			Name:       database.Label,
			FetchedAt:  time.Now(),
			CreatedAt:  *database.Created,
			Link:       fmt.Sprintf("https://cloud.linode.com/databases/%d", database.ID),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Linode",
		"account":   client.Name,
		"service":   "MySQL",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}

// InstancesCost calculates the cost for the given MySQL instance type and cluster size.
func InstancesCost(instanceType string, clusterSize int) (float64, bool) {
	// Calculate cost based on instance type
	if strings.Contains(instanceType, "Dedicated") {
		cost, ok := dedicatedCPUCosts[instanceType]
		if !ok {
			return 0, false
		}

		// Adjust cost based on the cluster size
		if clusterSize == 3 {
			cost *= 3

			return cost, true
		}

	} else if strings.Contains(instanceType, "Shared") {
		cost, ok := sharedCPUCosts[instanceType]
		if !ok {
			return 0, false
		}

		// Adjust cost for the cluster size
		if clusterSize == 3 {
			cost *= 2.333

			return cost, true
		}

	}

	return 0, false
}

// GetClusterSize retrieves the cluster size for a specific MySQL instance.
func GetClusterSize(ctx context.Context, client providers.ProviderClient, instanceID int) (int, error) {
	instance, err := client.LinodeClient.GetMySQLDatabase(ctx, instanceID)
	if err != nil {
		return 0, err
	}

	return instance.ClusterSize, nil
}
