package config

var (
	// these are our *global* config settings, to be shared by all packages.
	// each has corresponding public accessors below.
	// if anything requires a `Set` accessor, that indicates it perhaps
	// shouldn't be set here, because mutable vars shouldn't be global.
	apiKey     string
	regionCode string
)

// ApiKey is the key used for CIVO authentication
func ApiKey() string {
	return apiKey
}

// Region code used for CIVO client
func RegionCode() string {
	return regionCode
}
