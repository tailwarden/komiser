package databases

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

func Databases(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	databases, _, err := client.DigitalOceanClient.Databases.List(ctx, &godo.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, database := range databases {
		tags := make([]models.Tag, 0)
		for _, tag := range database.Tags {
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
			Service:    "Database",
			ResourceId: fmt.Sprintf("%s", database.ID),
			Region:     database.RegionSlug,
			Name:       database.Name,
			Tags:       tags,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://cloud.digitalocean.com/databases/%s", database.ID),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "DigitalOcean",
		"account":   client.Name,
		"service":   "Database",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
