package azure

import (
	"encoding/json"
	"net/http"

	. "github.com/mlabouardy/komiser/services/azure"
	. "github.com/narasago/komiser/services/cache"
)

type AzureHandler struct {
	cache    Cache
	multiple bool
	azure    Azure
}

func NewAzureHandler(cache Cache) *AzureHandler {
	azureHandler := AzureHandler{
		cache: cache,
		azure: Azure{},
	}

	return &azureHandler
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
