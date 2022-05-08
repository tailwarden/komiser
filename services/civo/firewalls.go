package civo

import (
	. "github.com/civo/civogo"
)

func (civo Civo) getFirewalls(apiKey, regionCode string) ([]Firewall, error) {
	client, err := getCivoClient(apiKey, regionCode)
	if err != nil {
		return nil, err
	}
	firewalls, err := client.ListFirewalls()
	if err != nil {
		return nil, err
	}
	return firewalls, nil
}

func (civo Civo) GetFirewallRulesCount(apiKey, regionCode string) (int, error) {
	firewalls, err := civo.getFirewalls(apiKey, regionCode)
	if err != nil {
		return 0, err
	}
	rulesCount := 0
	for _, firewall := range firewalls {
		rulesCount += firewall.RulesCount
	}
	return rulesCount, nil
}
