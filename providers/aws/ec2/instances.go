package ec2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	cloudwatchTypes "github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	awsUtils "github.com/tailwarden/komiser/providers/aws/utils"
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
	"github.com/tailwarden/komiser/utils"

	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
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

// AIM : simple use the list of SSM managed instances to generate respective resources. You may fetch metric and calculate the cost
// in aws s3 bucket price is calculated per request so we need to calculate the cost per month
//  but for ec2 it is calculated per hour so we need to calculate the cost per hour

func ConvertBytesToTerabytes(bytes int64) float64 {
	return float64(bytes) / 1099511627776
}

func getMangedEc2(ctx context.Context, client ProviderClient) ([]Resource, error) {

	resources := make([]Resource, 0)
	var config = ssm.DescribeInstanceInformationInput{
		MaxResults: aws.Int32(100),
	}
	ssmClient := ssm.NewFromConfig(*client.AWSClient)
	cloudwatchClient := cloudwatch.NewFromConfig(*client.AWSClient)
	pricingClient := pricing.NewFromConfig(*client.AWSClient)

	pricingOutput, err := pricingClient.GetProducts(ctx, &pricing.GetProductsInput{
		ServiceCode: aws.String("AmazonEC2"),
		Filters: []types.Filter{
			{
				Field: aws.String("regionCode"),
				Value: aws.String(client.AWSClient.Region),
				Type:  types.FilterTypeTermMatch,
			},
		},
	})

	if err != nil {
		log.Errorf("ERROR: Couldn't fetch pricing info for AWS EC2: %v", err)
		return resources, err
	}

	priceMap, err := awsUtils.GetPriceMap(pricingOutput, "group")
	if err != nil {
		log.Errorf("ERROR: Failed to calculate cost per month: %v", err)
		return resources, err
	}

	output, err := ssmClient.DescribeInstanceInformation(ctx, &config)
	if err != nil {
		return nil, err
	}

	for _, ec2 := range output.InstanceInformationList {
		metricesEc2sizebyOutput, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
			StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
			EndTime:    aws.Time(time.Now()),
			MetricName: aws.String("EC2SizeBytes"),
			Namespace:  aws.String("AWS/EC2"),
			Dimensions: []cloudwatchTypes.Dimension{
				{
					Name:  aws.String("InstanceName"),
					Value: ec2.Name,
				},
			},
			Unit:   cloudwatchTypes.StandardUnitBytes,
			Period: aws.Int32(3600),
			Statistics: []cloudwatchTypes.Statistic{
				cloudwatchTypes.StatisticAverage,
			},
		})

		if err != nil {
			log.Warnf("Couldn't fetch invocations metric for %s", *bucket.Name)
		}

		instanceType := ""
		if ec2.InstanceType != nil {
			instanceType = *ec2.InstanceType
		}



		sizeInTB := 0.0
		if ec2.Ins

		metricesUsesOutput, err := cloudwatchClient.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
			StartTime:  aws.Time(utils.BeginningOfMonth(time.Now())),
			EndTime:    aws.Time(time.Now()),
			MetricName: aws.String("AllRequests"),
			Namespace:  aws.String("AWS/EC2"),
			Dimensions: []cloudwatchTypes.Dimension{
				{
					Name:  aws.String("InstanceName"),
					Value: ec2.Name,
				},
			},
			Unit:   cloudwatchTypes.StandardUnitCount,
			Period: aws.Int32(3600),
			Statistics: []cloudwatchTypes.Statistic{
				cloudwatchTypes.StatisticAverage,
			},
		})

		if err != nil {
			log.Warnf("Couldn't fetch usage metric for %s", *bucket.Name)
		}

		

		sizeCharges := 0.0
		if metricesUsesOutput != nil && len(metricesUsesOutput.Datapoints) > 0 {
			sizeCharges = *metricesUsesOutput.Datapoints[0].Average
		}

		monthlyCost := sizeCharges * priceMap[*ec2.PlatformName]

		resources = append(resources, Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "EC2",
			Region:     client.AWSClient.Region,
			ResourceId: *ec2.InstanceId,
			Name:       *ec2.Name,
			Cost:       monthlyCost,
			CreatedAt:  *ec2.RegistrationDate,
			Tags: tags
			FetchedAt:  time.Now(),
			Link: 	 fmt.Sprintf("https://%s.console.aws.amazon.com/ec2/home?region=%s#InstanceDetails:instanceId=%s", client.AWSClient.Region, client.AWSClient.Region, *ec2.InstanceId),
		})

	}

	return nil, nil
}

func getEC2Relations(inst *etype.Instance, resourceArn string) (rel []models.Link) {
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

	// Get associated VPC
	rel = append(rel, models.Link{
		ResourceID: fmt.Sprintf("%s:vpc/%s", resourceArn, *inst.VpcId),
		Type:       "VPC",
		Name:       *inst.VpcId,
		Relation:   "USES",
	})

	// Get associated Subnet
	rel = append(rel, models.Link{
		ResourceID: fmt.Sprintf("%s:subnet/%s", resourceArn, *inst.SubnetId),
		Name:       *inst.SubnetId,
		Type:       "Subnet",
		Relation:   "USES",
	})

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
