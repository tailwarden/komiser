package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline"
)

func (aws AWS) ListDataPipelines(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := datapipeline.New(cfg)
		req := svc.ListPipelinesRequest(&datapipeline.ListPipelinesInput{})
		res, _ := req.Send(context.Background())
		if res != nil {
			sum += int64(len(res.PipelineIdList))
		}
	}
	return sum, nil
}
