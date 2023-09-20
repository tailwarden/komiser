package ec2

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/utils"
)

// This prices can vary strongly depending on the region
// For estimation, we just assume these prices based on us-east-1
var ebsPriceMap = map[string]float64{
	"standard": 0.05,
	"gp2":      0.1,
	"gp3":      0.08,
	"io1":      0.125,
	"io2":      0.125,
	"st1":      0.045,
	"st2":      0.015,
}

func Volumes(ctx context.Context, client ProviderClient) ([]Resource, error) {
	var config ec2.DescribeVolumesInput
	resources := make([]Resource, 0)
	ec2Client := ec2.NewFromConfig(*client.AWSClient)

	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return resources, err
	}

	accountId := stsOutput.Account
	region := client.AWSClient.Region

	for {
		output, err := ec2Client.DescribeVolumes(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, volume := range output.Volumes {
			tags := make([]Tag, 0)
			for _, tag := range volume.Tags {
				tags = append(tags, Tag{
					Key:   *tag.Key,
					Value: *tag.Value,
				})
			}
			startOfMonth := utils.BeginningOfMonth(time.Now())
			hourlyUsage := 0

			if volume.CreateTime.Before(startOfMonth) {
				hourlyUsage = int(time.Since(startOfMonth).Hours())
			} else {
				hourlyUsage = int(time.Since(*volume.CreateTime).Hours())
			}

			instanceMonths := float64(hourlyUsage) / 730.0
			instanceCost := 0.0
			hourlyCost, ok := ebsPriceMap[string(volume.VolumeType)]

			if !ok {
				log.WithFields(log.Fields{
					"service":    "EBS",
					"volumeId":   *volume.VolumeId,
					"volumeType": string(volume.VolumeType),
					"region":     region,
					"hourlyCost": hourlyCost,
				}).Warn("Volume type not supported for cost estimation, skipping.")
			} else {
				instanceCost = instanceMonths * hourlyCost
			}

			resourceArn := fmt.Sprintf("arn:aws:ec2:%s:%s:volume/%s", client.AWSClient.Region, *accountId, *volume.VolumeId)

			relations := getEBSRelations(ec2Client, volume)

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "EBS",
				Region:     client.AWSClient.Region,
				ResourceId: resourceArn,
				Relations:  relations,
				Cost:       instanceCost,
				Name:       *volume.VolumeId,
				CreatedAt:  *volume.CreateTime,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link:       fmt.Sprintf("https:/%s.console.aws.amazon.com/ec2/home?region=%s#VolumeDetails:volumeId=%s", client.AWSClient.Region, client.AWSClient.Region, *volume.VolumeId),
			})
		}
		if aws.ToString(output.NextToken) == "" {
			break
		}

		config.NextToken = output.NextToken
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "EBS",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}

func getEBSRelations(ec2Client *ec2.Client, volume types.Volume) (rel []models.Link) {
	// Get associated snapshots
	outputSnapshots, err := ec2Client.DescribeSnapshots(context.Background(), &ec2.DescribeSnapshotsInput{
		Filters: []types.Filter{
			types.Filter{
				Name:   aws.String("volume-id"),
				Values: []string{*volume.VolumeId},
			},
		},
	})
	if err != nil {
		return rel
	}

	for _, snapshot := range outputSnapshots.Snapshots {

		rel = append(rel, models.Link{
			ResourceID: *snapshot.VolumeId,
			Type:       "Elastic Block Storage",
			Name:       *snapshot.VolumeId,
			Relation:   "BACKUP",
		})
	}

	return rel
}
