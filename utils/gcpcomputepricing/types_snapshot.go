package gcpcomputepricing

// CalculateSnapshotCostData the snapshot data required for calculate pricing.
type CalculateSnapshotCostData struct {
	// A size of the storage used by the snapshot.
	// As snapshots share storage, this number is expected to change with snapshot creation/deletion.
	StorageBytes int64
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string
	// Location of the snapshot
	Region string
	// Name of GCP project.
	Project string
	// Actual pricing.
	Pricing *Pricing
}
