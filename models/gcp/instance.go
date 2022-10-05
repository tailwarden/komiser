package gcp

type Instance struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Status      string   `json:"status"`
	MachineType string   `json:"machineType"`
	CPUPlatform string   `json:"cpuPlatform"`
	Public      bool     `json:"public"`
	Zone        string   `json:"zone"`
	Tags        []string `json:"tags"`
	Region      string   `json:"region"`
}
