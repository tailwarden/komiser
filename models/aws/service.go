package aws

import "time"

type Service struct {
	Name      string
	CreatedAt time.Time
	Tags      []string
	Region    string
}
