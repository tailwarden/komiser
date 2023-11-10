package gcpcomputepricing

import (
	"testing"
)

func TestGet(t *testing.T) {
	pricing, err := Fetch()
	if err != nil {
		t.Fatal(err)
	}

	if pricing.Gcp.Compute.GCE.VmsOnDemand.CoresPerCore.E2.Vmimagee2Core.Regions["us-central1"].Prices[0].Nanos != 21811590 {
		t.Error("Broken")
	}
}
