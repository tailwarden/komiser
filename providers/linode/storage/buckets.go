package storage

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

func Buckets(ctx context.Context, client providers.ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)

	buckets, err := client.LinodeClient.ListObjectStorageBuckets(ctx, &linodego.ListOptions{})
	if err != nil {
		return resources, err
	}

	for _, bucket := range buckets {
		region := strings.Split(bucket.Hostname, ".")[1]

		resources = append(resources, models.Resource{
			Provider:   "Linode",
			Account:    client.Name,
			Service:    "Bucket",
			Region:     region,
			ResourceId: bucket.Hostname,
			Cost:       0,
			Name:       bucket.Label,
			FetchedAt:  time.Now(),
			CreatedAt:  *bucket.Created,
			Link: fmt.Sprintf("https://cloud.linode.com/object-storage"+
				"/buckets/%s/%s", region, bucket.Label),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Linode",
		"account":   client.Name,
		"service":   "Bucket",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
