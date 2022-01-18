package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2020-12-01/redis"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	. "github.com/mlabouardy/komiser/models/azure"
)

func getRedisClient(subscriptionID string) redis.Client {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	redisClient := redis.NewClient(subscriptionID)
	redisClient.Authorizer = a
	return redisClient
}

func (azure Azure) DescribeRedisInstances(subscriptionID string) ([]RedisInstance, error) {
	redisClient := getRedisClient(subscriptionID)
	ctx := context.Background()
	redisInstances := make([]RedisInstance, 0)
	for redisListItr, err := redisClient.ListBySubscriptionComplete(ctx); redisListItr.NotDone(); redisListItr.Next() {
		if err != nil {
			return redisInstances, err
		}
		redis := redisListItr.Value()
		redisInstances = append(redisInstances, RedisInstance{
			Name: *redis.Name,
			ID:   *redis.ID,
		})
	}
	return redisInstances, nil
}
