package compute

import (
	"context"
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func monthlyPriceGross(pricingsMonthly string, fallback float64) float64 {
	if pricingsMonthly == "" {
		return fallback
	}
	v, err := strconv.ParseFloat(pricingsMonthly, 64)
	if err != nil {
		return fallback
	}
	return v
}

func serverMonthlyCost(created time.Time, monthlyGross float64) float64 {
	now := time.Now()
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	start := created
	if created.Before(monthStart) {
		start = monthStart
	}
	if start.After(now) {
		return 0
	}
	hours := now.Sub(start).Hours()
	if hours < 0 {
		return 0
	}
	nextMonth := monthStart.AddDate(0, 1, 0)
	hoursInMonth := nextMonth.Sub(monthStart).Hours()
	if hoursInMonth <= 0 {
		return monthlyGross
	}
	if hours >= hoursInMonth {
		return monthlyGross
	}
	return monthlyGross * (hours / hoursInMonth)
}

func Servers(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	servers, err := client.HetznerClient.Server.All(ctx)
	if err != nil {
		return nil, err
	}

	// /servers embeds server_type WITHOUT pricings, so join with
	// /server-types (docs.hetzner.cloud/reference/cloud#tag/server-types)
	// which carries per-location monthly Gross.
	priceByTypeID := make(map[int64]map[string]float64)
	if types, err := client.HetznerClient.ServerType.All(ctx); err == nil {
		for _, t := range types {
			m := make(map[string]float64)
			for _, p := range t.Pricings {
				if p.Location != nil {
					m[p.Location.Name] = monthlyPriceGross(p.Monthly.Gross, 0)
				}
			}
			if len(t.Pricings) > 0 && len(m) == 0 {
				m[""] = monthlyPriceGross(t.Pricings[0].Monthly.Gross, 0)
			}
			priceByTypeID[t.ID] = m
		}
	} else {
		log.WithError(err).Warn("Hetzner server-types pricing lookup failed, costs may be 0")
	}

	for _, s := range servers {
		tags := make([]models.Tag, 0, len(s.Labels))
		for k, v := range s.Labels {
			tags = append(tags, models.Tag{Key: k, Value: v})
		}

		region := ""
		if s.Location != nil {
			region = s.Location.Name
		}

		monthlyGross := 0.0
		if s.ServerType != nil {
			// Prefer joined /server-types pricings; fall back to embedded (usually empty).
			if m, ok := priceByTypeID[s.ServerType.ID]; ok {
				if v, ok := m[region]; ok {
					monthlyGross = v
				} else if v, ok := m[""]; ok {
					monthlyGross = v
				}
			}
			if monthlyGross == 0 {
				for _, p := range s.ServerType.Pricings {
					if p.Location != nil && p.Location.Name == region {
						monthlyGross = monthlyPriceGross(p.Monthly.Gross, 0)
						break
					}
				}
			}
			if monthlyGross == 0 && len(s.ServerType.Pricings) > 0 {
				monthlyGross = monthlyPriceGross(s.ServerType.Pricings[0].Monthly.Gross, 0)
			}
		}

		resources = append(resources, models.Resource{
			Provider:   "Hetzner",
			Account:    client.Name,
			Service:    "Server",
			ResourceId: fmt.Sprintf("%d", s.ID),
			Region:     region,
			Name:       s.Name,
			Cost:       serverMonthlyCost(s.Created, monthlyGross),
			Tags:       tags,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://console.hetzner.cloud/servers/%d", s.ID),
		})
	}

	log.WithFields(log.Fields{
		"provider":  "Hetzner",
		"account":   client.Name,
		"service":   "Server",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
