package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
)

type AutoScalingGroupClient interface {
	DescribeAutoScalingGroups(
		ctx context.Context,
		params *autoscaling.DescribeAutoScalingGroupsInput,
		optFns ...func(*autoscaling.Options),
	) (*autoscaling.DescribeAutoScalingGroupsOutput, error)
}
