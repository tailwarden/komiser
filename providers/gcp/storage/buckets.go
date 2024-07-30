package storage

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	monitoringpb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"cloud.google.com/go/storage"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// GetBucketSize returns the size of a bucket in bytes
func GetBucketSize(ctx context.Context, client providers.ProviderClient, bucketName string) (int64, error) {
	monitoringClient, err := monitoring.NewMetricClient(ctx, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		log.WithError(err).Debug("failed to create monitoring client")
		return 0, err
	}

	beginningOfMonth := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.UTC).Unix()

	req := &monitoringpb.ListTimeSeriesRequest{
		Name:   "projects/" + client.GCPClient.Credentials.ProjectID,
		Filter: fmt.Sprintf(`metric.type="storage.googleapis.com/storage/total_bytes" resource.type="gcs_bucket" resource.label.bucket_name="%s"`, bucketName),
		Interval: &monitoringpb.TimeInterval{
			StartTime: &timestamp.Timestamp{Seconds: beginningOfMonth},
			EndTime:   &timestamp.Timestamp{Seconds: time.Now().Unix()},
		},
		Aggregation: &monitoringpb.Aggregation{
			AlignmentPeriod:  &duration.Duration{Seconds: 86400},
			PerSeriesAligner: monitoringpb.Aggregation_ALIGN_MEAN,
			GroupByFields:    []string{"resource.label.bucket_name"},
		},
		View: monitoringpb.ListTimeSeriesRequest_FULL,
	}

	res := monitoringClient.ListTimeSeries(ctx, req)

	for {
		series, err := res.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.WithError(err).Debug("failed to list time series")
			return 0, err
		}

		for _, point := range series.Points {
			return int64(point.Value.GetDoubleValue()), nil
		}
	}

	return 0, errors.New("no data found")
}

func getPricingForGCPBuckets() map[string]float64 {
	return map[string]float64{
		"STANDARD": 0.20,
		"NEARLINE": 0.10,
		"COLDLINE": 0.01,
		"ARCHIVE":  0.005,
	}
}

func Buckets(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	storageClient, err := storage.NewClient(ctx, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		return []models.Resource{}, err
	}

	buckets := storageClient.Buckets(ctx, client.GCPClient.Credentials.ProjectID)
	for {
		bucket, err := buckets.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			if strings.Contains(err.Error(), "SERVICE_DISABLED") {
				log.Warn(err.Error())
				return resources, nil
			} else {
				log.WithError(err).Errorf("failed to list buckets")
				return resources, err
			}
		}

		tags := make([]models.Tag, 0)
		if bucket.Labels != nil {
			for key, value := range bucket.Labels {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: value,
				})
			}
		}

		monthlyCost := 0.0

		bucketSize, err := GetBucketSize(ctx, client, bucket.Name)
		if err != nil {
			log.WithError(err).WithFields(log.Fields{
				"bucket": bucket.Name,
			}).Debug("failed to get bucket size, skipping cost calculation")
		}

		if bucketSize > 0 {
			bucketSizeInGB := float64(bucketSize) / 1024 / 1024 / 1024
			monthlyCost = float64(bucketSizeInGB) * getPricingForGCPBuckets()[strings.ToUpper(bucket.StorageClass)]
		}

		resources = append(resources, models.Resource{
			Provider:   "GCP",
			Account:    client.Name,
			Service:    "Bucket",
			Name:       bucket.Name,
			ResourceId: bucket.Name,
			Region:     strings.ToLower(bucket.Location),
			Tags:       tags,
			Cost:       monthlyCost,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://console.cloud.google.com/storage/browser/%s?project=%s", bucket.Name, client.GCPClient.Credentials.ProjectID),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "GCP",
		"account":   client.Name,
		"service":   "Bucket",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
