package dynamodb

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func Tables(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config dynamodb.ListTablesInput
	dynamodbClient := dynamodb.NewFromConfig(*client.AWSClient)
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

		resources = append(resources, Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "DynamoDB",
			ResourceId: resourceArn,
			Region:     client.AWSClient.Region,
			Name:       table,
			Cost:       0,
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
