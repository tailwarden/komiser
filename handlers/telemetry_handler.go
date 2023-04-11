package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *ApiHandler) TelemetryHandler(c *gin.Context) {
	response := struct {
		TelemetryEnabled bool `json:"telemetry_enabled"`
	}{
		TelemetryEnabled: handler.telemetry,
	}

	c.JSON(http.StatusOK, response)
}
