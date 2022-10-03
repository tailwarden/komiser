package aws

import "time"

type Cluster struct {
	Name      string
	ARN       string
	Tags      []string
	CreatedAt time.Time
}
