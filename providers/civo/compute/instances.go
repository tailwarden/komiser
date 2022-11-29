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

		resources = append(resources, models.Resource{
			Provider:   "Civo",
			Account:    client.Name,
			Service:    "Compute",
			Region:     resource.Region,
			ResourceId: resource.ID,
			Cost:       0,
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
