package gcp

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	redis "google.golang.org/api/redis/v1"
)

func (gcp GCP) GetRedisInstances() (int, error) {
	sum := 0

	src, err := google.DefaultTokenSource(oauth2.NoContext, redis.CloudPlatformScope)
	if err != nil {
		return sum, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := redis.New(client)
	if err != nil {
		return sum, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return sum, err
	}

	for _, project := range projects {
		uri := fmt.Sprintf("projects/%s/locations/-", project.ID)
		instances, err := svc.Projects.Locations.Instances.List(uri).Do()
		if err != nil {
			log.Println(err)
			return sum, err
		}
		sum += len(instances.Instances)
	}
	return sum, nil
}
