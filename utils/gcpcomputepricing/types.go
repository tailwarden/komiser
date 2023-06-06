package gcpcomputepricing

type Pricing struct {
	Gcp GCP `json:"gcp"`
}

type GCP struct {
	Compute Compute `json:"compute"`
}

type Compute struct {
	GCE              GCE              `json:"gce"`
	PersistentDisk   PersistentDisk   `json:"persistent_disk"`
	Gpus             GPUs             `json:"gpus"`
	LocalSSD         LocalSSD         `json:"local_ssd"`
	Hyperdisk        Hyperdisk        `json:"hyperdisk"`
	VMImageStorage   VMImageStorage   `json:"vm_image_storage"`
	CloudTPU         CloudTPU         `json:"cloud_tpu"`
	ComputeSolutions ComputeSolutions `json:"compute_solutions"`
}

type Subtype struct {
	Description string            `json:"description"`
	Regions     map[string]Region `json:"regions"`
}

type Region struct {
	Prices       []Price `json:"price"`
	Name         string  `json:"name"`
	InternalName string  `json:"internal_name"`
}

type Price struct {
	Val      uint64 `json:"val"`
	Currency string `json:"currency"`
	Nanos    uint64 `json:"nanos"`
}
