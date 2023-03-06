package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
)

type AutoScalingGroupClient interface {
	DescribeAutoScalingGroups(
		context.Context,
		*autoscaling.DescribeAutoScalingGroupsInput,
	) (*autoscaling.DescribeAutoScalingGroupsOutput, error)
}
