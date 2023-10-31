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

type node struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func Nodes(_ context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := []models.Resource{}

	projectIds, err := utils.GetProjects(client)
	if err != nil {
		return resources, err
	}

	for _, projectId := range projectIds {
		kubeIds := []string{}
		err = client.OVHClient.Get(fmt.Sprintf("/cloud/project/%s/kube", projectId), &kubeIds)
		if err != nil {
			return resources, err
		}

		for _, kubeId := range kubeIds {
			nodes := []node{}
			err = client.OVHClient.Get(fmt.Sprintf("/cloud/project/%s/kube/%s/node", projectId, kubeId), &nodes)
			if err != nil {
				return resources, err
			}

			for _, node := range nodes {
				resources = append(resources, models.Resource{
					Provider:   "OVH",
					Account:    client.Name,
					Service:    "Kube",
					Region:     client.OVHClient.Endpoint(),
					ResourceId: node.Id,
					Cost:       0,
					Name:       node.Name,
					FetchedAt:  time.Now(),
				})
			}
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
