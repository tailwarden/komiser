package gcpcomputepricing

type VMImageStorage struct {
	ImageStorage ImageStorage `json:"image_storage"`
}

type ImageStorage struct {
	Storageimage        Subtype `json:"storageimage"`
	Storagemachineimage Subtype `json:"storagemachineimage"`
}
