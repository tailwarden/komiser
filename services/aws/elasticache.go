package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
)

func (awsClient AWS) DescribeCacheClusters(cfg awsConfig.Config) (map[string]int, error) {
	output := make(map[string]int, 0)
	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return map[string]int{}, err
	}
	for _, region := range regions {

		cfg.Region = region.Name
		svc := elasticache.NewFromConfig(cfg)
		result, err := svc.DescribeCacheClusters(context.Background(), &elasticache.DescribeCacheClustersInput{})
		if err != nil {
			return output, err
		}

		for _, cluster := range result.CacheClusters {
			if output[*cluster.ReplicationGroupId] == 0 {
				output[*cluster.Engine]++
				output[*cluster.ReplicationGroupId]++
			}
		}
	}
	return output, nil
}
