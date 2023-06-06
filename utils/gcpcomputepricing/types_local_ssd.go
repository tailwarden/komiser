package gcpcomputepricing

type LocalSSD struct {
	Commit1Year LocalSSDCommit1Year `json:"commit_1_year"`
	Commit3Year LocalSSDCommit3Year `json:"commit_3_year"`
	OnDemand    LocalSSDOnDemand    `json:"on_demand"`
}

type LocalSSDCommit1Year struct {
	Commitmentlocalssd1Yv1 Subtype `json:"commitmentlocalssd1yv1"`
}

type LocalSSDCommit3Year struct {
	Commitmentlocalssd3Yv1 Subtype `json:"commitmentlocalssd3yv1"`
}

type LocalSSDOnDemand struct {
	Storagelocalssd            Subtype `json:"storagelocalssd"`
	Storagepreemptiblelocalssd Subtype `json:"storagepreemptiblelocalssd"`
}
