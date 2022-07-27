package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline"
)

func (awsClient AWS) ListDataPipelines(cfg awsConfig.Config) (int64, error) {
	var sum int64
	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := datapipeline.NewFromConfig(cfg)
		res, _ := svc.ListPipelines(context.Background(), &datapipeline.ListPipelinesInput{})
		if res != nil {
			sum += int64(len(res.PipelineIdList))
		}
	}
	return sum, nil
}
