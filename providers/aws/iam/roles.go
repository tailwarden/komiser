package instances

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	. "github.com/mlabouardy/komiser/models"
)

func Roles(ctx context.Context, cfg aws.Config, account string) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config iam.ListRolesInput
	iamClient := iam.NewFromConfig(cfg)
	output, err := iamClient.ListRoles(context.Background(), &config)
	if err != nil {
		return resources, err
	}

	for _, o := range output.Roles {
		tags := make([]Tag, 0)

		for _, t := range o.Tags {
			tags = append(tags, Tag{
				Key:   *t.Key,
				Value: *t.Value,
			})
		}

		resources = append(resources, Resource{
			Provider:  "AWS",
			Account:   account,
			Service:   "IAM Role",
			Region:    cfg.Region,
			Name:      *o.RoleName,
			Cost:      0,
			CreatedAt: *o.CreateDate,
			Tags:      tags,
			FetchedAt: time.Now(),
		})

		if aws.ToString(output.Marker) == "" {
			break
		}

		config.Marker = output.Marker
	}
	log.Printf("[%s] Fetched %d AWS IAM roles from %s\n", account, len(resources), cfg.Region)
	return resources, nil
}
