package ec2

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	. "github.com/mlabouardy/komiser/models"
	. "github.com/mlabouardy/komiser/providers"
)

func SecurityGroups(ctx context.Context, client ProviderClient) ([]Resource, error) {
	var config ec2.DescribeSecurityGroupsInput
	resources := make([]Resource, 0)
	ec2Client := ec2.NewFromConfig(*client.AWSClient)

	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return resources, err
	}

	accountId := stsOutput.Account

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

			resourceArn := fmt.Sprintf("arn:aws:ec2:%s:%s:instance/%s", client.AWSClient.Region, *accountId, sg.GroupId)

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "Security Group",
				Region:     client.AWSClient.Region,
				ResourceId: resourceArn,
				Cost:       0,
				Name:       *sg.GroupName,
				FetchedAt:  time.Now(),
				Tags:       tags,
			})
		}
		if aws.ToString(output.NextToken) == "" {
			break
		}

		config.NextToken = output.NextToken
	}
	log.Debugf("[%s] Fetched %d AWS Security groups from %s\n", client.Name, len(resources), client.AWSClient.Region)
	return resources, nil
}
