package droplets

import (
	"context"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/digitalocean/godo"
	. "github.com/mlabouardy/komiser/models"
	"github.com/mlabouardy/komiser/providers"
)

func Droplets(ctx context.Context, client providers.ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	droplets, _, _ := client.DigitalOceanClient.Droplets.List(ctx, &godo.ListOptions{})

	for _, droplet := range droplets {
		tags := make([]Tag, 0)
		for _, tag := range droplet.Tags {
			if strings.Contains(tag, ":") {
				parts := strings.Split(tag, ":")
				tags = append(tags, Tag{
					Key:   parts[0],
					Value: parts[1],
				})
			} else {
				tags = append(tags, Tag{
					Key:   tag,
					Value: tag,
				})
			}
		}

		resources = append(resources, Resource{
			Provider:   "DigitalOcean",
			Account:    client.Name,
			Service:    "Droplet",
			ResourceId: fmt.Sprint("%d", droplet.ID),
			Region:     droplet.Region.Name,
			Name:       droplet.Name,
			Cost:       0,
			Tags:       tags,
			FetchedAt:  time.Now(),
		})
	}

	log.Debugf("[%s] Fetched %d DigitalOcean Droplets\n", client.Name, len(resources))
	return resources, nil
}
