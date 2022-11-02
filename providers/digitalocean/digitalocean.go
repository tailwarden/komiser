package aws

import (
	"context"
	"log"

	"github.com/digitalocean/godo"
	. "github.com/mlabouardy/komiser/models"
	. "github.com/mlabouardy/komiser/providers/digitalocean/droplets"
	"github.com/uptrace/bun"
)

type FetchDataFunction func(ctx context.Context, client *godo.Client, account string) ([]Resource, error)

func listOfSupportedServices() []FetchDataFunction {
	return []FetchDataFunction{
		Droplets,
	}
}

func FetchDigitalOceanData(ctx context.Context, client *godo.Client, account string, db *bun.DB) {
	for _, function := range listOfSupportedServices() {
		resources, err := function(ctx, client, account)
		if err != nil {
			log.Println(err)
		} else {
			for _, resource := range resources {
				db.NewInsert().Model(&resource).Exec(context.Background())
			}
		}
	}
}
