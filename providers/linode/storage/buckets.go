package storage

import (
	"context"
	"fmt"
	// "strings"
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
		// tags := make([]Tag, 0)
		// for _, tag := range bucket.Tags {
		// 	if strings.Contains(tag, ":") {
		// 		parts := strings.Split(tag, ":")
		// 		tags = append(tags, models.Tag{
		// 			Key:   parts[0],
		// 			Value: parts[1],
		// 		})
		// 	} else {
		// 		tags = append(tags, models.Tag{
		// 			Key:   tag,
		// 			Value: tag,
		// 		})
		// 	}
		// }

		resources = append(resources, models.Resource{
			Provider:   "Linode",
			Account:    client.Name,
			Service:    "Bucket",
			// Region:     bucket.Region,
			// ResourceId: fmt.Sprintf("%s", bucket.ID),
			// Hostname:   fmt.Sprintf("%s", bucket.Hostname),
			Cost:       0,
			Name:       bucket.Label,
			FetchedAt:  time.Now(),
			CreatedAt:  *bucket.Created,
			// Tags:       tags,
			Link:       fmt.Sprintf("https://cloud.linode.com/object-storage/buckets/us-southeast-1/%s",  bucket.Label),
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
