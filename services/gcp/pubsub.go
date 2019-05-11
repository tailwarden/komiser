package gcp

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	pubsub "google.golang.org/api/pubsub/v1"
)

func (gcp GCP) GetTopics() (int, error) {
	sum := 0

	src, err := google.DefaultTokenSource(oauth2.NoContext, pubsub.PubsubScope)
	if err != nil {
		return sum, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := pubsub.New(client)
	if err != nil {
		return sum, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return sum, err
	}

	for _, project := range projects {
		uri := fmt.Sprintf("projects/%s", project.ID)
		topics, err := svc.Projects.Topics.List(uri).Do()
		if err != nil {
			log.Println(err)
			return 0, err
		}
		sum += len(topics.Topics)
	}

	return sum, nil
}
