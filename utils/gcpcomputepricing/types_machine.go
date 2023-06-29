package gcpcomputepricing

// Instance types
const (
	E2  = "e2"
	C3  = "c3"
	N2  = "n2"
	N2D = "n2d"
	T2A = "t2a"
	T2D = "t2d"
	N1  = "n1"
	C2  = "c2"
	C2D = "c2d"
	M1  = "m1"
	M2  = "m2"
	M3  = "m3"
)

// Instance commitment types
const (
	OnDemand                = "on-demand"
	Spot                    = "spot"
	Commitment1YearResource = "commitment-1-year-resource"
	Commitment3YearResource = "commitment-3-year-resource"
)

// CalculateMachineCostData then machine data required for calculate pricing.
type CalculateMachineCostData struct {
	// The type of the machine.
	MachineType string
	// Name of GCP project.
	Project string
	// Location of the machine.
	Zone string
	// Defines whether the instance is preemptible or on-demand.
	Commitment string
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string
	// Actual pricing.
	Pricing *Pricing
}
