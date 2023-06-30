package gcpcomputepricing

// Opts holds external data for cost calculations
// into GCP Compute price package.
type Opts struct {
	Type        string
	Commitment  string
	Region      string
	NumOfCPU    uint64
	NumOfMemory uint64
	DiskSize    uint64
}
