package storage

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/ovh/utils"
)

type volume struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func Volumes(_ context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := []models.Resource{}

	projectIds, err := utils.GetProjects(client)
	if err != nil {
		return resources, err
	}

	for _, projectId := range projectIds {
		volumes := []volume{}
		err = client.OVHClient.Get(fmt.Sprintf("/cloud/project/%s/volume", projectId), &volumes)
		if err != nil {
			return resources, err
		}

		for _, volume := range volumes {
			resources = append(resources, models.Resource{
				Provider:   "OVH",
				Account:    client.Name,
				Service:    "Volume",
				Region:     client.OVHClient.Endpoint(),
				ResourceId: volume.Id,
				Cost:       0,
				Name:       volume.Name,
				FetchedAt:  time.Now(),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "OVH",
		"account":   client.Name,
		"service":   "Volume",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
