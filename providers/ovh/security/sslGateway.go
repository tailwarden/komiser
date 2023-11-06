package security

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/ovh/utils"
)

type sslGateway struct {
	ServiceName string `json:"serviceName"`
	DisplayName string `json:"displayName"`
}

func SSLGateways(_ context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := []models.Resource{}

	projectIds, err := utils.GetProjects(client)
	if err != nil {
		return resources, err
	}

	for _, projectId := range projectIds {
		sslGateways := []sslGateway{}
		err = client.OVHClient.Get(fmt.Sprintf("/sslGateway/%s", projectId), &sslGateways)
		if err != nil {
			return resources, err
		}

		for _, sslGateway := range sslGateways {
			resources = append(resources, models.Resource{
				Provider:   "OVH",
				Account:    client.Name,
				Service:    "SSL",
				Region:     client.OVHClient.Endpoint(),
				ResourceId: sslGateway.ServiceName,
				Cost:       0,
				Name:       sslGateway.DisplayName,
				FetchedAt:  time.Now(),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "OVH",
		"account":   client.Name,
		"service":   "SSL",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
