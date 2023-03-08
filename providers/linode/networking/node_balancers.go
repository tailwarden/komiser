package networking

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

func NodeBalancers(ctx context.Context, client providers.ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)

	nodeBalancers, err := client.LinodeClient.ListNodeBalancers(ctx, &linodego.ListOptions{})
	if err != nil {
		return resources, err
	}

	for _, nodeBalancer := range nodeBalancers {
		tags := make([]Tag, 0)
		for _, tag := range nodeBalancer.Tags {
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
			Service:    "NodeBalancer",
			Region:     nodeBalancer.Region,
			ResourceId: fmt.Sprintf("%d", nodeBalancer.ID),
			Cost:       0,
			Name:       *nodeBalancer.Label,
			FetchedAt:  time.Now(),
			CreatedAt:  *nodeBalancer.Created,
			Tags:       tags,
			Link:       fmt.Sprintf("https://cloud.linode.com/nodebalancers/%d", nodeBalancer.ID),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Linode",
		"account":   client.Name,
		"service":   "NodeBalancer",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
