package instances

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models"
)

func SecurityGroups(ctx context.Context, cfg aws.Config, account string) ([]Resource, error) {
	var config ec2.DescribeSecurityGroupsInput
	resources := make([]Resource, 0)
	ec2Client := ec2.NewFromConfig(cfg)
	for {
		output, err := ec2Client.DescribeSecurityGroups(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, o := range output.SecurityGroups {
			tags := make([]Tag, 0)
			for _, tag := range o.Tags {
				tags = append(tags, Tag{
					Key:   *tag.Key,
					Value: *tag.Value,
				})
			}

			resources = append(resources, Resource{
				Provider:  "AWS",
				Account:   account,
				Service:   "Security Group",
				Region:    cfg.Region,
				Cost:      0,
				Name:      *o.GroupName,
				FetchedAt: time.Now(),
				Tags:      tags,
			})
		}
		if aws.ToString(output.NextToken) == "" {
			break
		}

		config.NextToken = output.NextToken
	}
	log.Printf("[%s] Fetched %d AWS Security groups from %s\n", account, len(resources), cfg.Region)
	return resources, nil
}
