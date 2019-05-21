package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
)

type IAMUser struct {
	Username         string    `json:"username"`
	ARN              string    `json:"arn"`
	CreateDate       time.Time `json:"createDate"`
	PasswordLastUsed time.Time `json:"passwordLastUsed"`
	UserId           string    `json:"userId"`
}

type Organization struct {
	Status string `json:"status"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

func (aws AWS) DescribeIAMRoles(cfg aws.Config) (int, error) {
	svc := iam.New(cfg)
	req := svc.ListRolesRequest(&iam.ListRolesInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return 0, err
	}
	return len(result.Roles), nil
}

func (aws AWS) DescribeIAMUser(cfg aws.Config) (IAMUser, error) {
	svc := iam.New(cfg)
	req := svc.GetUserRequest(&iam.GetUserInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return IAMUser{}, err
	}

	lastUsed := time.Now()
	if result.User.PasswordLastUsed != nil {
		lastUsed = *result.User.PasswordLastUsed
	}

	return IAMUser{
		Username:         *result.User.UserName,
		ARN:              *result.User.Arn,
		CreateDate:       *result.User.CreateDate,
		UserId:           *result.User.UserId,
		PasswordLastUsed: lastUsed,
	}, nil
}

func (aws AWS) DescribeIAMUsers(cfg aws.Config) (int, error) {
	svc := iam.New(cfg)
	req := svc.ListUsersRequest(&iam.ListUsersInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return 0, err
	}
	return len(result.Users), nil
}

func (aws AWS) DescribeIAMGroups(cfg aws.Config) (int, error) {
	svc := iam.New(cfg)
	req := svc.ListGroupsRequest(&iam.ListGroupsInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return 0, err
	}
	return len(result.Groups), nil
}

func (aws AWS) DescribeIAMPolicies(cfg aws.Config) (int, error) {
	svc := iam.New(cfg)
	req := svc.ListPoliciesRequest(&iam.ListPoliciesInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return 0, err
	}
	return len(result.Policies), nil
}

func (aws AWS) DescribeOrganization(cfg aws.Config) (Organization, error) {
	svc := organizations.New(cfg)
	req := svc.DescribeOrganizationRequest(&organizations.DescribeOrganizationInput{})
	result, err := req.Send(context.Background())
	if err != nil {
		return Organization{}, err
	}

	reqAccount := svc.DescribeAccountRequest(&organizations.DescribeAccountInput{
		AccountId: result.Organization.MasterAccountId,
	})
	resultAccount, err := reqAccount.Send(context.Background())
	if err != nil {
		return Organization{}, err
	}

	status, _ := resultAccount.Account.Status.MarshalValue()
	organization := Organization{
		Status: status,
		Name:   *resultAccount.Account.Name,
		Email:  *resultAccount.Account.Email,
	}

	return organization, nil
}
