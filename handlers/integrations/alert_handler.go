package integrations

import (
	"encoding/json"
	"log"
	"net/http"

	. "github.com/mlabouardy/komiser/handlers/aws"
	. "github.com/mlabouardy/komiser/handlers/azure"
	. "github.com/mlabouardy/komiser/handlers/gcp"
	. "github.com/mlabouardy/komiser/services/integrations/slack"
)

type AlertHandler struct {
	awsHandler   *AWSHandler
	gcpHandler   *GCPHandler
	azureHandler *AzureHandler
	slack        Slack
}

func NewAlertHandler(awsHandler *AWSHandler, gcpHandler *GCPHandler, azureHandler *AzureHandler) *AlertHandler {
	alertHandler := AlertHandler{
		awsHandler:   awsHandler,
		gcpHandler:   gcpHandler,
		azureHandler: azureHandler,
	}
	return &alertHandler
}

func (handler *AlertHandler) ConfigureSlack(token string, channel string) {
	handler.slack = Slack{
		Channel: channel,
		Token:   token,
	}
}

func (handler *AlertHandler) DailyNotifHandler() {
	log.Println("Sending daily cost alerts")
	handler.slack.SendDailyNotification(handler.awsHandler, handler.gcpHandler, handler.azureHandler)
}

func (handler *AlertHandler) ListIntegrationsHandler(w http.ResponseWriter, r *http.Request) {
	integrations := map[string]bool{
		"slack": false,
	}
	if handler.slack.Token != "" {
		integrations["slack"] = true
	}
	respondWithJSON(w, 200, integrations)
}

func (handler *AlertHandler) SetupSlackHandler(w http.ResponseWriter, r *http.Request) {
	var slack Slack
	err := json.NewDecoder(r.Body).Decode(&slack)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error while setting up Slack")
	}
	handler.ConfigureSlack(slack.Token, slack.Channel)
	respondWithJSON(w, 200, map[string]string{"message": "Slack has been enabled"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
