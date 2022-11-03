package instances

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	. "github.com/mlabouardy/komiser/models"
	. "github.com/mlabouardy/komiser/providers"
)

func Topics(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config sns.ListTopicsInput
	snsClient := sns.NewFromConfig(*client.AWSClient)

	for {
		output, err := snsClient.ListTopics(context.Background(), &config)
		if err != nil {
			return resources, err
		}

		for _, topic := range output.Topics {
			resources = append(resources, Resource{
				Provider:  "AWS",
				Account:   client.Name,
				Service:   "SNS",
				Region:    client.AWSClient.Region,
				Name:      *topic.TopicArn,
				Cost:      0,
				FetchedAt: time.Now(),
			})
		}

		if aws.ToString(config.NextToken) == "" {
			break
		}

		config.NextToken = output.NextToken
	}
	log.Printf("[%s] Fetched %d AWS SNS topics from %s\n", client.Name, len(resources), client.AWSClient.Region)
	return resources, nil
}
