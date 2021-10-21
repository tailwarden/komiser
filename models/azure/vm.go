package azure

type Vm struct {
	Name   string `json:"name"`
	Image  string `json:"image"`
	Region string `json:"region"`
	Disk   int32  `json:"disk"`
}
