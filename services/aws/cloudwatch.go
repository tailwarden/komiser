package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	models "github.com/mlabouardy/komiser/models/aws"
)

func (awsClient AWS) DescribeCloudWatchAlarms(cfg awsConfig.Config) (map[string]int, error) {
	output := make(map[string]int, 0)
	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return map[string]int{}, err
	}
	for _, region := range regions {
		alarms, err := awsClient.getCloudWatchAlarms(cfg, region.Name)
		if err != nil {
			return map[string]int{}, err
		}
		for _, alarm := range alarms {
			output[alarm.State]++
		}
	}
	return output, nil
}

func (awsClient AWS) getCloudWatchAlarms(cfg awsConfig.Config, region string) ([]models.Alarm, error) {
	cfg.Region = region
	svc := cloudwatch.NewFromConfig(cfg)
	result, err := svc.DescribeAlarms(context.Background(), &cloudwatch.DescribeAlarmsInput{})
	if err != nil {
		return []models.Alarm{}, err
	}
	listOfAlarms := make([]models.Alarm, 0)
	for _, alarm := range result.MetricAlarms {
		//alarmState, _ := alarm.StateValue
		listOfAlarms = append(listOfAlarms, models.Alarm{
			Name:  *alarm.AlarmName,
			State: "OK",
		})
	}
	return listOfAlarms, nil
}
