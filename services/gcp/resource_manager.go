package gcp

import (
	"context"
	"log"
	"strings"

	. "github.com/mlabouardy/komiser/models/gcp"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	cloudresourcemanager "google.golang.org/api/cloudresourcemanager/v1"
	iam "google.golang.org/api/iam/v1"
)

func (gcp GCP) GetProjects() ([]Project, error) {
	projects := make([]Project, 0)

	src, err := google.DefaultTokenSource(oauth2.NoContext, cloudresourcemanager.CloudPlatformReadOnlyScope)
	if err != nil {
		return projects, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := cloudresourcemanager.New(client)
	if err != nil {
		return projects, err
	}

	res, err := svc.Projects.List().Do()
	if err != nil {
		log.Println(err)
		return projects, err
	}

	for _, project := range res.Projects {
		projects = append(projects, Project{
			Name:       project.Name,
			ID:         project.ProjectId,
			CreateTime: project.CreateTime,
			Number:     project.ProjectNumber,
		})
	}
	return projects, nil
}

func (gcp GCP) GetIamUsers() (int, error) {
	src, err := google.DefaultTokenSource(oauth2.NoContext, iam.CloudPlatformScope)
	if err != nil {
		return 0, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := cloudresourcemanager.New(client)
	if err != nil {
		return 0, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return 0, err
	}

	users := make(map[string]struct{})

	for _, project := range projects {
		policy, err := svc.Projects.GetIamPolicy(project.ID, &cloudresourcemanager.GetIamPolicyRequest{}).Do()
		if err != nil {
			log.Println(err)
			return 0, err
		}
		for _, role := range policy.Bindings {
			for _, member := range role.Members {
				if strings.HasPrefix(member, "user") {
					users[strings.TrimPrefix(member, "user:")] = struct{}{}
				}
			}
		}
	}
	return len(users), nil
}
