package ec2

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func Snapshots(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	ec2Client := ec2.NewFromConfig(*client.AWSClient)

	input := ec2.DescribeSnapshotsInput{
		OwnerIds: []string{"self"},
	}

	for {
		output, err := ec2Client.DescribeSnapshots(ctx, &input)
		if err != nil {
			return resources, err
		}

		for _, snapshot := range output.Snapshots {
			tags := make([]Tag, 0)
			for _, tag := range snapshot.Tags {
				tags = append(tags, Tag{
					Key:   *tag.Key,
					Value: *tag.Value,
				})
			}

			resource := Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "EC2 Snapshot",
				ResourceId: aws.ToString(snapshot.SnapshotId),
				Region:     client.AWSClient.Region,
				Name:       aws.ToString(snapshot.SnapshotId),
				Tags:       tags,
				FetchedAt:  time.Now(),
				CreatedAt:  *snapshot.StartTime,
				Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/home?region=%s#napshotDetails:snapshotId=%s", client.AWSClient.Region, client.AWSClient.Region, *snapshot.SnapshotId),
				Metadata: map[string]string{
					"Description": aws.ToString(snapshot.Description),
				},
			}

			resources = append(resources, resource)
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}

	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "EC2 Snapshot",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
