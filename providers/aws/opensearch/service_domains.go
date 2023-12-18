package opensearch

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/opensearch"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
)

func ServiceDomains(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	openSearchClient := opensearch.NewFromConfig(*client.AWSClient)
	input := &opensearch.ListDomainNamesInput{}

	output, err := openSearchClient.ListDomainNames(ctx, input)
	if err != nil {
		return resources, err
	}

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "OpenSearch")
	if err != nil {
		log.Warnln("Couldn't fetch OpenSearch cost and usage:", err)
	}

	for _, domainName := range output.DomainNames {
		domainConfig, err := openSearchClient.DescribeDomain(ctx, &opensearch.DescribeDomainInput{
			DomainName: domainName.DomainName,
		})
		if err != nil {
			return resources, err
		}
		domain := domainConfig.DomainStatus

		resources = append(resources, Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "OpenSearch Service Domain",
			ResourceId: aws.ToString(domain.DomainId),
			Region:     client.AWSClient.Region,
			Name:       aws.ToString(domain.DomainName),
			Cost:       0,
			Metadata: map[string]string{
				"serviceCost": fmt.Sprint(serviceCost),
			},
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/aos/home?region=%s#/opensearch/domains/%s", client.AWSClient.Region, client.AWSClient.Region, aws.ToString(domain.DomainName)),
		})
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "OpenSearch Service Domain",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
