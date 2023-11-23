package codepipeline

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pipeline"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/tailwarden/komiser/utils"
)

func GetPipelines(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config pipeline.ListClustersInput
	pipelineClient := pipeline.NewFromConfig(*client.AWSClient)

	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return resources, err
	}

	accountId := stsOutput.Account

	for {
		output, err := pipelineClient.ListClusters(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, cluster := range output.Clusters {
			resourceArn := fmt.Sprintf("arn:aws:pipeline:%s:%s:cluster/%s", client.AWSClient.Region, *accountId, cluster)
			outputTags, err := pipelineClient.ListTagsForResource(ctx, &pipeline.ListTagsForResourceInput{
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

			outputDescribe, err := pipelineClient.DescribeCluster(ctx, &eks.DescribeClusterInput{
				Name: &cluster,
			})

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "codepipeline",
				ResourceId: resourceArn,
				Region:     client.AWSClient.Region,
				Name:       cluster,
				Tags:       tags,
				CreatedAt:  createdAt,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/codepipeline/home?region=%s#/clusters/%s", client.AWSClient.Region, client.AWSClient.Region, cluster),
			})
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}
å
		config.NextToken = output.NextToken
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "codepipeline",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}