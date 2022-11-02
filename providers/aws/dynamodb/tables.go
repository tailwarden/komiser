package instances

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	. "github.com/mlabouardy/komiser/models"
)

func Tables(ctx context.Context, cfg aws.Config, account string) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config dynamodb.ListTablesInput
	dynamodbClient := dynamodb.NewFromConfig(cfg)
	output, err := dynamodbClient.ListTables(context.Background(), &config)
	if err != nil {
		return resources, err
	}

	for _, table := range output.TableNames {
		resources = append(resources, Resource{
			Provider:  "AWS",
			Account:   account,
			Service:   "DynamoDB",
			Region:    cfg.Region,
			Name:      table,
			Cost:      0,
			FetchedAt: time.Now(),
		})
	}
	log.Printf("[%s] Fetched %d AWS DynamoDB tables from %s\n", account, len(resources), cfg.Region)
	return resources, nil
}
