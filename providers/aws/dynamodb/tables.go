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
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
)


func Tables(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config dynamodb.ListTablesInput
	dynamodbClient := dynamodb.NewFromConfig(*client.AWSClient)

	var monthlyCost float64 = 0.0
	// there is something strange going on when using pricing client with regions other than us-east-1
	// https://discord.com/channels/932683789384183808/1117721764957536318/1162338171435090032
	oldRegion := client.AWSClient.Region
	client.AWSClient.Region = "us-east-1"
	pricingClient := pricing.NewFromConfig(*client.AWSClient)
	client.AWSClient.Region = oldRegion
	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "Amazon DynamoDB")
	if err != nil {
		log.Warnln("Couldn't fetch Amazon DynamoDB cost and usage:", err)
	}

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
		log.Errorf("ERROR: Couldn't fetch pricing info for AWS DynamoDB: %v", err)
	}

	priceMap, err := awsUtils.GetPriceMap(pricingOutput, "group")

	if err != nil {
		log.Errorf("ERROR: Failed to fetch pricing map: %v", err)
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
			log.Errorf("ERROR: Failed to query DynamoDB table details: %v", err)
		}

		if tableDetails.Table != nil && tableDetails.Table.ProvisionedThroughput != nil {
			provisionedRCUs := tableDetails.Table.ProvisionedThroughput.ReadCapacityUnits
			provisionedWCUs := tableDetails.Table.ProvisionedThroughput.WriteCapacityUnits

			RCUCharges := awsUtils.GetCost(priceMap["DDB-ReadUnits"], awsUtils.Int64PtrToFloat64(provisionedRCUs))
			WCUCharges := awsUtils.GetCost(priceMap["DDB-WriteUnits"], awsUtils.Int64PtrToFloat64(provisionedWCUs))

			monthlyCost = RCUCharges + WCUCharges
		}

		resources = append(resources, Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "DynamoDB",
			ResourceId: resourceArn,
			Region:     client.AWSClient.Region,
			Name:       table,
			Cost:       monthlyCost,
			Metadata: map[string]string{
				"serviceCost": fmt.Sprint(serviceCost),
			},
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
