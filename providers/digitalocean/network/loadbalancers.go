package network

import (
	"context"
	"time"

	"github.com/digitalocean/godo"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

const (
	createdLayout string = "2006-01-02T15:04:05Z"
	// https://docs.digitalocean.com/products/networking/load-balancers/details/pricing/
	monthlyPrice float64 = 12
)

func LoadBalancers(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	lbs, _, err := client.DigitalOceanClient.LoadBalancers.List(ctx, &godo.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, lb := range lbs {
		currentTime := time.Now()
		currentMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, time.UTC)
		nextMonth := time.Date(currentTime.Year(), currentTime.Month()+1, 1, 0, 0, 0, 0, time.UTC)
		hoursInMonth := nextMonth.Sub(currentMonth).Hours()
		hourlyPrice := monthlyPrice / hoursInMonth

		created, err := time.Parse(createdLayout, lb.Created)
		if err != nil {
			return nil, err
		}

		duration := currentTime.Sub(created)
		if created.Before(currentMonth) {
			duration = currentTime.Sub(currentMonth)
		}

		monthlyCost := hourlyPrice * float64(duration.Hours())

		resources = append(resources, models.Resource{
			Provider:   "DigitalOcean",
			Account:    client.Name,
			Service:    "Load Balancer",
			Region:     lb.Region.Name,
			ResourceId: lb.ID,
			Cost:       monthlyCost,
			Name:       lb.Name,
			FetchedAt:  time.Now(),
			Link:       "https://cloud.digitalocean.com/networking/load_balancers",
		})
	}

	log.WithFields(log.Fields{
		"provider":  "DigitalOcean",
		"account":   client.Name,
		"service":   "Load Balancer",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
