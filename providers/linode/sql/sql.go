package sql

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/linode/linodego"
	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

// Cost data for Dedicated CPU instances
var dedicatedCPUCosts = map[string]map[string]float64{
	"Dedicated 4GB": {
		"1 Node": 65.00,
		"3 Node": 195.00,
	},
	"Dedicated 8GB": {
		"1 Node": 130.00,
		"3 Node": 390.00,
	},
	"Dedicated 16GB": {
		"1 Node": 260.00,
		"3 Node": 780.00,
	},
	"Dedicated 32GB": {
		"1 Node": 520.00,
		"3 Node": 1560.00,
	},
	"Dedicated 64GB": {
		"1 Node": 1040.00,
		"3 Node": 3120.00,
	},
	"Dedicated 96GB": {
		"1 Node": 1560.00,
		"3 Node": 4680.00,
	},
	"Dedicated 128GB": {
		"1 Node": 2080.00,
		"3 Node": 6240.00,
	},
	"Dedicated 256GB": {
		"1 Node": 4160.00,
		"3 Node": 12480.00,
	},
	"Dedicated 512GB": {
		"1 Node": 8320.00,
		"3 Node": 24960.00,
	},
}

// Cost data for Shared CPU instances
var sharedCPUCosts = map[string]map[string]float64{
	"Shared 1GB": {
		"1 Node":  15.00,
		"3 Node":  35.00,
	},
	"Shared 2GB": {
		"1 Node":  30.00,
		"3 Node":  70.00,
	},
	"Shared 4GB": {
		"1 Node":  60.00,
		"3 Node":  140.00,
	},
	"Shared 8GB": {
		"1 Node":  120.00,
		"3 Node":  280.00,
	},
	"Shared 16GB": {
		"1 Node":  240.00,
		"3 Node":  560.00,
	},
	"Shared 32GB": {
		"1 Node":  480.00,
		"3 Node":  1120.00,
	},
	"Shared 64GB": {
		"1 Node":  960.00,
		"3 Node":  2240.00,
	},
	"Shared 96GB": {
		"1 Node":  1440.00,
		"3 Node":  3360.00,
	},
	"Shared 128GB": {
		"1 Node":  1920.00,
		"3 Node":  4480.00,
	},
	"Shared 192GB": {
		"1 Node":  2880.00,
		"3 Node":  6720.00,
	},
	"Shared 256GB": {
		"1 Node":  3840.00,
		"3 Node":  8960.00,
	},
}

// Instances fetches SQL instances from the provider and returns them as resources.
func Instances(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	instances, err := client.SQLClient.GetInstances(ctx)
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

		// Calculate the cost based on the instance type and node count
		cost, ok := InstancesCost(instance.Type, instance.NodeCount)
		if !ok {
			log.Warnf("Failed to calculate cost for SQL instance: %s, Type: %s, NodeCount: %d", instance.ID, instance.Type, instance.NodeCount)
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

// InstancesCost calculates the cost for a SQL instance based on the instance type and node count.
func InstancesCost(instanceType string, nodeCount int) (float64, bool) {
	var costs map[string]map[string]float64

	if strings.HasPrefix(instanceType, "Dedicated") {
		costs = dedicatedCPUCosts
	} else if strings.HasPrefix(instanceType, "Shared") {
		costs = sharedCPUCosts
	} else {
		return 0, false
	}

	costMap, ok := costs[instanceType]
	if !ok {
		return 0, false
	}

	cost, ok := costMap[fmt.Sprintf("%d Node", nodeCount)]
	if !ok {
		return 0, false
	}

	return cost, true
}
