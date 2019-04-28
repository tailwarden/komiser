package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
)

func (aws AWS) DescribeCacheClusters(cfg aws.Config) (map[string]int, error) {
	output := make(map[string]int, 0)
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return map[string]int{}, err
	}
	for _, region := range regions {

		cfg.Region = region.Name
		svc := elasticache.New(cfg)
		req := svc.DescribeCacheClustersRequest(&elasticache.DescribeCacheClustersInput{})
		result, err := req.Send(context.Background())
		if err != nil {
			return output, err
		}

		for _, cluster := range result.CacheClusters {
			output[*cluster.Engine]++
		}
	}
	return output, nil
}
