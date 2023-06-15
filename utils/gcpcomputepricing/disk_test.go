package gcpcomputepricing

import "testing"

func TestGetPDStandardPrice(t *testing.T) {
	monthlyRate := diskGetter(t, Standard, 20)

	if monthlyRate != 880000000 {
		t.Errorf("Monthly rate should be 880000000, instead of %d", monthlyRate)
	}
}

func TestGetPDBalancedPrice(t *testing.T) {
	monthlyRate := diskGetter(t, Balanced, 50)

	if monthlyRate != 5500000000 {
		t.Errorf("Monthly rate should be 5500000000, instead of %d", monthlyRate)
	}
}

func TestGetPDSSDPrice(t *testing.T) {
	monthlyRate := diskGetter(t, SSD, 100)

	if monthlyRate != 18700000000 {
		t.Errorf("Monthly rate should be 18700000000, instead of %d", monthlyRate)
	}
}

func TestGetPDExtremePrice(t *testing.T) {
	monthlyRate := diskGetter(t, SSD, 150)

	if monthlyRate != 28050000000 {
		t.Errorf("Monthly rate should be 28050000000, instead of %d", monthlyRate)
	}
}

func diskGetter(t *testing.T, diskType string, diskSize uint64) uint64 {
	p, err := Fetch()
	if err != nil {
		t.Fatal(err)
	}

	monthlyRate, err := getDiskMonthly(
		p,
		Opts{
			Region:   "europe-north1",
			Type:     diskType,
			DiskSize: diskSize,
		},
		typeDiskGet,
	)
	if err != nil {
		t.Fatal(err)
	}

	return monthlyRate
}
