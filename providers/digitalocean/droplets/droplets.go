package droplets

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/digitalocean/godo"
	. "github.com/mlabouardy/komiser/models"
	. "github.com/mlabouardy/komiser/providers"
)

func Droplets(ctx context.Context, client ProviderClient) ([]Resource, error) {
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

	log.Printf("[%s] Fetched %d DigitalOcean Droplets\n", client.Name, len(resources))
	return resources, nil
}
