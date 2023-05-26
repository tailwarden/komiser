package utils

import "testing"

func TestGetRegionFromZone(t *testing.T) {
	scenarios := []struct {
		in  string
		out string
	}{
		{
			in:  "us-central1-a",
			out: "us-central1",
		},
		{
			in:  "southamerica-east1-b",
			out: "southamerica-east1",
		},
		{
			in:  "europe-central2-a",
			out: "europe-central2",
		},
		{
			in:  "me-west1-a",
			out: "me-west1",
		},
		{
			in:  "asia-south1-c",
			out: "asia-south1",
		},
	}

	for _, scenario := range scenarios {
		out := GcpGetRegionFromZone(scenario.in)
		if scenario.out != out {
			t.Errorf("Region should be %s, instead of %s", scenario.out, out)
		}
	}
}
