package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
)

func (aws AWS) DescribeRedshiftClusters(cfg aws.Config) (int64, error) {
	var total int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return total, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := redshift.New(cfg)
		req := svc.DescribeClustersRequest(&redshift.DescribeClustersInput{})
		res, err := req.Send(context.Background())
		if err != nil {
			return total, err
		}
		total += int64(len(res.Clusters))
	}
	return total, nil
}
