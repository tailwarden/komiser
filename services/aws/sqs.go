package aws

import (
	"context"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (aws AWS) DescribeQueues(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		queues, err := aws.getSQS(cfg, region.Name)
		if err != nil {
			return 0, err
		}
		sum += int64(len(queues))
	}
	return sum, nil
}

func (aws AWS) getSQS(cfg aws.Config, region string) ([]Queue, error) {
	cfg.Region = region
	svc := sqs.New(cfg)
	req := svc.ListQueuesRequest(&sqs.ListQueuesInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return []Queue{}, err
	}
	listOfQueues := make([]Queue, 0, len(result.QueueUrls))
	for _, queue := range result.QueueUrls {
		listOfQueues = append(listOfQueues, Queue{
			Name: queue,
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
		svc := sqs.New(cfg)
		req := svc.ListQueuesRequest(&sqs.ListQueuesInput{})
		result, err := req.Send(context.Background())
		if err != nil {
			return data, err
		}

		for _, queue := range result.QueueUrls {
			queueName := strings.Split(queue, "/")[len(strings.Split(queue, "/"))-1]

			cloudwatchClient := cloudwatch.New(cfg)
			reqCloudwatch := cloudwatchClient.GetMetricStatisticsRequest(&cloudwatch.GetMetricStatisticsInput{
				Namespace:  aws.String("AWS/SQS"),
				MetricName: aws.String("NumberOfMessagesSent"),
				StartTime:  aws.Time(time.Now().AddDate(0, 0, -7)),
				EndTime:    aws.Time(time.Now()),
				Period:     aws.Int64(86400),
				Statistics: []cloudwatch.Statistic{
					cloudwatch.StatisticSum,
				},
				Dimensions: []cloudwatch.Dimension{
					cloudwatch.Dimension{
						Name:  aws.String("QueueName"),
						Value: aws.String(queueName),
					},
				},
			})
			resultCloudWatch, err := reqCloudwatch.Send(context.Background())
			if err != nil {
				return data, err
			}

			for _, metric := range resultCloudWatch.Datapoints {
				key := (*metric.Timestamp).Format("2006-01-02")
				data[0].Datapoints[key] += *metric.Sum
			}

			reqCloudwatch2 := cloudwatchClient.GetMetricStatisticsRequest(&cloudwatch.GetMetricStatisticsInput{
				Namespace:  aws.String("AWS/SQS"),
				MetricName: aws.String("NumberOfMessagesDeleted"),
				StartTime:  aws.Time(time.Now().AddDate(0, 0, -7)),
				EndTime:    aws.Time(time.Now()),
				Period:     aws.Int64(86400),
				Statistics: []cloudwatch.Statistic{
					cloudwatch.StatisticSum,
				},
				Dimensions: []cloudwatch.Dimension{
					cloudwatch.Dimension{
						Name:  aws.String("QueueName"),
						Value: aws.String(queueName),
					},
				},
			})
			resultCloudWatch2, err := reqCloudwatch2.Send(context.Background())
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
