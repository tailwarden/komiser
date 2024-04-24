package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	//"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/models"
)

func (handler *ApiHandler) BulkUpdateTagsHandler(c *gin.Context) {
	var input BulkUpdateTag

	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resource := Resource{Tags: input.Tags}

	for _, resourceId := range input.Resources {
		_, err = models.HandleQuery(handler.db, handler.ctx, "UPDATE_TAGS", &resource, map[string]string{"id": fmt.Sprintf("%d", resourceId)})
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
	tags := make([]Tag, 0)

	resourceId := c.Param("id")

	err := json.NewDecoder(c.Request.Body).Decode(&tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resource := Resource{Tags: tags}

	_, err = models.HandleQuery(handler.db, handler.ctx, "UPDATE_TAGS", &resource, map[string]string{"id": string(resourceId)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while updating tags"})
		return
	}

	if handler.telemetry {
		handler.analytics.TrackEvent("tagging_resources", nil)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "tags has been successfuly updated"})
}
