package utils
import (
	"testing"
)

func TestNormalizeRegionNames(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"EU_CENTRAL_1", "eu-central-1"},
		{"US_EAST1", "us-east1"},
		{"southafricawest", "southafricawest"},
	}

	for _, tt := range tests {
		got := NormalizeRegionName(tt.input)

		if got != tt.expected {
			t.Errorf("got %s, want %s", got, tt.expected)
		}
	}
}