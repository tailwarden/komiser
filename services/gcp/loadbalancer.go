package gcp

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
	monitoring "google.golang.org/api/monitoring/v3"
)

func (gcp GCP) GetLoadBalancerRequests() ([]*monitoring.TimeSeries, error) {
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
		AggregationGroupByFields("resource.labels.project_id").
		AggregationPerSeriesAligner("ALIGN_SUM").
		Filter(`metric.type="loadbalancing.googleapis.com/https/request_count"`).
		IntervalEndTime(time.Now().Format("2006-01-02T15:04:05.000Z")).
		IntervalStartTime(time.Now().AddDate(0, 0, -7).Format("2006-01-02T15:04:05.000Z")).
		Do()

	if err != nil {
		log.Println(err)
		return []*monitoring.TimeSeries{}, err
	}

	return data.TimeSeries, nil
}

func (gcp GCP) GetTotalLoadBalancers() (int, error) {
	sum := 0

	src, err := google.DefaultTokenSource(oauth2.NoContext, compute.ComputeReadonlyScope)
	if err != nil {
		return sum, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := compute.New(client)
	if err != nil {
		return sum, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return sum, err
	}

	for _, project := range projects {
		urls, err := svc.UrlMaps.List(project.ID).Do()
		if err != nil {
			log.Println(err)
			return 0, err
		}
		sum += len(urls.Items)
	}

	return sum, nil
}
