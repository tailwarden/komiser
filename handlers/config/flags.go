package config

import (
	"flag"
)

// AddFlags adds flags applicable to all services.
// Remember to call `flag.Parse()` in your main or TestMain.
func AddFlags() error {
	flag.StringVar(&subscriptionID, "subscription", subscriptionID, "Subscription for tests.")
	flag.StringVar(&locationDefault, "location", locationDefault, "Default location for tests.")
	flag.StringVar(&cloudName, "cloud", cloudName, "Name of Azure cloud.")
	flag.StringVar(&baseGroupName, "baseGroupName", BaseGroupName(), "Specify prefix name of resource group for sample resources.")

	flag.BoolVar(&useDeviceFlow, "useDeviceFlow", useDeviceFlow, "Use device-flow grant type rather than client credentials.")
	flag.BoolVar(&keepResources, "keepResources", keepResources, "Keep resources created by samples.")

	return nil
}
