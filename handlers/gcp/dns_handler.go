package aws

import (
	"net/http"
)

func (handler *GCPHandler) DNSManagedZonesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_dns_zones")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetManagedZones()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "dns:CloudPlatformReadOnlyScope is missing")
		} else {
			handler.cache.Set("gcp_dns_zones", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) DNSARecordsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_dns_a_records")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetARecords()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "dns:CloudPlatformReadOnlyScope is missing")
		} else {
			handler.cache.Set("gcp_dns_a_records", response)
			respondWithJSON(w, 200, response)
		}
	}
}
