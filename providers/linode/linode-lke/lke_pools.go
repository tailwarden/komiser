package compute

import (
	"context"
	"fmt"
	// "strings"
	"time"

	"github.com/linode/linodego"
	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

type LinodeLKENodePool struct {
	NodePool *linodego.LKECluster `json:"node_pool"`
}

func LKENodePools(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	lkeClusters, err := client.LinodeClient.ListLKEClusters(ctx, &linodego.ListOptions{})
	if err != nil {
		return resources, err
	}

	for _, lkeCluster := range lkeClusters {
		nodePools, err := client.LinodeClient.ListLKENodePools(ctx, lkeCluster.ID, &linodego.ListOptions{})
		if err != nil {
			return resources, err
		}

		for _, nodePool := range nodePools {
			tags := make([]models.Tag, 0)
			// If Linode LKE node pools have tags, you can handle them here.
			// Replace the example tags below with your actual tag handling logic.
			if len(nodePool.Tags) > 0 {
				tags = append(tags, models.Tag{
					Key:   "example-key",
					Value: nodePool.Tags[0],
				})
			}

			resources = append(resources, models.Resource{
				Provider:   "Linode",
				Account:    client.Name,
				Service:    "Linode Kubernetes Engine",
				// Region:     nodePool.Region,
				ResourceId: fmt.Sprintf("%d", nodePool.ID),
				Cost:       0,
				// Name:       nodePool.Label,
				FetchedAt:  time.Now(),
				CreatedAt:  time.Time{}, // Update this with the actual created time.
				Tags:       tags,
				Link:       fmt.Sprintf("https://cloud.linode.com/kubernetes/clusters/%d", lkeCluster.ID),
				// Add any additional fields or data you want to collect here.
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "Linode",
		"account":   client.Name,
		"service":   "Linode Kubernetes Engine",
		"resources": len(resources),
	}).Info("Fetched LKE node pools")
	return resources, nil
}
