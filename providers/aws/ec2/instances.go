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
	. "github.com/mlabouardy/komiser/models"
	. "github.com/mlabouardy/komiser/providers"
)

type Ec2Product struct {
	Sku           string `json:sku`
	ProductFamily string `json:productFamily`
	Attributes    struct {
		Location        string `json:location`
		InstanceType    string `json:instanceType`
		Tenancy         string `json:tenancy`
		OperatingSystem string `json:operatingSystem`
		LicenseModel    string `json:licenseModel`
		UsageType       string `json:usagetype`
		PreInstalledSw  string `json:preInstalledSw`
	}
}

type PricingResult struct {
	Product Ec2Product `json:product`
	Terms   map[string]map[string]map[string]map[string]struct {
		PricePerUnit struct {
			USD string `json:USD`
		} `json:pricePerUnit`
	} `json:terms`
}

func GetRegionName(code string) string {
	regions := map[string]string{
		"us-east-2":      "US East (Ohio)",
		"us-east-1":      "US East (N. Virginia)",
		"us-west-1":      "US West (N. California)",
		"us-west-2":      "US West (Oregon)",
		"af-south-1":     "Africa (Cape Town)",
		"ap-east-1":      "Asia Pacific (Hong Kong)",
		"ap-south-2":     "Asia Pacific (Hyderabad)",
		"ap-southeast-3": "Asia Pacific (Jakarta)",
		"ap-south-1":     "Asia Pacific (Mumbai)",
		"ap-northeast-3": "Asia Pacific (Osaka)",
		"ap-northeast-2": "Asia Pacific (Seoul)",
		"ap-southeast-1": "Asia Pacific (Singapore)",
		"ap-southeast-2": "Asia Pacific (Sydney)",
		"ap-northeast-1": "Asia Pacific (Tokyo)",
		"ca-central-1":   "Canada (Central)",
		"eu-central-1":   "EU (Frankfurt)",
		"eu-west-1":      "EU (Ireland)",
		"eu-west-2":      "EU (London)",
		"eu-south-1":     "EU (Milan)",
		"eu-west-3":      "EU (Paris)",
		"eu-south-2":     "EU (Spain)",
		"eu-north-1":     "EU (Stockholm)",
		"eu-central-2":   "EU (Zurich)",
		"sa-east-1":      "South America (São Paulo)",
	}
	return regions[code]
}

func BeginningOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 0, -date.Day()+1)
}

func Instances(ctx context.Context, client ProviderClient) ([]Resource, error) {
	var nextToken string
	resources := make([]Resource, 0)
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
				tags := make([]Tag, 0)

				name := ""
				for _, tag := range instance.Tags {
					if *tag.Key == "Name" {
						name = *tag.Value
					}
					tags = append(tags, Tag{
						Key:   *tag.Key,
						Value: *tag.Value,
					})
				}

				startOfMonth := BeginningOfMonth(time.Now())
				hourlyUsage := 0
				if instance.LaunchTime.Before(startOfMonth) {
					hourlyUsage = int(time.Now().Sub(startOfMonth).Hours())
				} else {
					hourlyUsage = int(time.Now().Sub(*instance.LaunchTime).Hours())
				}

				pricingOutput, err := pricingClient.GetProducts(ctx, &pricing.GetProductsInput{
					ServiceCode: aws.String("AmazonEC2"),
					Filters: []types.Filter{
						types.Filter{
							Field: aws.String("operatingSystem"),
							Value: aws.String("linux"),
							Type:  types.FilterTypeTermMatch,
						},
						types.Filter{
							Field: aws.String("instanceType"),
							Value: aws.String(string(instance.InstanceType)),
							Type:  types.FilterTypeTermMatch,
						},
						types.Filter{
							Field: aws.String("location"),
							Value: aws.String(GetRegionName(client.AWSClient.Region)),
							Type:  types.FilterTypeTermMatch,
						},
						types.Filter{
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

					pricingResult := PricingResult{}
					json.Unmarshal([]byte(s), &pricingResult)

					hourlyCostRaw := pricingResult.Terms["OnDemand"][fmt.Sprintf("%s.JRTCKXETXF", pricingResult.Product.Sku)]["priceDimensions"][fmt.Sprintf("%s.JRTCKXETXF.6YS6EN2CT7", pricingResult.Product.Sku)].PricePerUnit.USD
					hourlyCost, _ = strconv.ParseFloat(hourlyCostRaw, 64)
				}

				monthlyCost := float64(hourlyUsage) * hourlyCost

				resourceArn := fmt.Sprintf("arn:aws:ec2:%s:%s:instance/%s", client.AWSClient.Region, *accountId, *instance.InstanceId)

				resources = append(resources, Resource{
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
						"instanceType": fmt.Sprintf("%s", instance.InstanceType),
						"state":        fmt.Sprintf("%s", instance.State.Name),
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
