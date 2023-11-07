package image

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/ovh/utils"
)

type image struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func Images(_ context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := []models.Resource{}

	projectIds, err := utils.GetProjects(client)
	if err != nil {
		return resources, err
	}

	for _, projectId := range projectIds {
		images := []image{}
		err = client.OVHClient.Get(fmt.Sprintf("/cloud/project/%s/image", projectId), &images)
		if err != nil {
			return resources, err
		}

		for _, image := range images {
			resources = append(resources, models.Resource{
				Provider:   "OVH",
				Account:    client.Name,
				Service:    "Image",
				Region:     client.OVHClient.Endpoint(),
				ResourceId: image.Id,
				Cost:       0,
				Name:       image.Name,
				FetchedAt:  time.Now(),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "OVH",
		"account":   client.Name,
		"service":   "Image",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
