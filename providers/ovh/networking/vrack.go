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

type vrack struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func Vracks(_ context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := []models.Resource{}

	projectIds, err := utils.GetProjects(client)
	if err != nil {
		return resources, err
	}

	for _, projectId := range projectIds {
		vracks := []vrack{}
		err = client.OVHClient.Get(fmt.Sprintf("/cloud/project/%s/vrack", projectId), &vracks)
		if err != nil {
			return resources, err
		}

		for _, vrack := range vracks {
			resources = append(resources, models.Resource{
				Provider:   "OVH",
				Account:    client.Name,
				Service:    "Vrack",
				Region:     client.OVHClient.Endpoint(),
				ResourceId: vrack.Id,
				Cost:       0,
				Name:       vrack.Name,
				FetchedAt:  time.Now(),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "OVH",
		"account":   client.Name,
		"service":   "Vrack",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
