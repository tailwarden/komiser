package digitalocean

import (
	"context"
	"github.com/tailwarden/komiser/providers/digitalocean/k8s"
	"log"

	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/digitalocean/droplets"
	"github.com/tailwarden/komiser/providers/digitalocean/network"
	"github.com/uptrace/bun"
)

func listOfSupportedServices() []providers.FetchDataFunction {
	return []providers.FetchDataFunction{
		droplets.Droplets,
		network.Firewalls,
		network.LoadBalancers,
		network.Vpcs,
		k8s.Clusters,
	}
}

func FetchResources(ctx context.Context, client providers.ProviderClient, db *bun.DB) {
	for _, fetchResources := range listOfSupportedServices() {
		resources, err := fetchResources(ctx, client)
		if err != nil {
			log.Printf("[%s][DigitalOcean] %s", client.Name, err)
		} else {
			for _, resource := range resources {
				db.NewInsert().Model(&resource).On("CONFLICT (resource_id) DO UPDATE").Set("link = EXCLUDED.link").Exec(context.Background())
			}
		}
	}
}
