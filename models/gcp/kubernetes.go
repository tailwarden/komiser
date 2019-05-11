package gcp

type Kubernetes struct {
	Status      string `json:"status"`
	MachineType string `json:"machineType"`
	Nodes       int64  `json:"nodes"`
	Zone        string `json:"zone"`
}
