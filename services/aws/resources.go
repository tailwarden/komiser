package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi"
)

func (aws AWS) DescribeResources(cfg aws.Config) ([]string, error) {
	listOfRegions := make([]string, 0)

	regions, err := aws.getRegions(cfg)
	if err != nil {
		return listOfRegions, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		client := resourcegroupstaggingapi.New(cfg)
		req := client.GetResourcesRequest(&resourcegroupstaggingapi.GetResourcesInput{})
		res, err := req.Send(context.Background())
		if err != nil {
			return listOfRegions, err
		}

		if len(res.ResourceTagMappingList) > 0 {
			listOfRegions = append(listOfRegions, region.Name)
		}
	}
	return listOfRegions, nil
}
