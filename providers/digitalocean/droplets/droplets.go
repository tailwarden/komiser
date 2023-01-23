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

const createdLayout = "2006-01-02T15:04:05Z" // 2020-07-21T18:37:44Z

func Droplets(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	droplets, _, err := client.DigitalOceanClient.Droplets.List(ctx, &godo.ListOptions{})
	if err != nil {
		return nil, err
	}

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

		hourlyPrice := droplet.Size.PriceHourly

		currentTime := time.Now()
		currentMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, time.UTC)

		created, err := time.Parse(createdLayout, droplet.Created)
		if err != nil {
			return nil, err
		}

		var duration time.Duration
		if created.Before(currentMonth) {
			duration = currentTime.Sub(currentMonth)
		} else {
			duration = currentTime.Sub(created)
		}

		monthlyCost := hourlyPrice * float64(duration.Hours())

		resources = append(resources, models.Resource{
			Provider:   "DigitalOcean",
			Account:    client.Name,
			Service:    "Droplet",
			ResourceId: fmt.Sprintf("%d", droplet.ID),
			Region:     droplet.Region.Name,
			Name:       droplet.Name,
			Cost:       monthlyCost,
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
