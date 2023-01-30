package elasticache

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func Clusters(ctx context.Context, client ProviderClient) ([]Resource, error) {
	var config elasticache.DescribeCacheClustersInput
	resources := make([]Resource, 0)
	elasticacheClient := elasticache.NewFromConfig(*client.AWSClient)

	for {
		output, err := elasticacheClient.DescribeCacheClusters(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, cluster := range output.CacheClusters {
			tagsResp, err := elasticacheClient.ListTagsForResource(ctx, &elasticache.ListTagsForResourceInput{
				ResourceName: cluster.ARN,
			})
			if err != nil {
				return resources, err
			}

			tags := make([]Tag, len(tagsResp.TagList))
			for i, t := range tagsResp.TagList {
				tags[i] = Tag{
					Key:   *t.Key,
					Value: *t.Value,
				}
			}

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "ElastiCache",
				Region:     client.AWSClient.Region,
				ResourceId: *cluster.ARN,
				Name:       *cluster.CacheClusterId,
				Cost:       0,
				CreatedAt:  *cluster.CacheClusterCreateTime,
				Tags:       tags,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https:/%s.console.aws.amazon.com/elasticache/home?region=%s#/%s/%s", client.AWSClient.Region, client.AWSClient.Region, *cluster.Engine, *cluster.CacheClusterId),
			})
		}

		if aws.ToString(output.Marker) == "" {
			break
		}

		config.Marker = output.Marker
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "ElastiCache",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
