package codebuild

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
)

func BuildProjects(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var listProjectsParams codebuild.ListProjectsInput
	resources := make([]models.Resource, 0)
	codebuildClient := codebuild.NewFromConfig(*client.AWSClient)
	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return resources, err
	}

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "CodeBuild")
	if err != nil {
		log.Warnln("Couldn't fetch CodeBuild cost and usage:", err)
	}
	accountId := stsOutput.Account

	for {
		output, err := codebuildClient.ListProjects(ctx, &listProjectsParams)
		if err != nil {
			return resources, err
		}

		for _, project := range output.Projects {
			resourceArn := fmt.Sprintf("arn:aws:codebuild:%s:%s:project/%s", client.AWSClient.Region, *accountId, project)
			tags := make([]models.Tag, 0)

			resources = append(resources, models.Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "CodeBuild",
				ResourceId: resourceArn,
				Region:     client.AWSClient.Region,
				Name:       project,
				Metadata: map[string]string{
					"serviceCost": fmt.Sprint(serviceCost),
				},
				Tags:      tags,
				FetchedAt: time.Now(),
				Link:      fmt.Sprintf("https://%s.console.aws.amazon.com/codesuite/codebuild/%s/projects/%s/details?region=%s", client.AWSClient.Region, *accountId, project, client.AWSClient.Region),
			})
		}
		if aws.ToString(output.NextToken) == "" {
			break
		}

		listProjectsParams.NextToken = output.NextToken
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "CodeBuild",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
