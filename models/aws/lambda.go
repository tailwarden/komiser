package aws

import "time"

type Lambda struct {
	Name    string
	Memory  int64
	Runtime string
}

type LambdaInvocationMetric struct {
	Region     string      `json:"region"`
	Datapoints []Datapoint `json:"datapoints"`
}

type Datapoint struct {
	Timestamp time.Time `json:"timestamp"`
	Value     float64   `json:"value"`
}
