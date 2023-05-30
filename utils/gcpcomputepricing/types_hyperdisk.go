package gcpcomputepricing

type Hyperdisk struct {
	HyperdiskVolumes HyperdiskVolumes `json:"hyperdisk_volumes"`
}

type HyperdiskVolumes struct {
	Extreme HyperdiskVolumesExtreme `json:"extreme"`
}

type HyperdiskVolumesExtreme struct {
	Capacity HyperdiskVolumesExtremeCapacity `json:"capacity"`
	IOps     HyperdiskVolumesExtremeIOps     `json:"iops"`
}

type HyperdiskVolumesExtremeCapacity struct {
	Storagehyperdiskextremecapacity Subtype `json:"storagehyperdiskextremecapacity"`
}

type HyperdiskVolumesExtremeIOps struct {
	Storagehyperdiskextremeiops Subtype `json:"storagehyperdiskextremeiops"`
}
