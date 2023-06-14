package gcpcomputepricing

import "testing"

// Nota bene:
// The monthly price per disc consists of the price per size (capacity) and
// the price per snapshot.

func TestGetPDStandard(t *testing.T) {
	monthlyRate := diskGetter(t, Standard, 20)

	// 1980000000 = 880000000 (capacity) + 1100000000 (snapshot)
	if monthlyRate != 1980000000 {
		t.Errorf("Monthly rate should be 1980000000, instead of %d", monthlyRate)
	}
}

func TestGetPDBalanced(t *testing.T) {
	monthlyRate := diskGetter(t, Balanced, 50)

	if monthlyRate != 5950000000 {
		t.Errorf("Monthly rate should be 5950000000, instead of %d", monthlyRate)
	}
}

func TestGetPDSSD(t *testing.T) {
	monthlyRate := diskGetter(t, SSD, 100)

	if monthlyRate != 18700000000 {
		t.Errorf("Monthly rate should be 18700000000, instead of %d", monthlyRate)
	}
}

func diskGetter(t *testing.T, diskType string, diskSize uint64) uint64 {
	p, err := Fetch()
	if err != nil {
		t.Fatal(err)
	}

	monthlyRate, err := CalculateDisk(p, Opts{
		Region:   "europe-north1",
		DiskType: diskType,
		DiskSize: diskSize,
	})
	if err != nil {
		t.Fatal(err)
	}

	return monthlyRate
}
