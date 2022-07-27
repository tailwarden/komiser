package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
)

func (aws AWS) GetGlueCrawlers(cfg awsConfig.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := glue.NewFromConfig(cfg)
		res, _ := svc.GetCrawlers(context.Background(), &glue.GetCrawlersInput{})
		if res != nil {
			sum += int64(len(res.Crawlers))
		}
	}
	return sum, nil
}

func (aws AWS) GetGlueJobs(cfg awsConfig.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := glue.NewFromConfig(cfg)
		res, _ := svc.GetJobs(context.Background(), &glue.GetJobsInput{})
		if res != nil {
			sum += int64(len(res.Jobs))
		}
	}
	return sum, nil
}
