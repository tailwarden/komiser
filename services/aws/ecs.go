package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (awsClient AWS) DescribeECS(cfg awsConfig.Config) (map[string]int, error) {
	countClusters := 0
	countTasks := 0
	countServices := 0
	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return map[string]int{}, err
	}
	for _, region := range regions {
		clusters, err := awsClient.getClusters(cfg, region.Name)
		if err != nil {
			return map[string]int{}, err
		}
		countClusters += len(clusters)
		for _, cluster := range clusters {
			tasks, err := awsClient.getTasks(cfg, cluster.Name, region.Name)
			if err != nil {
				return map[string]int{}, err
			}
			countTasks += len(tasks)
			services, err := awsClient.getServices(cfg, cluster.Name, region.Name)
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

func (awsClient AWS) getClusters(cfg awsConfig.Config, region string) ([]Cluster, error) {
	cfg.Region = region
	svc := ecs.NewFromConfig(cfg)
	result, err := svc.DescribeClusters(context.Background(), &ecs.DescribeClustersInput{})
	if err != nil {
		return []Cluster{}, err
	}
	listOfClusters := make([]Cluster, 0, len(result.Clusters))
	for _, cluster := range result.Clusters {
		tags := make([]string, 0)
		for _, t := range cluster.Tags {
			tags = append(tags, fmt.Sprintf("%s:%s", *t.Key, *t.Value))
		}
		listOfClusters = append(listOfClusters, Cluster{
			ARN:  *cluster.ClusterArn,
			Name: *cluster.ClusterName,
			Tags: tags,
		})
	}
	return listOfClusters, nil
}

func (awsClient AWS) getTasks(cfg aws.Config, cluster string, region string) ([]Task, error) {
	cfg.Region = region
	svc := ecs.NewFromConfig(cfg)
	result, err := svc.DescribeTasks(context.Background(), &ecs.DescribeTasksInput{
		Cluster: &cluster,
	})
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

func (awsClient AWS) getServices(cfg awsConfig.Config, cluster string, region string) ([]Service, error) {
	cfg.Region = region
	svc := ecs.NewFromConfig(cfg)
	result, err := svc.DescribeServices(context.Background(), &ecs.DescribeServicesInput{
		Cluster: &cluster,
	})
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
