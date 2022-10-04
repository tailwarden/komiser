package aws

import "time"

type Lambda struct {
	Name        string   `json:"name"`
	Memory      int64    `json:"memory"`
	Runtime     string   `json:"runtime"`
	FunctionArn string   `json:"arn"`
	Tags        []string `json:"tags"`
	Region      string   `json:"region"`
}

type LambdaInvocationMetric struct {
	Region     string      `json:"region"`
	Datapoints []Datapoint `json:"datapoints"`
}

type Datapoint struct {
	Timestamp time.Time `json:"timestamp"`
	Value     float64   `json:"value"`
}
