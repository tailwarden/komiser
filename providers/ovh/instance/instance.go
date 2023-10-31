package instance

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/ovh/utils"
)

type instance struct {
	Properties struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"properties"`
}

func Instances(_ context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := []models.Resource{}

	projectIds, err := utils.GetProjects(client)
	if err != nil {
		return resources, err
	}

	for _, projectId := range projectIds {
		instances := []instance{}
		err = client.OVHClient.Get(fmt.Sprintf("/cloud/project/%s/instance", projectId), &instances)
		if err != nil {
			return resources, err
		}

		for _, instance := range instances {
			resources = append(resources, models.Resource{
				Provider:   "OVH",
				Account:    client.Name,
				Service:    "Instance",
				Region:     client.OVHClient.Endpoint(),
				ResourceId: instance.Properties.Id,
				Cost:       0,
				Name:       instance.Properties.Name,
				FetchedAt:  time.Now(),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "OVH",
		"account":   client.Name,
		"service":   "Instance",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
