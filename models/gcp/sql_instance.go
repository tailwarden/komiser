package gcp

type SqlInstance struct {
	State  string `json:"state"`
	Region string `json:"region"`
	Kind   string `json:"kind"`
}
