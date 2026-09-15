package storage

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

// Hetzner volumes are billed per GB-month. API exposes no live price,
// so use list-price estimate consistent with Hetzner Cloud pricing pages.
const volumePricePerGBMonth = 0.044

func Volumes(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	volumes, err := client.HetznerClient.Volume.All(ctx)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	nextMonth := monthStart.AddDate(0, 1, 0)
	hoursInMonth := nextMonth.Sub(monthStart).Hours()

	for _, v := range volumes {		tags := make([]models.Tag, 0, len(v.Labels))
		for k, val := range v.Labels {
			tags = append(tags, models.Tag{Key: k, Value: val})
		}

		region := ""
		if v.Location != nil {
			region = v.Location.Name
		}

		monthlyGross := float64(v.Size) * volumePricePerGBMonth
		start := v.Created
		if start.Before(monthStart) {
			start = monthStart
		}
		cost := monthlyGross
		if hoursInMonth > 0 && now.After(start) {
			hours := now.Sub(start).Hours()
			if hours < hoursInMonth {
				cost = monthlyGross * (hours / hoursInMonth)
			}
		}

		resources = append(resources, models.Resource{
			Provider:   "Hetzner",
			Account:    client.Name,
			Service:    "Volume",
			ResourceId: fmt.Sprintf("%d", v.ID),
			Region:     region,
			Name:       v.Name,
			Cost:       cost,
			Tags:       tags,
			FetchedAt:  time.Now(),
			Link:       "https://console.hetzner.cloud/volumes",
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Hetzner",
		"account":   client.Name,
		"service":   "Volume",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
