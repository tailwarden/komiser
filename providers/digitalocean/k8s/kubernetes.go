package k8s

import (
	"context"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/digitalocean/godo"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Clusters(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	kubernetesClusters, _, err := client.DigitalOceanClient.Kubernetes.List(ctx, &godo.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, kubernetesCluster := range kubernetesClusters {
		tags := make([]models.Tag, 0)
		for _, tag := range kubernetesCluster.Tags {
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
			Provider:   "DigitalOcean",
			Account:    client.Name,
			Service:    "Kubernetes",
			ResourceId: fmt.Sprintf("%s", kubernetesCluster.ID),
			Region:     kubernetesCluster.RegionSlug,
			Name:       kubernetesCluster.Name,
			Tags:       tags,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://cloud.digitalocean.com/kubernetes/clusters/%s", kubernetesCluster.ID),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "DigitalOcean",
		"account":   client.Name,
		"service":   "Kubernetes",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
