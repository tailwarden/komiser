package lambda

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	cloudwatchTypes "github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/utils"
)

func Functions(ctx context.Context, client ProviderClient) ([]Resource, error) {
	var config lambda.ListFunctionsInput
	resources := make([]Resource, 0)
	cloudwatchClient := cloudwatch.NewFromConfig(*client.AWSClient)
	lambdaClient := lambda.NewFromConfig(*client.AWSClient)
	for {
		output, err := lambdaClient.ListFunctions(context.Background(), &config)
		if err != nil {
			return resources, err
		}

		for _, o := range output.Functions {
			metricsInvocationsOutput, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
				StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
				EndTime:    aws.Time(time.Now()),
				MetricName: aws.String("Invocations"),
				Namespace:  aws.String("AWS/Lambda"),
				Dimensions: []cloudwatchTypes.Dimension{
					{
						Name:  aws.String("FunctionName"),
						Value: o.FunctionName,
					},
				},
				Period: aws.Int32(3600),
				Statistics: []cloudwatchTypes.Statistic{
					cloudwatchTypes.StatisticSum,
				},
			})

			if err != nil {
				log.Warnf("Couldn't fetch invocations metric for %s", *o.FunctionName)
			}

			invocations := 0.0
			if metricsInvocationsOutput != nil && len(metricsInvocationsOutput.Datapoints) > 0 {
				invocations = *metricsInvocationsOutput.Datapoints[0].Sum
			}

			metricsDurationOutput, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
				StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
				EndTime:    aws.Time(time.Now()),
				MetricName: aws.String("Duration"),
				Namespace:  aws.String("AWS/Lambda"),
				Dimensions: []cloudwatchTypes.Dimension{
					{
						Name:  aws.String("FunctionName"),
						Value: o.FunctionName,
					},
				},
				Period: aws.Int32(3600),
				Statistics: []cloudwatchTypes.Statistic{
					cloudwatchTypes.StatisticAverage,
				},
			})
			if err != nil {
				log.Warnf("Couldn't fetch duration metric for %s", *o.FunctionName)
			}

			duration := 0.0
			if metricsDurationOutput != nil && len(metricsDurationOutput.Datapoints) > 0 {
				duration = *metricsDurationOutput.Datapoints[0].Average
			}

			computeCharges := (((invocations * duration) * (float64(*o.MemorySize))) / 1024) * 0.0000000083
			requestCharges := invocations * 0.2
			monthlyCost := computeCharges + requestCharges

			tags := make([]Tag, 0)
			tagsResp, err := lambdaClient.ListTags(context.Background(), &lambda.ListTagsInput{
				Resource: o.FunctionArn,
			})

			if err == nil {
				for key, value := range tagsResp.Tags {
					tags = append(tags, Tag{
						Key:   key,
						Value: value,
					})
				}
			}

			relations := getLambdaRelations(*client.AWSClient, o)

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "Lambda",
				ResourceId: *o.FunctionArn,
				Region:     client.AWSClient.Region,
				Name:       *o.FunctionName,
				Cost:       monthlyCost,
				Metadata: map[string]string{
					"runtime": string(o.Runtime),
				},
				Relations: relations,
				FetchedAt: time.Now(),
				Tags:      tags,
				Link:      fmt.Sprintf("https://%s.console.aws.amazon.com/lambda/home?region=%s#/functions/%s", client.AWSClient.Region, client.AWSClient.Region, *o.FunctionName),
			})
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
		"service":   "Lambda",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}

func getLambdaRelations(config aws.Config, lambda types.FunctionConfiguration) (rel []models.Link) {
	// Get associated IAM roles
	if lambda.Role != nil {
		iamClient := iam.NewFromConfig(config)
		roleOutput, err := iamClient.GetRole(context.Background(), &iam.GetRoleInput{
			RoleName: lambda.Role,
		})
		if err != nil {
			return rel
		}

		rel = append(rel, models.Link{
			ResourceID: *roleOutput.Role.RoleId,
			Type:       "IAM Role",
			Name:       *roleOutput.Role.Arn,
			Relation:   "USES",
		})
	}
	return rel
}
