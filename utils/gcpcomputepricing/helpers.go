package gcpcomputepricing

import (
	"strings"
)

// nameToTypeMatch reports whether the name string n contains any match of the string with type t.
func nameToTypeMatch(n, t string) bool {
	return strings.Contains(strings.ToLower(n), strings.ToLower(t))
}
