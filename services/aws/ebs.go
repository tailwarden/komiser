package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (awsClient AWS) DescribeVolumes(cfg awsConfig.Config) (map[string]interface{}, error) {
	totalVolumesSize := 0
	outputVolumesPerState := make(map[string]int, 0)
	outputVolumesPerFamily := make(map[string]int, 0)
	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return map[string]interface{}{}, err
	}
	for _, region := range regions {
		volumes, err := awsClient.getVolumes(cfg, region.Name)
		if err != nil {
			return map[string]interface{}{}, err
		}
		for _, volume := range volumes {
			outputVolumesPerFamily[volume.VolumeType]++
			outputVolumesPerState[volume.State]++
			totalVolumesSize += int(volume.Size)
		}
	}
	return map[string]interface{}{
		"total":  totalVolumesSize,
		"state":  outputVolumesPerState,
		"family": outputVolumesPerFamily,
	}, nil
}

func (awsClient AWS) DescribeVolumesPerState(cfg awsConfig.Config) (map[string]int, error) {
	output := make(map[string]int, 0)
	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return map[string]int{}, err
	}
	for _, region := range regions {
		volumes, err := awsClient.getVolumes(cfg, region.Name)
		if err != nil {
			return map[string]int{}, err
		}
		for _, volume := range volumes {
			output[volume.State]++
		}
	}
	return output, nil
}

func (awsClient AWS) getVolumes(cfg awsConfig.Config, region string) ([]Volume, error) {
	cfg.Region = region
	svc := ec2.NewFromConfig(cfg)
	result, err := svc.DescribeVolumes(context.Background(), &ec2.DescribeVolumesInput{})
	if err != nil {
		return []Volume{}, err
	}
	listOfVolumes := make([]Volume, 0)
	for _, volume := range result.Volumes {
		//	volumeType, _ := volume.VolumeType.MarshalValue()
		//	volumeState, _ := volume.State.MarshalValue()
		listOfVolumes = append(listOfVolumes, Volume{
			ID:         *volume.VolumeId,
			AZ:         *volume.AvailabilityZone,
			LaunchTime: *volume.CreateTime,
			Size:       int64(*volume.Size),
			State:      string(volume.State),
			VolumeType: string(volume.VolumeType),
		})
	}
	return listOfVolumes, nil
}
