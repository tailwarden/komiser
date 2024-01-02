package efs

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
)

func ElasticFileStorage(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config efs.DescribeFileSystemsInput
	efsClient := efs.NewFromConfig(*client.AWSClient)

	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return resources, err
	}

	accountId := stsOutput.Account

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "EFS")
	if err != nil {
		log.Warnln("Couldn't fetch EFS cost and usage:", err)
	}

	for {
		output, err := efsClient.DescribeFileSystems(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, filesystem := range output.FileSystems {
			if filesystem.Name != nil {

				resourceArn := fmt.Sprintf("arn:aws:efs:%s:%s:file-systems/%s", client.AWSClient.Region, *accountId, *filesystem.Name)
				outputTags, err := efsClient.ListTagsForResource(ctx, &efs.ListTagsForResourceInput{
					ResourceId: &resourceArn,
				})

				tags := make([]Tag, 0)

				if err == nil {
					for _, tag := range outputTags.Tags {
						tags = append(tags, Tag{
							Key:   *tag.Key,
							Value: *tag.Value,
						})
					}
				}

				monthlyCost := float64(filesystem.SizeInBytes.Value/1000000000) * 0.30

				resources = append(resources, Resource{
					Provider:   "AWS",
					Account:    client.Name,
					Service:    "EFS",
					ResourceId: resourceArn,
					Region:     client.AWSClient.Region,
					Name:       *filesystem.Name,
					Cost:       monthlyCost,
					Metadata: map[string]string{
						"serviceCost": fmt.Sprint(serviceCost),
					},
					Tags:       tags,
					FetchedAt:  time.Now(),
					Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/efs/home?region=%s#/file-systems/%s", client.AWSClient.Region, client.AWSClient.Region, *filesystem.Name),
				})
			}
		}

		if aws.ToString(output.NextMarker) == "" {
			break
		}

		config.Marker = output.NextMarker
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "EFS",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
