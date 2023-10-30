package network

import (
	"context"
	"time"

	"github.com/civo/civogo"
	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Firewalls(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	firewalls, err := client.CivoClient.ListFirewalls()
	if err != nil {
		return resources, err
	}

	for _, firewall := range firewalls {
		relation := getFirewallRelations(firewall)
		resources = append(resources, models.Resource{
			Provider:   "Civo",
			Account:    client.Name,
			Service:    "Firewall",
			Region:     client.CivoClient.Region,
			ResourceId: firewall.ID,
			Cost:       0,
			Name:       firewall.Name,
			Relations: relation,
			FetchedAt:  time.Now(),
			Link:       "https://dashboard.civo.com/firewalls",
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Civo",
		"account":   client.Name,
		"service":   "Firewall",
		"region":    client.CivoClient.Region,
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}


func getFirewallRelations(firewall civogo.Firewall) []models.Link {
	return []models.Link{
		{
			ResourceID: firewall.NetworkID,
			Type: "Network",
			Name: firewall.NetworkID, //cannot get the name of the network unless calling the network api
			Relation: "USES",
		},
	}
}