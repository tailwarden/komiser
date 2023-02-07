package network

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func LoadBalancers(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	lbs, err := client.CivoClient.ListLoadBalancers()
	if err != nil {
		return resources, err
	}

	for _, lb := range lbs {
		resources = append(resources, models.Resource{
			Provider:   "Civo",
			Account:    client.Name,
			Service:    "Load Balancer",
			Region:     client.CivoClient.Region,
			ResourceId: lb.ID,
			Cost:       10,
			Name:       lb.Name,
			FetchedAt:  time.Now(),
			Link:       "https://dashboard.civo.com/loadbalancers",
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Civo",
		"account":   client.Name,
		"service":   "Load Balancer",
		"region":    client.CivoClient.Region,
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
