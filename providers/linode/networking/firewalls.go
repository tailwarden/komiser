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

func Firewalls(ctx context.Context, client providers.ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)

	firewalls, err := client.LinodeClient.ListFirewalls(ctx, &linodego.ListOptions{})
	if err != nil {
		return resources, err
	}

	for _, firewall := range firewalls {
		tags := make([]Tag, 0)
		for _, tag := range firewall.Tags {
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
			Service:    "Firewall",
			Region:     "Global",
			ResourceId: fmt.Sprintf("%d", firewall.ID),
			Cost:       0,
			Name:       firewall.Label,
			FetchedAt:  time.Now(),
			CreatedAt:  *firewall.Created,
			Tags:       tags,
			Link:       fmt.Sprintf("https://cloud.linode.com/firewalls/%d", firewall.ID),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Linode",
		"account":   client.Name,
		"service":   "Firewall",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
