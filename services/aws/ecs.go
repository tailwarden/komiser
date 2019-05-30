package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	. "github.com/mlabouardy/komiser/models/aws"
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
			tasks, err := aws.getTasks(cfg, cluster, region.Name)
			if err != nil {
				return map[string]int{}, err
			}
			countTasks += len(tasks)
			services, err := aws.getServices(cfg, cluster, region.Name)
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

func (aws AWS) getClusters(cfg aws.Config, region string) ([]string, error) {
	cfg.Region = region
	svc := ecs.New(cfg)
	req := svc.ListClustersRequest(&ecs.ListClustersInput{})
	result, err := req.Send(context.Background())
	return result.ClusterArns, err
}

func (aws AWS) getTasks(cfg aws.Config, cluster string, region string) ([]string, error) {
	cfg.Region = region
	svc := ecs.New(cfg)
	req := svc.ListTasksRequest(&ecs.ListTasksInput{
		Cluster: &cluster,
	})
	result, err := req.Send(context.Background())
	return result.TaskArns, err
}

func (aws AWS) getServices(cfg aws.Config, cluster string, region string) ([]string, error) {
	cfg.Region = region
	svc := ecs.New(cfg)
	req := svc.ListServicesRequest(&ecs.ListServicesInput{
		Cluster: &cluster,
	})
	result, err := req.Send(context.Background())
	return result.ServiceArns, err
}
