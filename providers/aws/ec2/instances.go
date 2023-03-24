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
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/aws/aws-sdk-go-v2/service/pricing/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
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
					b, _ := json.Marshal(pricingOutput.PriceList[0])
					s, _ := strconv.Unquote(string(b))

					pricingResult := models.PricingResult{}
					err = json.Unmarshal([]byte(s), &pricingResult)
					if err != nil {
						log.WithError(err).Error("could not unmarshal")
					}

					hourlyCostRaw := pricingResult.Terms["OnDemand"][fmt.Sprintf("%s.JRTCKXETXF", pricingResult.Product.Sku)]["priceDimensions"][fmt.Sprintf("%s.JRTCKXETXF.6YS6EN2CT7", pricingResult.Product.Sku)].PricePerUnit.USD
					hourlyCost, _ = strconv.ParseFloat(hourlyCostRaw, 64)
				}

				monthlyCost := float64(hourlyUsage) * hourlyCost

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
