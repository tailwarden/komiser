package droplets

import (
	"context"
	"log"
	"time"

	"github.com/digitalocean/godo"
	. "github.com/mlabouardy/komiser/models"
	. "github.com/mlabouardy/komiser/providers"
)

func Droplets(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	droplets, _, _ := client.DigitalOceanClient.Droplets.List(ctx, &godo.ListOptions{})

	for _, droplet := range droplets {
		resources = append(resources, Resource{
			Provider:  "DigitalOcean",
			Account:   client.Name,
			Service:   "Droplet",
			Region:    droplet.Region.Name,
			Name:      droplet.Name,
			Cost:      0,
			FetchedAt: time.Now(),
		})
	}

	log.Printf("[%s] Fetched %d DigitalOcean Droplets\n", client.Name, len(resources))
	return resources, nil
}
