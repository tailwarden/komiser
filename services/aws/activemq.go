package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mq"
)

func (aws AWS) ListBrokers(cfg awsConfig.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := mq.NewFromConfig(cfg)
		res, _ := svc.ListBrokers(context.Background(), &mq.ListBrokersInput{})

		if res != nil {
			sum += int64(len(res.BrokerSummaries))
		}
	}
	return sum, nil
}
