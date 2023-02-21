package elb

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func BeginningOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 0, -date.Day()+1)
}

func LoadBalancers(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config elasticloadbalancingv2.DescribeLoadBalancersInput
	elbClient := elasticloadbalancingv2.NewFromConfig(*client.AWSClient)

	output, err := elbClient.DescribeLoadBalancers(ctx, &config)
	if err != nil {
		return resources, err
	}

	for _, loadbalancer := range output.LoadBalancers {
		resourceArn := *loadbalancer.LoadBalancerArn
		outputTags, err := elbClient.DescribeTags(ctx, &elasticloadbalancingv2.DescribeTagsInput{
			ResourceArns: []string{resourceArn},
		})
		if err != nil {
			return resources, err
		}

		tags := make([]Tag, 0)
		for _, tagDescription := range outputTags.TagDescriptions {
			for _, tag := range tagDescription.Tags {
				tags = append(tags, Tag{
					Key:   *tag.Key,
					Value: *tag.Value,
				})
			}
		}

		startOfMonth := BeginningOfMonth(time.Now())
		hourlyUsage := 0
		if (*loadbalancer.CreatedTime).Before(startOfMonth) {
			hourlyUsage = int(time.Since(startOfMonth).Hours())
		} else {
			hourlyUsage = int(time.Since(*loadbalancer.CreatedTime).Hours())
		}
		monthlyCost := float64(hourlyUsage) * 0.0225

		resources = append(resources, Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "ELB",
			ResourceId: resourceArn,
			Region:     client.AWSClient.Region,
			Name:       *loadbalancer.LoadBalancerName,
			Cost:       monthlyCost,
			Tags:       tags,
			CreatedAt:  *loadbalancer.CreatedTime,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/home?region=%s#/LoadBalancer:loadBalancerArn=%s", client.AWSClient.Region, client.AWSClient.Region, resourceArn),
		})
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
