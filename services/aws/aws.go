package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models/aws"
)

type AWS struct {
	Regions []string
}

func (aws AWS) getRegions(cfg awsConfig.Config) ([]Region, error) {
	if len(aws.Regions) > 0 {
		listOfRegions := make([]Region, 0)
		for _, region := range aws.Regions {
			listOfRegions = append(listOfRegions, Region{
				Name: region,
			})
		}
		return listOfRegions, nil
	}

	cfg.Region = "us-east-1"
	svc := ec2.NewFromConfig(cfg)
	regions, err := svc.DescribeRegions(context.Background(), &ec2.DescribeRegionsInput{})
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
