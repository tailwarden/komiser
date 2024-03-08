package sns

import (
	"context"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
	"github.com/tailwarden/komiser/utils"
)

func Topics(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	var config sns.ListTopicsInput
	snsClient := sns.NewFromConfig(*client.AWSClient)

	cloudwatchClient := cloudwatch.NewFromConfig(*client.AWSClient)

	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return resources, err
	}

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "Amazon Simple Notification Service")
	if err != nil {
		log.Warnln("Couldn't fetch Amazon Simple Notification Service cost and usage:", err)
	}

	accountId := stsOutput.Account

	for {
		output, err := snsClient.ListTopics(context.Background(), &config)
		if err != nil {
			return resources, err
		}

		for _, topic := range output.Topics {
			outputTags, err := snsClient.ListTagsForResource(ctx, &sns.ListTagsForResourceInput{
				ResourceArn: topic.TopicArn,
			})

			tags := make([]models.Tag, 0)

			if err == nil {
				for _, tag := range outputTags.Tags {
					tags = append(tags, models.Tag{
						Key:   *tag.Key,
						Value: *tag.Value,
					})
				}
			}

			resourceArnPrefix := fmt.Sprintf("arn:aws:sns:%s:%s:", client.AWSClient.Region, *accountId)
			topicName := strings.Replace(*topic.TopicArn, resourceArnPrefix, "", -1)

			metricsMessagesPublishedOutput, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
				StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
				EndTime:    aws.Time(time.Now()),
				MetricName: aws.String("NumberOfMessagesPublished"),
				Namespace:  aws.String("AWS/SNS"),
				Dimensions: []types.Dimension{
					{
						Name:  aws.String("TopicName"),
						Value: &topicName,
					},
				},
				Period: aws.Int32(3600),
				Statistics: []types.Statistic{
					types.StatisticSum,
				},
			})

			if err != nil {
				log.Warnf("Couldn't fetch invocations metric for %s", *topic.TopicArn)
			}

			requests := 0.0
			if metricsMessagesPublishedOutput != nil && len(metricsMessagesPublishedOutput.Datapoints) > 0 {
				requests = *metricsMessagesPublishedOutput.Datapoints[0].Sum
			}

			monthlyCost := (requests / 1000000) * 0.0000005

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "SNS",
				ResourceId: *topic.TopicArn,
				Region:     client.AWSClient.Region,
				Name:       *topic.TopicArn,
				Cost:       monthlyCost,
				Metadata: map[string]string{
					"serviceCost": fmt.Sprint(serviceCost),
				},
				Tags:      tags,
				FetchedAt: time.Now(),
				Link:      fmt.Sprintf("https://%s.console.aws.amazon.com/sns/v3/home?region=%s#/topic/%s", client.AWSClient.Region, client.AWSClient.Region, *topic.TopicArn),
			})
		}

		if aws.ToString(config.NextToken) == "" {
			break
		}

		config.NextToken = output.NextToken
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "SNS",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
