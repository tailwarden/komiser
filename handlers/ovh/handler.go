package ovh

import (
	"encoding/json"
	"net/http"

	. "github.com/mlabouardy/komiser/services/cache"
	. "github.com/mlabouardy/komiser/services/ovh"
)

type OVHHandler struct {
	cache Cache
	ovh   OVH
}

func NewOVHHandler(cache Cache, endpoint string) *OVHHandler {
	ovhHandler := OVHHandler{
		cache: cache,
		ovh: OVH{
			Endpoint: endpoint,
		},
	}
	return &ovhHandler
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
