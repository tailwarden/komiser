package systemsmanager

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func getMangedEc2(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	ssmClient := ssm.NewFromConfig(*client.AWSClient)

	output, err := ssmClient.DescribeInstanceInformation(ctx, &ssm.DescribeInstanceInformationInput{
		MaxResults: aws.Int32(100),
	})

	if err != nil {
		return nil, err
	}

	for _, ec2instance := range output.InstanceInformationList {
		running, err := isRunning(ctx, *ec2instance.InstanceId, client)
		if err != nil {
			return nil, err
		}

		if running {
			instanceType, err := getInstanceType(ctx, *ec2instance.InstanceId, client)
			if err != nil {
				return nil, err
			}

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "EC2",
				Region:     client.AWSClient.Region,
				ResourceId: *ec2instance.InstanceId,
				Name:       *ec2instance.Name,
				Cost:       0.0, // No cost calculation in this version
				CreatedAt:  *ec2instance.RegistrationDate,
				Tags:       nil, // No tags in this version
				Link: fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/home?region=%s#InstanceDetails:instanceId=%s",
					client.AWSClient.Region, client.AWSClient.Region, *ec2instance.InstanceId),
			})
		}
	}

	return resources, nil
}

func getInstanceType(ctx context.Context, instanceID string, client providers.ProviderClient) (instanceType string, err error) {
	config := ec2.DescribeInstancesInput{
		InstanceIds: []string{instanceID},
		MaxResults:  aws.Int32(1),
	}
	ec2Client := ec2.NewFromConfig(*client.AWSClient)

	output, err := ec2Client.DescribeInstances(ctx, &config)
	if err != nil {
		return "", err
	}

	for _, reservations := range output.Reservations {
		for _, instance := range reservations.Instances {
			instanceType := string(instance.InstanceType)
			return instanceType, nil
		}
	}
	return "", nil
}

func isRunning(ctx context.Context, instanceID string, client providers.ProviderClient) (running bool, err error) {
	config := ec2.DescribeInstancesInput{
		InstanceIds: []string{instanceID},
		MaxResults:  aws.Int32(1),
	}
	ec2Client := ec2.NewFromConfig(*client.AWSClient)

	output, err := ec2Client.DescribeInstances(ctx, &config)
	if err != nil {
		return false, err
	}

	for _, reservations := range output.Reservations {
		for _, instance := range reservations.Instances {
			if instance.State.Name != ec2.InstanceStateNameStopped {
				return true, nil
			}
		}
	}
	return false, nil
}
