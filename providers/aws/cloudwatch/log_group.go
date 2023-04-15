package cloudwatch

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	log "github.com/sirupsen/logrus"

	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func LogGroups(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	cloudWatchLogsClient := cloudwatchlogs.NewFromConfig(*client.AWSClient)
	input := &cloudwatchlogs.DescribeLogGroupsInput{}
	for {
		output, err := cloudWatchLogsClient.DescribeLogGroups(ctx, input)
		if err != nil {
			return resources, err
		}
		for _, group := range output.LogGroups {
			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "CloudWatch Log Group",
				ResourceId: aws.ToString(group.Arn),
				Region:     client.AWSClient.Region,
				Name:       aws.ToString(group.LogGroupName),
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/cloudwatch/home?region=%s#logsV2:log-groups/log-group/%s", client.AWSClient.Region, client.AWSClient.Region, aws.ToString(group.LogGroupName)),
			})
		}
		if output.NextToken == nil {
			break
		}
		input.NextToken = output.NextToken
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "CloudWatch Log Group",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
