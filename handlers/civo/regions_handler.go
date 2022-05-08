package civo

import (
	"fmt"
	"net/http"

	"github.com/mlabouardy/komiser/handlers/civo/config"
)

func (handler *CivoHandler) RegionsHandler(w http.ResponseWriter, r *http.Request) {
	err := config.ParseEnvironment()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't parse environment")
	}
	apiKey := config.ApiKey()
	regionCode := config.RegionCode()
	key := fmt.Sprintf("civo.%s.regions", apiKey)
	response, found := handler.cache.Get(key)
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.civo.GetRegionsCount(apiKey, regionCode)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Compute:GetRegionsCount is missing")
		} else {
			handler.cache.Set(key, response)
			respondWithJSON(w, 200, response)
		}
	}
}
