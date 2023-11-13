package kube

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/ovh/utils"
)

func Clusters(_ context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := []models.Resource{}

	projectIds, err := utils.GetProjects(client)
	if err != nil {
		return resources, err
	}

	for _, projectId := range projectIds {
		clusterIds := []string{}
		err = client.OVHClient.Get(fmt.Sprintf("/cloud/project/%s/kube", projectId), &clusterIds)
		if err != nil {
			return resources, err
		}

		for _, clusterId := range clusterIds {
			resources = append(resources, models.Resource{
				Provider:   "OVH",
				Account:    client.Name,
				Service:    "Kube",
				Region:     client.OVHClient.Endpoint(),
				ResourceId: clusterId,
				Cost:       0,
				Name:       clusterId,
				FetchedAt:  time.Now(),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "OVH",
		"account":   client.Name,
		"service":   "Kube",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
