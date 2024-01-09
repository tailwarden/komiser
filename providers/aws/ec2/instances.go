package ec2

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	etype "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/aws/aws-sdk-go-v2/service/pricing/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
	"github.com/tailwarden/komiser/utils"
)

func Instances(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	var nextToken string
	resources := make([]models.Resource, 0)
	ec2Client := ec2.NewFromConfig(*client.AWSClient)

	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return resources, err
	}

	accountId := stsOutput.Account

	serviceCost, err := awsUtils.GetCostAndUsage(ctx, client.AWSClient.Region, "EC2")
	if err != nil {
		log.Warnln("Couldn't fetch EC2 cost and usage:", err)
	}

	oldRegion := client.AWSClient.Region
	client.AWSClient.Region = "us-east-1"
	pricingClient := pricing.NewFromConfig(*client.AWSClient)
	client.AWSClient.Region = oldRegion

	for {
		output, err := ec2Client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
			NextToken: &nextToken,
		})
		if err != nil {
			return resources, err
		}

		for _, reservations := range output.Reservations {
			for _, instance := range reservations.Instances {
				tags := make([]models.Tag, 0)

				name := ""
				for _, tag := range instance.Tags {
					if *tag.Key == "Name" {
						name = *tag.Value
					}
					tags = append(tags, models.Tag{
						Key:   *tag.Key,
						Value: *tag.Value,
					})
				}

				monthlyCost := 0.0

				if instance.State.Name != "stopped" {
					// no need to calc usage and fetch price if stopped

					startOfMonth := utils.BeginningOfMonth(time.Now())
					hourlyUsage := 0
					if instance.LaunchTime.Before(startOfMonth) {
						hourlyUsage = int(time.Since(startOfMonth).Hours())
					} else {
						hourlyUsage = int(time.Since(*instance.LaunchTime).Hours())
					}

					pricingOutput, err := pricingClient.GetProducts(ctx, &pricing.GetProductsInput{
						ServiceCode: aws.String("AmazonEC2"),
						Filters: []types.Filter{
							{
								Field: aws.String("operatingSystem"),
								Value: aws.String("linux"),
								Type:  types.FilterTypeTermMatch,
							},
							{
								Field: aws.String("instanceType"),
								Value: aws.String(string(instance.InstanceType)),
								Type:  types.FilterTypeTermMatch,
							},
							{
								Field: aws.String("regionCode"),
								Value: aws.String(client.AWSClient.Region),
								Type:  types.FilterTypeTermMatch,
							},
							{
								Field: aws.String("capacitystatus"),
								Value: aws.String("Used"),
								Type:  types.FilterTypeTermMatch,
							},
						},
						MaxResults: aws.Int32(1),
					})
					if err != nil {
						log.Warnf("Couldn't fetch invocations metric for %s", name)
					}

					hourlyCost := 0.0

					if pricingOutput != nil && len(pricingOutput.PriceList) > 0 {

						pricingResult := models.PricingResult{}
						err := json.Unmarshal([]byte(pricingOutput.PriceList[0]), &pricingResult)
						if err != nil {
							log.Fatalf("Failed to unmarshal JSON: %v", err)
						}

						for _, onDemand := range pricingResult.Terms.OnDemand {
							for _, priceDimension := range onDemand.PriceDimensions {
								hourlyCost, err = strconv.ParseFloat(priceDimension.PricePerUnit.USD, 64)
								if err != nil {
									log.Fatalf("Failed to parse hourly cost: %v", err)
								}
								break
							}
							break
						}

						//log.Printf("Hourly cost EC2: %f", hourlyCost)

					}

					monthlyCost = float64(hourlyUsage) * hourlyCost

				}

				relations := getEC2Relations(&instance, fmt.Sprintf("arn:aws:ec2:%s:%s", client.AWSClient.Region, *accountId))
				resourceArn := fmt.Sprintf("arn:aws:ec2:%s:%s:instance/%s", client.AWSClient.Region, *accountId, *instance.InstanceId)

				resources = append(resources, models.Resource{
					Provider:   "AWS",
					Account:    client.Name,
					Service:    "EC2",
					Region:     client.AWSClient.Region,
					Name:       name,
					ResourceId: resourceArn,
					CreatedAt:  *instance.LaunchTime,
					FetchedAt:  time.Now(),
					Cost:       monthlyCost,
					Tags:       tags,
					Relations:  relations,
					Metadata: map[string]string{
						"instanceType": string(instance.InstanceType),
						"state":        string(instance.State.Name),
						"serviceCost": fmt.Sprint(serviceCost),
					},
					Link: fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/home?region=%s#InstanceDetails:instanceId=%s", client.AWSClient.Region, client.AWSClient.Region, *instance.InstanceId),
				})
			}
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}

		nextToken = *output.NextToken
	}
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "EC2",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}

func getEC2Relations(inst *etype.Instance, resourceArn string) []models.Link {

	var rel []models.Link
	// Get associated security groups
	for _, sgrp := range inst.SecurityGroups {
		rel = append(rel, models.Link{
			ResourceID: *sgrp.GroupId,
			Type:       "Security Group",
			Name:       *sgrp.GroupName,
			Relation:   "USES",
		})
	}

	// Get associated volumes
	for _, blk := range inst.BlockDeviceMappings {
		id := fmt.Sprintf("%s:volume/%s", resourceArn, *blk.Ebs.VolumeId)
		rel = append(rel, models.Link{
			ResourceID: id,
			Type:       "Elastic Block Storage",
			Name:       *blk.DeviceName,
			Relation:   "USES",
		})
	}

	if inst.VpcId != nil {
		// Get associated VPC
		rel = append(rel, models.Link{
			ResourceID: fmt.Sprintf("%s:vpc/%s", resourceArn, *inst.VpcId),
			Type:       "VPC",
			Name:       *inst.VpcId,
			Relation:   "USES",
		})
	}

	if inst.SubnetId != nil {
		// Get associated Subnet
		rel = append(rel, models.Link{
			ResourceID: fmt.Sprintf("%s:subnet/%s", resourceArn, *inst.SubnetId),
			Name:       *inst.SubnetId,
			Type:       "Subnet",
			Relation:   "USES",
		})
	}

	// Get associated Keypair
	if inst.KeyName != nil {
		rel = append(rel, models.Link{
			ResourceID: *inst.KeyName,
			Name:       *inst.KeyName,
			Type:       "Key Pair",
			Relation:   "USES",
		})
	}

	// Get associated IAM roles
	if inst.IamInstanceProfile != nil {
		rel = append(rel, models.Link{
			ResourceID: *inst.IamInstanceProfile.Id,
			Name:       *inst.IamInstanceProfile.Arn,
			Type:       "IAM Role",
			Relation:   "USES",
		})
	}

	// Get associated network interfaces
	for _, ei := range inst.NetworkInterfaces {
		rel = append(rel, models.Link{
			ResourceID: *ei.NetworkInterfaceId,
			Name:       *ei.NetworkInterfaceId,
			Type:       "Network Interface",
			Relation:   "USES",
		})
	}
	return rel
}
