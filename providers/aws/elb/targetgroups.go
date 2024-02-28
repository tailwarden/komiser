package elb

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
)

func TargetGroups(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	var config elasticloadbalancingv2.DescribeTargetGroupsInput
	elbtgClient := elasticloadbalancingv2.NewFromConfig(*client.AWSClient)

	output, err := elbtgClient.DescribeTargetGroups(ctx, &config)

	if err != nil {
		return resources, err
	}

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "ELB")
	if err != nil {
		log.Warnln("Couldn't fetch ELB cost and usage:", err)
	}

	for _, targetgroup := range output.TargetGroups {
		resourceArn := *targetgroup.TargetGroupArn
		outputTags, err := elbtgClient.DescribeTags(ctx, &elasticloadbalancingv2.DescribeTagsInput{
			ResourceArns: []string{resourceArn},
		})
		if err != nil {
			return resources, err
		}

		tags := make([]models.Tag, 0)
		for _, tagDescription := range outputTags.TagDescriptions {
			for _, tag := range tagDescription.Tags {
				tags = append(tags, models.Tag{
					Key:   *tag.Key,
					Value: *tag.Value,
				})
			}
		}

		resources = append(resources, models.Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "Target Group",
			ResourceId: resourceArn,
			Region:     client.AWSClient.Region,
			Name:       *targetgroup.TargetGroupName,
			Tags:       tags,
			FetchedAt:  time.Now(),
			Metadata: map[string]string{
				"serviceCost": fmt.Sprint(serviceCost),
			},
			Link: fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/home?region=%s#TargetGroup:targetGroupArn=%s", client.AWSClient.Region, client.AWSClient.Region, resourceArn),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "Target Group",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
