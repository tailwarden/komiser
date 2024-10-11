package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
	"github.com/tailwarden/komiser/models"
)

func hitCustomWebhook(endpoint string, secret string, viewName string, resources int, cost float64, alertType string) {
	var payloadJSON []byte
	var err error
	payload := models.CustomWebhookPayload{
		Komiser:   Version,
		View:      viewName,
		Timestamp: time.Now().Unix(),
	}

	switch alertType {
	case "BUDGET":
		payload.Message = "Cost alert"
		payload.Data = cost
	case "USAGE":
		payload.Message = "Usage alert"
		payload.Data = float64(resources)
	default:
		log.Error("Invalid Alert Type")
		return
	}

	payloadJSON, err = json.Marshal(payload)
	if err != nil {
		log.Error("Couldn't encode JSON payload:", err)
		return
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payloadJSON))
	if err != nil {
		log.Error("Couldn't create HTTP request for custom webhook endpoint:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if len(secret) > 0 {
		req.Header.Set("Authorization", secret)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Couldn't make HTTP request for custom webhook endpoint:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Error("Custom Webhook with endpoint " + endpoint + " returned back a status code of " + string(rune(resp.StatusCode)) + " . Expected Status Code: 200")
		return
	}
}

// createSlackAttachment creates a `slack.Attachment`.
// This attachment can then be added to the WebhookMessage for the alert.
func createSlackAttachment(viewName string, port int, viewId int, resources int, cost float64, hostname string, alertType string) slack.Attachment {
	// if the hostname is empty i.e. not defined in config.toml
	// default to localhost and the runtime port value
	if hostname == "" {
		hostname = fmt.Sprintf("http://localhost:%d", port)
	}

	attachment := slack.Attachment{
		Color:         "danger",
		AuthorName:    "Komiser",
		AuthorSubname: "by Tailwarden",
		AuthorLink:    "https://tailwarden.com",
		AuthorIcon:    "https://cdn.komiser.io/images/komiser-logo.jpeg",
		Footer:        "Komiser",
		Actions: []slack.AttachmentAction{
			{
				Name: "open",
				Text: "Open view",
				Type: "button",
				URL:  fmt.Sprintf("%s/inventory?view=%d", hostname, viewId),
			},
		},
		Fields: []slack.AttachmentField{
			{
				Title: "View",
				Value: viewName,
			},
		},
		FooterIcon: "https://github.com/tailwarden/komiser",
		Ts:         json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
	}

	if alertType == "BUDGET" {
		attachment.Text = "Cost alert :warning:"
		attachment.Fields = append(attachment.Fields, slack.AttachmentField{
			Title: "Cost",
			Value: fmt.Sprintf("%.2f$", cost),
		})
	} else if alertType == "USAGE" {
		attachment.Text = "Usage alert :warning:"
		attachment.Fields = append(attachment.Fields, slack.AttachmentField{
			Title: "Resources",
			Value: fmt.Sprintf("%d", resources),
		})
	}
	return attachment
}

func hitSlackWebhook(viewName string, port int, viewId int, resources int, cost float64, webhookUrl string, hostname string, alertType string) {

	attachment := createSlackAttachment(viewName, port, viewId, resources, cost, hostname, alertType)

	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}

	err := slack.PostWebhook(webhookUrl, &msg)
	if err != nil {
		log.Warn(err)
	}

}
