package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
)

func (aws AWS) GetGlueCrawlers(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := glue.New(cfg)
		req := svc.GetCrawlersRequest(&glue.GetCrawlersInput{})
		res, _ := req.Send(context.Background())
		if res != nil {
			sum += int64(len(res.Crawlers))
		}
	}
	return sum, nil
}

func (aws AWS) GetGlueJobs(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := glue.New(cfg)
		req := svc.GetJobsRequest(&glue.GetJobsInput{})
		res, _ := req.Send(context.Background())
		if res != nil {
			sum += int64(len(res.Jobs))
		}
	}
	return sum, nil
}
