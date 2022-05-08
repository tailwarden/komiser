package config

import "os"

// ParseEnvironment loads a sibling `.env` file then looks through all environment
// variables to set global configuration.
func ParseEnvironment() error {
	// Retrieves the CIVO Api Key to be used for accessing Civo go lang APIs
	apiKey = os.Getenv("CIVO_API_KEY")

	//Retrieves the CIVO Region Code to be used for accessing Civo go lang APIs
	regionCode = os.Getenv("CIVO_REGION_CODE")

	return nil

}
