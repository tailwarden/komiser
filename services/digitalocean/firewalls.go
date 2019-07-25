package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"
)

type FirewallRule struct {
	Protocol string `json:"protocol"`
	Port     string `json:"port"`
	ID       string `json:"id"`
	Status   string `json:"status"`
}

func (dg DigitalOcean) DescribeFirewalls(client *godo.Client) (int, error) {
	firewalls, _, err := client.Firewalls.List(context.TODO(), &godo.ListOptions{})
	if err != nil {
		return 0, err
	}
	return len(firewalls), nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (dg DigitalOcean) DescribeUnsecureFirewalls(client *godo.Client) ([]FirewallRule, error) {
	rules := make([]FirewallRule, 0)

	firewalls, _, err := client.Firewalls.List(context.TODO(), &godo.ListOptions{})
	if err != nil {
		return rules, err
	}

	for _, firewall := range firewalls {
		for _, rule := range firewall.InboundRules {
			if contains(rule.Sources.Addresses, "0.0.0.0/0") {
				rules = append(rules, FirewallRule{
					Protocol: rule.Protocol,
					Port:     rule.PortRange,
					ID:       firewall.ID,
					Status:   firewall.Status,
				})
			}
		}
	}

	return rules, nil
}
