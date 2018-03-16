package services

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elb"
	"github.com/aws/aws-sdk-go-v2/service/elbv2"
	. "github.com/mlabouardy/komiser/models"
)

func (aws AWS) DescribeElasticLoadBalancerPerFamily(cfg aws.Config) (map[string]int, error) {
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

func (aws AWS) getClassicElasticLoadBalancers(cfg aws.Config, region string) ([]LoadBalancer, error) {
	cfg.Region = region
	svc := elb.New(cfg)
	req := svc.DescribeLoadBalancersRequest(&elb.DescribeLoadBalancersInput{})
	result, err := req.Send()
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
	result, err := req.Send()
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
