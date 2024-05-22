package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
)

func (handler *ApiHandler) StatsHandler(c *gin.Context) {
	regions, err := handler.ctrl.CountRegionsFromResources(c)
	if err != nil {
		logrus.WithError(err).Error("scan failed")
	}

	resources, err := handler.ctrl.CountResources(c, "", "")
	if err != nil {
		logrus.WithError(err).Error("scan failed")
	}

	cost, err := handler.ctrl.SumResourceCost(c)
	if err != nil {
		logrus.WithError(err).Error("scan failed")
	}

	output := struct {
		Resources int     `json:"resources"`
		Regions   int     `json:"regions"`
		Costs     float64 `json:"costs"`
	}{
		Resources: resources.Total,
		Regions:   regions.Total,
		Costs:     cost.Total,
	}

	if handler.telemetry {
		handler.analytics.TrackEvent("global_stats", map[string]interface{}{
			"costs":     cost.Total,
			"regions":   regions.Total,
			"resources": resources.Total,
		})
	}

	c.JSON(http.StatusOK, output)
}

func (handler *ApiHandler) FilterStatsHandler(c *gin.Context) {
	var filters []models.Filter

	err := json.NewDecoder(c.Request.Body).Decode(&filters)
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{"message": err.Error()})
		return
	}

	view := new(models.View)
	view.Filters = filters

	regionCount, resourceCount, costCount, err := handler.ctrl.StatsWithFilter(c, *view, []int64{}, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	output := struct {
		Resources int     `json:"resources"`
		Regions   int     `json:"regions"`
		Costs     float64 `json:"costs"`
	}{
		Resources: resourceCount.Total,
		Regions:   regionCount.Total,
		Costs:     costCount.Total,
	}

	c.JSON(http.StatusOK, output)
}

func (handler *ApiHandler) ListRegionsHandler(c *gin.Context) {
	outputs, err := handler.ctrl.ListRegions(c)
	if err != nil {
		logrus.WithError(err).Error("scan failed")
	}

	regions := make([]string, 0)
	for _, o := range outputs {
		regions = append(regions, o.Region)
	}

	c.JSON(http.StatusOK, regions)
}

func (handler *ApiHandler) ListProvidersHandler(c *gin.Context) {
	if handler.db == nil {
		c.JSON(http.StatusInternalServerError, []string{})
		return
	}

	outputs, err := handler.ctrl.ListProviders(c)
	if err != nil {
		logrus.WithError(err).Error("scan failed")
	}

	providers := make([]string, 0)
	for _, o := range outputs {
		providers = append(providers, o.Provider)
	}

	c.JSON(http.StatusOK, providers)
}

func (handler *ApiHandler) ListServicesHandler(c *gin.Context) {
	outputs, err := handler.ctrl.ListServices(c)
	if err != nil {
		logrus.WithError(err).Error("scan failed")
	}

	services := make([]string, 0)
	for _, o := range outputs {
		services = append(services, o.Service)
	}

	c.JSON(http.StatusOK, services)
}

func (handler *ApiHandler) ListAccountsHandler(c *gin.Context) {
	if handler.db == nil {
		c.JSON(http.StatusInternalServerError, []string{})
		return
	}

	outputs, err := handler.ctrl.ListAccountNames(c)
	if err != nil {
		logrus.WithError(err).Error("scan failed")
	}

	accounts := make([]string, 0)

	for _, o := range outputs {
		accounts = append(accounts, o.Account)
	}

	c.JSON(http.StatusOK, accounts)
}
