package aws

import (
	"encoding/json"
	"net/http"

	. "github.com/mlabouardy/komiser/services/cache"
	. "github.com/mlabouardy/komiser/services/gcp"
)

type GCPHandler struct {
	cache Cache
	gcp   GCP
}

func NewGCPHandler(cache Cache, dataset string) *GCPHandler {
	gcpHandler := GCPHandler{
		cache: cache,
		gcp: GCP{
			BigQueryDataset: dataset,
		},
	}
	return &gcpHandler
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
