package systemsmanager

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
)

func GetManagedEc2(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	ssmClient := ssm.NewFromConfig(*client.AWSClient)
	ec2Client := ec2.NewFromConfig(*client.AWSClient)

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "SystemsManager")
	if err != nil {
		log.Warnln("Couldn't fetch SystemsManager cost and usage:", err)
	}

	var nexttoken *string

	for {
		ssmOutput, err := ssmClient.DescribeInstanceInformation(ctx, &ssm.DescribeInstanceInformationInput{
			NextToken: nexttoken,
		})
		if err != nil {
			return resources, err
		}

		instanceIds := make([]string, 0, len(ssmOutput.InstanceInformationList))
		for _, ec2instance := range ssmOutput.InstanceInformationList {
			instanceIds = append(instanceIds, *ec2instance.InstanceId)
		}
		ec2Output, err := ec2Client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
			InstanceIds: instanceIds,
		})
		if err != nil {
			return resources, err
		}

		account, accountID, err := fetchID(ctx, client)
		if err != nil {
			return resources, err
		}

		for _, ec2instance := range ec2Output.Reservations {
			for _, instance := range ec2instance.Instances {

				if instance.State.Name == types.InstanceStateNameRunning {
					tags := make([]models.Tag, 0)
					for _, tag := range instance.Tags {
						tags = append(tags, models.Tag{
							Key:   *tag.Key,
							Value: *tag.Value,
						})
					}

					resources = append(resources, models.Resource{
						Provider:   "AWS",
						Account:    account,
						AccountId:  accountID,
						Service:    "SSM Instance",
						Region:     client.AWSClient.Region,
						ResourceId: *instance.InstanceId,
						Name:       string(instance.InstanceType),
						CreatedAt:  *instance.LaunchTime,
						FetchedAt:  time.Now(),
						Tags:       tags,
						Metadata: map[string]string{
							"serviceCost": fmt.Sprint(serviceCost),
						},
						Link: fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/home?region=%s#InstanceDetails:instanceId=%s",
							client.AWSClient.Region, client.AWSClient.Region, *instance.InstanceId),
					})
				}

			}
		}
		if ssmOutput.NextToken == nil {
			break
		}
		nexttoken = ssmOutput.NextToken
	}

	return resources, nil
}

func fetchID(ctx context.Context, client providers.ProviderClient) (accountID string, userID string, err error) {
	svc := sts.NewFromConfig(*client.AWSClient)
	result, err := svc.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		fmt.Println("Got error retrieving account ID:")
		return "", "", err
	}

	// if the accountID and userID are the same, then the account is a root account
	return *result.Account, *result.UserId, nil
}
