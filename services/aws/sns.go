package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (aws AWS) DescribeSNSTopics(cfg awsConfig.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		topics, err := aws.getSNSTopics(cfg, region.Name)
		if err != nil {
			return 0, err
		}
		sum += int64(len(topics))
	}
	return sum, nil
}

func (aws AWS) getSNSTopics(cfg awsConfig.Config, region string) ([]Topic, error) {
	cfg.Region = region
	svc := sns.NewFromConfig(cfg)
	result, err := svc.ListTopics(context.Background(), &sns.ListTopicsInput{})
	if err != nil {
		return []Topic{}, err
	}
	listOfTopics := make([]Topic, 0, len(result.Topics))
	for _, topic := range result.Topics {
		listOfTopics = append(listOfTopics, Topic{
			ARN: *topic.TopicArn,
		})
	}
	return listOfTopics, nil
}
