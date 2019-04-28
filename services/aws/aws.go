package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models/aws"
)

type AWS struct{}

func (aws AWS) getRegions(cfg aws.Config) ([]Region, error) {
	svc := ec2.New(cfg)
	req := svc.DescribeRegionsRequest(&ec2.DescribeRegionsInput{})
	regions, err := req.Send(context.Background())
	if err != nil {
		return []Region{}, err
	}
	listOfRegions := make([]Region, 0, len(regions.Regions))
	for _, region := range regions.Regions {
		listOfRegions = append(listOfRegions, Region{
			Name: *region.RegionName,
		})
	}
	return listOfRegions, nil
}
