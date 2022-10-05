package aws

import (
	"context"
	"fmt"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type AWSDynamoDBTable struct {
	Name   string   `json:"name"`
	Region string   `json:"region"`
	ID     string   `json:"id"`
	Tags   []string `json:"tags"`
}

func (awsClient AWS) DescribeDynamoDBTables(cfg awsConfig.Config) ([]AWSDynamoDBTable, error) {
	listOfTables := make([]AWSDynamoDBTable, 0)
	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return listOfTables, err
	}
	for _, region := range regions {
		tables, err := awsClient.getDynamoDBTables(cfg, region.Name)
		if err != nil {
			return listOfTables, err
		}

		for _, table := range tables {
			listOfTables = append(listOfTables, table)
		}
	}
	return listOfTables, nil
}

func (awsClient AWS) getDynamoDBTables(cfg awsConfig.Config, region string) ([]AWSDynamoDBTable, error) {
	cfg.Region = region
	svc := dynamodb.NewFromConfig(cfg)
	result, err := svc.ListTables(context.Background(), &dynamodb.ListTablesInput{})
	if err != nil {
		return []AWSDynamoDBTable{}, err
	}
	listOfTables := make([]AWSDynamoDBTable, 0)
	for _, table := range result.TableNames {

		tagsResp, err := svc.ListTagsOfResource(context.Background(), &dynamodb.ListTagsOfResourceInput{
			ResourceArn: &table,
		})
		if err != nil {
			return []AWSDynamoDBTable{}, err
		}

		tags := make([]string, 0)
		for _, tag := range tagsResp.Tags {
			tags = append(tags, fmt.Sprintf("%s:%s", *tag.Key, *tag.Value))
		}

		listOfTables = append(listOfTables, AWSDynamoDBTable{
			Name:   table,
			ID:     table,
			Region: region,
			Tags:   tags,
		})
	}
	return listOfTables, nil
}
