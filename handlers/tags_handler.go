package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tailwarden/komiser/models"
)

func (handler *ApiHandler) BulkUpdateTagsHandler(c *gin.Context) {
	var input models.BulkUpdateTag

	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, resourceId := range input.Resources {
		_, err = handler.ctrl.UpdateTags(c, input.Tags, fmt.Sprint(resourceId))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error while updating tags"})
			return
		}
	}

	if handler.telemetry {
		handler.analytics.TrackEvent("tagging_resources", nil)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "tags has been successfuly updated"})
}

func (handler *ApiHandler) UpdateTagsHandler(c *gin.Context) {
	tags := make([]models.Tag, 0)

	resourceId := c.Param("id")

	err := json.NewDecoder(c.Request.Body).Decode(&tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = handler.ctrl.UpdateTags(c, tags, resourceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while updating tags"})
		return
	}

	if handler.telemetry {
		handler.analytics.TrackEvent("tagging_resources", nil)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "tags has been successfuly updated"})
}
