package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (aws AWS) DescribeCloudWatchAlarms(cfg aws.Config) (map[string]int, error) {
	output := make(map[string]int, 0)
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return map[string]int{}, err
	}
	for _, region := range regions {
		alarms, err := aws.getCloudWatchAlarms(cfg, region.Name)
		if err != nil {
			return map[string]int{}, err
		}
		for _, alarm := range alarms {
			output[alarm.State]++
		}
	}
	return output, nil
}

func (aws AWS) getCloudWatchAlarms(cfg aws.Config, region string) ([]Alarm, error) {
	cfg.Region = region
	svc := cloudwatch.New(cfg)
	req := svc.DescribeAlarmsRequest(&cloudwatch.DescribeAlarmsInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return []Alarm{}, err
	}
	listOfAlarms := make([]Alarm, 0)
	for _, alarm := range result.MetricAlarms {
		alarmState, _ := alarm.StateValue.MarshalValue()
		listOfAlarms = append(listOfAlarms, Alarm{
			Name:  *alarm.AlarmName,
			State: alarmState,
		})
	}
	return listOfAlarms, nil
}
