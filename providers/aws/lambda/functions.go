package instances

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	. "github.com/mlabouardy/komiser/models"
	. "github.com/mlabouardy/komiser/providers"
)

func Functions(ctx context.Context, client ProviderClient) ([]Resource, error) {
	var config lambda.ListFunctionsInput
	resources := make([]Resource, 0)
	lambdaClient := lambda.NewFromConfig(*client.AWSClient)
	for {
		output, err := lambdaClient.ListFunctions(context.Background(), &config)
		if err != nil {
			return resources, err
		}

		for _, o := range output.Functions {
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

			resources = append(resources, Resource{
				Provider: "AWS",
				Account:  client.Name,
				Service:  "Lambda",
				Region:   client.AWSClient.Region,
				Name:     *o.FunctionName,
				Cost:     0,
				Metadata: map[string]string{
					"runtime": fmt.Sprintf("%s", o.Runtime),
				},
				FetchedAt: time.Now(),
				Tags:      tags,
			})
		}

		if aws.ToString(output.NextMarker) == "" {
			break
		}

		config.Marker = output.NextMarker
	}
	log.Printf("[%s] Fetched %d AWS Lambda functions from %s\n", client.Name, len(resources), client.AWSClient.Region)
	return resources, nil
}
