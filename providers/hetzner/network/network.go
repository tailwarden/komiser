package network

import (
	"context"
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func prorated(created time.Time, monthlyGross float64) float64 {
	now := time.Now()
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	start := created
	if start.Before(monthStart) {
		start = monthStart
	}
	if !now.After(start) {
		return 0
	}
	nextMonth := monthStart.AddDate(0, 1, 0)
	hoursInMonth := nextMonth.Sub(monthStart).Hours()
	if hoursInMonth <= 0 {
		return monthlyGross
	}
	hours := now.Sub(start).Hours()
	if hours >= hoursInMonth {
		return monthlyGross
	}
	return monthlyGross * (hours / hoursInMonth)
}

func LoadBalancers(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	lbs, err := client.HetznerClient.LoadBalancer.All(ctx)
	if err != nil {
		return nil, err
	}

	// Join with /load-balancer-types for per-location monthly Gross.
	priceByLBTypeID := make(map[int64]map[string]float64)
	if types, err := client.HetznerClient.LoadBalancerType.All(ctx); err == nil {
		for _, t := range types {
			m := make(map[string]float64)
			for _, p := range t.Pricings {
				if p.Location != nil {
					if v, err := strconv.ParseFloat(p.Monthly.Gross, 64); err == nil {
						m[p.Location.Name] = v
					}
				}
			}
			if len(t.Pricings) > 0 && len(m) == 0 {
				if v, err := strconv.ParseFloat(t.Pricings[0].Monthly.Gross, 64); err == nil {
					m[""] = v
				}
			}
			priceByLBTypeID[t.ID] = m
		}
	} else {
		log.WithError(err).Warn("Hetzner load-balancer-types pricing lookup failed, costs may be 0")
	}

	for _, lb := range lbs {
		tags := make([]models.Tag, 0, len(lb.Labels))
		for k, v := range lb.Labels {
			tags = append(tags, models.Tag{Key: k, Value: v})
		}

		region := ""
		if lb.Location != nil {
			region = lb.Location.Name
		}

		monthlyGross := 0.0
		if lb.LoadBalancerType != nil {
			if m, ok := priceByLBTypeID[lb.LoadBalancerType.ID]; ok {
				if v, ok := m[region]; ok {
					monthlyGross = v
				} else if v, ok := m[""]; ok {
					monthlyGross = v
				}
			}
			if monthlyGross == 0 {
				for _, p := range lb.LoadBalancerType.Pricings {
					if p.Location != nil && p.Location.Name == region {
						if v, err := strconv.ParseFloat(p.Monthly.Gross, 64); err == nil {
							monthlyGross = v
						}
						break
					}
				}
			}
			if monthlyGross == 0 && len(lb.LoadBalancerType.Pricings) > 0 {
				if v, err := strconv.ParseFloat(lb.LoadBalancerType.Pricings[0].Monthly.Gross, 64); err == nil {
					monthlyGross = v
				}
			}
		}

		resources = append(resources, models.Resource{
			Provider:   "Hetzner",
			Account:    client.Name,
			Service:    "Load Balancer",
			ResourceId: fmt.Sprintf("%d", lb.ID),
			Region:     region,
			Name:       lb.Name,
			Cost:       prorated(lb.Created, monthlyGross),
			Tags:       tags,
			FetchedAt:  time.Now(),
			Link:       "https://console.hetzner.cloud/load-balancers",
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Hetzner",
		"account":   client.Name,
		"service":   "Load Balancer",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}

func Firewalls(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	firewalls, err := client.HetznerClient.Firewall.All(ctx)
	if err != nil {
		return nil, err
	}

	for _, fw := range firewalls {
		tags := make([]models.Tag, 0, len(fw.Labels))
		for k, v := range fw.Labels {
			tags = append(tags, models.Tag{Key: k, Value: v})
		}

		resources = append(resources, models.Resource{
			Provider:   "Hetzner",
			Account:    client.Name,
			Service:    "Firewall",
			ResourceId: fmt.Sprintf("%d", fw.ID),
			Region:     "",
			Name:       fw.Name,
			Cost:       0,
			Tags:       tags,
			FetchedAt:  time.Now(),
			Link:       "https://console.hetzner.cloud/firewalls",
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Hetzner",
		"account":   client.Name,
		"service":   "Firewall",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}

func FloatingIPs(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	ips, err := client.HetznerClient.FloatingIP.All(ctx)
	if err != nil {
		return nil, err
	}

	for _, ip := range ips {
		tags := make([]models.Tag, 0, len(ip.Labels))
		for k, v := range ip.Labels {
			tags = append(tags, models.Tag{Key: k, Value: v})
		}

		region := ""
		if ip.HomeLocation != nil {
			region = ip.HomeLocation.Name
		}

		name := ip.Name
		if name == "" {
			name = ip.IP.String()
		}

		resources = append(resources, models.Resource{
			Provider:   "Hetzner",
			Account:    client.Name,
			Service:    "Floating IP",
			ResourceId: fmt.Sprintf("%d", ip.ID),
			Region:     region,
			Name:       name,
			Cost:       0,
			Tags:       tags,
			FetchedAt:  time.Now(),
			Link:       "https://console.hetzner.cloud/floating-ips",
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Hetzner",
		"account":   client.Name,
		"service":   "Floating IP",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}

func Networks(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	networks, err := client.HetznerClient.Network.All(ctx)
	if err != nil {
		return nil, err
	}

	for _, n := range networks {
		tags := make([]models.Tag, 0, len(n.Labels))
		for k, v := range n.Labels {
			tags = append(tags, models.Tag{Key: k, Value: v})
		}

		resources = append(resources, models.Resource{
			Provider:   "Hetzner",
			Account:    client.Name,
			Service:    "Network",
			ResourceId: fmt.Sprintf("%d", n.ID),
			Region:     "",
			Name:       n.Name,
			Cost:       0,
			Tags:       tags,
			FetchedAt:  time.Now(),
			Link:       "https://console.hetzner.cloud/networks",
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Hetzner",
		"account":   client.Name,
		"service":   "Network",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
