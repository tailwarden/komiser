package sql

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/linode/linodego"
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

// Instances fetches SQL instances from the provider and returns them as resources.
func Instances(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	// Fetch instances from the Linode provider
	instances, err := GetInstances(ctx, client)
	if err != nil {
		return resources, err
	}

	for _, instance := range instances {
		tags := make([]models.Tag, 0)
		for _, tag := range instance.Tags {
			if strings.Contains(tag, ":") {
				parts := strings.Split(tag, ":")
				tags = append(tags, models.Tag{
					Key:   parts[0],
					Value: parts[1],
				})
			} else {
				tags = append(tags, models.Tag{
					Key:   tag,
					Value: tag,
				})
			}
		}

		// Calculate the cost based on the instance type
		cost, ok := InstancesCost(instance.Type)
		if !ok {
			log.Warnf("Failed to calculate cost for SQL instance: %s, Type: %s", instance.ID, instance.Type)
		}

		resources = append(resources, models.Resource{
			Provider:   "Linode",
			Account:    client.Name,
			Service:    "SQL",
			Region:     instance.Region,
			ResourceId: instance.ID,
			Cost:       cost,
			Name:       instance.Label,
			FetchedAt:  time.Now(),
			CreatedAt:  instance.Created,
			Tags:       tags,
			Link:       fmt.Sprintf("https://cloud.linode.com/databases/%s", instance.ID),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Linode",
		"account":   client.Name,
		"service":   "SQL",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}

// InstancesCost calculates the cost for a SQL instance based on the instance type.
func InstancesCost(instanceType string) (float64, bool) {
	var cost float64

	if strings.HasPrefix(instanceType, "Dedicated") {
		cost, ok := dedicatedCPUCosts[instanceType]
		if !ok {
			return 0, false
		}
	} else if strings.HasPrefix(instanceType, "Shared") {
		cost, ok := sharedCPUCosts[instanceType]
		if !ok {
			return 0, false
		}
	} else {
		return 0, false
	}

	return cost, true
}

// GetInstances fetches SQL instances from the Linode provider.
func GetInstances(ctx context.Context, client providers.ProviderClient) ([]linodego.Instance, error) {
	// Use the Linode provider client to fetch instances
	instances, err := client.LinodeClient.ListInstances(ctx, &linodego.ListOptions{})
	if err != nil {
		return nil, err
	}
	return instances, nil
}
