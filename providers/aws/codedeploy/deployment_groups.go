package codedeploy

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
)

func DeploymentGroups(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var listApplicationParams codedeploy.ListApplicationsInput
	resources := make([]models.Resource, 0)
	codedeployClient := codedeploy.NewFromConfig(*client.AWSClient)
	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return resources, err
	}
	accountId := stsOutput.Account

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "CodeDeploy")
	if err != nil {
		log.Warnln("Couldn't fetch CodeDeploy cost and usage:", err)
	}

	for {
		output, err := codedeployClient.ListApplications(ctx, &listApplicationParams)
		if err != nil {
			return resources, err
		}
		for _, application := range output.Applications {
			var listDeploymentGroupParams codedeploy.ListDeploymentGroupsInput
			listDeploymentGroupParams.ApplicationName = &application
			for {
				listDeploymentGroupOutput, err := codedeployClient.ListDeploymentGroups(ctx, &listDeploymentGroupParams)
				if err != nil {
					return resources, nil
				}
				for _, deploymentGroup := range listDeploymentGroupOutput.DeploymentGroups {
					// logic for arn
					resourceArn := fmt.Sprintf("arn:aws:codedeploy:%s:%s:deploymentgroup:%s/%s", client.AWSClient.Region, *accountId, application, deploymentGroup)
					tags := make([]models.Tag, 0)
					resources = append(resources, models.Resource{
						Provider:   "AWS",
						Account:    client.Name,
						Service:    "CodeDeploy",
						ResourceId: resourceArn,
						Region:     client.AWSClient.Region,
						Name:       deploymentGroup,
						Metadata: map[string]string{
							"serviceCost": fmt.Sprint(serviceCost),
						},
						Tags:      tags,
						FetchedAt: time.Now(),
					})
				}
				if aws.ToString(listDeploymentGroupOutput.NextToken) == "" {
					break
				}
				listDeploymentGroupParams.NextToken = listDeploymentGroupOutput.NextToken
			}

		}

		if aws.ToString(output.NextToken) == "" {
			break
		}
		listApplicationParams.NextToken = output.NextToken
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "CodeDeploy",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
