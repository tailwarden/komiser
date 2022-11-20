package lambda

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

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
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "Lambda",
				ResourceId: *o.FunctionArn,
				Region:     client.AWSClient.Region,
				Name:       *o.FunctionName,
				Cost:       0,
				Metadata: map[string]string{
					"runtime": fmt.Sprintf("%s", o.Runtime),
				},
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
	}).Debugf("Fetched resources")
	return resources, nil
}
