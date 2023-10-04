package lambda

import (
        "context"
        "fmt"
        "strings"
        "time"
        
        log "github.com/sirupsen/logrus"
        
        "github.com/aws/aws-sdk-go-v2/service/lambda"
        "github.com/tailwarden/komiser/models"
        "github.com/tailwarden/komiser/providers"
)

func EventSourceMappings(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
        var config lambda.ListEventSourceMappingsInput
        resources := make([]models.Resource, 0)
        lambdaClient := lambda.NewFromConfig(*client.AWSClient)

        result, err := lambdaClient.ListEventSourceMappings(context.Background(), &config)
        if err != nil {
                log.Errorf("ERROR: Failed to fetch EventSourceMappings: %v", err)
                return resources, err
        }

        for _, mapping := range result.EventSourceMappings {

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

        log.WithFields(log.Fields{
                "provider":  "AWS",
                "account":   client.Name,
                "region":    client.AWSClient.Region,
                "service":   "EventSourceMapping",
                "resources": len(resources),
        }).Info("Fetched resources")
        return resources, nil
}