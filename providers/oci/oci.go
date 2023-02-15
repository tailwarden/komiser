package oci

import (
	"context"
	"github.com/tailwarden/komiser/providers/oci/iam"
	"github.com/tailwarden/komiser/providers/oci/oracledatabase"
	"log"

	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/oci/compute"
	"github.com/uptrace/bun"
)

func listOfSupportedServices() []providers.FetchDataFunction {
	return []providers.FetchDataFunction{
		compute.Instances,
		iam.Policies,
		oracledatabase.AutonomousDatabases,
	}
}

func FetchResources(ctx context.Context, client providers.ProviderClient, db *bun.DB) {
	for _, fetchResources := range listOfSupportedServices() {
		resources, err := fetchResources(ctx, client)
		if err != nil {
			log.Printf("[%s][OCI] %s", client.Name, err)
		} else {
			for _, resource := range resources {
				db.NewInsert().Model(&resource).Exec(context.Background())
			}
		}
	}
}
