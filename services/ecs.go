package services

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	. "github.com/mlabouardy/komiser/models"
)

func (aws AWS) DescribeECS(cfg aws.Config) (map[string]int, error) {
	countClusters := 0
	countTasks := 0
	countServices := 0
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return map[string]int{}, err
	}
	for _, region := range regions {
		clusters, err := aws.getClusters(cfg, region.Name)
		if err != nil {
			return map[string]int{}, err
		}
		countClusters += len(clusters)
		for _, cluster := range clusters {
			tasks, err := aws.getTasks(cfg, cluster.Name, region.Name)
			if err != nil {
				return map[string]int{}, err
			}
			countTasks += len(tasks)
			services, err := aws.getServices(cfg, cluster.Name, region.Name)
			if err != nil {
				return map[string]int{}, err
			}
			countServices += len(services)
		}
	}
	return map[string]int{
		"clusters": countClusters,
		"tasks":    countTasks,
		"services": countServices,
	}, nil
}

func (aws AWS) getClusters(cfg aws.Config, region string) ([]Cluster, error) {
	cfg.Region = region
	svc := ecs.New(cfg)
	req := svc.DescribeClustersRequest(&ecs.DescribeClustersInput{})
	result, err := req.Send()
	if err != nil {
		return []Cluster{}, err
	}
	listOfClusters := make([]Cluster, 0, len(result.Clusters))
	for _, cluster := range result.Clusters {
		listOfClusters = append(listOfClusters, Cluster{
			Name: *cluster.ClusterName,
		})
	}
	return listOfClusters, nil
}

func (aws AWS) getTasks(cfg aws.Config, cluster string, region string) ([]Task, error) {
	cfg.Region = region
	svc := ecs.New(cfg)
	req := svc.DescribeTasksRequest(&ecs.DescribeTasksInput{
		Cluster: &cluster,
	})
	result, err := req.Send()
	if err != nil {
		return []Task{}, err
	}
	listOfTasks := make([]Task, 0, len(result.Tasks))
	for _, task := range result.Tasks {
		listOfTasks = append(listOfTasks, Task{
			ARN: *task.TaskArn,
		})
	}
	return listOfTasks, nil
}

func (aws AWS) getServices(cfg aws.Config, cluster string, region string) ([]Service, error) {
	cfg.Region = region
	svc := ecs.New(cfg)
	req := svc.DescribeServicesRequest(&ecs.DescribeServicesInput{
		Cluster: &cluster,
	})
	result, err := req.Send()
	if err != nil {
		return []Service{}, err
	}
	listOfServices := make([]Service, 0, len(result.Services))
	for _, service := range result.Services {
		listOfServices = append(listOfServices, Service{
			Name: *service.ServiceName,
		})
	}
	return listOfServices, nil
}
