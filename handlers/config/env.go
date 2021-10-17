package config

import (
	"log"
	"os"
	"strconv"
)

// ParseEnvironment loads a sibling `.env` file then looks through all environment
// variables to set global configuration.
func ParseEnvironment() error {

	// AZURE_GROUP_NAME and `config.GroupName()` are deprecated.
	// Use AZURE_BASE_GROUP_NAME and `config.GenerateGroupName()` instead.
	groupName = os.Getenv("AZURE_GROUP_NAME")
	baseGroupName = os.Getenv("AZURE_BASE_GROUP_NAME")

	locationDefault = os.Getenv("AZURE_LOCATION_DEFAULT")

	var err error
	useDeviceFlow, err = strconv.ParseBool(os.Getenv("AZURE_USE_DEVICEFLOW"))
	if err != nil {
		log.Printf("invalid value specified for AZURE_USE_DEVICEFLOW, disabling\n")
		useDeviceFlow = false
	}
	keepResources, err = strconv.ParseBool(os.Getenv("AZURE_SAMPLES_KEEP_RESOURCES"))
	if err != nil {
		log.Printf("invalid value specified for AZURE_SAMPLES_KEEP_RESOURCES, discarding\n")
		keepResources = false
	}

	// these must be provided by environment
	// clientID
	clientID = os.Getenv("AZURE_CLIENT_ID")

	// clientSecret
	clientSecret = os.Getenv("AZURE_CLIENT_SECRET")

	// tenantID (AAD)
	tenantID = os.Getenv("AZURE_TENANT_ID")

	// subscriptionID (ARM)
	subscriptionID = os.Getenv("AZURE_SUBSCRIPTION_ID")

	return nil
}
