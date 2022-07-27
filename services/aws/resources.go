package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi"
)

func (aws AWS) DescribeResources(cfg awsConfig.Config) ([]string, error) {
	listOfRegions := make([]string, 0)

	regions, err := aws.getRegions(cfg)
	if err != nil {
		return listOfRegions, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		client := resourcegroupstaggingapi.NewFromConfig(cfg)
		res, err := client.GetResources(context.Background(), &resourcegroupstaggingapi.GetResourcesInput{})
		if err != nil {
			return listOfRegions, err
		}

		if len(res.ResourceTagMappingList) > 0 {
			listOfRegions = append(listOfRegions, region.Name)
		}
	}
	return listOfRegions, nil
}
