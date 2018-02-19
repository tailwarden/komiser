package models

type VPC struct {
	ID        string
	State     string
	CidrBlock string
	Tags      []string
}
