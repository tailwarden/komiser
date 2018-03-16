package services

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models"
)

func (aws AWS) DescribeVolumesTotalSize(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		volumes, err := aws.getVolumes(cfg, region.Name)
		if err != nil {
			return 0, err
		}
		for _, volume := range volumes {
			sum += volume.Size
		}
	}
	return sum, nil
}

func (aws AWS) DescribeVolumesPerFamily(cfg aws.Config) (map[string]int, error) {
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
			output[volume.VolumeType]++
		}
	}
	return output, nil
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
	result, err := req.Send()
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
		})
	}
	return listOfVolumes, nil
}
