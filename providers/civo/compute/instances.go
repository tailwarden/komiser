package compute

import (
	"context"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/mlabouardy/komiser/models"
	"github.com/mlabouardy/komiser/providers"
)

func Instances(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	instances, err := client.CivoClient.ListAllInstances()
	if err != nil {
		return resources, err
	}

	for _, resource := range instances {
		tags := make([]models.Tag, 0)

		for _, tag := range resource.Tags {
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

		hourlyPrice := getInstancePrice(int(resource.RAMMegabytes/1024), resource.CPUCores)

		currentTime := time.Now()
		currentMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, time.UTC)
		var duration time.Duration
		if resource.CreatedAt.Before(currentMonth) {
			duration = currentTime.Sub(currentMonth)
		} else {
			duration = currentTime.Sub(resource.CreatedAt)
		}

		monthlyCost := hourlyPrice * float64(duration.Hours())

		resources = append(resources, models.Resource{
			Provider:   "Civo",
			Account:    client.Name,
			Service:    "Compute",
			Region:     client.CivoClient.Region,
			ResourceId: resource.ID,
			Cost:       monthlyCost,
			Name:       resource.Hostname,
			FetchedAt:  time.Now(),
			CreatedAt:  resource.CreatedAt,
			Tags:       tags,
			Link:       fmt.Sprintf("https://dashboard.civo.com/instances/%s", resource.ID),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Civo",
		"account":   client.Name,
		"service":   "Compute",
		"region":    client.CivoClient.Region,
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}

func getInstancePrice(ramSize int, cpuCores int) float64 {
	hourlyRate := 0.0
	if ramSize <= 2 {
		if cpuCores <= 1 {
			hourlyRate = 0.01
		} else if cpuCores <= 2 {
			hourlyRate = 0.02
		} else {
			hourlyRate = 0.04
		}
	} else if ramSize <= 4 {
		if cpuCores <= 1 {
			hourlyRate = 0.02
		} else if cpuCores <= 2 {
			hourlyRate = 0.04
		} else {
			hourlyRate = 0.08
		}
	} else if ramSize <= 8 {
		if cpuCores <= 1 {
			hourlyRate = 0.04
		} else if cpuCores <= 2 {
			hourlyRate = 0.08
		} else {
			hourlyRate = 0.16
		}
	} else if ramSize <= 16 {
		if cpuCores <= 1 {
			hourlyRate = 0.08
		} else if cpuCores <= 2 {
			hourlyRate = 0.16
		} else {
			hourlyRate = 0.32
		}
	} else {
		if cpuCores <= 1 {
			hourlyRate = 0.16
		} else if cpuCores <= 2 {
			hourlyRate = 0.32
		} else {
			hourlyRate = 0.64
		}
	}
	return hourlyRate
}
