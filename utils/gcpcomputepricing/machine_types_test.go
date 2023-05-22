package gcpcomputepricing

import (
	"testing"
)

func TestGetE2standard2OnDemand(t *testing.T) {
	hourlyRate := getter(t, 2, 8)

	if hourlyRate != 67011420 {
		t.Errorf("Hourly rate should be 67011420, instead of %d", hourlyRate)
	}
}

func TestGetE2standard4OnDemand(t *testing.T) {
	hourlyRate := getter(t, 4, 16)

	if hourlyRate != 134022840 {
		t.Errorf("Hourly rate should be 134022840, instead of %d", hourlyRate)
	}
}

func TestGetE2standard8OnDemand(t *testing.T) {
	hourlyRate := getter(t, 8, 32)

	if hourlyRate != 268045680 {
		t.Errorf("Hourly rate should be 268045680, instead of %d", hourlyRate)
	}
}

func TestGetE2standard16OnDemand(t *testing.T) {
	hourlyRate := getter(t, 16, 64)

	if hourlyRate != 536091360 {
		t.Errorf("Hourly rate should be 536091360, instead of %d", hourlyRate)
	}
}

func TestGetE2standard32OnDemand(t *testing.T) {
	hourlyRate := getter(t, 32, 128)

	if hourlyRate != 1072182720 {
		t.Errorf("Hourly rate should be 1072182720, instead of %d", hourlyRate)
	}
}

func getter(t *testing.T, cpu, memory uint64) uint64 {
	p, err := Fetch()
	if err != nil {
		t.Fatal(err)
	}

	hourlyRate, err := CalculateMachine(p, Opts{
		Type:        E2,
		Commitment:  OnDemand,
		Region:      "us-west1",
		NumOfCPU:    cpu,
		NumOfMemory: memory,
	})
	if err != nil {
		t.Fatal(err)
	}

	return hourlyRate
}
