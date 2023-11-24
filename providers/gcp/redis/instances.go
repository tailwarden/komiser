package redis

import (
	"context"
	"encoding/json"
	"math"
	"net/http"

	"fmt"
	"regexp"
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

type RedisPrice []struct {
	Val      int     `json:"val"`
	Currency string  `json:"currnecy"`
	Nanos    float64 `json:"nanos"`
}

type RegionBasedPricing struct {
	Regions map[string]struct {
		Price RedisPrice `json:"price"`
	} `json:"regions"`
}

type GcpDatabasePricing struct {
	Gcp struct {
		Databases struct {
			CloudMemorystore struct {
				Redis struct {
					Basic    map[string]RegionBasedPricing `json:"basic"`
					Standard map[string]RegionBasedPricing `json:"standard"`
				} `json:"redis"`
			} `json:"cloud_memorystore"`
		} `json:"databases"`
	} `json:"gcp"`
}

func calculateRedisCost(redis *redispb.Instance, pricing GcpDatabasePricing) float64 {
	var priceMap map[string]RegionBasedPricing
	var priceKey string

	prices := []int32{4, 10, 35, 100}
	capacityTier := getCapacityTier(redis.MemorySizeGb, prices)

	if redis.Tier == redispb.Instance_BASIC {
		priceMap = pricing.Gcp.Databases.CloudMemorystore.Redis.Basic
		priceKey = fmt.Sprintf("Rediscapacitybasicm%ddefault", capacityTier)
	} else if redis.Tier == redispb.Instance_STANDARD_HA {
		priceMap = pricing.Gcp.Databases.CloudMemorystore.Redis.Standard
		if redis.ReadReplicasMode == redispb.Instance_READ_REPLICAS_DISABLED {
			priceKey = fmt.Sprintf("Rediscapacitystandardm%ddefault", capacityTier)
		} else {
			priceKey = fmt.Sprintf("Rediscapacitystandardnodem%d", capacityTier)
		}
	}

	pricePerHrPerGbInNanos := priceMap[priceKey].Regions[redis.LocationId].Price[0].Nanos
	pricePerHrPerGbInDollars := pricePerHrPerGbInNanos / math.Pow(10, 9)

	now := time.Now()
	startTime := getStartTime(redis.GetCreateTime().AsTime(), now)

	hours := now.Sub(startTime).Hours()

	cost := hours * pricePerHrPerGbInDollars

	if redis.ReadReplicasMode == redispb.Instance_READ_REPLICAS_ENABLED {
		cost *= float64(redis.ReplicaCount)
	}

	return cost
}

func getCapacityTier(memorySizeGb int32, prices []int32) int {
	capacityTier := 5
	for idx, price := range prices {
		if memorySizeGb <= price {
			capacityTier = idx + 1
		}
	}
	return capacityTier
}

func getStartTime(createTime, now time.Time) time.Time {
	firstOfCurrentMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	if createTime.After(firstOfCurrentMonth) {
		return createTime
	}
	return firstOfCurrentMonth
}

func Instances(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	res, err := http.Get("https://www.gstatic.com/cloud-site-ux/pricing/data/gcp-databases.json")
	if err != nil {
		return nil, err
	}

	var pricing GcpDatabasePricing
	err = json.NewDecoder(res.Body).Decode(&pricing)
	if err != nil {
		return nil, err
	}

	resources := make([]models.Resource, 0)

	regions, err := utils.FetchGCPRegionsInRealtime(client.GCPClient.Credentials.ProjectID, option.WithCredentials(client.GCPClient.Credentials))
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
