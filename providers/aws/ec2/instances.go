package instances

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models"
)

func Instances(ctx context.Context, cfg aws.Config, account string) ([]Resource, error) {
	var nextToken string
	resources := make([]Resource, 0)
	ec2Client := ec2.NewFromConfig(cfg)
	for {
		output, err := ec2Client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
			NextToken: &nextToken,
		})
		if err != nil {
			return resources, err
		}

		for _, reservations := range output.Reservations {
			for _, instance := range reservations.Instances {
				tags := make([]Tag, 0)

				name := ""
				for _, tag := range instance.Tags {
					if *tag.Key == "Name" {
						name = *tag.Value
					}
					tags = append(tags, Tag{
						Key:   *tag.Key,
						Value: *tag.Value,
					})
				}

				cost := 0.0

				resources = append(resources, Resource{
					Provider:  "AWS",
					Account:   account,
					Service:   "EC2",
					Region:    cfg.Region,
					Name:      name,
					CreatedAt: *instance.LaunchTime,
					FetchedAt: time.Now(),
					Cost:      cost,
					Tags:      tags,
					Metadata: map[string]string{
						"instanceType": fmt.Sprintf("%s", instance.InstanceType),
						"state":        fmt.Sprintf("%s", instance.State.Name),
					},
				})
			}
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}

		nextToken = *output.NextToken
	}
	log.Printf("[%s] Fetched %d AWS EC2 instances from %s\n", account, len(resources), cfg.Region)
	return resources, nil
}
