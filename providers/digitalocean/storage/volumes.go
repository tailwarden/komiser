package storage

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

func Volumes(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	volumes, _, err := client.DigitalOceanClient.Storage.ListVolumes(ctx, &godo.ListVolumeParams{})
	if err != nil {
		return nil, err
	}

	for _, volume := range volumes {
		tags := make([]models.Tag, 0)
		for _, tag := range volume.Tags {
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
			Service:    "Volume",
			ResourceId: fmt.Sprintf("%s", volume.ID),
			Region:     volume.Region.Name,
			Name:       volume.Name,
			Tags:       tags,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://cloud.digitalocean.com/volumes/%s", volume.ID),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "DigitalOcean",
		"account":   client.Name,
		"service":   "Volume",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
