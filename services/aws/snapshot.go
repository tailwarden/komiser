package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (awsClient AWS) DescribeSnapshots(cfg awsConfig.Config) (map[string]int, error) {
	tablesTotalSize := 0
	tablesTotalNumber := 0
	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return map[string]int{}, err
	}
	for _, region := range regions {
		snapshots, err := awsClient.getSnapshots(cfg, region.Name)
		if err != nil {
			return map[string]int{}, err
		}
		tablesTotalNumber += len(snapshots)
		for _, snapshot := range snapshots {
			tablesTotalSize += int(snapshot.VolumeSize)
		}
	}
	return map[string]int{
		"total": tablesTotalNumber,
		"size":  tablesTotalSize,
	}, nil
}

func (awsClient AWS) getSnapshots(cfg awsConfig.Config, region string) ([]Snapshot, error) {
	cfg.Region = region
	svc := ec2.NewFromConfig(cfg)
	result, err := svc.DescribeSnapshots(context.Background(), &ec2.DescribeSnapshotsInput{
		OwnerIds: []string{"self"},
	})
	if err != nil {
		return []Snapshot{}, err
	}
	listOfSnapshots := make([]Snapshot, 0)
	for _, snapshot := range result.Snapshots {
		listOfSnapshots = append(listOfSnapshots, Snapshot{
			State:      string(snapshot.State),
			VolumeSize: int64(*snapshot.VolumeSize),
		})
	}
	return listOfSnapshots, nil
}
