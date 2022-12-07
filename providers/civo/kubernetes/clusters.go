package kubernetes

import (
	"context"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/mlabouardy/komiser/models"
	"github.com/mlabouardy/komiser/providers"
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

		resources = append(resources, models.Resource{
			Provider:   "Civo",
			Account:    client.Name,
			Service:    "Kubernetes",
			Region:     client.CivoClient.Region,
			ResourceId: cluster.ID,
			Cost:       0,
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
