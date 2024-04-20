package internal

import (
	"context"
	"time"

	"github.com/go-co-op/gocron"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func scheduleJobs(ctx context.Context, cfg *models.Config, clients []providers.ProviderClient, regions []string, port int, telemetry bool) {
	cron := gocron.NewScheduler(time.UTC)
	setupResourceFetchingJob(ctx, cron, clients, regions, telemetry)
	setupAlertCheckingJob(ctx, cron, cfg, port, telemetry)
	setupReportingJob(ctx, cron, cfg)
	cron.StartAsync()
}

func setupResourceFetchingJob(ctx context.Context, cron *gocron.Scheduler, clients []providers.ProviderClient, regions []string, telemetry bool) {
	_, err := cron.Every(1).Hours().Do(func() {
		log.Info("Fetching resources workflow has started")
		fetchResources(ctx, clients, regions, telemetry)
	})
	handleError(err, "setting up resource fetching cron job failed")
}

func setupAlertCheckingJob(ctx context.Context, cron *gocron.Scheduler, cfg *models.Config, port int, telemetry bool) {
	_, err := cron.Every(1).Hours().Do(func() {
		alerts, err := listAlerts(ctx)
		if err != nil {
			log.WithError(err).Error("failed to list alerts")
		}
		if len(alerts) > 0 {
			log.Info("Checking Alerts")
			checkingAlerts(ctx, *cfg, telemetry, port, alerts)
		}
	})
	handleError(err, "failed to setup alert checking cron job")
}

func setupReportingJob(ctx context.Context, cron *gocron.Scheduler, cfg *models.Config) {
	_, err := cron.Every(1).Friday().At("09:00").Do(func() {
		if len(cfg.Slack.Webhook) > 0 && cfg.Slack.Reporting {
			log.Info("Sending weekly reporting")
			sendTagsCoverageReport(ctx, *cfg)
			sendCostBreakdownReport(ctx, *cfg)
		}
	})
	handleError(err, "failed to setup cron job")
}
