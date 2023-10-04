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

func LinodeInstancesAndInstanceDisks(ctx context.Context, client providers.ProviderClient) ([]Resource, error) {
	linodeInstances, err := client.LinodeClient.ListInstances(ctx, &linodego.ListOptions{
		PageOptions: &linodego.PageOptions{},
		PageSize:    0,
		Filter:      "",
	})
	if err != nil {
		return nil, err
	}

	resources := make([]Resource, 0)
	for _, instance := range linodeInstances {

		tags := make([]Tag, 0)
		for _, tag := range instance.Tags {
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
			Service:    "Linode Instance",
			Region:     instance.Region,
			ResourceId: fmt.Sprintf("%d", instance.ID),
			Cost:       0,
			Name:       instance.Label,
			FetchedAt:  time.Now(),
			CreatedAt:  *instance.Created,
			Tags:       tags,
			Link:       fmt.Sprintf("https://cloud.linode.com/linodes/%d", instance.ID),
		})

		instanceDisks, err := client.LinodeClient.ListInstanceDisks(ctx, instance.ID, &linodego.ListOptions{})
		if err != nil {
			return resources, err
		}

		for _, disk := range instanceDisks {
			resources = append(resources, models.Resource{
				Provider:   "Linode",
				Account:    client.Name,
				Service:    "Linode Instance Disk",
				Region:     instance.Region,
				ResourceId: fmt.Sprintf("%d", disk.ID),
				Cost:       0,
				Name:       disk.Label,
				FetchedAt:  time.Now(),
				CreatedAt:  *disk.Created,
				Link:       fmt.Sprintf("https://cloud.linode.com/linodes/%d/storage", instance.ID),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "Linode",
		"account":   client.Name,
		"service":   "Linode Instance and Instance Disk",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
