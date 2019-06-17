package ovh

import (
	"net/http"
)

func (handler *OVHHandler) DescribeSSLCertificatesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ovh.cloud.ssl.certifcates")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.ovh.GetSSLCertificates()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			handler.cache.Set("ovh.cloud.ssl.certificates", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *OVHHandler) DescribeSSLGatewaysHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ovh.cloud.ssl.gateways")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.ovh.GetSSLGateways()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			handler.cache.Set("ovh.cloud.ssl.gateways", response)
			respondWithJSON(w, 200, response)
		}
	}
}
