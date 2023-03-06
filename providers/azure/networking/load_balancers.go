package networking

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func LoadBalancers(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	loadBalancerClient, err := armnetwork.NewLoadBalancersClient(
		client.AzureClient.SubscriptionId,
		client.AzureClient.Credentials,
		&arm.ClientOptions{},
	)
	if err != nil {
		return resources, err
	}

	pager := loadBalancerClient.NewListAllPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return resources, err
		}

		for _, loadBalancer := range page.LoadBalancerListResult.Value {
			tags := make([]models.Tag, 0)

			for key, value := range loadBalancer.Tags {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: *value,
				})
			}

			resources = append(resources, models.Resource{
				Provider:   "Azure",
				Account:    client.Name,
				Service:    "Load Balancer",
				Region:     *loadBalancer.Location,
				ResourceId: *loadBalancer.ID,
				Cost:       0,
				Name:       *loadBalancer.Name,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link:       fmt.Sprintf("https://portal.azure.com/#resource%s", *loadBalancer.ID),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "Azure",
		"account":   client.Name,
		"service":   "Load Balancer",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
