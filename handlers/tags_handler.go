package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
		_, err = handler.db.NewUpdate().Model(&resource).Column("tags").Where("id = ?", resourceId).Exec(handler.ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error while updating tags"})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{"message": "tags has been successfuly updated"})
}

func (handler *ApiHandler) UpdateTagsHandler(c *gin.Context) {
	tags := make([]Tag, 0)

	resourceId := c.Param("id")

	id, err := strconv.Atoi(resourceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "resource id should be an integer"})
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(&tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resource := Resource{Tags: tags}

	_, err = handler.db.NewUpdate().Model(&resource).Column("tags").Where("id = ?", id).Exec(handler.ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while updating tags"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "tags has been successfuly updated"})
}
