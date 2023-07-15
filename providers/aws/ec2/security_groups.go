package ec2

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func SecurityGroups(ctx context.Context, client ProviderClient) ([]Resource, error) {
	var config ec2.DescribeSecurityGroupsInput
	resources := make([]Resource, 0)
	ec2Client := ec2.NewFromConfig(*client.AWSClient)

	for {
		output, err := ec2Client.DescribeSecurityGroups(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, sg := range output.SecurityGroups {
			tags := make([]Tag, 0)
			for _, tag := range sg.Tags {
				tags = append(tags, Tag{
					Key:   *tag.Key,
					Value: *tag.Value,
				})
			}

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "Security Group",
				Region:     client.AWSClient.Region,
				ResourceId: *sg.GroupId,
				Cost:       0,
				Name:       *sg.GroupName,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link:       fmt.Sprintf("https:/%s.console.aws.amazon.com/ec2/home?region=%s#SecurityGroup:groupId=%s", client.AWSClient.Region, client.AWSClient.Region, *sg.GroupId),
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
		"service":   "Security Group",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
