package droplets

import (
	"context"
	"log"
	"time"

	"github.com/digitalocean/godo"
	. "github.com/mlabouardy/komiser/models"
)

func Droplets(ctx context.Context, client *godo.Client, account string) ([]Resource, error) {
	resources := make([]Resource, 0)
	droplets, _, _ := client.Droplets.List(ctx, &godo.ListOptions{})

	for _, droplet := range droplets {
		resources = append(resources, Resource{
			Provider:  "DigitalOcean",
			Account:   account,
			Service:   "Droplet",
			Region:    droplet.Region.Name,
			Name:      droplet.Name,
			Cost:      0,
			FetchedAt: time.Now(),
		})
	}

	log.Printf("[%s] Fetched %d DigitalOcean Droplets\n", account, len(resources))
	return resources, nil
}
