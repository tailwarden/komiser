package droplets

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

func Droplets(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	droplets, _, _ := client.DigitalOceanClient.Droplets.List(ctx, &godo.ListOptions{})

	for _, droplet := range droplets {
		tags := make([]models.Tag, 0)
		for _, tag := range droplet.Tags {
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
			Service:    "Droplet",
			ResourceId: fmt.Sprintf("%d", droplet.ID),
			Region:     droplet.Region.Name,
			Name:       droplet.Name,
			Cost:       0,
			Tags:       tags,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://cloud.digitalocean.com/droplets/%d/graphs", droplet.ID),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "DigitalOcean",
		"account":   client.Name,
		"service":   "Droplet",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
