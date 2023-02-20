package functions

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Namespaces(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	namespaces, _, err := client.DigitalOceanClient.Functions.ListNamespaces(ctx)
	if err != nil {
		return nil, err
	}

	for _, namespace := range namespaces {
		// triggers are what will invoke functions and cause to charge on account.
		// add triggers to namespaces [containing namespaces] too, as triggers are sub-resources to namespaces.
		allTriggers, err := Triggers(ctx, client, namespace)
		if err != nil {
			return nil, err
		}
		if len(allTriggers) > 0 {
			resources = append(resources, allTriggers...)
		}

		resources = append(resources, models.Resource{
			Provider:   "DigitalOcean",
			Account:    client.Name,
			Service:    "Namespace",
			ResourceId: fmt.Sprintf("%s", namespace.UUID),
			Region:     namespace.Region,
			Name:       namespace.Namespace,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://cloud.digitalocean.com/functions/%s", namespace.UUID),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "DigitalOcean",
		"account":   client.Name,
		"service":   "Namespace",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
