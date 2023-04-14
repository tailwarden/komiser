package redis

import (
	"context"
	"fmt"
	"regexp"
	"time"

	redis "cloud.google.com/go/redis/apiv1"
	"cloud.google.com/go/redis/apiv1/redispb"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func Redis(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	regions, err := listGCPRegions(client.GCPClient.Credentials.ProjectID, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		logrus.WithError(err).Errorf("failed to list zones to fetch redis")
		return resources, err
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

			resources = append(resources, models.Resource{
				Provider:   "GCP",
				Account:    client.Name,
				ResourceId: redisInstanceName,
				Service:    "Redis",
				Name:       redis.DisplayName,
				Region:     regionName,
				CreatedAt:  redis.CreateTime.AsTime(),
				Cost:       0,
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

func listGCPRegions(projectId string, creds option.ClientOption) ([]string, error) {
	var regions []string

	ctx := context.Background()
	computeService, err := compute.NewService(ctx, creds)
	if err != nil {
		log.WithError(err).Debug("failed to create new service for fetching GCP regions for redis instance")
		return nil, err
	}

	regionList, err := computeService.Regions.List(projectId).Do()
	if err != nil {
		log.WithError(err).Debug("failed to list regions for fetching GCP regions for redis instance")
		return nil, err
	}

	for _, region := range regionList.Items {
		regions = append(regions, region.Name)
	}
	return regions, nil
}
