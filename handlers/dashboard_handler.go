package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/utils"
)

func (handler *ApiHandler) DashboardStatsHandler(c *gin.Context) {
	regions := struct {
		Count int `bun:"count" json:"total"`
	}{}

	err := handler.db.NewRaw("SELECT COUNT(*) as count FROM (SELECT DISTINCT region FROM resources) AS temp").Scan(handler.ctx, &regions)
	if err != nil {
		logrus.WithError(err).Error("scan failed")
	}

	resources := struct {
		Count int `bun:"count" json:"total"`
	}{}

	err = handler.db.NewRaw("SELECT COUNT(*) as count FROM resources").Scan(handler.ctx, &resources)
	if err != nil {
		logrus.WithError(err).Error("scan failed")
	}

	cost := struct {
		Sum float64 `bun:"sum" json:"total"`
	}{}

	err = handler.db.NewRaw("SELECT SUM(cost) as sum FROM resources").Scan(handler.ctx, &cost)
	if err != nil {
		logrus.WithError(err).Error("scan failed")
	}

	accounts := struct {
		Count int `bun:"count" json:"total"`
	}{}

	err = handler.db.NewRaw("SELECT COUNT(*) as count FROM (SELECT DISTINCT account FROM resources) AS temp").Scan(handler.ctx, &accounts)
	if err != nil {
		logrus.WithError(err).Error("scan failed")
	}

	output := struct {
		Resources int     `json:"resources"`
		Regions   int     `json:"regions"`
		Costs     float64 `json:"costs"`
		Accounts  int     `json:"accounts"`
	}{
		Resources: resources.Count,
		Regions:   regions.Count,
		Costs:     cost.Sum,
		Accounts:  accounts.Count,
	}

	c.JSON(http.StatusOK, output)
}

func (handler *ApiHandler) ResourcesBreakdownStatsHandler(c *gin.Context) {
	input := models.InputResources{}

	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	groups := make([]models.OutputResources, 0)

	if len(input.Exclude) > 0 {
		s, _ := json.Marshal(input.Exclude)
		err = handler.db.NewRaw(fmt.Sprintf("SELECT %s as label, COUNT(*) as total FROM resources WHERE %s NOT IN (%s) GROUP BY %s ORDER by total desc;", input.Filter, input.Filter, strings.Trim(string(s), "[]"), input.Filter)).Scan(handler.ctx, &groups)
		if err != nil {
			logrus.WithError(err).Error("scan failed")
		}
	} else {
		err = handler.db.NewRaw(fmt.Sprintf("SELECT %s as label, COUNT(*) as total FROM resources GROUP BY %s ORDER by total desc;", input.Filter, input.Filter)).Scan(handler.ctx, &groups)
		if err != nil {
			logrus.WithError(err).Error("scan failed")
		}
	}

	segments := groups

	if len(groups) > 3 {
		segments = groups[:4]
		if len(groups) > 4 {
			sum := 0
			for i := 4; i < len(groups); i++ {
				sum += groups[i].Total
			}

			segments = append(segments, models.OutputResources{
				Label: "Other",
				Total: sum,
			})
		}
	}
	c.JSON(http.StatusOK, segments)
}

func (handler *ApiHandler) LocationBreakdownStatsHandler(c *gin.Context) {
	groups := make([]models.OutputResources, 0)

	err := handler.db.NewRaw("SELECT region as label, COUNT(*) as total FROM resources GROUP BY region ORDER by total desc;").Scan(handler.ctx, &groups)
	if err != nil {
		logrus.WithError(err).Error("scan failed")
	}

	locations := make([]models.OutputLocations, 0)

	for _, group := range groups {

		location := utils.GetLocationFromRegion(group.Label)

		if location.Label != "" {
			locations = append(locations, models.OutputLocations{
				Name:      location.Name,
				Label:     location.Label,
				Latitude:  location.Latitude,
				Longitude: location.Longitude,
				Resources: group.Total,
			})
		}
	}

	c.JSON(http.StatusOK, locations)
}

func (handler *ApiHandler) CostBreakdownHandler(c *gin.Context) {
	input := models.InputCostBreakdown{}

	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	groups := make([]models.OutputCostBreakdownRaw, 0)

	query := `SELECT provider, account, region, service, sum(cost) as total, strftime("%Y-%m-%d", fetched_at) as period FROM resources WHERE`
	if input.Granularity == "MONTHLY" {
		query = `SELECT provider, account, region, service, sum(cost) as total, strftime("%Y-%m", fetched_at) as period FROM resources WHERE`
	}

	if len(input.Exclude) > 0 {
		s, _ := json.Marshal(input.Exclude)
		err = handler.db.NewRaw(fmt.Sprintf(`%s %s NOT IN (%s) AND DATE(fetched_at) BETWEEN '%s' AND '%s' GROUP BY %s;`, query, input.Group, strings.Trim(string(s), "[]"), input.Start, input.End, input.Group)).Scan(handler.ctx, &groups)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		err = handler.db.NewRaw(fmt.Sprintf(`%s DATE(fetched_at) BETWEEN '%s' AND '%s' GROUP BY %s;`, query, input.Start, input.End, input.Group)).Scan(handler.ctx, &groups)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}

	data := make(map[string][]models.Datapoint, 0)

	for _, group := range groups {
		if len(data[group.Period]) == 0 {
			data[group.Period] = make([]models.Datapoint, 0)
		}

		name := group.Provider
		switch input.Group {
		case "account":
			name = group.Account
		case "region":
			name = group.Region
		case "service":
			name = group.Service
		}

		data[group.Period] = append(data[group.Period], models.Datapoint{
			Name:   name,
			Amount: group.Total,
		})
	}

	output := make([]models.OutputCostBreakdown, 0)

	for period, datapoints := range data {
		sort.Slice(datapoints, func(i, j int) bool {
			return datapoints[i].Amount > datapoints[j].Amount
		})

		listOfDatapoints := datapoints
		if len(datapoints) > 3 {
			listOfDatapoints = datapoints[:4]
			if len(datapoints) > 4 {
				sum := 0.0
				for i := 4; i < len(datapoints); i++ {
					sum += datapoints[i].Amount
				}

				listOfDatapoints = append(listOfDatapoints, models.Datapoint{
					Name:   "Other",
					Amount: sum,
				})
			}
		}

		output = append(output, models.OutputCostBreakdown{
			Date:       period,
			Datapoints: listOfDatapoints,
		})
	}

	sort.Slice(output, func(i, j int) bool {
		dateFormat := "2006-01-02"
		if input.Granularity == "MONTHLY" {
			dateFormat = "2006-01"
		}

		firstDate, _ := time.Parse(dateFormat, output[i].Date)
		secondDate, _ := time.Parse(dateFormat, output[j].Date)

		return firstDate.Before(secondDate)
	})

	c.JSON(http.StatusOK, output)
}
