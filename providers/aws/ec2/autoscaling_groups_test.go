package ec2_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/stretchr/testify/assert"
	"github.com/tailwarden/komiser/mocks"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers/aws/ec2"
)

func TestDiscoverReturnsNoErrorIfHappyPath(t *testing.T) {

	// Arrange
	ctx := context.Background()
	var queryInput autoscaling.DescribeAutoScalingGroupsInput
	asgClient := mocks.NewAutoScalingGroupClient(t)
	asgClient.On(
		"DescribeAutoScalingGroups",
		ctx,
		&queryInput,
	).Return(&autoscaling.DescribeAutoScalingGroupsOutput{}, nil)

	discoverer := ec2.ASGDiscoverer{
		Client:      asgClient,
		Ctx:         ctx,
		AccountName: "test",
		Region:      "us-east-1",
	}

	// Act
	_, err := discoverer.Discover()

	// Assert
	assert.NoError(t, err)
}

func TestDiscoverReturnsResourcesAsExpectedWhenDiscovered(t *testing.T) {

	// Arrange
	asg1 := newASG("foo")
	asg2 := newASG("bar")
	ctx := context.Background()
	var queryInput autoscaling.DescribeAutoScalingGroupsInput
	asgClient := mocks.NewAutoScalingGroupClient(t)
	asgClient.On(
		"DescribeAutoScalingGroups",
		ctx,
		&queryInput,
	).Return(&autoscaling.DescribeAutoScalingGroupsOutput{
		AutoScalingGroups: []types.AutoScalingGroup{asg1, asg2},
	}, nil)

	discoverer := ec2.ASGDiscoverer{
		Client:      asgClient,
		Ctx:         ctx,
		AccountName: "my-account",
		Region:      "my-region",
	}

	// Act
	resources, err := discoverer.Discover()

	// Assert
	expectedTags := make([]models.Tag, 0)
	expectedTags = append(expectedTags, models.Tag{
		Key:   "Name",
		Value: "my-asg-foo",
	})

	assert.NoError(t, err)
	assert.Len(t, resources, 2)
	assert.Equal(t, "AWS", resources[0].Provider)
	assert.Equal(t, "my-account", resources[0].Account)
	assert.Equal(t, "AutoScalingGroup", resources[0].Service)
	assert.Equal(t, "arn:aws:autoscaling:us-east-1:123456789012:autoScalingGroup:12345678-1234-1234-1234-123456789012:autoScalingGroupName/my-asg-foo", resources[0].ResourceId)
	assert.Equal(t, float64(0), resources[0].Cost)
	assert.Equal(t, expectedTags, resources[0].Tags)
	assert.Equal(t, "https://my-region.console.aws.amazon.com/ec2/home?region=my-region#AutoScalingGroupDetails:id=my-asg-foo", resources[0].Link)
}

func TestDiscoverReturnsErrorIfCannotDiscoverASGs(t *testing.T) {

	// Arrange
	ctx := context.Background()
	var queryInput autoscaling.DescribeAutoScalingGroupsInput
	asgClient := mocks.NewAutoScalingGroupClient(t)
	asgClient.On(
		"DescribeAutoScalingGroups",
		ctx,
		&queryInput,
	).Return(&autoscaling.DescribeAutoScalingGroupsOutput{}, errors.New("Could not discover ASGs"))

	discoverer := ec2.ASGDiscoverer{
		Client:      asgClient,
		Ctx:         ctx,
		AccountName: "my-account",
		Region:      "my-region",
	}

	// Act
	resources, err := discoverer.Discover()

	// Assert
	assert.Error(t, err)
	assert.ErrorContains(t, err, "Could not discover ASGs")
	assert.Len(t, resources, 0)
}

func getStringPointer(str string) *string {
	return &str
}

func newASG(seed string) types.AutoScalingGroup {

	tag := types.TagDescription{
		Key:   getStringPointer("Name"),
		Value: getStringPointer("my-asg-" + seed),
	}

	return types.AutoScalingGroup{
		AutoScalingGroupName: getStringPointer("my-asg-" + seed),
		AvailabilityZones:    []string{"us-east-1a", "us-east-1b"},
		CreatedTime:          aws.Time(time.Now()),
		Tags: []types.TagDescription{
			tag,
		},
		AutoScalingGroupARN: getStringPointer("arn:aws:autoscaling:us-east-1:123456789012:autoScalingGroup:12345678-1234-1234-1234-123456789012:autoScalingGroupName/my-asg-" + seed),
	}
}
