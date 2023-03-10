package handlers

import "net/http"

func (handler *ApiHandler) TelemetryHandler(w http.ResponseWriter, r *http.Request) {
	response := struct {
		TelemetryEnabled bool `json:"telemetry_enabled"`
	}{
		TelemetryEnabled: handler.telemetryEnabled,
	}

	respondWithJSON(w, 200, response)
}
