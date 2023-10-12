package compute

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/tailwarden/komiser/providers"
	"golang.org/x/oauth2/google"
)

func TestInstances(t *testing.T) {
	t.Skip("Only for local development because it is using a Google Cloud connection")
	// Replace the empty string with a SA or credentials file location
	data, err := os.ReadFile("")
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
