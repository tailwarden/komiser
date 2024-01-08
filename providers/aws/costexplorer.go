package aws

import (
	"context"
	"encoding/json"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/providers"
)

const costExplorerCacheFile = "cost_explorer_cache.json"

// readCostExplorerCache reads cost explorer cache data from the file.
func readCostExplorerCache() ([]byte, error) {
	file, err := os.ReadFile(costExplorerCacheFile)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// writeCostExplorerCache writes cost explorer cache data to the file.
func writeCostExplorerCache(data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(costExplorerCacheFile, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func getCostexplorerOutput(ctx context.Context, client providers.ProviderClient, start, end string) ([]*costexplorer.GetCostAndUsageOutput, error) {
	costexplorerOutputList := []*costexplorer.GetCostAndUsageOutput{}
	costexplorerClient := costexplorer.NewFromConfig(*client.AWSClient)
	var nextPageToken *string
	for {
		costexplorerOutput, err := costexplorerClient.GetCostAndUsage(ctx, &costexplorer.GetCostAndUsageInput{
			Granularity: "DAILY",
			Metrics:     []string{"UnblendedCost"},
			TimePeriod: &types.DateInterval{
				Start: aws.String(start),
				End:   aws.String(end),
			},
			GroupBy: []types.GroupDefinition{
				{
					Key:  aws.String("SERVICE"),
					Type: "DIMENSION",
				},
				{
					Key:  aws.String("REGION"),
					Type: "DIMENSION",
				},
			},
			NextPageToken: nextPageToken,
		})
		if err != nil {
			log.Warn("Couldn't fetch cost and usage data:", err)
			return nil, err
		}

		costexplorerOutputList = append(costexplorerOutputList, costexplorerOutput)

		if aws.ToString(costexplorerOutput.NextPageToken) == "" {
			break
		}

		nextPageToken = costexplorerOutput.NextPageToken
	}
	return costexplorerOutputList, nil
}
