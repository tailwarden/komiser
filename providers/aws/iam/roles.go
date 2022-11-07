package instances

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	. "github.com/mlabouardy/komiser/models"
	. "github.com/mlabouardy/komiser/providers"
)

func Roles(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config iam.ListRolesInput
	iamClient := iam.NewFromConfig(*client.AWSClient)
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
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "IAM Role",
			ResourceId: *o.Arn,
			Region:     client.AWSClient.Region,
			Name:       *o.RoleName,
			Cost:       0,
			CreatedAt:  *o.CreateDate,
			Tags:       tags,
			FetchedAt:  time.Now(),
		})

		if aws.ToString(output.Marker) == "" {
			break
		}

		config.Marker = output.Marker
	}
	log.Printf("[%s] Fetched %d AWS IAM roles from %s\n", client.Name, len(resources), client.AWSClient.Region)
	return resources, nil
}
