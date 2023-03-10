package network

import (
	"context"
	"time"

	"github.com/digitalocean/godo"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Vpcs(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	regionsMap := make(map[string]string)

	regions, _, err := client.DigitalOceanClient.Regions.List(ctx, &godo.ListOptions{})
	if err != nil {
		log.Warnf("[%s][DigitalOcean] Couldn't fetch the list of regions: %s", client.Name, err)
	}
	for _, region := range regions {
		regionsMap[region.Slug] = region.Name
	}

	vpcs, _, err := client.DigitalOceanClient.VPCs.List(ctx, &godo.ListOptions{})
	if err != nil {
		return resources, err
	}

	for _, vpc := range vpcs {
		region := vpc.RegionSlug
		if val, ok := regionsMap[region]; ok {
			region = val
		}
		resources = append(resources, models.Resource{
			Provider:   "DigitalOcean",
			Account:    client.Name,
			Service:    "VPC",
			Region:     region,
			ResourceId: vpc.ID,
			Cost:       0,
			Name:       vpc.Name,
			FetchedAt:  time.Now(),
			Link:       "https://cloud.digitalocean.com/networking/vpc",
		})
	}

	log.WithFields(log.Fields{
		"provider":  "DigitalOcean",
		"account":   client.Name,
		"service":   "VPC",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
