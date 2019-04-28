package aws

import "time"

// EC2 Instance
type EC2 struct {
	ID           string    `json:"id"`
	InstanceType string    `json:"instanceType"`
	LaunchTime   time.Time `json:"launchTime"`
	State        string    `json:"state"`
	Tags         []string  `json:"tags"`
	Public       bool      `json:"public"`
}
