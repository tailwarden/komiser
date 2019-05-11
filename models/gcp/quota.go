package gcp

type Quota struct {
	Metric string  `json:"metric"`
	Limit  float64 `json:"limit"`
	Usage  float64 `json:"usage"`
}
