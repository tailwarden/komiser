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

type LinodeInstance struct {
	Instance  *linodego.Instance
	NodeCount int
}

func Linodes(ctx context.Context, client providers.ProviderClient, linodeInstances []LinodeInstance) ([]Resource, error) {
    resources := make([]Resource, 0)

    for _, linodeInstance := range linodeInstances {
        instance := linodeInstance.Instance

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
            Service:    "Linode",
            Region:     instance.Region,
            ResourceId: fmt.Sprintf("%d", instance.ID),
            Cost:       0,
            Name:       instance.Label,
            FetchedAt:  time.Now(),
            CreatedAt:  *instance.Created,
            Tags:       tags,
            Link:       fmt.Sprintf("https://cloud.linode.com/linodes/%d", instance.ID),
            NodeCount:  linodeInstance.NodeCount, // Include the NodeCount value
        })
    }

    log.WithFields(log.Fields{
        "provider":  "Linode",
        "account":   client.Name,
        "service":   "Linode",
        "resources": len(resources),
    }).Info("Fetched resources")
    return resources, nil
}
