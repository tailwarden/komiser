package s3

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3control"
	log "github.com/sirupsen/logrus"

	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func AccessPoint(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	mySession := session.Must(session.NewSession())

	svc := s3control.New(mySession)
	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return resources, err
	}
	AccountId := stsOutput.Account
	Name := stsOutput.Arn

	svc.CreateAccessPointRequest(&s3control.CreateAccessPointInput{
		AccountId: AccountId,
		Name:      Name,
	})
	input := s3control.ListAccessPointsInput{
		AccountId: AccountId,
	}

	listOfAccessPoints, err := svc.ListAccessPoints(&input)

	AccesspointList := []string{}

	for _, accessPoints := range listOfAccessPoints.AccessPointList {
		AccesspointList = append(AccesspointList, *accessPoints.Name)
	}
	resources = append(resources, Resource{
		Provider:  "AWS",
		Account:   client.Name,
		Service:   "S3-AccessPoint",
		Region:    client.AWSClient.Region,
		FetchedAt: time.Now(),
	})
	log.WithFields(log.Fields{
		"provider":  "AWS",
		"account":   client.Name,
		"region":    client.AWSClient.Region,
		"service":   "S3-AccessPoint",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, err
	// return listOfAccessPoints.AccessPointList, err

}
