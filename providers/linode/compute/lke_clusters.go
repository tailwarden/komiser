package compute

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/linode/linodego"
	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func LKEClusters(ctx context.Context, client providers.ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)

	lkeClusters, err := client.LinodeClient.ListLKEClusters(ctx, &linodego.ListOptions{})
	if err != nil {
		return resources, err
	}

	for _, lkeCluster := range lkeClusters {
		tags := make([]Tag, 0)
		for _, tag := range lkeCluster.Tags {
			if strings.Contains(tag, ":") {
				parts := strings.Split(tag, ":")
				tags = append(tags, models.Tag{
					Key:   parts[0],
					Value: parts[1],
				})
			} else {
				tags = append(tags, models.Tag{
					Key:   tag,
					Value: tag,
				})
			}
		}

		resources = append(resources, models.Resource{
			Provider:   "Linode",
			Account:    client.Name,
			Service:    "LKE",
			Region:     lkeCluster.Region,
			ResourceId: fmt.Sprintf("%d", lkeCluster.ID),
			Cost:       0,
			Name:       lkeCluster.Label,
			FetchedAt:  time.Now(),
			CreatedAt:  *lkeCluster.Created,
			Tags:       tags,
			Link:       fmt.Sprintf("https://cloud.linode.com/kubernetes/clusters/%d", lkeCluster.ID),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Linode",
		"account":   client.Name,
		"service":   "LKE",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
