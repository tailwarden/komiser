package gcp

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	bigquery "google.golang.org/api/bigquery/v2"
	monitoring "google.golang.org/api/monitoring/v3"
)

func (gcp GCP) GetBigQueryTables() (int, error) {
	sum := 0

	src, err := google.DefaultTokenSource(oauth2.NoContext, bigquery.CloudPlatformReadOnlyScope)
	if err != nil {
		return sum, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := bigquery.New(client)
	if err != nil {
		return sum, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return sum, err
	}

	for _, project := range projects {
		datasets, err := svc.Datasets.List(project.ID).Do()
		if err != nil {
			log.Println(err)
			return sum, err
		}
		for _, dataset := range datasets.Datasets {
			tables, err := svc.Tables.List(project.ID, dataset.DatasetReference.DatasetId).Do()
			if err != nil {
				log.Println(err)
				return sum, err
			}
			sum += len(tables.Tables)
		}
	}
	return sum, nil
}

func (gcp GCP) GetBigQueryDatasets() (int, error) {
	sum := 0

	src, err := google.DefaultTokenSource(oauth2.NoContext, bigquery.CloudPlatformReadOnlyScope)
	if err != nil {
		return sum, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := bigquery.New(client)
	if err != nil {
		return sum, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return sum, err
	}

	for _, project := range projects {
		datasets, err := svc.Datasets.List(project.ID).Do()
		if err != nil {
			log.Println(err)
			return sum, err
		}
		sum += len(datasets.Datasets)
	}

	return sum, nil
}

func (gcp GCP) GetBigQueryScannedStatements() ([]*monitoring.TimeSeries, error) {
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
		AggregationGroupByFields("project", "resource.labels.statement_type").
		AggregationPerSeriesAligner("ALIGN_RATE").
		Filter(`metric.type="bigquery.googleapis.com/query/statement_scanned_bytes_billed"`).
		IntervalEndTime(time.Now().Format("2006-01-02T15:04:05.000Z")).
		IntervalStartTime(time.Now().AddDate(0, 0, -7).Format("2006-01-02T15:04:05.000Z")).
		Do()

	if err != nil {
		log.Println(err)
		return []*monitoring.TimeSeries{}, err
	}

	return data.TimeSeries, nil
}

func (gcp GCP) GetBigQueryStoredBytes() ([]*monitoring.TimeSeries, error) {
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
		AggregationCrossSeriesReducer("REDUCE_SUM").
		AggregationGroupByFields("resource.labels.dataset_id").
		AggregationPerSeriesAligner("ALIGN_SUM").
		Filter(`metric.type="bigquery.googleapis.com/storage/stored_bytes"`).
		IntervalEndTime(time.Now().Format("2006-01-02T15:04:05.000Z")).
		IntervalStartTime(time.Now().AddDate(0, 0, -7).Format("2006-01-02T15:04:05.000Z")).
		Do()

	if err != nil {
		log.Println(err)
		return []*monitoring.TimeSeries{}, err
	}

	return data.TimeSeries, nil
}
