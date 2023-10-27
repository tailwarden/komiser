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

type sshKey struct {
	Properties struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"properties"`
}

func SSHKeys(_ context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := []models.Resource{}

	projectIds, err := utils.GetProjects(client)
	if err != nil {
		return resources, err
	}

	for _, projectId := range projectIds {
		sshKeys := []sshKey{}
		err = client.OVHClient.Get(fmt.Sprintf("/v2/cloud/project/%s/sshkey", projectId), &sshKeys)
		if err != nil {
			return resources, err
		}

		for _, sshKey := range sshKeys {
			resources = append(resources, models.Resource{
				Provider:   "OVH",
				Account:    client.Name,
				Service:    "Instance",
				Region:     client.OVHClient.Endpoint(),
				ResourceId: sshKey.Properties.Id,
				Cost:       0,
				Name:       sshKey.Properties.Name,
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
