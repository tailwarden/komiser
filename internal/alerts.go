package internal

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
)

func checkingAlerts(ctx context.Context, cfg models.Config, telemetry bool, port int, alerts []models.Alert) {
	for _, alert := range alerts {
		var view models.View
		err := db.NewRaw(fmt.Sprintf("SELECT * FROM views WHERE id = %s", alert.ViewId)).Scan(ctx, &view)
		if err != nil {
			log.WithError(err).Error("scan failed")
		}

		stats, err := getViewStats(ctx, view.Filters)
		if err != nil {
			log.Error("Couldn't get stats for view:", view.Name)
		} else {
			if alert.Type == "BUDGET" && alert.Budget <= stats.Costs {
				if telemetry {
					analytics.TrackEvent("sending_alerts", map[string]interface{}{
						"type": "budget",
					})
				}
				if alert.IsSlack {
					log.Info("Sending Slack budget alert for view:", view.Name)
					hitSlackWebhook(view.Name, port, int(view.Id), 0, stats.Costs, cfg.Slack.Webhook, cfg.Slack.Host, alert.Type)
				} else {
					log.Info("Sending Custom Webhook budget alert for view:", view.Name)
					hitCustomWebhook(alert.Endpoint, alert.Secret, view.Name, 0, stats.Costs, alert.Type)
				}
			}
			if alert.Type == "USAGE" && alert.Usage <= stats.Resources {
				if telemetry {
					analytics.TrackEvent("sending_alerts", map[string]interface{}{
						"type": "usage",
					})
				}
				if alert.IsSlack {
					log.Info("Sending Slack usage alert for view:", view.Name)
					hitSlackWebhook(view.Name, port, int(view.Id), stats.Resources, 0, cfg.Slack.Webhook, cfg.Slack.Host, alert.Type)
				} else {
					log.Info("Sending Custom Webhook usage alert for view:", view.Name)
					hitCustomWebhook(alert.Endpoint, alert.Secret, view.Name, stats.Resources, 0, alert.Type)
				}
			}
		}
	}
}

func listAlerts(ctx context.Context) ([]models.Alert, error) {
	alerts := make([]models.Alert, 0)

	err := db.NewRaw("SELECT * FROM alerts").Scan(ctx, &alerts)
	if err != nil {
		return alerts, err
	}
	return alerts, nil
}
