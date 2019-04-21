package aws

import (
	"net/http"
)

func (handler *AWSHandler) CostAndUsageHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("cost_usage_history")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeCostAndUsage(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ce:GetCostAndUsage is missing")
		} else {
			handler.cache.Set("cost_usage_history", response.History)
			respondWithJSON(w, 200, response.History)
		}
	}
}

func (handler *AWSHandler) CurrentCostHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("cost_usage_total")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeCostAndUsage(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ce:GetCostAndUsage is missing")
		} else {
			handler.cache.Set("cost_usage_total", response.Total)
			respondWithJSON(w, 200, response.Total)
		}
	}
}

func (handler *AWSHandler) CostAndUsagePerInstanceTypeHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("cost_per_instance_type")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeCostAndUsagePerInstanceType(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ce:GetCostAndUsage is missing")
		} else {
			handler.cache.Set("cost_per_instance_type", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) DescribeForecastPriceHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("cost_forecast")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeForecastPrice(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ce:GetCostForecast is missing")
		} else {
			handler.cache.Set("cost_forecast", response)
			respondWithJSON(w, 200, response)
		}
	}
}
