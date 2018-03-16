package services

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func (aws AWS) DescribeIAMRolesTotal(cfg aws.Config) (int, error) {
	svc := iam.New(cfg)
	req := svc.ListRolesRequest(&iam.ListRolesInput{})
	result, err := req.Send()
	if err != nil {
		return 0, err
	}
	return len(result.Roles), nil
}

func (aws AWS) DescribeIAMUsersTotal(cfg aws.Config) (int, error) {
	svc := iam.New(cfg)
	req := svc.ListUsersRequest(&iam.ListUsersInput{})
	result, err := req.Send()
	if err != nil {
		return 0, err
	}
	return len(result.Users), nil
}

func (aws AWS) DescribeIAMGroupsTotal(cfg aws.Config) (int, error) {
	svc := iam.New(cfg)
	req := svc.ListGroupsRequest(&iam.ListGroupsInput{})
	result, err := req.Send()
	if err != nil {
		return 0, err
	}
	return len(result.Groups), nil
}

func (aws AWS) DescribeIAMPoliciesTotal(cfg aws.Config) (int, error) {
	svc := iam.New(cfg)
	req := svc.ListPoliciesRequest(&iam.ListPoliciesInput{})
	result, err := req.Send()
	if err != nil {
		return 0, err
	}
	return len(result.Policies), nil
}
