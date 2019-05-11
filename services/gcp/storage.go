package gcp

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	monitoring "google.golang.org/api/monitoring/v3"
	storage "google.golang.org/api/storage/v1"
)

func (gcp GCP) GetTotalBuckets() (int, error) {
	sum := 0

	src, err := google.DefaultTokenSource(oauth2.NoContext, storage.CloudPlatformReadOnlyScope)
	if err != nil {
		return sum, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := storage.New(client)
	if err != nil {
		return sum, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return sum, err
	}

	for _, project := range projects {
		buckets, err := svc.Buckets.List(project.ID).Do()
		if err != nil {
			log.Println(err)
			return 0, err
		}
		sum += len(buckets.Items)
	}

	return sum, nil
}

func (gcp GCP) GetBucketSize() ([]*monitoring.TimeSeries, error) {
	src, err := google.DefaultTokenSource(oauth2.NoContext, monitoring.MonitoringReadScope)
	if err != nil {
		return []*monitoring.TimeSeries{}, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := monitoring.New(client)
	if err != nil {
		return []*monitoring.TimeSeries{}, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return []*monitoring.TimeSeries{}, err
	}

	uri := fmt.Sprintf("projects/%s", projects[0].ID)
	data, err := svc.Projects.TimeSeries.
		List(uri).
		AggregationAlignmentPeriod("86400s").
		AggregationGroupByFields("project").
		AggregationPerSeriesAligner("ALIGN_MEAN").
		Filter(`metric.type="storage.googleapis.com/storage/total_bytes"`).
		IntervalEndTime(time.Now().Format("2006-01-02T15:04:05.000Z")).
		IntervalStartTime(time.Now().AddDate(0, 0, -7).Format("2006-01-02T15:04:05.000Z")).
		Do()

	if err != nil {
		log.Println(err)
		return []*monitoring.TimeSeries{}, err
	}

	return data.TimeSeries, nil
}

func (gcp GCP) GetBucketObjects() ([]*monitoring.TimeSeries, error) {
	src, err := google.DefaultTokenSource(oauth2.NoContext, monitoring.MonitoringReadScope)
	if err != nil {
		return []*monitoring.TimeSeries{}, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := monitoring.New(client)
	if err != nil {
		return []*monitoring.TimeSeries{}, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return []*monitoring.TimeSeries{}, err
	}

	uri := fmt.Sprintf("projects/%s", projects[0].ID)
	data, err := svc.Projects.TimeSeries.
		List(uri).
		AggregationAlignmentPeriod("86400s").
		AggregationGroupByFields("project").
		AggregationPerSeriesAligner("ALIGN_SUM").
		Filter(`metric.type="storage.googleapis.com/storage/object_count"`).
		IntervalEndTime(time.Now().Format("2006-01-02T15:04:05.000Z")).
		IntervalStartTime(time.Now().AddDate(0, 0, -7).Format("2006-01-02T15:04:05.000Z")).
		Do()

	if err != nil {
		log.Println(err)
		return []*monitoring.TimeSeries{}, err
	}

	return data.TimeSeries, nil
}
