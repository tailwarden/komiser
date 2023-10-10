package alerting

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Alerts(_ context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	projectIds := []string{}
	err := client.OVHClient.Get("/v2/cloud/project", projectIds)
	if err != nil {
		return resources, err
	}

	for _, projectId := range projectIds {
		alertingIds := []string{}
		err = client.OVHClient.Get(fmt.Sprintf("/v2/cloud/project/%s/alerting", projectId), &alertingIds)
		if err != nil {
			return resources, err
		}

		for _, alertingId := range alertingIds {
			alertIds := []string{}
			err = client.OVHClient.Get(fmt.Sprintf("/v2/cloud/project/%s/alerting/%s/alert", projectId, alertingId), &alertIds)
			if err != nil {
				return resources, err
			}

			for _, alertId := range alertIds {
				resources = append(resources, models.Resource{
					Provider:   "OVH",
					Account:    client.Name,
					Service:    "Alerting",
					Region:     client.OVHClient.Endpoint(),
					ResourceId: alertId,
					Cost:       0,
					Name:       alertingId,
					FetchedAt:  time.Now(),
				})
			}
		}
	}

	log.WithFields(log.Fields{
		"provider":  "OVH",
		"account":   client.Name,
		"service":   "Alerting",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
