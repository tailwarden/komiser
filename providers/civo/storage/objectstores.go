package storage

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func ObjectStores(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	stores, err := client.CivoClient.ListObjectStores()
	if err != nil {
		return resources, err
	}

	for _, store := range stores.Items {

		monthlyCost := float64(store.MaxSize/500) * 5

		resources = append(resources, models.Resource{
			Provider:   "Civo",
			Account:    client.Name,
			Service:    "Object Store",
			Region:     client.CivoClient.Region,
			ResourceId: store.ID,
			Cost:       monthlyCost,
			Name:       store.Name,
			FetchedAt:  time.Now(),
			Link:       "https://dashboard.civo.com/object-stores",
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Civo",
		"account":   client.Name,
		"service":   "Object Store",
		"region":    client.CivoClient.Region,
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
