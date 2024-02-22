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

func Databases(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var config lightsail.GetRelationalDatabasesInput
	resources := make([]models.Resource, 0)
	lsClient := lightsail.NewFromConfig(*client.AWSClient)

	output, err := lsClient.GetRelationalDatabases(ctx, &config)
	if err != nil {
		return resources, err
	}

	for _, instance := range output.RelationalDatabases {

		instanceName := *instance.MasterDatabaseName

		resources = append(resources, models.Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "Lightsail Databases",
			Region:     client.AWSClient.Region,
			ResourceId: *instance.Arn,
			Name:       instanceName,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://lightsail.aws.amazon.com/ls/webapp/%s/databases/%s/connect", client.AWSClient.Region, *instance.MasterDatabaseName),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "Lightsail Databases",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
