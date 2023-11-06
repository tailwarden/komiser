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

type container struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func Containers(_ context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := []models.Resource{}

	projectIds, err := utils.GetProjects(client)
	if err != nil {
		return resources, err
	}

	for _, projectId := range projectIds {
		containers := []container{}
		err = client.OVHClient.Get(fmt.Sprintf("/cloud/project/%s/storage", projectId), &containers)
		if err != nil {
			return resources, err
		}

		for _, container := range containers {
			resources = append(resources, models.Resource{
				Provider:   "OVH",
				Account:    client.Name,
				Service:    "Container",
				Region:     client.OVHClient.Endpoint(),
				ResourceId: container.Id,
				Cost:       0,
				Name:       container.Name,
				FetchedAt:  time.Now(),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "OVH",
		"account":   client.Name,
		"service":   "Container",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
