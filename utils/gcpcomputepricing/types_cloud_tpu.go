package gcpcomputepricing

type CloudTPU struct {
	Compute CloudTPUCompute `json:"compute"`
}

type CloudTPUCompute struct {
	TPU TPU `json:"tpu"`
}

type TPU struct {
	Tpuv2Preemptiblesecondsdefault    Subtype `json:"tpuv2preemptiblesecondsdefault"`
	Tpuv2Preemptiblesecondspoddefault Subtype `json:"tpuv2preemptiblesecondspoddefault"`
	Tpuv2Secondsdefault               Subtype `json:"tpuv2secondsdefault"`
	Tpuv2Secondspoddefault            Subtype `json:"tpuv2secondspoddefault"`
	Tpuv3Preemptiblesecondsdefault    Subtype `json:"tpuv3preemptiblesecondsdefault"`
	Tpuv3Preemptiblesecondspoddefault Subtype `json:"tpuv3preemptiblesecondspoddefault"`
	Tpuv3Secondsdefault               Subtype `json:"tpuv3secondsdefault"`
	Tpuv3Secondspoddefault            Subtype `json:"tpuv3secondspoddefault"`
	Tpuv4Preemptiblesecondspoddefault Subtype `json:"tpuv4preemptiblesecondspoddefault"`
	Tpuv4Secondspoddefault            Subtype `json:"tpuv4secondspoddefault"`
}
