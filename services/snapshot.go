package services

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models"
)

func (aws AWS) DescribeSnapshotsTotal(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		snapshots, err := aws.getSnapshots(cfg, region.Name)
		if err != nil {
			return 0, err
		}
		sum += int64(len(snapshots))
	}
	return sum, nil
}

func (aws AWS) DescribeSnapshotsSize(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		snapshots, err := aws.getSnapshots(cfg, region.Name)
		if err != nil {
			return 0, err
		}
		for _, snapshot := range snapshots {
			sum += snapshot.VolumeSize
		}
	}
	return sum, nil
}

func (aws AWS) getSnapshots(cfg aws.Config, region string) ([]Snapshot, error) {
	cfg.Region = region
	svc := ec2.New(cfg)
	req := svc.DescribeSnapshotsRequest(&ec2.DescribeSnapshotsInput{})
	result, err := req.Send()
	if err != nil {
		return []Snapshot{}, err
	}
	listOfSnapshots := make([]Snapshot, 0)
	for _, snapshot := range result.Snapshots {
		snapshotState, _ := snapshot.State.MarshalValue()
		listOfSnapshots = append(listOfSnapshots, Snapshot{
			State:      snapshotState,
			VolumeSize: *snapshot.VolumeSize,
		})
	}
	return listOfSnapshots, nil
}
