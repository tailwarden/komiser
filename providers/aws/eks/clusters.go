package eks

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	. "github.com/mlabouardy/komiser/models"
	. "github.com/mlabouardy/komiser/providers"
)

func KubernetesClusters(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config eks.ListClustersInput
	eksClient := eks.NewFromConfig(*client.AWSClient)

	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return resources, err
	}

	accountId := stsOutput.Account

	for {
		output, err := eksClient.ListClusters(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, cluster := range output.Clusters {
			resourceArn := fmt.Sprintf("arn:aws:eks:%s:%s:cluster/%s", client.AWSClient.Region, *accountId, cluster)
			outputTags, err := eksClient.ListTagsForResource(ctx, &eks.ListTagsForResourceInput{
				ResourceArn: &resourceArn,
			})

			tags := make([]Tag, 0)

			if err == nil {
				for key, value := range outputTags.Tags {
					tags = append(tags, Tag{
						Key:   key,
						Value: value,
					})
				}
			}

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "EKS",
				ResourceId: resourceArn,
				Region:     client.AWSClient.Region,
				Name:       cluster,
				Cost:       0,
				Tags:       tags,
				FetchedAt:  time.Now(),
			})
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}

		config.NextToken = output.NextToken
	}
	log.Debugf("[%s] Fetched %d AWS EKS clusters from %s\n", client.Name, len(resources), client.AWSClient.Region)
	return resources, nil
}
