package gcp

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	monitoring "google.golang.org/api/monitoring/v3"
)

func (gcp GCP) GetBillableReceivedLogs() ([]*monitoring.TimeSeries, error) {
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
		AggregationGroupByFields("project", "resource.labels.resource_type").
		AggregationPerSeriesAligner("ALIGN_SUM").
		Filter(`metric.type="logging.googleapis.com/billing/monthly_bytes_ingested"`).
		IntervalEndTime(time.Now().Format("2006-01-02T15:04:05.000Z")).
		IntervalStartTime(time.Now().AddDate(0, 0, -7).Format("2006-01-02T15:04:05.000Z")).
		Do()

	if err != nil {
		log.Println(err)
		return []*monitoring.TimeSeries{}, err
	}

	return data.TimeSeries, nil
}
