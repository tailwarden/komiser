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

type sslCert struct {
	Properties struct {
		ServiceName string `json:"serviceName"`
		CN          string `json:"commonName"`
	} `json:"properties"`
}

func SSLCertificates(_ context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := []models.Resource{}

	projectIds, err := utils.GetProjects(client)
	if err != nil {
		return resources, err
	}

	for _, projectId := range projectIds {
		sslCertIds := []sslCert{}
		err := client.OVHClient.Get(fmt.Sprintf("/ssl/%s", projectId), &sslCertIds)
		if err != nil {
			return resources, err
		}

		for _, sslCert := range sslCertIds {
			resources = append(resources, models.Resource{
				Provider:   "OVH",
				Account:    client.Name,
				Service:    "SSL",
				Region:     client.OVHClient.Endpoint(),
				ResourceId: sslCert.Properties.ServiceName,
				Cost:       0,
				Name:       sslCert.Properties.CN,
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
