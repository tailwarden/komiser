package redshift

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
)

func Resources(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	var config redshift.DescribeClustersInput
	redshiftClient := redshift.NewFromConfig(*client.AWSClient)

	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return resources, err
	}

	accountId := stsOutput.Account

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "Redshift")
	if err != nil {
		log.Warnln("Couldn't fetch Redshift cost and usage:", err)
	}

	for {
		output, err := redshiftClient.DescribeClusters(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, cluster := range output.Clusters {
			resourceArn := fmt.Sprintf("arn:aws:redshift:%s:%s:cluster/%s", client.AWSClient.Region, *accountId, *cluster.ClusterIdentifier)
			outputTags := cluster.Tags

			tags := make([]models.Tag, 0)

			for _, tag := range outputTags {
				tags = append(tags, models.Tag{
					Key:   *tag.Key,
					Value: *tag.Value,
				})
			}

			monthlyCost := float64(0)

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "Redshift Cluster",
				ResourceId: resourceArn,
				Region:     client.AWSClient.Region,
				Name:       *cluster.ClusterIdentifier,
				Cost:       monthlyCost,
				Metadata: map[string]string{
					"serviceCost": fmt.Sprint(serviceCost),
				},
				Tags:      tags,
				FetchedAt: time.Now(),
				Link:      fmt.Sprintf("https://%s.console.aws.amaxon.com/redshift/home?region=%s/clusters/%s", client.AWSClient.Region, client.AWSClient.Region, *cluster.ClusterIdentifier),
			})

		}

		if aws.ToString(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "Redshift Cluster",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil

}
