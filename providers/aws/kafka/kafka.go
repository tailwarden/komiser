package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/kafka"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Kafka(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	wAclClient := kafka.NewFromConfig(*client.AWSClient)

	clusters, err := wAclClient.ListClustersV2(ctx, &kafka.ListClustersV2Input{})
	if err != nil {
		return resources, err
	}

	for _, cluster := range clusters.ClusterInfoList {

		tags := make([]models.Tag, 0)
		for key, value := range cluster.Tags {
			tags = append(tags, models.Tag{
				Key:   key,
				Value: value,
			})
		}
		resources = append(resources, models.Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "Kafka",
			ResourceId: *cluster.ClusterArn,
			Region:     client.AWSClient.Region,
			Name:       *cluster.ClusterArn,
			Cost:       0,
			Tags:       tags,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/msk/home?region=%s#/cluster/%s/view?tabId=metrics", client.AWSClient.Region, client.AWSClient.Region, *cluster.ClusterArn),
		})
	}

	log.WithFields(log.Fields{
		"provider":    "AWS",
		"account":     client.Name,
		"region":      client.AWSClient.Region,
		"service":     "Kafka",
		"resources":   len(resources),
		"serviceCost": fmt.Sprint(0),
	}).Info("Fetched resources")
	return resources, nil
}
