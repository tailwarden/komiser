package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	models "github.com/mlabouardy/komiser/models/aws"
)

func (awsClient AWS) DescribeDynamoDBTables(cfg awsConfig.Config) (map[string]interface{}, error) {
	outputThroughput := make(map[string]int, 0)
	sumTables := 0
	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return map[string]interface{}{}, err
	}
	for _, region := range regions {
		tables, err := awsClient.getDynamoDBTables(cfg, region.Name)
		if err != nil {
			return map[string]interface{}{}, err
		}
		sumTables += len(tables)
		for _, table := range tables {
			cfg.Region = region.Name
			svc := dynamodb.NewFromConfig(cfg)
			result, err := svc.DescribeTable(context.Background(), &dynamodb.DescribeTableInput{
				TableName: &table.Name,
			})
			if err != nil {
				return map[string]interface{}{}, err
			}
			outputThroughput["readCapacity"] += int(*result.Table.ProvisionedThroughput.ReadCapacityUnits)
			outputThroughput["writeCapacity"] += int(*result.Table.ProvisionedThroughput.WriteCapacityUnits)
		}
	}
	return map[string]interface{}{
		"throughput": outputThroughput,
		"total":      sumTables,
	}, nil
}

func (awsClient AWS) getDynamoDBTables(cfg awsConfig.Config, region string) ([]models.Table, error) {
	cfg.Region = region
	svc := dynamodb.NewFromConfig(cfg)
	result, err := svc.ListTables(context.Background(), &dynamodb.ListTablesInput{})
	if err != nil {
		return []models.Table{}, err
	}
	listOfTables := make([]models.Table, 0)
	for _, table := range result.TableNames {
		listOfTables = append(listOfTables, models.Table{
			Name: table,
		})
	}
	return listOfTables, nil
}
