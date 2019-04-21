package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (aws AWS) DescribeDynamoDBTables(cfg aws.Config) (map[string]interface{}, error) {
	outputThroughput := make(map[string]int, 0)
	sumTables := 0
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return map[string]interface{}{}, err
	}
	for _, region := range regions {
		tables, err := aws.getDynamoDBTables(cfg, region.Name)
		if err != nil {
			return map[string]interface{}{}, err
		}
		sumTables += len(tables)
		for _, table := range tables {
			cfg.Region = region.Name
			svc := dynamodb.New(cfg)
			req := svc.DescribeTableRequest(&dynamodb.DescribeTableInput{
				TableName: &table.Name,
			})
			result, err := req.Send(context.Background())
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

func (aws AWS) getDynamoDBTables(cfg aws.Config, region string) ([]Table, error) {
	cfg.Region = region
	svc := dynamodb.New(cfg)
	req := svc.ListTablesRequest(&dynamodb.ListTablesInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return []Table{}, err
	}
	listOfTables := make([]Table, 0)
	for _, table := range result.TableNames {
		listOfTables = append(listOfTables, Table{
			Name: table,
		})
	}
	return listOfTables, nil
}
