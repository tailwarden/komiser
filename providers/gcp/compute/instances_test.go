package compute

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/tailwarden/komiser/providers"
	"golang.org/x/oauth2/google"
)

func TestInstances(t *testing.T) {
	data, err := ioutil.ReadFile("")
	if err != nil {
		t.Fatal(err)
	}

	creds, err := google.CredentialsFromJSON(context.Background(), data, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		t.Fatal(err)
	}
	resource, err := Instances(context.Background(), providers.ProviderClient{
		GCPClient: &providers.GCPClient{Credentials: creds},
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resource)
}

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
		out := getRegionFromZone(scenario.in)
		if scenario.out != out {
			t.Errorf("Region should be %s, instead of %s", scenario.out, out)
		}
	}
}
