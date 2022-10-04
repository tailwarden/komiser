package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	. "github.com/mlabouardy/komiser/models/aws"
)

type ECSData struct {
	Services []Service
	Clusters []Cluster
	Tasks    []Task
}

func (awsClient AWS) DescribeECS(cfg awsConfig.Config) (ECSData, error) {
	ecsData := ECSData{
		Services: make([]Service, 0),
		Clusters: make([]Cluster, 0),
		Tasks:    make([]Task, 0),
	}

	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return ecsData, err
	}
	for _, region := range regions {
		clusters, err := awsClient.getClusters(cfg, region.Name)
		if err != nil {
			return ecsData, err
		}

		for _, cluster := range clusters {
			ecsData.Clusters = append(ecsData.Clusters, cluster)
		}

		for _, cluster := range clusters {
			tasks, err := awsClient.getTasks(cfg, cluster.Name, region.Name)
			if err != nil {
				return ecsData, err
			}

			for _, t := range tasks {
				ecsData.Tasks = append(ecsData.Tasks, t)
			}

			services, err := awsClient.getServices(cfg, cluster.Name, region.Name)
			if err != nil {
				return ecsData, err
			}
			for _, s := range services {
				ecsData.Services = append(ecsData.Services, s)
			}
		}
	}
	return ecsData, nil
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
			ARN:    *cluster.ClusterArn,
			Name:   *cluster.ClusterName,
			Tags:   tags,
			Region: region,
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
		tags := make([]string, 0)
		for _, t := range task.Tags {
			tags = append(tags, fmt.Sprintf("%s:%s", *t.Key, *t.Value))
		}
		listOfTasks = append(listOfTasks, Task{
			ARN:       *task.TaskArn,
			CreatedAt: *task.CreatedAt,
			Tags:      tags,
			Region:    region,
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
		tags := make([]string, 0)
		for _, t := range service.Tags {
			tags = append(tags, fmt.Sprintf("%s:%s", *t.Key, *t.Value))
		}
		listOfServices = append(listOfServices, Service{
			Name:      *service.ServiceName,
			CreatedAt: *service.CreatedAt,
			Tags:      tags,
			Region:    region,
		})
	}
	return listOfServices, nil
}
