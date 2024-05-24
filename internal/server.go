package internal

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	v1 "github.com/tailwarden/komiser/internal/api/v1"
	"github.com/tailwarden/komiser/models"
)

func loggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := ctx.Request.Method
		reqUri := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()

		log.WithFields(log.Fields{
			"method":  reqMethod,
			"uri":     reqUri,
			"status":  statusCode,
			"latency": latencyTime,
			"ip":      clientIP,
		}).Info("HTTP request")

		ctx.Next()
	}
}

func runServer(address string, port int, telemetry bool, cfg models.Config, configPath string, accounts []models.Account) error {
	log.Infof("Komiser version: %s, commit: %s, buildt: %s", Version, Commit, Buildtime)

	r := v1.Endpoints(context.Background(), telemetry, analytics, db, cfg, configPath, accounts)

	r.Use(loggingMiddleware())

	if err := r.Run(fmt.Sprintf("%s:%d", address, port)); err != nil {
		return err
	}

	log.Infof("Server started on %s:%d", address, port)

	return nil
}
