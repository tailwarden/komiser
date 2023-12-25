package sqs

import (
	"context"
	"fmt"
	"path"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
	"github.com/tailwarden/komiser/utils"
)

func Queues(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	cloudwatchClient := cloudwatch.NewFromConfig(*client.AWSClient)

	var config sqs.ListQueuesInput
	sqsClient := sqs.NewFromConfig(*client.AWSClient)

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "Amazon Simple Queue Service")
	if err != nil {
		log.Warnln("Couldn't fetch Amazon Simple Queue Service cost and usage:", err)
	}
	for {
		output, err := sqsClient.ListQueues(context.Background(), &config)
		if err != nil {
			return resources, err
		}

		for _, queue := range output.QueueUrls {
			queueName := path.Base(queue)

			metricsNbOfMessagesSentOutput, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
				StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
				EndTime:    aws.Time(time.Now()),
				MetricName: aws.String("NumberOfMessagesSent"),
				Namespace:  aws.String("AWS/SQS"),
				Dimensions: []types.Dimension{
					types.Dimension{
						Name:  aws.String("QueueName"),
						Value: aws.String(queueName),
					},
				},
				Period: aws.Int32(3600),
				Statistics: []types.Statistic{
					types.StatisticSum,
				},
			})

			if err != nil {
				log.Warnf("Couldn't fetch invocations metric for %s", queueName)
			}

			nbOfMessagesSent := 0.0
			if metricsNbOfMessagesSentOutput != nil && len(metricsNbOfMessagesSentOutput.Datapoints) > 0 {
				nbOfMessagesSent = *metricsNbOfMessagesSentOutput.Datapoints[0].Sum
			}

			metricsNbOfMessagesReceivedOutput, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
				StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
				EndTime:    aws.Time(time.Now()),
				MetricName: aws.String("NumberOfMessagesReceived"),
				Namespace:  aws.String("AWS/SQS"),
				Dimensions: []types.Dimension{
					types.Dimension{
						Name:  aws.String("QueueName"),
						Value: aws.String(queueName),
					},
				},
				Period: aws.Int32(3600),
				Statistics: []types.Statistic{
					types.StatisticSum,
				},
			})

			if err != nil {
				log.Warnf("Couldn't fetch invocations metric for %s", queueName)
			}

			nbOfMessagesReceived := 0.0
			if metricsNbOfMessagesReceivedOutput != nil && len(metricsNbOfMessagesReceivedOutput.Datapoints) > 0 {
				nbOfMessagesReceived = *metricsNbOfMessagesReceivedOutput.Datapoints[0].Sum
			}

			metricsNbOfMessagesDeletedOutput, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
				StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
				EndTime:    aws.Time(time.Now()),
				MetricName: aws.String("NumberOfMessagesDeleted"),
				Namespace:  aws.String("AWS/SQS"),
				Dimensions: []types.Dimension{
					types.Dimension{
						Name:  aws.String("QueueName"),
						Value: aws.String(queueName),
					},
				},
				Period: aws.Int32(3600),
				Statistics: []types.Statistic{
					types.StatisticSum,
				},
			})

			if err != nil {
				log.Warnf("Couldn't fetch invocations metric for %s", queueName)
			}

			nbOfMessagesDeleted := 0.0
			if metricsNbOfMessagesDeletedOutput != nil && len(metricsNbOfMessagesDeletedOutput.Datapoints) > 0 {
				nbOfMessagesDeleted = *metricsNbOfMessagesDeletedOutput.Datapoints[0].Sum
			}

			monthlyCost := 0.0

			if (nbOfMessagesSent + nbOfMessagesReceived + nbOfMessagesDeleted) > 1000000 {
				monthlyCost = ((nbOfMessagesSent + nbOfMessagesReceived + nbOfMessagesDeleted) - 1000000) * 0.40
			}

			outputTags, err := sqsClient.ListQueueTags(ctx, &sqs.ListQueueTagsInput{
				QueueUrl: &queue,
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

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "SQS",
				ResourceId: queue,
				Region:     client.AWSClient.Region,
				Name:       queueName,
				Cost:       monthlyCost,
				Metadata: map[string]string{
					"serviceCost": fmt.Sprint(serviceCost),
				},
				Tags:       tags,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/sqs/v2/home?region=%s#/queues/%s", client.AWSClient.Region, client.AWSClient.Region, queue),
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
		"service":   "SQS",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
