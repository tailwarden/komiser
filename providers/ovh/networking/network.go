package networking

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/ovh/utils"
)

type network struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func PublicNetworks(_ context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := []models.Resource{}

	projectIds, err := utils.GetProjects(client)
	if err != nil {
		return resources, err
	}

	for _, projectId := range projectIds {
		networks := []network{}
		err = client.OVHClient.Get(fmt.Sprintf("/cloud/project/%s/network/public", projectId), &networks)
		if err != nil {
			return resources, err
		}

		for _, network := range networks {
			resources = append(resources, models.Resource{
				Provider:   "OVH",
				Account:    client.Name,
				Service:    "Network",
				Region:     client.OVHClient.Endpoint(),
				ResourceId: network.Id,
				Cost:       0,
				Name:       network.Name,
				FetchedAt:  time.Now(),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "OVH",
		"account":   client.Name,
		"service":   "Network",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}

func PrivateNetworks(_ context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := []models.Resource{}

	projectIds, err := utils.GetProjects(client)
	if err != nil {
		return resources, err
	}

	for _, projectId := range projectIds {
		networks := []network{}
		err = client.OVHClient.Get(fmt.Sprintf("/cloud/project/%s/network/private", projectId), &networks)
		if err != nil {
			return resources, err
		}

		for _, network := range networks {
			resources = append(resources, models.Resource{
				Provider:   "OVH",
				Account:    client.Name,
				Service:    "Network",
				Region:     client.OVHClient.Endpoint(),
				ResourceId: network.Id,
				Cost:       0,
				Name:       network.Name,
				FetchedAt:  time.Now(),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "OVH",
		"account":   client.Name,
		"service":   "Network",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
