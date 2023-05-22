package gcpcomputepricing

type GPUs struct {
	GpusCommit1Year GpusCommit1Year `json:"gpus_commit_1_year"`
	GpusCommit3Year GpusCommit3Year `json:"gpus_commit_3_year"`
	GpusOnDemand    GpusOnDemand    `json:"gpus_on_demand"`
	GpusPreemptible GpusPreemptible `json:"gpus_preemptible"`
	GpusReservation GpusReservation `json:"gpus_reservation"`
}

type GpusCommit1Year struct {
	L4   GpusCommit1YearL4   `json:"l4"`
	A100 GpusCommit1YearA100 `json:"a100"`
	K80  GpusCommit1YearK80  `json:"k80"`
	P100 GpusCommit1YearP100 `json:"p100"`
	P4   GpusCommit1YearP4   `json:"p4"`
	T4   GpusCommit1YearT4   `json:"t4"`
	V100 GpusCommit1YearV100 `json:"v100"`
}

type GpusCommit1YearL4 struct {
	Commitmentgpunvidial41Yv1 Subtype `json:"commitmentgpunvidial41yv1"`
}

type GpusCommit1YearA100 struct {
	Commitmentgpunvidiateslaa1001Yv1 Subtype `json:"commitmentgpunvidiateslaa1001yv1"`
}

type GpusCommit1YearK80 struct {
	Commitmentgpunvidiateslak801Yv1 Subtype `json:"commitmentgpunvidiateslak801yv1"`
}

type GpusCommit1YearP100 struct {
	Commitmentgpunvidiateslap1001Yv1 Subtype `json:"commitmentgpunvidiateslap1001yv1"`
}

type GpusCommit1YearP4 struct {
	Commitmentgpunvidiateslap41Yv1 Subtype `json:"commitmentgpunvidiateslap41yv1"`
}

type GpusCommit1YearT4 struct {
	Commitmentgpunvidiateslat41Yv1 Subtype `json:"commitmentgpunvidiateslat41yv1"`
}

type GpusCommit1YearV100 struct {
	Commitmentgpunvidiateslav1001Yv1 Subtype `json:"commitmentgpunvidiateslav1001yv1"`
}

type GpusCommit3Year struct {
	L4   GpusCommit3YearL4   `json:"l4"`
	A100 GpusCommit3YearA100 `json:"a100"`
	K80  GpusCommit3YearK80  `json:"k80"`
	P100 GpusCommit3YearP100 `json:"p100"`
	P4   GpusCommit3YearP4   `json:"p4"`
	T4   GpusCommit3YearT4   `json:"t4"`
	V100 GpusCommit3YearV100 `json:"v100"`
}

type GpusCommit3YearL4 struct {
	Commitmentgpunvidial43Yv1 Subtype `json:"commitmentgpunvidial43yv1"`
}

type GpusCommit3YearA100 struct {
	Commitmentgpunvidiateslaa1003Yv1 Subtype `json:"commitmentgpunvidiateslaa1003yv1"`
}

type GpusCommit3YearK80 struct {
	Commitmentgpunvidiateslak803Yv1 Subtype `json:"commitmentgpunvidiateslak803yv1"`
}

type GpusCommit3YearP100 struct {
	Commitmentgpunvidiateslap1003Yv1 Subtype `json:"commitmentgpunvidiateslap1003yv1"`
}

type GpusCommit3YearP4 struct {
	Commitmentgpunvidiateslap43Yv1 Subtype `json:"commitmentgpunvidiateslap43yv1"`
}

type GpusCommit3YearT4 struct {
	Commitmentgpunvidiateslat43Yv1 Subtype `json:"commitmentgpunvidiateslat43yv1"`
}

type GpusCommit3YearV100 struct {
	Commitmentgpunvidiateslav1003Yv1 Subtype `json:"commitmentgpunvidiateslav1003yv1"`
}

type GpusOnDemand struct {
	L4   GpusOnDemandL4   `json:"l4"`
	A100 GpusOnDemandA100 `json:"a100"`
	K80  GpusOnDemandK80  `json:"k80"`
	P100 GpusOnDemandP100 `json:"p100"`
	P4   GpusOnDemandP4   `json:"p4"`
	T4   GpusOnDemandT4   `json:"t4"`
	V100 GpusOnDemandV100 `json:"v100"`
}

type GpusOnDemandL4 struct {
	Gpunvidial4 Subtype `json:"gpunvidial4"`
}

type GpusOnDemandA100 struct {
	Gpunvidiateslaa100 Subtype `json:"gpunvidiateslaa100"`
}

type GpusOnDemandK80 struct {
	Gpunvidiateslak80 Subtype `json:"gpunvidiateslak80"`
}

type GpusOnDemandP100 struct {
	Gpunvidiateslap100 Subtype `json:"gpunvidiateslap100"`
}

type GpusOnDemandP4 struct {
	Gpunvidiateslap4 Subtype `json:"gpunvidiateslap4"`
}

type GpusOnDemandT4 struct {
	Gpunvidiateslat4 Subtype `json:"gpunvidiateslat4"`
}

type GpusOnDemandV100 struct {
	Gpunvidiateslav100 Subtype `json:"gpunvidiateslav100"`
}

type GpusPreemptible struct {
	L4   GpusPreemptibleL4   `json:"l4"`
	A100 GpusPreemptibleA100 `json:"a100"`
	K80  GpusPreemptibleK80  `json:"k80"`
	P100 GpusPreemptibleP100 `json:"p100"`
	P4   GpusPreemptibleP4   `json:"p4"`
	T4   GpusPreemptibleT4   `json:"t4"`
	V100 GpusPreemptibleV100 `json:"v100"`
}

type GpusPreemptibleL4 struct {
	Gpupreemptiblenvidial4 Subtype `json:"gpupreemptiblenvidial4"`
}

type GpusPreemptibleA100 struct {
	Gpupreemptiblenvidiateslaa100 Subtype `json:"gpupreemptiblenvidiateslaa100"`
}

type GpusPreemptibleK80 struct {
	Gpupreemptiblenvidiateslak80 Subtype `json:"gpupreemptiblenvidiateslak80"`
}

type GpusPreemptibleP100 struct {
	Gpupreemptiblenvidiateslap100 Subtype `json:"gpupreemptiblenvidiateslap100"`
}

type GpusPreemptibleP4 struct {
	Gpupreemptiblenvidiateslap4 Subtype `json:"gpupreemptiblenvidiateslap4"`
}

type GpusPreemptibleT4 struct {
	Gpupreemptiblenvidiateslat4 Subtype `json:"gpupreemptiblenvidiateslat4"`
}

type GpusPreemptibleV100 struct {
	Gpupreemptiblenvidiateslav100 Subtype `json:"gpupreemptiblenvidiateslav100"`
}

type GpusReservation struct {
	A100 GpusReservationA100 `json:"a100"`
}

type GpusReservationA100 struct {
	Reservationgpunvidiateslaa100 Subtype `json:"reservationgpunvidiateslaa100"`
}
