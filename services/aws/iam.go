package aws

import (
	"context"
	"fmt"
	"time"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
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

func (aws AWS) DescribeIAMRoles(cfg awsConfig.Config) (int, error) {
	svc := iam.NewFromConfig(cfg)
	result, err := svc.ListRoles(context.Background(), &iam.ListRolesInput{})
	if err != nil {
		return 0, err
	}
	return len(result.Roles), nil
}

func (aws AWS) DescribeIAMUser(cfg awsConfig.Config) (IAMUser, error) {
	svc := iam.NewFromConfig(cfg)
	result, err := svc.GetUser(context.Background(), &iam.GetUserInput{})
	if err != nil {
		return IAMUser{}, err
	}

	fmt.Println(*result.User)

	return IAMUser{
		Username:         *result.User.UserName,
		ARN:              *result.User.Arn,
		CreateDate:       *result.User.CreateDate,
		UserId:           *result.User.UserId,
		PasswordLastUsed: *result.User.PasswordLastUsed,
	}, nil
}

func (aws AWS) DescribeIAMUsers(cfg awsConfig.Config) (int, error) {
	svc := iam.NewFromConfig(cfg)
	fmt.Println(svc)
	result, err := svc.ListUsers(context.Background(), &iam.ListUsersInput{})
	if err != nil {
		return 0, err
	}
	return len(result.Users), nil
}

func (aws AWS) DescribeIAMGroups(cfg awsConfig.Config) (int, error) {
	svc := iam.NewFromConfig(cfg)
	result, err := svc.ListGroups(context.Background(), &iam.ListGroupsInput{})
	if err != nil {
		return 0, err
	}
	return len(result.Groups), nil
}

func (aws AWS) DescribeIAMPolicies(cfg awsConfig.Config) (int, error) {
	svc := iam.NewFromConfig(cfg)
	result, err := svc.ListPolicies(context.Background(), &iam.ListPoliciesInput{})
	if err != nil {
		return 0, err
	}
	return len(result.Policies), nil
}

func (aws AWS) DescribeOrganization(cfg awsConfig.Config) (Organization, error) {
	svc := organizations.NewFromConfig(cfg)
	result, err := svc.DescribeOrganization(context.Background(), &organizations.DescribeOrganizationInput{})
	if err != nil {
		return Organization{}, err
	}

	resultAccount, err := svc.DescribeAccount(context.Background(), &organizations.DescribeAccountInput{
		AccountId: result.Organization.MasterAccountId,
	})
	if err != nil {
		return Organization{}, err
	}

	organization := Organization{
		Status: string(resultAccount.Account.Status),
		Name:   *resultAccount.Account.Name,
		Email:  *resultAccount.Account.Email,
	}

	return organization, nil
}
