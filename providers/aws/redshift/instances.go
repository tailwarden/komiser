package redshift

import (
  "context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/service/redshift"
	. "github.com/mlabouardy/komiser/models"
	. "github.com/mlabouardy/komiser/providers"
)

func RedshiftClusters(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	if client.AWSClient.Region == "us-east-1" {
		redshiftClient := redshift.NewFromConfig(*client.AWSClient)
		describeClustersInput := &redshift.DescribeClustersInput{}
		output, err := redshiftClient.DescribeClusters(context.Background(), describeClustersInput)
		if err != nil {
			return resources, err
		}

		for _, cluster := range output.Clusters {
			resourceArn := fmt.Sprintf("arn:aws:redshift:%s:%s:cluster:%s", client.AWSClient.Region, client.AccountID, *cluster.ClusterIdentifier)

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "Redshift",
				Region:     client.AWSClient.Region,
				ResourceId: resourceArn,
				Name:       *cluster.ClusterIdentifier,
				Cost:       0, // You can calculate the cost if needed
				CreatedAt:  *cluster.ClusterCreateTime,
				Tags:       []Tag{}, // Redshift does not have tags, but you can add them if necessary
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://console.aws.amazon.com/redshift/home?region=%s#cluster-details?cluster=%s", client.AWSClient.Region, *cluster.ClusterIdentifier),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "Redshift",
		"resources": len(resources),
	}).Debugf("Fetched Redshift clusters")

	return resources, nil
}
