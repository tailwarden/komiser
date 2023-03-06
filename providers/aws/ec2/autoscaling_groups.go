package ec2

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

// This is to comply with current architecture as described in CONTRIBUTING.md
func AutoScalingGroups(ctx context.Context, clientProvider ProviderClient) ([]Resource, error) {
	client := autoscaling.NewFromConfig(*clientProvider.AWSClient)

	d := ASGDiscoverer{
		client:      client,
		ctx:         ctx,
		accountName: clientProvider.Name,
		region:      clientProvider.AWSClient.Region,
	}

	return d.Discover()
}

// As I see it, this could be a struct that implements the Discoverer interface
// This would allow us to test it in isolation
type ASGDiscoverer struct {
	client      *autoscaling.Client
	ctx         context.Context
	accountName string
	region      string
}

// This could possibly be the only method the interface requires
// A good job has been done at making []Resource standard across providers and services
// Maybe we could add a factory method to the Resource struct to make it easier to create

// I am not thrilled about the way I deal with the account name and region.
// I would rather see if there's a way to fetch them from the client itself.
// Else, I would create a simple struct containing those two values and pass it around
func (d ASGDiscoverer) Discover() ([]Resource, error) {

	resources := make([]Resource, 0)
	var queryInput autoscaling.DescribeAutoScalingGroupsInput

	for {
		output, err := d.client.DescribeAutoScalingGroups(d.ctx, &queryInput)
		if err != nil {
			return resources, err
		}

		for _, asg := range output.AutoScalingGroups {
			tags := make([]Tag, 0)
			for _, tag := range asg.Tags {
				tags = append(tags, Tag{
					Key:   *tag.Key,
					Value: *tag.Value,
				})
			}

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    d.accountName,
				Service:    "AutoScalingGroup",
				Region:     d.region,
				ResourceId: *asg.AutoScalingGroupARN,
				Cost:       0,
				Name:       *asg.AutoScalingGroupName,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link: fmt.Sprintf(
					"https:/%s.console.aws.amazon.com/vpc/home?region=%s#SubnetDetails:subnetId=%s",
					d.region,
					d.region,
					*asg.AutoScalingGroupName,
				),
			})
		}
		if aws.ToString(output.NextToken) == "" {
			break
		}

		queryInput.NextToken = output.NextToken
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   d.accountName,
		"region":    d.region,
		"service":   "AutoScalingGroup",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
