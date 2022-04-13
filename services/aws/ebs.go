package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (aws AWS) DescribeVolumes(cfg aws.Config) (map[string]interface{}, error) {
	totalVolumesSize := 0
	outputVolumesPerState := make(map[string]int, 0)
	outputVolumesPerFamily := make(map[string]int, 0)
	outputVolumesPerEncryption := make(map[bool]int, 0)
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return map[string]interface{}{}, err
	}
	for _, region := range regions {
		volumes, err := aws.getVolumes(cfg, region.Name)
		if err != nil {
			return map[string]interface{}{}, err
		}
		for _, volume := range volumes {
			outputVolumesPerFamily[volume.VolumeType]++
			outputVolumesPerState[volume.State]++
			outputVolumesPerEncryption[volume.Encrypted]++
			totalVolumesSize += int(volume.Size)
		}
	}
	return map[string]interface{}{
		"total":     totalVolumesSize,
		"state":     outputVolumesPerState,
		"family":    outputVolumesPerFamily,
		"encrypted": outputVolumesPerEncryption,
	}, nil
}

func (aws AWS) DescribeVolumesPerState(cfg aws.Config) (map[string]int, error) {
	output := make(map[string]int, 0)
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return map[string]int{}, err
	}
	for _, region := range regions {
		volumes, err := aws.getVolumes(cfg, region.Name)
		if err != nil {
			return map[string]int{}, err
		}
		for _, volume := range volumes {
			output[volume.State]++
		}
	}
	return output, nil
}

func (aws AWS) getVolumes(cfg aws.Config, region string) ([]Volume, error) {
	cfg.Region = region
	svc := ec2.New(cfg)
	req := svc.DescribeVolumesRequest(&ec2.DescribeVolumesInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return []Volume{}, err
	}
	listOfVolumes := make([]Volume, 0)
	for _, volume := range result.Volumes {
		volumeType, _ := volume.VolumeType.MarshalValue()
		volumeState, _ := volume.State.MarshalValue()
		listOfVolumes = append(listOfVolumes, Volume{
			ID:         *volume.VolumeId,
			AZ:         *volume.AvailabilityZone,
			LaunchTime: *volume.CreateTime,
			Size:       *volume.Size,
			State:      volumeState,
			VolumeType: volumeType,
			Encrypted:  *volume.Encrypted,
		})
	}
	return listOfVolumes, nil
}
