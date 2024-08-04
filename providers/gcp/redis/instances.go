package redis

import (
	"context"

	"fmt"
	"regexp"
	"strings"
	"time"

	redis "cloud.google.com/go/redis/apiv1"
	"cloud.google.com/go/redis/apiv1/redispb"
	"github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	"github.com/tailwarden/komiser/utils"
)



func Instances(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	pricing, err := FetchPricing()
	if err != nil {
		return nil, err
	}

	resources := make([]models.Resource, 0)

	regions, err := utils.FetchGCPRegionsInRealtime(client.GCPClient.Credentials.ProjectID, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		if strings.Contains(err.Error(), "SERVICE_DISABLED") {
			logrus.Warn(err.Error())
			return resources, nil
		} else {
			logrus.WithError(err).Errorf("failed to list zones to fetch redis")
			return resources, err
		}
	}

	redisClient, err := redis.NewCloudRedisRESTClient(ctx, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		logrus.WithError(err).Errorf("failed to create redis client")
		return resources, err
	}

RegionsLoop:
	for _, regionName := range regions {
		req := &redispb.ListInstancesRequest{
			Parent: "projects/" + client.GCPClient.Credentials.ProjectID + "/locations/" + regionName,
		}

		redisInstances := redisClient.ListInstances(ctx, req)

		for {
			redis, err := redisInstances.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				if err.Error() == "googleapi: Error 403: Location "+regionName+" is not found or access is unauthorized." {
					continue RegionsLoop
				} else {
					logrus.WithError(err).Errorf("failed to list redis for region " + regionName)
					return resources, err
				}
			}
			re := regexp.MustCompile(`instances\/(.+)$`)
			redisInstanceName := re.FindStringSubmatch(redis.Name)[1]

			cost := calculateRedisCost(redis, pricing)

			resources = append(resources, models.Resource{
				Provider:   "GCP",
				Account:    client.Name,
				ResourceId: redisInstanceName,
				Service:    "Redis",
				Name:       redis.DisplayName,
				Region:     regionName,
				CreatedAt:  redis.CreateTime.AsTime(),
				Cost:       cost,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://console.cloud.google.com/memorystore/redis/locations/%s/instances/%s/details/overview?project=%s", regionName, redisInstanceName, client.GCPClient.Credentials.ProjectID),
			})

		}

	}

	logrus.WithFields(logrus.Fields{
		"provider":  "GCP",
		"account":   client.Name,
		"service":   "Redis",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
