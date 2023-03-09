package clusters

import (
	"testing"
)

func TestNormalizeRegionNames(t *testing.T) {
	got := normalizeRegionName("EU_CENTRAL_1")
	if got != "eu-central-1" {
		t.Errorf("NormalizeRegionName('EU_CENTRAL_1') = %s; want 'eu-central-1'", got)
	}

	got = normalizeRegionName("US_EAST1")
	if got != "us-east1" {
		t.Errorf("NormalizeRegionName('US_EAST1') = %s; want 'us-east1'", got)
	}
}
