package network

import (
	"context"
	"time"

	"github.com/civo/civogo"
	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func LoadBalancers(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	lbs, err := client.CivoClient.ListLoadBalancers()
	if err != nil {
		return resources, err
	}

	for _, lb := range lbs {
		relations := getLoadBalancerRelations(lb)
		resources = append(resources, models.Resource{
			Provider:   "Civo",
			Account:    client.Name,
			Service:    "Load Balancer",
			Region:     client.CivoClient.Region,
			ResourceId: lb.ID,
			Cost:       10,
			Name:       lb.Name,
			Relations:  relations,
			FetchedAt:  time.Now(),
			Link:       "https://dashboard.civo.com/loadbalancers",
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Civo",
		"account":   client.Name,
		"service":   "Load Balancer",
		"region":    client.CivoClient.Region,
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}

func getLoadBalancerRelations(lb civogo.LoadBalancer) []models.Link {
	var rel []models.Link

	if len(lb.FirewallID) > 0 {
		rel = append(rel, models.Link{
			ResourceID: lb.FirewallID,
			Type:       "Firewall",
			Name:       lb.FirewallID, //cannot get the name of the network unless calling the network api
			Relation:   "USES",
		})
	}

	if len(lb.FirewallID) > 0 {
		rel = append(rel, models.Link{
			ResourceID: lb.ClusterID,
			Type:       "Cluster",
			Name:       lb.ClusterID,
			Relation:   "USES",
		})
	}
	return rel
}
