package gcpcomputepricing

import "testing"

func TestGetSnapshotStandardPrice(t *testing.T) {
	monthlyRate := snapshotGetter(t, 4848259584)

	if monthlyRate != 220000000 {
		t.Errorf("Monthly rate should be 220000000, instead od %d", monthlyRate)
	}
}

func snapshotGetter(t *testing.T, storageBytes uint64) uint64 {
	p, err := Fetch()
	if err != nil {
		t.Fatal(err)
	}

	monthlyRate, err := getSnapshotMonthly(
		p,
		Opts{
			Region: "europe-north1",
		},
		typeSnapshotGetter,
		storageBytes,
	)
	if err != nil {
		t.Fatal(err)
	}

	return monthlyRate
}
