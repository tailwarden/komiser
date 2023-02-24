package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/utils"
)

func (handler *ApiHandler) DashboardStatsHandler(w http.ResponseWriter, r *http.Request) {
	regions := struct {
		Count int `bun:"count" json:"total"`
	}{}

	handler.db.NewRaw("SELECT COUNT(*) as count FROM (SELECT DISTINCT region FROM resources) AS temp").Scan(handler.ctx, &regions)

	resources := struct {
		Count int `bun:"count" json:"total"`
	}{}

	handler.db.NewRaw("SELECT COUNT(*) as count FROM resources").Scan(handler.ctx, &resources)

	cost := struct {
		Sum float64 `bun:"sum" json:"total"`
	}{}

	handler.db.NewRaw("SELECT SUM(cost) as sum FROM resources").Scan(handler.ctx, &cost)

	accounts := struct {
		Count int `bun:"count" json:"total"`
	}{}

	handler.db.NewRaw("SELECT COUNT(*) as count FROM (SELECT DISTINCT account FROM resources) AS temp").Scan(handler.ctx, &accounts)

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

	respondWithJSON(w, 200, output)
}

func (handler *ApiHandler) ResourcesBreakdownStatsHandler(w http.ResponseWriter, r *http.Request) {
	input := models.InputResources{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	groups := make([]models.OutputResources, 0)

	if len(input.Exclude) > 0 {
		s, _ := json.Marshal(input.Exclude)
		handler.db.NewRaw(fmt.Sprintf("SELECT %s as label, COUNT(*) as total FROM resources WHERE %s NOT IN (%s) GROUP BY %s ORDER by total desc;", input.Filter, input.Filter, strings.Trim(string(s), "[]"), input.Filter)).Scan(handler.ctx, &groups)
	} else {
		handler.db.NewRaw(fmt.Sprintf("SELECT %s as label, COUNT(*) as total FROM resources GROUP BY %s ORDER by total desc;", input.Filter, input.Filter)).Scan(handler.ctx, &groups)
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
	respondWithJSON(w, 200, segments)
}

func (handler *ApiHandler) LocationBreakdownStatsHandler(w http.ResponseWriter, r *http.Request) {
	groups := make([]models.OutputResources, 0)

	handler.db.NewRaw("SELECT region as label, COUNT(*) as total FROM resources GROUP BY region ORDER by total desc;").Scan(handler.ctx, &groups)

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

	respondWithJSON(w, 200, locations)
}

func (handler *ApiHandler) CostBreakdownHandler(w http.ResponseWriter, r *http.Request) {
	input := models.InputCostBreakdown{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
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
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		err = handler.db.NewRaw(fmt.Sprintf(`%s DATE(fetched_at) BETWEEN '%s' AND '%s' GROUP BY %s;`, query, input.Start, input.End, input.Group)).Scan(handler.ctx, &groups)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
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

	respondWithJSON(w, 200, output)
}
