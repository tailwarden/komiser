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

type ip struct {
	Id string `json:"id"`
	IP string `json:"ip"`
}

func IPs(_ context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := []models.Resource{}

	projectIds, err := utils.GetProjects(client)
	if err != nil {
		return resources, err
	}

	for _, projectId := range projectIds {
		ips := []ip{}
		err = client.OVHClient.Get(fmt.Sprintf("/cloud/project/%s/ip", projectId), &ips)
		if err != nil {
			return resources, err
		}

		for _, ip := range ips {
			resources = append(resources, models.Resource{
				Provider:   "OVH",
				Account:    client.Name,
				Service:    "IP",
				Region:     client.OVHClient.Endpoint(),
				ResourceId: ip.Id,
				Cost:       0,
				Name:       ip.IP,
				FetchedAt:  time.Now(),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "OVH",
		"account":   client.Name,
		"service":   "IP",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}

func FailoverIPs(_ context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := []models.Resource{}

	projectIds, err := utils.GetProjects(client)
	if err != nil {
		return resources, err
	}

	for _, projectId := range projectIds {
		ips := []ip{}
		err = client.OVHClient.Get(fmt.Sprintf("/cloud/project/%s/ip/failover", projectId), &ips)
		if err != nil {
			return resources, err
		}

		for _, ip := range ips {
			resources = append(resources, models.Resource{
				Provider:   "OVH",
				Account:    client.Name,
				Service:    "IP",
				Region:     client.OVHClient.Endpoint(),
				ResourceId: ip.Id,
				Cost:       0,
				Name:       ip.IP,
				FetchedAt:  time.Now(),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "OVH",
		"account":   client.Name,
		"service":   "IP",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
