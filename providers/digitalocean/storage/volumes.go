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

const createdLayout = "2006-01-02T15:04:05Z"

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
		
		var hourlyPrice float64
		sizeInGB := volume.SizeGigaBytes
		if sizeInGB <= 100 {
			hourlyPrice = 0.015
		} else if sizeInGB <= 500 {
			hourlyPrice = 0.075
		} else {
			hourlyPrice = 0.150
		}

		currentTime := time.Now()
		currentMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, time.UTC)

		created, err := time.Parse(createdLayout, volume.CreatedAt.String())
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
			Service:    "Volume",
			ResourceId: volume.ID,
			Region:     volume.Region.Name,
			Name:       volume.Name,
			Cost:       monthlyCost,
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
