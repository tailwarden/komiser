package redis

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"

	"cloud.google.com/go/redis/apiv1/redispb"
)

const (
	M1GbLimit = 4
	M2GbLimit = 10
	M3GbLimit = 35
	M4GbLimit = 100
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

func FetchPricing() (*GcpDatabasePricing, error) {
	res, err := http.Get("https://www.gstatic.com/cloud-site-ux/pricing/data/gcp-databases.json")
	if err != nil {
		return nil, err
	}

	var pricing GcpDatabasePricing
	err = json.NewDecoder(res.Body).Decode(&pricing)
	if err != nil {
		return nil, err
	}

	return &pricing, nil
}

func calculateRedisCost(redis *redispb.Instance, pricing *GcpDatabasePricing) float64 {
	var priceMap map[string]RegionBasedPricing
	var priceKey string

	prices := []int32{M1GbLimit, M2GbLimit, M3GbLimit, M4GbLimit}
	capacityTier := getCapacityTier(redis.MemorySizeGb, prices)

	if redis.Tier == redispb.Instance_BASIC {
		priceMap = pricing.Gcp.Databases.CloudMemorystore.Redis.Basic
		priceKey = fmt.Sprintf("rediscapacitybasicm%ddefault", capacityTier)
	} else if redis.Tier == redispb.Instance_STANDARD_HA {
		priceMap = pricing.Gcp.Databases.CloudMemorystore.Redis.Standard
		if redis.ReadReplicasMode == redispb.Instance_READ_REPLICAS_DISABLED {
			priceKey = fmt.Sprintf("rediscapacitystandardm%ddefault", capacityTier)
		} else {
			priceKey = fmt.Sprintf("rediscapacitystandardnodem%d", capacityTier)
		}
	}

	location := strings.Join(strings.Split(redis.LocationId, "-")[:2], "-")
	pricePerHrPerGbInNanos := priceMap[priceKey].Regions[location].Price[0].Nanos
	pricePerHrPerGbInDollars := pricePerHrPerGbInNanos / math.Pow(10, 9)

	now := time.Now().UTC()
	startTime := getStartTime(redis.GetCreateTime().AsTime(), now)

	hours := now.Sub(startTime).Hours()

	cost := hours * pricePerHrPerGbInDollars

	if redis.ReadReplicasMode == redispb.Instance_READ_REPLICAS_ENABLED {
		cost *= float64(redis.ReplicaCount)
	}

	return cost
}

func getStartTime(createTime, now time.Time) time.Time {
	firstOfCurrentMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	if createTime.After(firstOfCurrentMonth) {
		return createTime
	}
	return firstOfCurrentMonth
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
