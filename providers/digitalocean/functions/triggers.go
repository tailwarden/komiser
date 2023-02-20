package functions

import (
	"context"
	"fmt"
	"github.com/digitalocean/godo"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Triggers(ctx context.Context, client providers.ProviderClient, namespace godo.FunctionsNamespace) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	triggers, _, err := client.DigitalOceanClient.Functions.ListTriggers(ctx, namespace.Namespace)
	if err != nil {
		return nil, err
	}

	for _, trigger := range triggers {
		resources = append(resources, models.Resource{
			Provider:   "DigitalOcean",
			Account:    client.Name,
			Service:    "Trigger",
			ResourceId: fmt.Sprintf("%s", trigger.Name),
			Region:     namespace.Region,
			Name:       trigger.Name,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://cloud.digitalocean.com/functions/%s/_/%s/triggers", namespace.Namespace, trigger.Function),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "DigitalOcean",
		"account":   client.Name,
		"service":   "Trigger",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
