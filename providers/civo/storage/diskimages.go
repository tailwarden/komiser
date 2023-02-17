package storage

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func DiskImages(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	diskImages, err := client.CivoClient.ListDiskImages()
	if err != nil {
		return resources, err
	}

	for _, resource := range diskImages {
		resources = append(resources, models.Resource{
			Provider:   "Civo",
			Account:    client.Name,
			Service:    "DiskImage",
			Region:     client.CivoClient.Region,
			ResourceId: resource.ID,
			Name:       resource.Name,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://dashboard.civo.com/diskimages/%s", resource.ID),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Civo",
		"account":   client.Name,
		"service":   "DiskImage",
		"region":    client.CivoClient.Region,
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
