package handlers

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (handler *ApiHandler) NewFeedbackHandler(c *gin.Context) {

	err := c.Request.ParseMultipartForm(1000 << 20)
	if err != nil {
		logrus.WithError(err).Error("Unable to parse form")
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

	_ = writer.WriteField("Email", email)
	_ = writer.WriteField("Description", description)

	imagePart, err := writer.CreateFormFile("files[0]", filepath.Base("temp-image"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create form file"})
		return
	}
	_, _ = io.Copy(imagePart, imageFile)

	writer.Close()

	var url = ""

	resp, err := http.Post(url, writer.FormDataContentType(), &requestBody)
	if err != nil {
		logrus.WithError(err).Error("scan failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "scan failed"})
		return
	}
	defer resp.Body.Close()

	c.JSON(http.StatusOK, resp.Body)
}
