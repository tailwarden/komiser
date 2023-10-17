package dynamodb

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/aws/aws-sdk-go-v2/service/pricing/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
	// "github.com/tailwarden/komiser/utils"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func int64PtrToFloat64(i *int64) float64 {
    if i == nil {
        return 0.0  // or any default value you prefer
    }
    return float64(*i)
}


func Tables(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config dynamodb.ListTablesInput
	dynamodbClient := dynamodb.NewFromConfig(*client.AWSClient)
	pricingClient := pricing.NewFromConfig(*client.AWSClient)

	pricingOutput, err := pricingClient.GetProducts(ctx, &pricing.GetProductsInput{
	    ServiceCode: aws.String("AmazonDynamoDB"),
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


	output, err := dynamodbClient.ListTables(ctx, &config)
	if err != nil {
		return resources, err
	}

	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return resources, err
	}

	accountId := stsOutput.Account

	for _, table := range output.TableNames {
		resourceArn := fmt.Sprintf("arn:aws:dynamodb:%s:%s:table/%s", client.AWSClient.Region, *accountId, table)
		outputTags, err := dynamodbClient.ListTagsOfResource(ctx, &dynamodb.ListTagsOfResourceInput{
			ResourceArn: &resourceArn,
		})

		tags := make([]Tag, 0)

		if err == nil {
			for _, tag := range outputTags.Tags {
				tags = append(tags, Tag{
					Key:   *tag.Key,
					Value: *tag.Value,
				})
			}
		}

		tableDetails, err := dynamodbClient.DescribeTable(ctx, &dynamodb.DescribeTableInput{
			TableName: aws.String(table),
		})

		if err != nil {
			return resources, err
		}

		var provisionedRCUs *int64
		var provisionedWCUs *int64

		if tableDetails != nil && tableDetails.Table != nil && tableDetails.Table.ProvisionedThroughput != nil {
			provisionedRCUs = tableDetails.Table.ProvisionedThroughput.ReadCapacityUnits
			provisionedWCUs = tableDetails.Table.ProvisionedThroughput.WriteCapacityUnits
		}

		log.Errorf("ERROR: Failed to calculate cost per month: %v", err)

		RCUCharges := awsUtils.GetCost(priceMap["AWS-DynamoDB-ProvisionedReadCapacityUnits"], int64PtrToFloat64(provisionedRCUs))
		PWUCharges := awsUtils.GetCost(priceMap["AWS-DynamoDB-ProvisionedWriteCapacityUnits"], int64PtrToFloat64(provisionedWCUs))
		monthlyCost := RCUCharges + PWUCharges

		resources = append(resources, Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "DynamoDB",
			ResourceId: resourceArn,
			Region:     client.AWSClient.Region,
			Name:       table,
			Cost:       monthlyCost,
			Tags:       tags,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/dynamodbv2/home?region=%s#table?initialTagKey=&name=%s", client.AWSClient.Region, client.AWSClient.Region, table),
		})
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "DynamoDB",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
