package instances

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	. "github.com/mlabouardy/komiser/models"
	. "github.com/mlabouardy/komiser/providers"
)

func Tables(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config dynamodb.ListTablesInput
	dynamodbClient := dynamodb.NewFromConfig(*client.AWSClient)
	output, err := dynamodbClient.ListTables(context.Background(), &config)
	if err != nil {
		return resources, err
	}

	for _, table := range output.TableNames {
		resources = append(resources, Resource{
			Provider:  "AWS",
			Account:   client.Name,
			Service:   "DynamoDB",
			Region:    client.AWSClient.Region,
			Name:      table,
			Cost:      0,
			FetchedAt: time.Now(),
		})
	}
	log.Printf("[%s] Fetched %d AWS DynamoDB tables from %s\n", client.Name, len(resources), client.AWSClient.Region)
	return resources, nil
}
