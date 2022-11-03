package instances

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	. "github.com/mlabouardy/komiser/models"
	. "github.com/mlabouardy/komiser/providers"
)

func Queues(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config sqs.ListQueuesInput
	sqsClient := sqs.NewFromConfig(*client.AWSClient)

	for {
		output, err := sqsClient.ListQueues(context.Background(), &config)
		if err != nil {
			return resources, err
		}

		for _, queue := range output.QueueUrls {
			resources = append(resources, Resource{
				Provider:  "AWS",
				Account:   client.Name,
				Service:   "SQS",
				Region:    client.AWSClient.Region,
				Name:      queue,
				Cost:      0,
				FetchedAt: time.Now(),
			})
		}

		if aws.ToString(config.NextToken) == "" {
			break
		}

		config.NextToken = output.NextToken
	}
	log.Printf("[%s] Fetched %d AWS SQS queues from %s\n", client.Name, len(resources), client.AWSClient.Region)
	return resources, nil
}
