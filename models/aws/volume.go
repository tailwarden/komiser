package aws

import "time"

type Volume struct {
	ID         string
	AZ         string
	LaunchTime time.Time
	Size       int64
	State      string
	VolumeType string
	Encrypted  bool
}
