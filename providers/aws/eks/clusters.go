package eks

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func BeginningOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 0, -date.Day()+1)
}

func KubernetesClusters(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config eks.ListClustersInput
	eksClient := eks.NewFromConfig(*client.AWSClient)

	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return resources, err
	}

	accountId := stsOutput.Account

	for {
		output, err := eksClient.ListClusters(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, cluster := range output.Clusters {
			resourceArn := fmt.Sprintf("arn:aws:eks:%s:%s:cluster/%s", client.AWSClient.Region, *accountId, cluster)
			outputTags, err := eksClient.ListTagsForResource(ctx, &eks.ListTagsForResourceInput{
				ResourceArn: &resourceArn,
			})

			tags := make([]Tag, 0)

			if err == nil {
				for key, value := range outputTags.Tags {
					tags = append(tags, Tag{
						Key:   key,
						Value: value,
					})
				}
			}

			outputDescribe, err := eksClient.DescribeCluster(ctx, &eks.DescribeClusterInput{
				Name: &cluster,
			})

			monthlyCost := 0.0
			if err == nil {
				startOfMonth := BeginningOfMonth(time.Now())
				hourlyUsage := 0
				if (*outputDescribe.Cluster.CreatedAt).Before(startOfMonth) {
					hourlyUsage = int(time.Now().Sub(startOfMonth).Hours())
				} else {
					hourlyUsage = int(time.Now().Sub(*outputDescribe.Cluster.CreatedAt).Hours())
				}
				monthlyCost = float64(hourlyUsage) * 0.10
			}

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "EKS",
				ResourceId: resourceArn,
				Region:     client.AWSClient.Region,
				Name:       cluster,
				Cost:       monthlyCost,
				Tags:       tags,
				CreatedAt:  *outputDescribe.Cluster.CreatedAt,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/eks/home?region=%s#/clusters/%s", client.AWSClient.Region, client.AWSClient.Region, cluster),
			})
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}

		config.NextToken = output.NextToken
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "EKS",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
