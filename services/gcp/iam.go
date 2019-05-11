package gcp

import (
	"context"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	iam "google.golang.org/api/iam/v1"
)

func (gcp GCP) GetIamRoles() (int, error) {
	src, err := google.DefaultTokenSource(oauth2.NoContext, iam.CloudPlatformScope)
	if err != nil {
		return 0, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := iam.New(client)
	if err != nil {
		return 0, err
	}

	roles, err := svc.Roles.List().Do()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return len(roles.Roles), nil
}

func (gcp GCP) GetServiceAccounts() (int, error) {
	sum := 0

	src, err := google.DefaultTokenSource(oauth2.NoContext, iam.CloudPlatformScope)
	if err != nil {
		return sum, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := iam.New(client)
	if err != nil {
		return sum, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return sum, err
	}

	for _, project := range projects {
		serviceAccounts, err := svc.Projects.ServiceAccounts.List("projects/" + project.ID).Do()
		if err != nil {
			log.Println(err)
			return sum, err
		}

		sum += len(serviceAccounts.Accounts)
	}

	return sum, nil
}
