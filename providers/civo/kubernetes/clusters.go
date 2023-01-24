package kubernetes

import (
	"context"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Clusters(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	clusters, err := client.CivoClient.ListKubernetesClusters()
	if err != nil {
		return resources, err
	}

	for _, cluster := range clusters.Items {
		tags := make([]models.Tag, 0)

		for _, tag := range cluster.Tags {
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

		monthlyCost := 0.0
		for _, instance := range cluster.Instances {
			hourlyPrice := 0.0

			for _, pool := range cluster.Pools {
				for _, poolInstance := range pool.Instances {
					if poolInstance.ID == instance.ID {
						if strings.Contains(pool.Size, "g4s") {
							// general purpose
							hourlyPrice = float64(instance.RAMMegabytes/1024) * 0.007440
						} else if strings.Contains(pool.Size, "g4p") {
							// performance purpose
							hourlyPrice = float64(instance.RAMMegabytes/1024) * 0.119048
						} else if strings.Contains(pool.Size, "g4c") {
							// CPU optimized
							hourlyPrice = float64(instance.RAMMegabytes/1024) * 0.190476
						} else if strings.Contains(pool.Size, "g4r") {
							// CPU optimized
							hourlyPrice = float64(instance.RAMMegabytes/1024) * 0.107143
						}
					}
				}
			}

			currentTime := time.Now()
			currentMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, time.UTC)
			var duration time.Duration
			if instance.CreatedAt.Before(currentMonth) {
				duration = currentTime.Sub(currentMonth)
			} else {
				duration = currentTime.Sub(instance.CreatedAt)
			}

			instanceMonthlyCost := hourlyPrice * float64(duration.Hours())

			monthlyCost += instanceMonthlyCost
		}

		resources = append(resources, models.Resource{
			Provider:   "Civo",
			Account:    client.Name,
			Service:    "Kubernetes",
			Region:     client.CivoClient.Region,
			ResourceId: cluster.ID,
			Cost:       monthlyCost,
			Name:       cluster.Name,
			Tags:       tags,
			FetchedAt:  time.Now(),
			CreatedAt:  cluster.CreatedAt,
			Link:       fmt.Sprintf("https://dashboard.civo.com/kubernetes/%s", cluster.ID),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Civo",
		"account":   client.Name,
		"service":   "Kubernetes",
		"region":    client.CivoClient.Region,
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
