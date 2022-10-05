package integrations

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	. "github.com/mlabouardy/komiser/handlers/aws"
	. "github.com/mlabouardy/komiser/handlers/azure"
	azureConfig "github.com/mlabouardy/komiser/handlers/azure/config"
	. "github.com/mlabouardy/komiser/handlers/digitalocean"
	. "github.com/mlabouardy/komiser/handlers/gcp"
	. "github.com/mlabouardy/komiser/services/ini"
	. "github.com/mlabouardy/komiser/services/integrations/slack"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

type AccountHandler struct {
	awsHandler          *AWSHandler
	gcpHandler          *GCPHandler
	azureHandler        *AzureHandler
	digitaloceanHandler *DigitalOceanHandler
	slack               Slack
}

func NewAccountHandler(awsHandler *AWSHandler, gcpHandler *GCPHandler, azureHandler *AzureHandler, digitaloceanHandler *DigitalOceanHandler) *AccountHandler {
	accountHandler := AccountHandler{
		awsHandler:          awsHandler,
		gcpHandler:          gcpHandler,
		azureHandler:        azureHandler,
		digitaloceanHandler: digitaloceanHandler,
	}
	return &accountHandler
}

func (handler *AccountHandler) ListCloudAccountsHandler(w http.ResponseWriter, r *http.Request) {
	accounts := make(map[string][]string, 0)

	sections, err := OpenFile(config.DefaultSharedCredentialsFilename())
	if err == nil {
		for _, section := range sections.List() {
			accounts["AWS"] = append(accounts["AWS"], section)
		}
	}

	_, err = google.DefaultTokenSource(oauth2.NoContext, compute.ComputeReadonlyScope)
	if err == nil {
		accounts["GCP"] = append(accounts["GCP"], "Default")
	}

	err = azureConfig.ParseEnvironment()
	if err == nil {
		subscriptionID := azureConfig.SubscriptionID()
		accounts["AZURE"] = append(accounts["AZURE"], subscriptionID)
	}

	if os.Getenv("DIGITALOCEAN_ACCESS_TOKEN") != "" {
		accounts["DIGITALOCEAN"] = append(accounts["DIGITALOCEAN"], "Default")
	}

	respondWithJSON(w, 200, accounts)
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
