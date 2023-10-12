package handlers

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (handler *ApiHandler) NewFeedbackHandler(c *gin.Context) {

	url := os.Getenv("DISCORD_FEEDBACK_URL")
	err := c.Request.ParseMultipartForm(100 << 20)
	if err != nil {
		log.WithError(err).Error("Unable to parse form")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse form"})
		return
	}
	email := c.PostForm("email")
	description := c.PostForm("description")
	imageFile, _, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get image file"})
		return
	}
	defer imageFile.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	payloadField, _ := writer.CreateFormField("payload_json")
	payloadJSON := fmt.Sprintf(`{
		"content": "",
		"embeds": [
			{
				"title": "Feedback",
				"description": "%s",
				"fields": [
					{
						"name": "Email",
						"value": "%s",
						"inline": true
					}
				]
			}
		]
	}`, description, email)

	_, _ = payloadField.Write([]byte(payloadJSON))

	imagePart, err := writer.CreateFormFile("files[0]", filepath.Base("temp-image"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create form file"})
		return
	}
	_, _ = io.Copy(imagePart, imageFile)

	writer.Close()

	resp, err := http.Post(url, writer.FormDataContentType(), &requestBody)
	if err != nil {
		log.WithError(err).Error("scan failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed sending data to discord"})
		return
	}
	defer resp.Body.Close()

	c.JSON(http.StatusOK, gin.H{"Response": "Feedback submitted"})
}
