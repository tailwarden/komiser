package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (aws AWS) DescribeSnapshots(cfg aws.Config) (map[string]int, error) {
	tablesTotalSize := 0
	tablesTotalNumber := 0
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return map[string]int{}, err
	}
	for _, region := range regions {
		snapshots, err := aws.getSnapshots(cfg, region.Name)
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

func (aws AWS) getSnapshots(cfg aws.Config, region string) ([]Snapshot, error) {
	cfg.Region = region
	svc := ec2.New(cfg)
	req := svc.DescribeSnapshotsRequest(&ec2.DescribeSnapshotsInput{
		OwnerIds: []string{"self"},
	})
	result, err := req.Send(context.Background())
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
