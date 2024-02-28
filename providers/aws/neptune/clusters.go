package neptune

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/neptune"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Clusters(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var config neptune.DescribeDBClustersInput
	resources := make([]models.Resource, 0)
	neptuneClient := neptune.NewFromConfig(*client.AWSClient)

	output, err := neptuneClient.DescribeDBClusters(ctx, &config)
	if err != nil {
		return resources, err
	}

	for _, cluster := range output.DBClusters {
		clusterName := ""
		if cluster.DatabaseName != nil {
			clusterName = *cluster.DatabaseName
		}
		resources = append(resources, models.Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "Neptune Clusters",
			Region:     client.AWSClient.Region,
			ResourceId: *cluster.DBClusterArn,
			Name:       clusterName,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/neptune/home?region=%s#database-details:id=%s;resource-type=cluster;tab=connectivity", client.AWSClient.Region, client.AWSClient.Region, clusterName),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "Neptune Clusters",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
