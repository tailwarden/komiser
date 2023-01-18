package ec2

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func Volumes(ctx context.Context, client ProviderClient) ([]Resource, error) {
	var config ec2.DescribeVolumesInput
	resources := make([]Resource, 0)
	ec2Client := ec2.NewFromConfig(*client.AWSClient)

	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return resources, err
	}

	accountId := stsOutput.Account

	for {
		output, err := ec2Client.DescribeVolumes(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, volume := range output.Volumes {
			tags := make([]Tag, 0)
			for _, tag := range volume.Tags {
				tags = append(tags, Tag{
					Key:   *tag.Key,
					Value: *tag.Value,
				})
			}

			resourceArn := fmt.Sprintf("arn:aws:ec2:%s:%s:volume/%s", client.AWSClient.Region, *accountId, *volume.VolumeId)

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "EBS",
				Region:     client.AWSClient.Region,
				ResourceId: resourceArn,
				Cost:       0,
				Name:       *volume.VolumeId,
				CreatedAt:  *volume.CreateTime,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link:       fmt.Sprintf("https:/%s.console.aws.amazon.com/ec2/home?region=%s#VolumeDetails:volumeId=%s", client.AWSClient.Region, client.AWSClient.Region, *volume.VolumeId),
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
		"service":   "EBS",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
