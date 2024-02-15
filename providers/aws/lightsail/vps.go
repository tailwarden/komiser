package lightsail

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func VPS(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var config lightsail.GetInstancesInput
	resources := make([]models.Resource, 0)
	lsClient := lightsail.NewFromConfig(*client.AWSClient)

	for {
		output, err := lsClient.GetInstances(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, instance := range output.Instances {

			instanceName := *instance.Name

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "Lightsail VPS",
				Region:     client.AWSClient.Region,
				ResourceId: *instance.Arn,
				Name:       instanceName,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://lightsail.aws.amazon.com/ls/webapp/%s/instances/%s/connect", client.AWSClient.Region, *instance.Name),
				Metadata: map[string]string{
					"BlueprintName": *instance.BlueprintName,
				},
			})
		}

		if aws.ToString(output.NextPageToken) == "" {
			break
		}

		config.PageToken = output.NextPageToken
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "Lightsail VPS",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
