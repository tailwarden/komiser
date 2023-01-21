package network

import (
	"context"
	"strings"
	"time"

	"github.com/digitalocean/godo"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Firewalls(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	regions := make([]string, 0)

	firewalls, _, err := client.DigitalOceanClient.Firewalls.List(ctx, &godo.ListOptions{})
	if err != nil {
		return resources, err
	}

	for _, firewall := range firewalls {
		for _, id := range firewall.DropletIDs {
			droplet, _, _ := client.DigitalOceanClient.Droplets.Get(ctx, id)
			regions = append(regions, droplet.Region.Name)
		}
		resources = append(resources, models.Resource{
			Provider:   "DigitalOcean",
			Account:    client.Name,
			Service:    "Firewall",
			Region:     strings.Join(regions, ", "),
			ResourceId: firewall.ID,
			Cost:       0,
			Name:       firewall.Name,
			FetchedAt:  time.Now(),
			Link:       "https://cloud.digitalocean.com/networking/firewalls",
		})
	}

	log.WithFields(log.Fields{
		"provider":  "DigitalOcean",
		"account":   client.Name,
		"service":   "Firewall",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
