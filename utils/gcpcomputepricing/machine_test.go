package gcpcomputepricing

import (
	"testing"
)

func TestCalculateMachineHourly(t *testing.T) {
	var tests = []struct {
		name   string
		inputs []interface{}
	}{
		{
			"TestGetE2standard2OnDemand",
			[]interface{}{E2, 2, 8},
		},
		{
			"TestGetE2standard4OnDemand",
			[]interface{}{E2, 4, 16},
		},
		{
			"TestGetE2standard8OnDemand",
			[]interface{}{E2, 8, 32},
		},
		{
			"TestGetE2standard16OnDemand",
			[]interface{}{E2, 16, 64},
		},
		{
			"TestGetE2standard32OnDemand",
			[]interface{}{E2, 32, 128},
		},
		{
			"TestGetC3standard2OnDemand",
			[]interface{}{C3, 2, 8},
		},
				{
			"TestGetC3standard4OnDemand",
			[]interface{}{C3, 4, 16},
		},
				{
			"TestGetC3standard8OnDemand",
			[]interface{}{C3, 8, 32},
		},
				{
			"TestGetC3standard16OnDemand",
			[]interface{}{C3, 16, 64},
		},
				{
			"TestGetC3standard32OnDemand",
			[]interface{}{C3, 32, 128},
		},
		{
			"TestGetN2standard2OnDemand",
			[]interface{}{N2, 2, 8},
		},
				{
			"TestGetN2standard4OnDemand",
			[]interface{}{N2, 4, 16},
		},
				{
			"TestGetN2standard8OnDemand",
			[]interface{}{N2, 8, 32},
		},
				{
			"TestGetN2standard16OnDemand",
			[]interface{}{N2, 16, 64},
		},
				{
			"TestGetN2standard32OnDemand",
			[]interface{}{N2, 32, 128},
		},
		{
			"TestGetN2Dstandard2OnDemand",
			[]interface{}{N2D, 2, 8},
		},
				{
			"TestGetN2Dstandard4OnDemand",
			[]interface{}{N2D, 4, 16},
		},
				{
			"TestGetN2Dstandard8OnDemand",
			[]interface{}{N2D, 8, 32},
		},
				{
			"TestGetN2Dstandard16OnDemand",
			[]interface{}{N2D, 16, 64},
		},
				{
			"TestGetN2Dstandard32OnDemand",
			[]interface{}{N2D, 32, 128},
		},
		{
			"TestGetT2Astandard2OnDemand",
			[]interface{}{T2A, 2, 8},
		},
				{
			"TestGetT2Astandard4OnDemand",
			[]interface{}{T2A, 4, 16},
		},
				{
			"TestGetT2Astandard8OnDemand",
			[]interface{}{T2A, 8, 32},
		},
		{
			"TestGetT2Astandard16OnDemand",
			[]interface{}{T2A, 16, 64},
		},
				{
			"TestGetT2Astandard32OnDemand",
			[]interface{}{T2A, 32, 128},
		},
		{
			"TestGetT2Dstandard2OnDemand",
			[]interface{}{T2D, 2, 8},
		},
				{
			"TestGetT2Dstandard4OnDemand",
			[]interface{}{T2D, 4, 16},
		},
				{
			"TestGetT2Dstandard8OnDemand",
			[]interface{}{T2D, 8, 32},
		},
				{
			"TestGetT2Dstandard16OnDemand",
			[]interface{}{T2D, 16, 64},
		},
				{
			"TestGetT2Dstandard32OnDemand",
			[]interface{}{T2D, 32, 128},
		},
		{
			"TestGetN1standard2OnDemand",
			[]interface{}{N1, 2, 8},
		},
				{
			"TestGetN1standard4OnDemand",
			[]interface{}{N1, 4, 16},
		},
				{
			"TestGetN1standard8OnDemand",
			[]interface{}{N1, 8, 32},
		},
				{
			"TestGetN1standard16OnDemand",
			[]interface{}{N1, 16, 64},
		},
				{
			"TestGetN1standard32OnDemand",
			[]interface{}{N1, 32, 128},
		},
		{
			"TestGetC2Dstandard2OnDemand",
			[]interface{}{C2D, 2, 8},
		},
				{
			"TestGetC2Dstandard4OnDemand",
			[]interface{}{C2D, 4, 16},
		},
				{
			"TestGetC2Dstandard8OnDemand",
			[]interface{}{C2D, 8, 32},
		},
				{
			"TestGetC2Dstandard16OnDemand",
			[]interface{}{C2D, 16, 64},
		},
				{
			"TestGetC2Dstandard32OnDemand",
			[]interface{}{C2D, 32, 128},
		},
		{
			"TestGetM3standard2OnDemand",
			[]interface{}{N2D, 2, 8},
		},
				{
			"TestGetM3standard4OnDemand",
			[]interface{}{N2D, 4, 16},
		},
				{
			"TestGetM3standard8OnDemand",
			[]interface{}{N2D, 8, 32},
		},
				{
			"TestGetM3standard16OnDemand",
			[]interface{}{N2D, 16, 64},
		},
				{
			"TestGetM3standard32OnDemand",
			[]interface{}{N2D, 32, 128},
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
				NumOfCPU:    uint64(tt.inputs[1].(int)),
				NumOfMemory: uint64(tt.inputs[2].(int)),
			})

			if err != nil {
				t.Fatal(err)
			}

			if got <= 0 {
				t.Errorf("Hourly rate should be greater than 0, but is %d", got)
			}
		})

	}
}
