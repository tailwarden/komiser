package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

func (aws AWS) DescribeEKSClusters(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return sum, err
	}
	for _, region := range regions {
		if region.Name != "sa-east-1" && region.Name != "ca-central-1" && region.Name != "us-west-1" {
			cfg.Region = region.Name
			svc := eks.New(cfg)
			req := svc.ListClustersRequest(&eks.ListClustersInput{})
			res, err := req.Send(context.Background())
			if err != nil {
				return sum, err
			}

			sum += int64(len(res.Clusters))
		}
	}
	return sum, nil
}
