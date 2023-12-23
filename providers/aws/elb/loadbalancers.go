package elb

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
	"github.com/tailwarden/komiser/utils"
)

type ELBOptions struct {
	elbType  string
	elbPrice float64
}

func LoadBalancers(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)

	elbValues := map[string]ELBOptions{
		"application": {"Application", 0.0225},
		"network":     {"Network", 0.0225},
		"gateway":     {"Gateway", 0.0125},
	}

	var config elasticloadbalancingv2.DescribeLoadBalancersInput
	elbClient := elasticloadbalancingv2.NewFromConfig(*client.AWSClient)

	output, err := elbClient.DescribeLoadBalancers(ctx, &config)
	if err != nil {
		return resources, err
	}

	var configListeners elasticloadbalancingv2.DescribeListenersInput
	var configRules elasticloadbalancingv2.DescribeRulesInput

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "ELB")
	if err != nil {
		log.Warnln("Couldn't fetch ELB cost and usage:", err)
	}

	for _, loadbalancer := range output.LoadBalancers {
		resourceArn := *loadbalancer.LoadBalancerArn
		resourceType := string(loadbalancer.Type)

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

		startOfMonth := utils.BeginningOfMonth(time.Now())
		hourlyUsage := 0
		if (*loadbalancer.CreatedTime).Before(startOfMonth) {
			hourlyUsage = int(time.Since(startOfMonth).Hours())
		} else {
			hourlyUsage = int(time.Since(*loadbalancer.CreatedTime).Hours())
		}
		monthlyCost := float64(hourlyUsage) * elbValues[resourceType].elbPrice

		resources = append(resources, Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    fmt.Sprintf("ELB %s", elbValues[resourceType].elbType),
			ResourceId: resourceArn,
			Region:     client.AWSClient.Region,
			Name:       *loadbalancer.LoadBalancerName,
			Cost:       monthlyCost,
			Metadata: map[string]string{
				"serviceCost": fmt.Sprint(serviceCost),
			},
			Tags:       tags,
			CreatedAt:  *loadbalancer.CreatedTime,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/home?region=%s#LoadBalancer:loadBalancerArn=%s", client.AWSClient.Region, client.AWSClient.Region, resourceArn),
		})

		configListeners.LoadBalancerArn = &resourceArn
		output, err := elbClient.DescribeListeners(ctx, &configListeners)
		if err != nil {
			return resources, err
		}

		for _, listener := range output.Listeners {
			listenerArn := *listener.ListenerArn
			outputTags, err := elbClient.DescribeTags(ctx, &elasticloadbalancingv2.DescribeTagsInput{
				ResourceArns: []string{listenerArn},
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

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    fmt.Sprintf("ELB %s Listener", elbValues[resourceType].elbType),
				ResourceId: listenerArn,
				Region:     client.AWSClient.Region,
				Name:       *loadbalancer.LoadBalancerName,
				Tags:       tags,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/home?region=%s#ELBListenerV2:listenerArn=%s", client.AWSClient.Region, client.AWSClient.Region, listenerArn),
			})

			configRules.ListenerArn = &listenerArn
			output, err := elbClient.DescribeRules(ctx, &configRules)
			if err != nil {
				return resources, err
			}

			for rule_number, rule := range output.Rules {
				ruleArn := *rule.RuleArn
				outputTags, err := elbClient.DescribeTags(ctx, &elasticloadbalancingv2.DescribeTagsInput{
					ResourceArns: []string{ruleArn},
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
				resources = append(resources, Resource{
					Provider:   "AWS",
					Account:    client.Name,
					Service:    fmt.Sprintf("ELB %s Listener Rule", elbValues[resourceType].elbType),
					ResourceId: ruleArn,
					Region:     client.AWSClient.Region,
					Name:       *loadbalancer.LoadBalancerName + fmt.Sprintf(" Rule %d", rule_number+1),
					Tags:       tags,
					FetchedAt:  time.Now(),
					Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/home?region=%s#ListenerRuleDetails:ruleArn=%s", client.AWSClient.Region, client.AWSClient.Region, ruleArn),
				})
			}

		}

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
