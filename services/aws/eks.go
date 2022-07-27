package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

func (awsClient AWS) DescribeEKSClusters(cfg awsConfig.Config) (int64, error) {
	var sum int64
	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return sum, err
	}
	for _, region := range regions {
		if region.Name != "sa-east-1" && region.Name != "ca-central-1" && region.Name != "us-west-1" {
			cfg.Region = region.Name
			svc := eks.NewFromConfig(cfg)
			res, err := svc.ListClusters(context.Background(), &eks.ListClustersInput{})
			if err != nil {
				return sum, err
			}

			sum += int64(len(res.Clusters))
		}
	}
	return sum, nil
}
