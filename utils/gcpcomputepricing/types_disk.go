package gcpcomputepricing

// Persistent Disk types
const (
	// Standard persistent disks
	Standard = "pd-standard"
	// Performance (SSD) persistent disks
	SSD = "pd-ssd"
	// Balanced persistent disks
	Balanced = "pd-balanced"
	// Extreme persistent disks
	Extreme = "pd-extreme"
)

// CalculateDiskCostData the disk data required for calculate pricing.
type CalculateDiskCostData struct {
	// The type of the disk.
	DiskType string
	// The disk size in GiB.
	Size int64
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string
	// Name of GCP project.
	Project string
	// Location of the disk.
	Zone string
	// Actual pricing.
	Pricing *Pricing
}
