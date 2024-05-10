package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
	"github.com/tailwarden/komiser/models"
)

func sendTagsCoverageReport(ctx context.Context, cfg models.Config) {
	tags := make([]struct {
		Total int        `bun:"total"`
		Label models.Tag `bun:"label"`
	}, 0)

	err := db.NewRaw("SELECT count(*) as total, value as label FROM resources CROSS JOIN json_each(tags) GROUP BY value ORDER BY total DESC").Scan(ctx, &tags)
	if err != nil {
		log.WithError(err).Error("scan failed")
	}

	fields := make([]slack.AttachmentField, 0)

	for _, tag := range tags {
		fields = append(fields, slack.AttachmentField{
			Title: fmt.Sprintf("%s:%s", tag.Label.Key, tag.Label.Value),
			Value: fmt.Sprintf("%d", tag.Total),
			Short: true,
		})
	}

	output := struct {
		Total int `bun:"total"`
	}{}

	err = db.NewRaw("SELECT COUNT(*) as total FROM resources where json_array_length(tags) = 0;").Scan(ctx, &output)
	if err != nil {
		log.WithError(err).Error("scan failed")
	}

	currentTime := time.Now()

	attachment := slack.Attachment{
		Color:         "good",
		AuthorName:    "Komiser",
		AuthorSubname: "by Tailwarden",
		AuthorLink:    "https://tailwarden.com",
		AuthorIcon:    "https://cdn.komiser.io/images/komiser-logo.jpeg",
		Text:          fmt.Sprintf("On %s %d: *%d* of your resources are untagged. Below list of most used key/value pairs:", currentTime.Month(), currentTime.Day(), output.Total),
		Footer:        "Komiser",
		Fields:        fields,
		FooterIcon:    "https://github.com/tailwarden/komiser",
		Ts:            json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}

	err = slack.PostWebhook(cfg.Slack.Webhook, &msg)
	if err != nil {
		log.Warn(err)
	}
}

func sendCostBreakdownReport(ctx context.Context, cfg models.Config) {
	groups := make([]models.OutputCostByField, 0)
	currentTime := time.Now()

	for _, field := range []string{"service", "provider", "account", "region"} {
		err := db.NewRaw(fmt.Sprintf("SELECT %s as label, SUM(cost) as total FROM resources GROUP BY %s ORDER by total desc;", field, field)).Scan(ctx, &groups)
		if err != nil {
			log.WithError(err).Error("scan failed")
		}

		segments := groups

		if len(groups) > 3 {
			segments = groups[:4]
			if len(groups) > 4 {
				sum := 0.0
				for i := 4; i < len(groups); i++ {
					sum += groups[i].Total
				}

				segments = append(segments, models.OutputCostByField{
					Label: "Others",
					Total: sum,
				})
			}
		}

		fields := make([]slack.AttachmentField, 0)
		for _, segment := range segments {
			fields = append(fields, slack.AttachmentField{
				Title: segment.Label,
				Value: fmt.Sprintf("%.2f", segment.Total),
				Short: true,
			})
		}

		attachment := slack.Attachment{
			Color:         "good",
			AuthorName:    "Komiser",
			AuthorSubname: "by Tailwarden",
			AuthorLink:    "https://tailwarden.com",
			AuthorIcon:    "https://cdn.komiser.io/images/komiser-logo.jpeg",
			Text:          fmt.Sprintf("On %s %d: cost breakdown by cloud %s", currentTime.Month(), currentTime.Day(), field),
			Footer:        "Komiser",
			Fields:        fields,
			FooterIcon:    "https://github.com/tailwarden/komiser",
			Ts:            json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
		}
		msg := slack.WebhookMessage{
			Attachments: []slack.Attachment{attachment},
		}

		err = slack.PostWebhook(cfg.Slack.Webhook, &msg)
		if err != nil {
			log.Warn(err)
		}
	}
}
