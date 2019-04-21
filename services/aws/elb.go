package aws

import (
	"context"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/elb"
	"github.com/aws/aws-sdk-go-v2/service/elbv2"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (aws AWS) DescribeElasticLoadBalancer(cfg aws.Config) (map[string]int, error) {
	output := make(map[string]int, 0)
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return map[string]int{}, err
	}
	for _, region := range regions {
		elbsv1, err := aws.getClassicElasticLoadBalancers(cfg, region.Name)
		if err != nil {
			return map[string]int{}, err
		}
		elbsv2, err := aws.getElasticLoadBalancersV2(cfg, region.Name)
		if err != nil {
			return map[string]int{}, err
		}
		for _, elb := range elbsv1 {
			output[elb.Type]++
		}
		for _, elb := range elbsv2 {
			output[elb.Type]++
		}
	}
	return output, nil
}

type ELBMetric struct {
	Region     string
	Datapoints []Datapoint
}

func (awsClient AWS) GetELBRequests(cfg aws.Config) ([]ELBMetric, error) {
	metrics := make([]ELBMetric, 0)

	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return metrics, err
	}

	for _, region := range regions {
		cfg.Region = region.Name
		cloudwatchClient := cloudwatch.New(cfg)
		reqCloudwatch := cloudwatchClient.GetMetricStatisticsRequest(&cloudwatch.GetMetricStatisticsInput{
			Namespace:  aws.String("AWS/ELB"),
			MetricName: aws.String("RequestCount"),
			StartTime:  aws.Time(time.Now().AddDate(0, -6, 0)),
			EndTime:    aws.Time(time.Now()),
			Period:     aws.Int64(86400),
			Statistics: []cloudwatch.Statistic{
				cloudwatch.StatisticSum,
			},
		})

		resultCloudWatch, err := reqCloudwatch.Send(context.Background())
		if err != nil {
			return metrics, err
		}

		series := make([]Datapoint, 0)

		for _, metric := range resultCloudWatch.Datapoints {
			series = append(series, Datapoint{
				Timestamp: *metric.Timestamp,
				Value:     *metric.Sum,
			})
		}

		sort.Slice(series, func(i, j int) bool {
			return series[i].Timestamp.Sub(series[j].Timestamp) < 0
		})

		metrics = append(metrics, ELBMetric{
			Region:     region.Name,
			Datapoints: series,
		})
	}

	return metrics, nil
}

func (aws AWS) getClassicElasticLoadBalancers(cfg aws.Config, region string) ([]LoadBalancer, error) {
	cfg.Region = region
	svc := elb.New(cfg)
	req := svc.DescribeLoadBalancersRequest(&elb.DescribeLoadBalancersInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return []LoadBalancer{}, err
	}
	listOfElasticLoadBalancers := make([]LoadBalancer, 0)
	for _, lb := range result.LoadBalancerDescriptions {
		listOfElasticLoadBalancers = append(listOfElasticLoadBalancers, LoadBalancer{
			DNSName: *lb.DNSName,
			Type:    "classic",
		})
	}
	return listOfElasticLoadBalancers, nil
}

func (aws AWS) getElasticLoadBalancersV2(cfg aws.Config, region string) ([]LoadBalancer, error) {
	cfg.Region = region
	svc := elbv2.New(cfg)
	req := svc.DescribeLoadBalancersRequest(&elbv2.DescribeLoadBalancersInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return []LoadBalancer{}, err
	}
	listOfElasticLoadBalancers := make([]LoadBalancer, 0)
	for _, lb := range result.LoadBalancers {
		lbType, _ := lb.Type.MarshalValue()
		listOfElasticLoadBalancers = append(listOfElasticLoadBalancers, LoadBalancer{
			DNSName: *lb.DNSName,
			State:   lb.State.String(),
			Type:    lbType,
		})
	}
	return listOfElasticLoadBalancers, nil
}
