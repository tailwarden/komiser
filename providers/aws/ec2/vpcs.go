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
)

func Vpcs(ctx context.Context, client ProviderClient) ([]Resource, error) {
	var config ec2.DescribeVpcsInput
	resources := make([]Resource, 0)
	ec2Client := ec2.NewFromConfig(*client.AWSClient)

	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return resources, err
	}

	accountId := stsOutput.Account

	for {
		output, err := ec2Client.DescribeVpcs(ctx, &config)
		if err != nil {
			return resources, err
		}

		for _, vpc := range output.Vpcs {
			tags := make([]Tag, 0)
			for _, tag := range vpc.Tags {
				tags = append(tags, Tag{
					Key:   *tag.Key,
					Value: *tag.Value,
				})
			}

			resourceArn := fmt.Sprintf("arn:aws:ec2:%s:%s:vpc/%s", client.AWSClient.Region, *accountId, *vpc.VpcId)

			relations := getVPCRelations(ec2Client, vpc)

			resources = append(resources, Resource{
				Provider:   "AWS",
				Account:    client.Name,
				Service:    "VPC",
				Region:     client.AWSClient.Region,
				ResourceId: resourceArn,
				Cost:       0,
				Name:       *vpc.VpcId,
				Relations:  relations,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link:       fmt.Sprintf("https:/%s.console.aws.amazon.com/vpc/home?region=%s#VpcDetails:VpcId=%s", client.AWSClient.Region, client.AWSClient.Region, *vpc.VpcId),
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
		"service":   "VPC",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}

func getVPCRelations(ec2Client *ec2.Client, vpc types.Vpc) (rel []models.Link) {
	// Get associated subnets
	outputSubnets, err := ec2Client.DescribeSubnets(context.Background(), &ec2.DescribeSubnetsInput{
		Filters: []types.Filter{
			types.Filter{
				Name:   aws.String("vpc-id"),
				Values: []string{*vpc.VpcId},
			},
		},
	})
	if err != nil {
		return rel
	}

	for _, subnet := range outputSubnets.Subnets {
		rel = append(rel, models.Link{
			ResourceID: *subnet.SubnetArn,
			Type:       "Subnet",
			Name:       *subnet.SubnetArn,
			Relation:   "USES",
		})
	}

	outputIGWs, err := ec2Client.DescribeInternetGateways(context.Background(), &ec2.DescribeInternetGatewaysInput{
		Filters: []types.Filter{
			types.Filter{
				Name:   aws.String("vpc-id"),
				Values: []string{*vpc.VpcId},
			},
		},
	})
	if err != nil {
		return rel
	}

	for _, igw := range outputIGWs.InternetGateways {
		rel = append(rel, models.Link{
			ResourceID: *igw.InternetGatewayId,
			Type:       "Internet Gateway",
			Name:       *igw.InternetGatewayId,
			Relation:   "ATTACHED",
		})
	}

	return rel
}
