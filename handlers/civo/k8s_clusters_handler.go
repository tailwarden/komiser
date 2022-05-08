package civo

import (
	"fmt"
	"net/http"

	"github.com/mlabouardy/komiser/handlers/civo/config"
)

func (handler *CivoHandler) K8sClustersHandler(w http.ResponseWriter, r *http.Request) {
	err := config.ParseEnvironment()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't parse environment")
	}
	apiKey := config.ApiKey()
	regionCode := config.RegionCode()
	key := fmt.Sprintf("civo.%s.k8s.clusters", apiKey)
	response, found := handler.cache.Get(key)
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.civo.GetK8sClustersCount(apiKey, regionCode)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Compute:GetK8sClustersCount is missing")
		} else {
			handler.cache.Set(key, response)
			respondWithJSON(w, 200, response)
		}
	}
}
