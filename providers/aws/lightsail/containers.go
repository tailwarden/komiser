package lightsail

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Containers(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var config lightsail.GetContainerServicesInput
	resources := make([]models.Resource, 0)
	lsClient := lightsail.NewFromConfig(*client.AWSClient)

	output, err := lsClient.GetContainerServices(ctx, &config)
	if err != nil {
		return resources, err
	}

	for _, instance := range output.ContainerServices {

		instanceName := *instance.ContainerServiceName

		resources = append(resources, models.Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "Lightsail Containers",
			Region:     client.AWSClient.Region,
			ResourceId: *instance.Arn,
			Name:       instanceName,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://lightsail.aws.amazon.com/ls/webapp/%s/containers/%s/connect", client.AWSClient.Region, *instance.ContainerServiceName),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "Lightsail Containers",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
