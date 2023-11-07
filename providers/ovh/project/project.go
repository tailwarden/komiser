package project

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/ovh/utils"
)

type project struct {
	ProjectId   string `json:"project_id"`
	ProjectName string `json:"projectName"`
}

func Projects(_ context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := []models.Resource{}

	projectIds, err := utils.GetProjects(client)
	if err != nil {
		return resources, err
	}

	for _, projectId := range projectIds {
		projects := []project{}
		err = client.OVHClient.Get(fmt.Sprintf("/cloud/project/%s", projectId), &projects)
		if err != nil {
			return resources, err
		}

		for _, project := range projects {
			resources = append(resources, models.Resource{
				Provider:   "OVH",
				Account:    client.Name,
				Service:    "Project",
				Region:     client.OVHClient.Endpoint(),
				ResourceId: project.ProjectId,
				Cost:       0,
				Name:       project.ProjectName,
				FetchedAt:  time.Now(),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "OVH",
		"account":   client.Name,
		"service":   "Project",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
