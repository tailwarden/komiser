package aws

import (
	"context"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (aws AWS) DescribeQueues(cfg aws.Config) ([]Queue, error) {
	queues := make([]Queue, 0)
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return queues, err
	}
	for _, region := range regions {
		queuesResp, err := aws.getSQS(cfg, region.Name)
		if err != nil {
			return queues, err
		}

		for _, queue := range queuesResp {
			queues = append(queues, queue)
		}
	}
	return queues, nil
}

func (aws AWS) getSQS(cfg aws.Config, region string) ([]Queue, error) {
	cfg.Region = region
	svc := sqs.NewFromConfig(cfg)
	result, err := svc.ListQueues(context.Background(), &sqs.ListQueuesInput{})
	if err != nil {
		return []Queue{}, err
	}
	listOfQueues := make([]Queue, 0, len(result.QueueUrls))
	for _, queue := range result.QueueUrls {
		listOfQueues = append(listOfQueues, Queue{
			Name:   queue,
			ID:     queue,
			Region: region,
		})
	}
	return listOfQueues, nil
}

type SQSMetric struct {
	Metric     string
	Datapoints map[string]float64
}

func (awsClient AWS) GetNumberOfMessagesSentAndDeletedSQS(cfg aws.Config) ([]SQSMetric, error) {
	data := []SQSMetric{
		SQSMetric{
			Metric:     "Sent",
			Datapoints: map[string]float64{},
		},
		SQSMetric{
			Metric:     "Deleted",
			Datapoints: map[string]float64{},
		},
	}

	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return data, err
	}

	for _, region := range regions {
		cfg.Region = region.Name
		svc := sqs.NewFromConfig(cfg)
		result, err := svc.ListQueues(context.Background(), &sqs.ListQueuesInput{})
		if err != nil {
			return data, err
		}

		for _, queue := range result.QueueUrls {
			queueName := strings.Split(queue, "/")[len(strings.Split(queue, "/"))-1]

			cloudwatchClient := cloudwatch.NewFromConfig(cfg)
			resultCloudWatch, err := cloudwatchClient.GetMetricStatistics(context.Background(), &cloudwatch.GetMetricStatisticsInput{
				Namespace:  aws.String("AWS/SQS"),
				MetricName: aws.String("NumberOfMessagesSent"),
				StartTime:  aws.Time(time.Now().AddDate(0, 0, -7)),
				EndTime:    aws.Time(time.Now()),
				Period:     aws.Int32(86400),
				Statistics: []types.Statistic{
					types.StatisticSum,
				},
				Dimensions: []types.Dimension{
					types.Dimension{
						Name:  aws.String("QueueName"),
						Value: aws.String(queueName),
					},
				},
			})
			if err != nil {
				return data, err
			}

			for _, metric := range resultCloudWatch.Datapoints {
				key := (*metric.Timestamp).Format("2006-01-02")
				data[0].Datapoints[key] += *metric.Sum
			}

			resultCloudWatch2, err := cloudwatchClient.GetMetricStatistics(context.Background(), &cloudwatch.GetMetricStatisticsInput{
				Namespace:  aws.String("AWS/SQS"),
				MetricName: aws.String("NumberOfMessagesDeleted"),
				StartTime:  aws.Time(time.Now().AddDate(0, 0, -7)),
				EndTime:    aws.Time(time.Now()),
				Period:     aws.Int32(86400),
				Statistics: []types.Statistic{
					types.StatisticSum,
				},
				Dimensions: []types.Dimension{
					types.Dimension{
						Name:  aws.String("QueueName"),
						Value: aws.String(queueName),
					},
				},
			})
			if err != nil {
				return data, err
			}

			for _, metric := range resultCloudWatch2.Datapoints {
				key := (*metric.Timestamp).Format("2006-01-02")
				data[1].Datapoints[key] += *metric.Sum
			}

		}
	}

	return data, nil
}
