package services

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	. "github.com/mlabouardy/komiser/models"
)

func (aws AWS) DescribeQueuesTotal(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		queues, err := aws.getSQS(cfg, region.Name)
		if err != nil {
			return 0, err
		}
		sum += int64(len(queues))
	}
	return sum, nil
}

func (aws AWS) getSQS(cfg aws.Config, region string) ([]Queue, error) {
	cfg.Region = region
	svc := sqs.New(cfg)
	req := svc.ListQueuesRequest(&sqs.ListQueuesInput{})
	result, err := req.Send()
	if err != nil {
		return []Queue{}, err
	}
	listOfQueues := make([]Queue, 0, len(result.QueueUrls))
	for _, queue := range result.QueueUrls {
		listOfQueues = append(listOfQueues, Queue{
			Name: queue,
		})
	}
	return listOfQueues, nil
}
