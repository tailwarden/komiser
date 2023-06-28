package ecs

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func TaskSet(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config ecs.ListTasksInput
	ecsClient := ecs.NewFromConfig(*client.AWSClient)
	stsClient := sts.NewFromConfig(*client.AWSClient)
	stsOutput, err := ecsClient.ListTaskDefinitions(context.Background(), &config)
	if err != nil {
		return resources, err
	}

	accountId := stsOutput

}
