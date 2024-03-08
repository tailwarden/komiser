package wafv2

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func WebAcls(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	wAclClient := wafv2.NewFromConfig(*client.AWSClient)

	wacls, err := wAclClient.ListWebACLs(ctx, &wafv2.ListWebACLsInput{})
	if err != nil {
		return resources, err
	}

	for _, acl := range wacls.WebACLs {
		outputTags, err := wAclClient.ListTagsForResource(ctx, &wafv2.ListTagsForResourceInput{
			ResourceARN: acl.ARN,
		})

		tags := make([]models.Tag, 0)

		if err == nil {
			for _, tag := range outputTags.TagInfoForResource.TagList {
				tags = append(tags, models.Tag{
					Key:   *tag.Key,
					Value: *tag.Value,
				})
			}
		}

		resources = append(resources, models.Resource{
			Provider:   "AWS",
			Account:    client.Name,
			Service:    "Web Acl",
			ResourceId: *acl.ARN,
			Region:     client.AWSClient.Region,
			Name:       *acl.Name,
			Cost:       0,
			Tags:       tags,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://%s.console.aws.amazon.com/wafv2/homev2/web-acl/%s/%s?region=%s", client.AWSClient.Region, *acl.Name, *acl.Id, client.AWSClient.Region),
		})
	}

	log.WithFields(log.Fields{
		"provider":    "AWS",
		"account":     client.Name,
		"region":      client.AWSClient.Region,
		"service":     "Web Acl",
		"resources":   len(resources),
		"serviceCost": fmt.Sprint(0),
	}).Info("Fetched resources")
	return resources, nil
}
