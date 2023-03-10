package storage

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

func Volumes(ctx context.Context, client providers.ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)

	volumes, err := client.LinodeClient.ListVolumes(ctx, &linodego.ListOptions{})
	if err != nil {
		return resources, err
	}

	for _, volume := range volumes {
		tags := make([]Tag, 0)
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
			Provider:   "Linode",
			Account:    client.Name,
			Service:    "Volume",
			Region:     volume.Region,
			ResourceId: fmt.Sprintf("%d", volume.ID),
			Cost:       0,
			Name:       volume.Label,
			FetchedAt:  time.Now(),
			CreatedAt:  *volume.Created,
			Tags:       tags,
			Link:       "https://cloud.linode.com/volumes",
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Linode",
		"account":   client.Name,
		"service":   "Volume",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
