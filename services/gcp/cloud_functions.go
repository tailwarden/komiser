package gcp

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	cloudfunctions "google.golang.org/api/cloudfunctions/v1"
)

func (gcp GCP) CloudFunctions() (map[string]int, error) {
	listOfFunctions := make(map[string]int, 0)

	src, err := google.DefaultTokenSource(oauth2.NoContext, cloudfunctions.CloudPlatformScope)
	if err != nil {
		return listOfFunctions, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := cloudfunctions.New(client)
	if err != nil {
		return listOfFunctions, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return listOfFunctions, err
	}

	for _, project := range projects {
		uri := fmt.Sprintf("projects/%s/locations/-", project.ID)

		functions, err := svc.Projects.Locations.Functions.List(uri).Do()
		if err != nil {
			log.Println(err)
			return listOfFunctions, err
		}
		for _, function := range functions.Functions {
			listOfFunctions[function.Runtime]++
		}
	}

	return listOfFunctions, nil
}
