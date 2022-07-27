package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
)

func (aws AWS) DescribeRedshiftClusters(cfg awsConfig.Config) (int64, error) {
	var total int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return total, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := redshift.NewFromConfig(cfg)
		res, err := svc.DescribeClusters(context.Background(), &redshift.DescribeClustersInput{})
		if err != nil {
			return total, err
		}
		total += int64(len(res.Clusters))
	}
	return total, nil
}
