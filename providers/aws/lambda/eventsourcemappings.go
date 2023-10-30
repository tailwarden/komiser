package lambda

import (
        "context"
        "fmt"
        "strings"
        "time"

        log "github.com/sirupsen/logrus"

        "github.com/aws/aws-sdk-go-v2/aws"
        "github.com/aws/aws-sdk-go-v2/service/lambda"
        "github.com/tailwarden/komiser/models"
        "github.com/tailwarden/komiser/providers"
)

func EventSourceMappings(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
        resources := make([]models.Resource, 0)
        lambdaClient := lambda.NewFromConfig(*client.AWSClient)
        mappingsPerPage := aws.Int32(50)
        params := &lambda.ListEventSourceMappingsInput{
                MaxItems: mappingsPerPage,
        }

        paginator := lambda.NewListEventSourceMappingsPaginator(lambdaClient, params, func(options *lambda.ListEventSourceMappingsPaginatorOptions) {
                options.Limit = *mappingsPerPage
        })

        for paginator.HasMorePages() {
                
                output, err := paginator.NextPage(ctx)
                if err != nil {
                        log.Errorf("ERROR: Error occurred while retrieving EventSourceMappings page: %v", err)
                        return resources, err
                }

                for _, mapping := range output.EventSourceMappings {

                        lambdaNameSplit := strings.Split(*mapping.FunctionArn, ":")
                        lambdaName := lambdaNameSplit[len(lambdaNameSplit)-1]

                        resources = append(resources, models.Resource{
                                Provider:   "AWS",
                                Account:    client.Name,
                                Service:    "EventSourceMapping",
                                ResourceId: *mapping.UUID,
                                Region:     client.AWSClient.Region,
                                Name:       *mapping.UUID,
                                Cost:       0.0,
                                Metadata: map[string]string{
                                        "lambda": *mapping.FunctionArn,
                                        "source": *mapping.EventSourceArn,
                                },
                                FetchedAt: time.Now(),
                                Link:      fmt.Sprintf("https://%s.console.aws.amazon.com/lambda/home?region=%s#/functions/%s", client.AWSClient.Region, client.AWSClient.Region, lambdaName),
                        })
                }
        }

        log.WithFields(log.Fields{
                "provider":  "AWS",
                "account":   client.Name,
                "region":    client.AWSClient.Region,
                "service":   "EventSourceMapping",
                "resources": len(resources),
        }).Info("Fetched resources")
        return resources, nil
}