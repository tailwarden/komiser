package cvm

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	tccvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

const (
	searchParams  string = "id%3D17&rid=17&"
	createdLayout string = "2006-01-02T15:04:05Z"
)

func Instances(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)
	zonesMap := make(map[string]string)

	zones, err := client.TencentClient.DescribeZones(tccvm.NewDescribeZonesRequest())
	if err != nil {
		log.Warnf("[%s][Tencent] Couldn't fetch the list of zones: %s", client.Name, err)
	}

	for _, zone := range zones.Response.ZoneSet {
		zonesMap[*zone.Zone] = *zone.ZoneName
	}

	instanceRequest := tccvm.NewDescribeInstancesRequest()
	instanceRequest.Limit = common.Int64Ptr(100)

	instances, err := client.TencentClient.DescribeInstances(instanceRequest)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		log.Errorf("[%s][Tencent] An API error has returned: %s", client.Name, err)
	}
	if err != nil {
		return nil, err
	}

	for _, instance := range instances.Response.InstanceSet {
		tags := make([]models.Tag, 0, len(instance.Tags))
		for _, tag := range instance.Tags {
			tags = append(tags, models.Tag{
				Key:   *tag.Key,
				Value: *tag.Value,
			})
		}

		zone := *instance.Placement.Zone
		if val, ok := zonesMap[zone]; ok {
			zone = val
		}

		priceRequest := tccvm.NewInquiryPriceRunInstancesRequest()
		priceRequest.InstanceName = instance.InstanceName
		priceRequest.ImageId = instance.ImageId
		priceRequest.Placement = instance.Placement

		price, err := client.TencentClient.InquiryPriceRunInstances(priceRequest)
		if err != nil {
			log.Warnf("[%s][Tencent] Couldn't fetch the price of instance: %s", client.Name, err)
		}

		currentTime := time.Now()
		currentMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, time.UTC)

		created, err := time.Parse(createdLayout, *instance.CreatedTime)
		if err != nil {
			return nil, err
		}

		duration := currentTime.Sub(created)
		if created.Before(currentMonth) {
			duration = currentTime.Sub(currentMonth)
		}

		totalPrice := *price.Response.Price.InstancePrice.UnitPriceDiscount + *price.Response.Price.BandwidthPrice.UnitPriceDiscount
		monthlyCost := totalPrice * float64(duration.Hours())

		resources = append(resources, models.Resource{
			Provider:   "Tencent",
			Account:    client.Name,
			Service:    "Instance",
			ResourceId: *instance.InstanceId,
			Region:     zone,
			Name:       *instance.InstanceName,
			Cost:       monthlyCost,
			Tags:       tags,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://console.intl.cloud.tencent.com/cvm/instance/detail?searchParams=%sid=%s", searchParams, *instance.InstanceId),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Tencent",
		"account":   client.Name,
		"service":   "Instance",
		"region":    client.TencentClient.GetRegion(),
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
