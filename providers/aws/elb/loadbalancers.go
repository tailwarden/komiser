package elb

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	. "github.com/mlabouardy/komiser/models"
	. "github.com/mlabouardy/komiser/providers"
)

func LoadBalancers(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config elasticloadbalancing.DescribeLoadBalancersInput
	elbClient := elasticloadbalancing.NewFromConfig(*client.AWSClient)

	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return resources, err
	}

	accountId := stsOutput.Account

	for {
		output, err := elbClient.DescribeLoadBalancers(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, loadbalancer := range output.LoadBalancerDescriptions {
			resourceArn := fmt.Sprintf("arn:aws:elasticloadbalancing:%s:%s:loadbalancer/app/%s", client.AWSClient.Region, *accountId, *loadbalancer.LoadBalancerName)
			outputTags, err := elbClient.DescribeTags(ctx, &elasticloadbalancing.DescribeTagsInput{
				LoadBalancerNames: config.LoadBalancerNames,
			})

			tags := make([]Tag, 0)
			if err == nil {
				for _, tagDescription := range outputTags.TagDescriptions {
					for _, tag := range tagDescription.Tags {
						tags = append(tags, Tag{
							Key:   *tag.Key,
							Value: *tag.Value,
						})
					}
				}
			}

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "ELB",
				ResourceId: resourceArn,
				Region:     client.AWSClient.Region,
				Name:       *loadbalancer.LoadBalancerName,
				Cost:       0,
				Tags:       tags,
				CreatedAt:  *loadbalancer.CreatedTime,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/home?region=%s#/LoadBalancer:loadBalancerArn=%s", client.AWSClient.Region, client.AWSClient.Region, resourceArn),
			})
		}

		if aws.ToString(output.NextMarker) == "" {
			break
		}

		config.Marker = output.NextMarker
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "ELB",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
