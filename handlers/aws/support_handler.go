package aws

import (
	"net/http"
)

func (handler *AWSHandler) SupportOpenTicketsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("support_open_tickets")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.OpenSupportTickets(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "support:DescribeCases is missing")
		} else {
			handler.cache.Set("support_open_tickets", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) SupportTicketsInLastSixMonthsHandlers(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("support_history")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.TicketsInLastSixMonthsTickets(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "support:DescribeCases is missing")
		} else {
			handler.cache.Set("support_history", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) DescribeServiceLimitsChecks(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("service_limits")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeServiceLimitsChecks(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "support:DescribeTrustedAdvisorChecks is missing")
		} else {
			handler.cache.Set("service_limits", response)
			respondWithJSON(w, 200, response)
		}
	}
}
