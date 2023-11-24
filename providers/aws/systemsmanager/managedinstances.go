package systemsmanager

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func getMangedEc2(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	ssmClient := ssm.NewFromConfig(*client.AWSClient)

	output, err := ssmClient.DescribeInstanceInformation(ctx, &ssm.DescribeInstanceInformationInput{
		MaxResults: aws.Int32(50),
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, ec2instance := range output.InstanceInformationList {
		running, err := isRunning(ctx, *ec2instance.InstanceId, client)
		if err != nil {
			log.Fatal(err)
		}

		if true == running {
			fmt.Println("running")
			instance, err := getInstanceType(ctx, *ec2instance.InstanceId, client)
			if err != nil {
				return nil, err
			}

			if len(instance.Reservations) > 0 && len(instance.Reservations[0].Instances) > 0 {
				currentInstance := instance.Reservations[0].Instances[0]

				tags := make([]models.Tag, 0)
				for _, tag := range currentInstance.Tags {
					tags = append(tags, models.Tag{
						Key:   *tag.Key,
						Value: *tag.Value,
					})
				}
				account, account_id := fetchid()

				resources = append(resources, models.Resource{
					Provider:   "AWS",
					Account:    account,
					AccountId:  account_id,
					Service:    "EC2",
					Region:     client.AWSClient.Region,
					ResourceId: *currentInstance.InstanceId,
					Name:       string(currentInstance.InstanceType),
					CreatedAt:  *currentInstance.LaunchTime,
					FetchedAt:  time.Now(),
					Tags:       tags,
					Link: fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/home?region=%s#InstanceDetails:instanceId=%s",
						client.AWSClient.Region, client.AWSClient.Region, *ec2instance.InstanceId),
				})
			}
		}
	}

	return resources, nil
}

func getInstanceType(ctx context.Context, instanceID string, client providers.ProviderClient) (*ec2.DescribeInstancesOutput, error) {
	config := ec2.DescribeInstancesInput{
		InstanceIds: []string{instanceID},
	}
	ec2Client := ec2.NewFromConfig(*client.AWSClient)

	output, err := ec2Client.DescribeInstances(ctx, &config)
	if err != nil {
		return output, err
	}

	return output, nil
}

func isRunning(ctx context.Context, instanceID string, client providers.ProviderClient) (running bool, err error) {
	fmt.Println("status check for instance id: ", instanceID)
	config := ec2.DescribeInstancesInput{
		InstanceIds: []string{instanceID},
	}
	ec2Client := ec2.NewFromConfig(*client.AWSClient)

	output, err := ec2Client.DescribeInstances(ctx, &config)
	if err != nil {
		return false, err
	}

	for _, reservations := range output.Reservations {
		for _, instance := range reservations.Instances {
			if instance.State.Name == types.InstanceStateNameRunning {
				return true, nil
			}
		}
	}
	return false, nil
}

func fetchid() (accountID string, userid string) {
	svc := sts.NewFromConfig(*newclient())
	result, err := svc.GetCallerIdentity(context.Background(), &sts.GetCallerIdentityInput{})
	if err != nil {
		fmt.Println("Got error retrieving account ID:")
		log.Println(err.Error())
	}

	// if the accountID and userid same then the account is root account
	accountID = *result.Account
	userid = *result.UserId

	return accountID, userid
}
