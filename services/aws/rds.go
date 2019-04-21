package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (aws AWS) DescribeRDSInstances(cfg aws.Config) (map[string]int, error) {
	output := make(map[string]int, 0)
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return map[string]int{}, err
	}
	for _, region := range regions {
		instances, err := aws.getRDSInstances(cfg, region.Name)
		if err != nil {
			return map[string]int{}, err
		}
		for _, instance := range instances {
			output[instance.Engine]++
		}
	}
	return output, nil
}

func (aws AWS) getRDSInstances(cfg aws.Config, region string) ([]DBInstance, error) {
	cfg.Region = region
	svc := rds.New(cfg)
	req := svc.DescribeDBInstancesRequest(&rds.DescribeDBInstancesInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return []DBInstance{}, err
	}
	listOfInstances := make([]DBInstance, 0)
	for _, instance := range result.DBInstances {
		listOfInstances = append(listOfInstances, DBInstance{
			Status:           *instance.DBInstanceStatus,
			StorageType:      *instance.StorageType,
			AllocatedStorage: *instance.AllocatedStorage,
			InstanceClass:    *instance.DBInstanceClass,
			Engine:           *instance.Engine,
		})
	}
	return listOfInstances, nil
}
