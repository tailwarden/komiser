package gcpcomputepricing

import (
	"testing"
)

func TestCalculateMachineHourly(t *testing.T) {
	var tests = []struct {
		name   string
		inputs []interface{}
		ans    uint64
	}{
		{
			"TestGetE2standard2OnDemand",
			[]interface{}{E2, uint64(2), uint64(8)},
			67011420,
		},
		{
			"TestGetE2standard4OnDemand",
			[]interface{}{E2, uint64(4), uint64(16)},
			134022840,
		},
		{
			"TestGetE2standard8OnDemand",
			[]interface{}{E2, uint64(8), uint64(32)},
			268045680,
		},
		{
			"TestGetE2standard16OnDemand",
			[]interface{}{E2, uint64(16), uint64(64)},
			536091360,
		},
		{
			"TestGetE2standard32OnDemand",
			[]interface{}{E2, uint64(32), uint64(128)},
			1072182720,
		},
		{
			"TestGetN2standard2OnDemand",
			[]interface{}{N2, uint64(2), uint64(8)},
			97118000,
		},
				{
			"TestGetN2standard4OnDemand",
			[]interface{}{N2, uint64(4), uint64(16)},
			194236000,
		},
				{
			"TestGetN2standard8OnDemand",
			[]interface{}{N2, uint64(8), uint64(32)},
			388472000,
		},
				{
			"TestGetN2standard16OnDemand",
			[]interface{}{N2, uint64(16), uint64(64)},
			776944000,
		},
				{
			"TestGetN2standard32OnDemand",
			[]interface{}{N2, uint64(32), uint64(128)},
			1553888000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := Fetch()

			if err != nil {
				t.Fatal(err)
			}
			got, err := calculateMachineHourly(p, Opts{
				Type:        tt.inputs[0].(string),
				Commitment:  OnDemand,
				Region:      "us-west1",
				NumOfCPU:    tt.inputs[1].(uint64),
				NumOfMemory: tt.inputs[2].(uint64),
			})
			exp := tt.ans

			if err != nil {
				t.Fatal(err)
			}

			if got != exp {
				t.Errorf("Hourly rate should be %d, instead of %d", exp, got)
			}
		})

	}
}
