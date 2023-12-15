package kinesis

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	cloudwatchTypes "github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	pricingTypes "github.com/aws/aws-sdk-go-v2/service/pricing/types"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
	"github.com/tailwarden/komiser/utils"
)

type KinesisClient interface {
	ListStreamConsumers(ctx context.Context, params *kinesis.ListStreamConsumersInput, optFns ...func(*kinesis.Options)) (*kinesis.ListStreamConsumersOutput, error)
}

type PricingClient interface {
	GetProducts(ctx context.Context, params *pricing.GetProductsInput, optFns ...func(*pricing.Options)) (*pricing.GetProductsOutput, error)
}

type CloudwatchClient interface {
	GetMetricStatistics(ctx context.Context, params *cloudwatch.GetMetricStatisticsInput, optFns ...func(*cloudwatch.Options)) (*cloudwatch.GetMetricStatisticsOutput, error)
}

func Streams(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	kinesisClient := kinesis.NewFromConfig(*client.AWSClient)
	cloudwatchClient := cloudwatch.NewFromConfig(*client.AWSClient)

	tempRegion := client.AWSClient.Region
	client.AWSClient.Region = "us-east-1"
	pricingClient := pricing.NewFromConfig(*client.AWSClient)
	client.AWSClient.Region = tempRegion

	priceMap, err := retrievePriceMap(ctx, pricingClient, client.AWSClient.Region)
	if err != nil {
		return resources, err
	}

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "Amazon Kinesis")
	if err != nil {
		log.Warnln("Couldn't fetch Amazon Kinesis cost and usage:", err)
	}
	var config kinesis.ListStreamsInput
	for {
		output, err := kinesisClient.ListStreams(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, stream := range output.StreamSummaries {
			tags := make([]Tag, 0)
			tagsResp, err := kinesisClient.ListTagsForStream(context.Background(), &kinesis.ListTagsForStreamInput{
				StreamARN: stream.StreamARN,
			})
			if err == nil {
				for _, t := range tagsResp.Tags {
					tags = append(tags, Tag{
						Key:   aws.ToString(t.Key),
						Value: aws.ToString(t.Value),
					})
				}
			} else {
				log.Warn("Failed to fetch tags for kinesis streams")
			}
			summaryOutput, err := kinesisClient.DescribeStreamSummary(ctx, &kinesis.DescribeStreamSummaryInput{
				StreamARN: stream.StreamARN,
			})
			if err != nil {
				return resources, err
			}
			totalPutRecords := retrievePutRecords(ctx, cloudwatchClient, stream.StreamName)
			cost, err := calculateCostOfKinesisDataStream(summaryOutput.StreamDescriptionSummary, totalPutRecords, priceMap)
			if err != nil {
				return resources, err
			}
			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "Kinesis Stream",
				ResourceId: *stream.StreamARN,
				Region:     client.AWSClient.Region,
				Name:       *stream.StreamName,
				Cost:       cost,
				Metadata: map[string]string{
					"serviceCost": fmt.Sprint(serviceCost),
				},
				CreatedAt:  *stream.StreamCreationTimestamp,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/kinesis/home?region=%s#/streams/details/%s", client.AWSClient.Region, client.AWSClient.Region, *stream.StreamName),
			})
			consumers, err := getStreamConsumers(ctx, kinesisClient, stream, client.Name, client.AWSClient.Region)
			if err != nil {
				return resources, err
			}
			resources = append(resources, consumers...)
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
		"service":   "Kinesis Stream",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}

func getStreamConsumers(ctx context.Context, kinesisClient KinesisClient, stream types.StreamSummary, clientName, region string) ([]Resource, error) {
	resources := make([]Resource, 0)
	config := kinesis.ListStreamConsumersInput{
		StreamARN: aws.String(aws.ToString(stream.StreamARN)),
	}

	for {
		output, err := kinesisClient.ListStreamConsumers(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, consumer := range output.Consumers {
			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    clientName,
				Service:    "Kinesis EFO Consumer",
				ResourceId: *consumer.ConsumerARN,
				Region:     region,
				Name:       *consumer.ConsumerName,
				Cost:       0,
				CreatedAt:  *consumer.ConsumerCreationTimestamp,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/kinesis/home?region=%s#/streams/details/%s/registeredConsumers/%s", region, region, aws.ToString(stream.StreamName), *consumer.ConsumerName),
			})
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}

	return resources, nil
}

func retrievePriceMap(ctx context.Context, pricingClient PricingClient, region string) (map[string][]awsUtils.PriceDimensions, error) {
	pricingOutput, err := pricingClient.GetProducts(ctx, &pricing.GetProductsInput{
		ServiceCode: aws.String("AmazonKinesis"),
		Filters: []pricingTypes.Filter{
			{
				Field: aws.String("regionCode"),
				Value: aws.String(region),
				Type:  pricingTypes.FilterTypeTermMatch,
			},
		},
	})
	if err != nil {
		log.Errorf("ERROR: Couldn't fetch pricing info for AWS Kinesis: %v", err)
		return nil, err
	}
	priceMap, err := awsUtils.GetPriceMap(pricingOutput, "group")
	if err != nil {
		log.Errorf("ERROR: Failed to calculate cost per month: %v", err)
		return nil, err
	}
	return priceMap, nil
}

func retrievePutRecords(ctx context.Context, cloudwatchClient CloudwatchClient, streamName *string) float64 {
	output, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
		StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
		EndTime:    aws.Time(time.Now()),
		MetricName: aws.String("PutRecords.SuccessfulRecords"),
		Namespace:  aws.String("AWS/Kinesis"),
		Dimensions: []cloudwatchTypes.Dimension{
			{
				Name:  aws.String("StreamName"),
				Value: streamName,
			},
		},
		Period: aws.Int32(60 * 60 * 24 * 30), // 30 days
		Statistics: []cloudwatchTypes.Statistic{
			cloudwatchTypes.StatisticSum,
		},
	})
	if err != nil {
		log.Warnf("Couldn't fetch metrics for %s: %v", *streamName, err)
		return 0.0
	}
	if output == nil || len(output.Datapoints) == 0 {
		log.Warnf("Couldn't fetch metrics for %s: %v", *streamName, err)
		return 0.0
	}
	return *output.Datapoints[0].Sum
}

func calculateCostOfKinesisDataStream(summary *types.StreamDescriptionSummary, totalPutRecords float64, priceMap map[string][]awsUtils.PriceDimensions) (float64, error) {
	if summary.StreamModeDetails.StreamMode == types.StreamModeProvisioned {
		startOfMonth := utils.BeginningOfMonth(time.Now())
		hourlyUsage := int32(0)
		if (*summary.StreamCreationTimestamp).Before(startOfMonth) {
			hourlyUsage = int32(time.Since(startOfMonth).Hours())
		} else {
			hourlyUsage = int32(time.Since(*summary.StreamCreationTimestamp).Hours())
		}
		nbShards := aws.ToInt32(summary.OpenShardCount)
		shardCost := awsUtils.GetCost(priceMap["Provisioned shard hour"], float64(hourlyUsage*nbShards))

		putRecordsCost := awsUtils.GetCost(priceMap["Payload Units"], totalPutRecords)

		return shardCost + putRecordsCost, nil
	}
	return 0.0, nil
}
