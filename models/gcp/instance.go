package gcp

type Instance struct {
	Name        string `json:"name"`
	Status      string `json:"status"`
	MachineType string `json:"machineType"`
	CPUPlatform string `json:"cpuPlatform"`
	Public      bool   `json:"public"`
	Zone        string `json:"zone"`
}
