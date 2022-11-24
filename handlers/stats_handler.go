package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	. "github.com/mlabouardy/komiser/models"
)

func (handler *ApiHandler) StatsHandler(w http.ResponseWriter, r *http.Request) {
	regions := struct {
		Count int `bun:"count" json:"total"`
	}{}

	handler.db.NewRaw("SELECT COUNT(*) FROM (SELECT DISTINCT region FROM resources) AS temp").Scan(handler.ctx, &regions)

	resources := struct {
		Count int `bun:"count" json:"total"`
	}{}

	handler.db.NewRaw("SELECT COUNT(*) FROM resources").Scan(handler.ctx, &resources)

	cost := struct {
		Sum int `bun:"sum" json:"total"`
	}{}

	handler.db.NewRaw("SELECT SUM(count) FROM resources").Scan(handler.ctx, &cost)

	output := struct {
		Resources int `json:"resources"`
		Regions   int `json:"regions"`
		Costs     int `json:"costs"`
	}{
		Resources: resources.Count,
		Regions:   regions.Count,
		Costs:     cost.Sum,
	}

	respondWithJSON(w, 200, output)
}

func (handler *ApiHandler) FilterStatsHandler(w http.ResponseWriter, r *http.Request) {
	var filters []Filter

	err := json.NewDecoder(r.Body).Decode(&filters)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	filterWithTags := false
	whereQueries := make([]string, 0)
	for _, filter := range filters {
		if filter.Field == "name" || filter.Field == "region" || filter.Field == "service" || filter.Field == "provider" || filter.Field == "account" {
			switch filter.Operator {
			case "IS":
				for i := 0; i < len(filter.Values); i++ {
					filter.Values[i] = fmt.Sprintf("'%s'", filter.Values[i])
				}
				query := fmt.Sprintf("(%s IN (%s))", filter.Field, strings.Join(filter.Values, ","))
				whereQueries = append(whereQueries, query)
			case "IS_NOT":
				for i := 0; i < len(filter.Values); i++ {
					filter.Values[i] = fmt.Sprintf("'%s'", filter.Values[i])
				}
				query := fmt.Sprintf("(%s NOT IN (%s))", filter.Field, strings.Join(filter.Values, ","))
				whereQueries = append(whereQueries, query)
			case "CONTAINS":
				queries := make([]string, 0)
				specialChar := "%"
				for i := 0; i < len(filter.Values); i++ {
					queries = append(queries, fmt.Sprintf("(%s LIKE '%s%s%s')", filter.Field, specialChar, filter.Values[i], specialChar))
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(%s)", strings.Join(queries, " OR ")))
			case "NOT_CONTAINS":
				queries := make([]string, 0)
				specialChar := "%"
				for i := 0; i < len(filter.Values); i++ {
					queries = append(queries, fmt.Sprintf("(%s NOT LIKE '%s%s%s')", filter.Field, specialChar, filter.Values[i], specialChar))
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(%s)", strings.Join(queries, " AND ")))
			case "IS_EMPTY":
				whereQueries = append(whereQueries, fmt.Sprintf("((coalesce(%s, '') = ''))", filter.Field))
			case "IS_NOT_EMPTY":
				whereQueries = append(whereQueries, fmt.Sprintf("((coalesce(%s, '') != ''))", filter.Field))
			default:
				respondWithError(w, http.StatusBadRequest, "Operation is invalid or not supported")
				return
			}
		} else if strings.HasPrefix(filter.Field, "tag:") {
			filterWithTags = true
			key := strings.ReplaceAll(filter.Field, "tag:", "")
			switch filter.Operator {
			case "CONTAINS":
			case "IS":
				for i := 0; i < len(filter.Values); i++ {
					filter.Values[i] = fmt.Sprintf("'%s'", filter.Values[i])
				}
				query := fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' IN (%s)))", key, strings.Join(filter.Values, ","))
				whereQueries = append(whereQueries, query)
			case "NOT_CONTAINS":
			case "IS_NOT":
				for i := 0; i < len(filter.Values); i++ {
					filter.Values[i] = fmt.Sprintf("'%s'", filter.Values[i])
				}
				query := fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' NOT IN (%s)))", key, strings.Join(filter.Values, ","))
				whereQueries = append(whereQueries, query)
			case "IS_EMPTY":
				whereQueries = append(whereQueries, fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' = ''))", key))
			case "IS_NOT_EMPTY":
				whereQueries = append(whereQueries, fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' != ''))", key))
			default:
				respondWithError(w, http.StatusBadRequest, "Operation is invalid or not supported")
				return
			}
		} else if filter.Field == "tags" {
			switch filter.Operator {
			case "IS_EMPTY":
				whereQueries = append(whereQueries, "jsonb_array_length(tags) = 0")
			case "IS_NOT_EMPTY":
				whereQueries = append(whereQueries, "jsonb_array_length(tags) != 0")
			default:
				respondWithError(w, http.StatusBadRequest, "Operation is invalid or not supported")
				return
			}
		} else {
			respondWithError(w, http.StatusBadRequest, "Field is invalid or not supported")
			return
		}
	}

	whereClause := strings.Join(whereQueries, " AND ")

	if filterWithTags {
		regions := struct {
			Count int `bun:"count" json:"total"`
		}{}

		query := fmt.Sprintf("FROM resources CROSS JOIN jsonb_array_elements(tags) AS res WHERE %s", whereClause)
		handler.db.NewRaw(fmt.Sprintf("SELECT COUNT(*) FROM (SELECT DISTINCT region %s) AS temp", query)).Scan(handler.ctx, &regions)

		resources := struct {
			Count int `bun:"count" json:"total"`
		}{}

		handler.db.NewRaw(fmt.Sprintf("SELECT COUNT(*) %s", query)).Scan(handler.ctx, &resources)

		cost := struct {
			Sum int `bun:"sum" json:"total"`
		}{}

		handler.db.NewRaw(fmt.Sprintf("SELECT SUM(count) %s", query)).Scan(handler.ctx, &cost)

		output := struct {
			Resources int `json:"resources"`
			Regions   int `json:"regions"`
			Costs     int `json:"costs"`
		}{
			Resources: resources.Count,
			Regions:   regions.Count,
			Costs:     cost.Sum,
		}

		respondWithJSON(w, 200, output)
	} else {
		query := fmt.Sprintf("FROM resources WHERE %s", whereClause)

		regions := struct {
			Count int `bun:"count" json:"total"`
		}{}
		handler.db.NewRaw(fmt.Sprintf("SELECT COUNT(*) FROM (SELECT DISTINCT region %s) AS temp", query)).Scan(handler.ctx, &regions)

		resources := struct {
			Count int `bun:"count" json:"total"`
		}{}

		handler.db.NewRaw(fmt.Sprintf("SELECT COUNT(*) %s", query)).Scan(handler.ctx, &resources)

		cost := struct {
			Sum int `bun:"sum" json:"total"`
		}{}

		handler.db.NewRaw(fmt.Sprintf("SELECT SUM(count) %s", query)).Scan(handler.ctx, &cost)

		output := struct {
			Resources int `json:"resources"`
			Regions   int `json:"regions"`
			Costs     int `json:"costs"`
		}{
			Resources: resources.Count,
			Regions:   regions.Count,
			Costs:     cost.Sum,
		}

		respondWithJSON(w, 200, output)
	}
}

func (handler *ApiHandler) ListRegionsHandler(w http.ResponseWriter, r *http.Request) {
	type Output struct {
		Region string `bun:"region" json:"region"`
	}

	outputs := make([]Output, 0)

	handler.db.NewRaw("SELECT DISTINCT(region) FROM resources").Scan(handler.ctx, &outputs)

	regions := make([]string, 0)

	for _, o := range outputs {
		regions = append(regions, o.Region)
	}

	respondWithJSON(w, 200, regions)
}

func (handler *ApiHandler) ListProvidersHandler(w http.ResponseWriter, r *http.Request) {
	type Output struct {
		Provider string `bun:"provider" json:"provider"`
	}

	outputs := make([]Output, 0)

	handler.db.NewRaw("SELECT DISTINCT(provider) FROM resources").Scan(handler.ctx, &outputs)

	providers := make([]string, 0)

	for _, o := range outputs {
		providers = append(providers, o.Provider)
	}

	respondWithJSON(w, 200, providers)
}

func (handler *ApiHandler) ListServicesHandler(w http.ResponseWriter, r *http.Request) {
	type Output struct {
		Service string `bun:"service" json:"service"`
	}

	outputs := make([]Output, 0)

	handler.db.NewRaw("SELECT DISTINCT(service) FROM resources").Scan(handler.ctx, &outputs)

	services := make([]string, 0)

	for _, o := range outputs {
		services = append(services, o.Service)
	}

	respondWithJSON(w, 200, services)
}

func (handler *ApiHandler) ListAccountsHandler(w http.ResponseWriter, r *http.Request) {
	type Output struct {
		Account string `bun:"account" json:"account"`
	}

	outputs := make([]Output, 0)

	handler.db.NewRaw("SELECT DISTINCT(account) FROM resources").Scan(handler.ctx, &outputs)

	accounts := make([]string, 0)

	for _, o := range outputs {
		accounts = append(accounts, o.Account)
	}

	respondWithJSON(w, 200, accounts)
}

func (handler *ApiHandler) ResourcesCounterHandler(w http.ResponseWriter, r *http.Request) {
	output := struct {
		Count int `bun:"count" json:"total"`
	}{}

	handler.db.NewRaw("SELECT COUNT(*) FROM resources").Scan(handler.ctx, &output)

	respondWithJSON(w, 200, output)
}

func (handler *ApiHandler) CostCounterHandler(w http.ResponseWriter, r *http.Request) {
	output := struct {
		Sum int `bun:"sum" json:"total"`
	}{}

	handler.db.NewRaw("SELECT SUM(count) FROM resources").Scan(handler.ctx, &output)

	respondWithJSON(w, 200, output)
}

func (handler *ApiHandler) EnableTrackingHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, map[string]bool{"tracking": !handler.noTracking})
}
