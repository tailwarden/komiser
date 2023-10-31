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

type sshKey struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func SSHKeys(_ context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := []models.Resource{}

	projectIds, err := utils.GetProjects(client)
	if err != nil {
		return resources, err
	}

	for _, projectId := range projectIds {
		sshKeys := []sshKey{}
		err = client.OVHClient.Get(fmt.Sprintf("/cloud/project/%s/sshkey", projectId), &sshKeys)
		if err != nil {
			return resources, err
		}

		for _, sshKey := range sshKeys {
			resources = append(resources, models.Resource{
				Provider:   "OVH",
				Account:    client.Name,
				Service:    "SSH",
				Region:     client.OVHClient.Endpoint(),
				ResourceId: sshKey.Id,
				Cost:       0,
				Name:       sshKey.Name,
				FetchedAt:  time.Now(),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "OVH",
		"account":   client.Name,
		"service":   "SSH",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
