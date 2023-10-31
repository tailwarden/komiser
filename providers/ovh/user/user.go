package user

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/ovh/utils"
)

type user struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func Users(_ context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := []models.Resource{}

	projectIds, err := utils.GetProjects(client)
	if err != nil {
		return resources, err
	}

	for _, projectId := range projectIds {
		users := []user{}
		err = client.OVHClient.Get(fmt.Sprintf("/cloud/project/%s/user", projectId), &users)
		if err != nil {
			return resources, err
		}

		for _, user := range users {
			resources = append(resources, models.Resource{
				Provider:   "OVH",
				Account:    client.Name,
				Service:    "User",
				Region:     client.OVHClient.Endpoint(),
				ResourceId: user.Id,
				Cost:       0,
				Name:       user.Username,
				FetchedAt:  time.Now(),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "OVH",
		"account":   client.Name,
		"service":   "User",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
