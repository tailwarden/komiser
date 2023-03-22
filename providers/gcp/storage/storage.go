package storage

import (
	"context"
	"time"

	"cloud.google.com/go/storage"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func Buckets(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	storageClient, err := storage.NewClient(ctx, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		return []models.Resource{}, err
	}

	buckets := storageClient.Buckets(ctx, client.GCPClient.Credentials.ProjectID)
	for {
		bucket, err := buckets.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.WithError(err).Errorf("failed to list buckets")
			return resources, err
		}

		tags := make([]models.Tag, 0)
		if bucket.Labels != nil {
			for key, value := range bucket.Labels {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: value,
				})
			}
		}

		resources = append(resources, models.Resource{
			Provider:   "GCP",
			Account:    client.Name,
			Service:    "Bucket",
			ResourceId: bucket.Name,
			Region:     bucket.Location,
			Tags:       tags,
			FetchedAt:  time.Now(),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "GCP",
		"account":   client.Name,
		"service":   "Bucket",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
