package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func main() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	for _, region := range getRegions(cfg) {
		cfg.Region = region
		ec2Svc := ec2.New(cfg)
		params := &ec2.DescribeInstancesInput{}
		req := ec2Svc.DescribeInstancesRequest(params)
		result, err := req.Send()
		if err != nil {
			log.Fatal(err)
		}

		for _, reservation := range result.Reservations {
			for _, instance := range reservation.Instances {
				fmt.Println(instance.Tags)
			}
		}
	}
}

func getRegions(cfg aws.Config) []string {
	svc := ec2.New(cfg)
	req := svc.DescribeRegionsRequest(&ec2.DescribeRegionsInput{})
	regions, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	listOfRegions := make([]string, 0, len(regions.Regions))
	for _, region := range regions.Regions {
		fmt.Println(*region.RegionName)
		listOfRegions = append(listOfRegions, *region.RegionName)
	}
	return listOfRegions
}
