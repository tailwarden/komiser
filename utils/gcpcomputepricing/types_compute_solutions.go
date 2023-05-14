package gcpcomputepricing

type ComputeSolutions struct {
	GoogleCloudVmwareEngine GoogleCloudVmwareEngine `json:"google_cloud_vmware_engine"`
}

type GoogleCloudVmwareEngine struct {
	Standard72Hourlyv2  Subtype `json:"standard72hourlyv2"`
	Standard72Hourly    Subtype `json:"standard72hourly"`
	Vmwareengineucs12Mo Subtype `json:"vmwareengineucs12mo"`
	Vmwareengineucs36Mo Subtype `json:"vmwareengineucs36mo"`
}
