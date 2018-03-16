package services

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	. "github.com/mlabouardy/komiser/models"
)

func (aws AWS) DescribeDynamoDBTablesTotal(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		tables, err := aws.getDynamoDBTables(cfg, region.Name)
		if err != nil {
			return 0, err
		}
		sum += int64(len(tables))
	}
	return sum, nil
}

func (aws AWS) DescribeDynamoDBTablesProvisionedThroughput(cfg aws.Config) (map[string]int, error) {
	output := make(map[string]int, 0)
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return map[string]int{}, err
	}
	for _, region := range regions {
		tables, err := aws.getDynamoDBTables(cfg, region.Name)
		if err != nil {
			return map[string]int{}, err
		}
		for _, table := range tables {
			cfg.Region = region.Name
			svc := dynamodb.New(cfg)
			req := svc.DescribeTableRequest(&dynamodb.DescribeTableInput{
				TableName: &table.Name,
			})
			result, err := req.Send()
			if err != nil {
				return map[string]int{}, err
			}
			output["readCapacity"] += int(*result.Table.ProvisionedThroughput.ReadCapacityUnits)
			output["writeCapacity"] += int(*result.Table.ProvisionedThroughput.WriteCapacityUnits)
		}
	}
	return output, nil
}

func (aws AWS) getDynamoDBTables(cfg aws.Config, region string) ([]Table, error) {
	cfg.Region = region
	svc := dynamodb.New(cfg)
	req := svc.ListTablesRequest(&dynamodb.ListTablesInput{})
	result, err := req.Send()
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
