package compute

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/civo/civogo"
	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Instances(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	instances, err := client.CivoClient.ListAllInstances()
	if err != nil {
		return resources, err
	}

	for _, resource := range instances {
		tags := make([]models.Tag, 0)

		for _, tag := range resource.Tags {
			if strings.Contains(tag, ":") {
				parts := strings.Split(tag, ":")
				tags = append(tags, models.Tag{
					Key:   parts[0],
					Value: parts[1],
				})
			} else {
				tags = append(tags, models.Tag{
					Key:   tag,
					Value: tag,
				})
			}
		}

		hourlyPrice := float64(resource.RAMMegabytes/1024) * 0.007440

		currentTime := time.Now()
		currentMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, time.UTC)
		var duration time.Duration
		if resource.CreatedAt.Before(currentMonth) {
			duration = currentTime.Sub(currentMonth)
		} else {
			duration = currentTime.Sub(resource.CreatedAt)
		}

		relations := getComputeRelations(resource) 

		monthlyCost := hourlyPrice * float64(duration.Hours())

		resources = append(resources, models.Resource{
			Provider:   "Civo",
			Account:    client.Name,
			Service:    "Compute",
			Region:     client.CivoClient.Region,
			ResourceId: resource.ID,
			Cost:       monthlyCost,
			Name:       resource.Hostname,
			FetchedAt:  time.Now(),
			CreatedAt:  resource.CreatedAt,
			Tags:       tags,
			Relations: relations,
			Link:       fmt.Sprintf("https://dashboard.civo.com/instances/%s", resource.ID),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Civo",
		"account":   client.Name,
		"service":   "Compute",
		"region":    client.CivoClient.Region,
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}

func getComputeRelations(compute civogo.Instance) []models.Link {
	return []models.Link{
		{
			ResourceID: compute.NetworkID,
			Type: "Network",
			Name: compute.NetworkID, //cannot get the name of the network unless calling the network api
			Relation: "USES",
		},
	}
}
