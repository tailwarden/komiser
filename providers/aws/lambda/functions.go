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
	lambdaTypes "github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/aws/aws-sdk-go-v2/service/pricing/types"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
	"github.com/tailwarden/komiser/utils"
)

const (
	freeTierInvocations = 1000000
	freeTierDuration    = 400000
)

func Functions(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var config lambda.ListFunctionsInput
	resources := make([]models.Resource, 0)
	cloudwatchClient := cloudwatch.NewFromConfig(*client.AWSClient)
	lambdaClient := lambda.NewFromConfig(*client.AWSClient)

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "AWS Lambda")
	if err != nil {
		log.Warnln("Couldn't fetch AWS Lambda cost and usage:", err)
	}

	tempRegion := client.AWSClient.Region
	client.AWSClient.Region = "us-east-1"
	pricingClient := pricing.NewFromConfig(*client.AWSClient)
	client.AWSClient.Region = tempRegion

	pricingOutput, err := pricingClient.GetProducts(ctx, &pricing.GetProductsInput{
		ServiceCode: aws.String("AWSLambda"),
		Filters: []types.Filter{
			{
				Field: aws.String("regionCode"),
				Value: aws.String(client.AWSClient.Region),
				Type:  types.FilterTypeTermMatch,
			},
		},
	})
	if err != nil {
		log.Errorf("ERROR: Couldn't fetch pricing info for AWS Lambda: %v", err)
		return resources, err
	}

	priceMap, err := awsUtils.GetPriceMap(pricingOutput, "group")
	if err != nil {
		log.Errorf("ERROR: Failed to calculate cost per month: %v", err)
		return resources, err
	}

	for {
		output, err := lambdaClient.ListFunctions(context.Background(), &config)
		if err != nil {
			return resources, err
		}

		for _, o := range output.Functions {
			archSuffix := ""
			if o.Architectures[0] == lambdaTypes.ArchitectureArm64 {
				archSuffix = "-ARM"
			}

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
				log.Warnf("Couldn't fetch invocations metric for %s: %v", *o.FunctionName, err)
			}

			invocations := 0.0
			if metricsInvocationsOutput != nil && len(metricsInvocationsOutput.Datapoints) > 0 {
				invocations = *metricsInvocationsOutput.Datapoints[0].Sum
			}
			if invocations > freeTierInvocations {
				invocations -= freeTierInvocations
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
				log.Warnf("Couldn't fetch duration metric for %s: %v", *o.FunctionName, err)
			}

			duration := 0.0
			if metricsDurationOutput != nil && len(metricsDurationOutput.Datapoints) > 0 {
				duration = *metricsDurationOutput.Datapoints[0].Average
			}
			totalDuration := ((invocations * duration) * (float64(*o.MemorySize))) / (1024 * 1024)
			if totalDuration < freeTierDuration {
				totalDuration -= freeTierDuration
			}

			computeCharges := awsUtils.GetCost(priceMap["AWS-Lambda-Duration"+archSuffix], totalDuration)
			requestCharges := awsUtils.GetCost(priceMap["AWS-Lambda-Requests"+archSuffix], invocations)
			monthlyCost := computeCharges + requestCharges

			tags := make([]models.Tag, 0)
			tagsResp, err := lambdaClient.ListTags(context.Background(), &lambda.ListTagsInput{
				Resource: o.FunctionArn,
			})

			if err == nil {
				for key, value := range tagsResp.Tags {
					tags = append(tags, models.Tag{
						Key:   key,
						Value: value,
					})
				}
			}

			relations := getLambdaRelations(*client.AWSClient, o)

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "Lambda",
				ResourceId: *o.FunctionArn,
				Region:     client.AWSClient.Region,
				Name:       *o.FunctionName,
				Cost:       monthlyCost,
				Metadata: map[string]string{
					"runtime":     string(o.Runtime),
					"serviceCost": fmt.Sprint(serviceCost),
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

func getLambdaRelations(config aws.Config, lambda lambdaTypes.FunctionConfiguration) (rel []models.Link) {
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
