package aws

import "time"

type Task struct {
	ARN       string
	Tags      []string
	CreatedAt time.Time
	Region    string
}
