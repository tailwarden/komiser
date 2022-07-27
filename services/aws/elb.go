package aws

import (
	"context"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (awsClient AWS) DescribeElasticLoadBalancer(cfg aws.Config) (map[string]int, error) {
	output := make(map[string]int, 0)
	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return map[string]int{}, err
	}
	for _, region := range regions {
		elbsv1, err := awsClient.getClassicElasticLoadBalancers(cfg, region.Name)
		if err != nil {
			return map[string]int{}, err
		}
		elbsv2, err := awsClient.getElasticLoadBalancersV2(cfg, region.Name)
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
		cloudwatchClient := cloudwatch.NewFromConfig(cfg)
		resultCloudWatch, err := cloudwatchClient.GetMetricStatistics(context.Background(), &cloudwatch.GetMetricStatisticsInput{
			Namespace:  aws.String("AWS/ELB"),
			MetricName: aws.String("RequestCount"),
			StartTime:  aws.Time(time.Now().AddDate(0, -6, 0)),
			EndTime:    aws.Time(time.Now()),
			Period:     aws.Int32(86400),
			Statistics: []types.Statistic{
				types.StatisticSum,
			},
		})
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

func (awsClient AWS) getClassicElasticLoadBalancers(cfg aws.Config, region string) ([]LoadBalancer, error) {
	cfg.Region = region
	svc := elasticloadbalancing.NewFromConfig(cfg)
	result, err := svc.DescribeLoadBalancers(context.Background(), &elasticloadbalancing.DescribeLoadBalancersInput{})
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

func (awsClient AWS) getElasticLoadBalancersV2(cfg aws.Config, region string) ([]LoadBalancer, error) {
	cfg.Region = region
	svc := elasticloadbalancingv2.NewFromConfig(cfg)
	result, err := svc.DescribeLoadBalancers(context.Background(), &elasticloadbalancingv2.DescribeLoadBalancersInput{})
	if err != nil {
		return []LoadBalancer{}, err
	}
	listOfElasticLoadBalancers := make([]LoadBalancer, 0)
	for _, lb := range result.LoadBalancers {
		listOfElasticLoadBalancers = append(listOfElasticLoadBalancers, LoadBalancer{
			DNSName: *lb.DNSName,
			State:   *lb.State.Reason,
			Type:    string(lb.Type),
		})
	}
	return listOfElasticLoadBalancers, nil
}
