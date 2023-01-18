package network

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Networks(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	networks, err := client.CivoClient.ListNetworks()
	if err != nil {
		return resources, err
	}

	for _, network := range networks {
		resources = append(resources, models.Resource{
			Provider:   "Civo",
			Account:    client.Name,
			Service:    "Network",
			Region:     client.CivoClient.Region,
			ResourceId: network.ID,
			Cost:       0,
			Name:       network.Name,
			FetchedAt:  time.Now(),
			Link:       "https://dashboard.civo.com/networks",
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Civo",
		"account":   client.Name,
		"service":   "Network",
		"region":    client.CivoClient.Region,
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
