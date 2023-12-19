package servicecatalog

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
)

func Products(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	serviceCatalogsClient := servicecatalog.NewFromConfig(*client.AWSClient)

	input := &servicecatalog.SearchProductsAsAdminInput{}

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "ServiceCatalog")
	if err != nil {
		log.Warnln("Couldn't fetch ServiceCatalog cost and usage:", err)
	}

	for {
		output, err := serviceCatalogsClient.SearchProductsAsAdmin(ctx, input)
		if err != nil {
			return resources, err
		}

		for _, product := range output.ProductViewDetails {
			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "Service Catalog Application",
				ResourceId: aws.ToString(product.ProductViewSummary.ProductId),
				Region:     client.AWSClient.Region,
				Name:       aws.ToString(product.ProductViewSummary.Name),
				Cost:       0,
				Metadata: map[string]string{
					"serviceCost": fmt.Sprint(serviceCost),
				},
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/servicecatalog/home?region=%s#/admin-products/%s", client.AWSClient.Region, client.AWSClient.Region, aws.ToString(product.ProductViewSummary.ProductId)),
			})
		}

		if output.NextPageToken == nil {
			break
		}

		input.PageToken = output.NextPageToken
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "Service Catalog",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
