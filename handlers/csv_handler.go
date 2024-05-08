package handlers

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/uptrace/bun/dialect"
)

func (handler *ApiHandler) DownloadInventoryCSV(c *gin.Context) {
	resources, err := handler.ctrl.ListResources(c)
	if err != nil {
		logrus.WithError(err).Error("Could not read from DB")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cloud not read from DB"})
		return
	}

	if handler.telemetry {
		handler.analytics.TrackEvent("exporting_csv", nil)
	}

	respondWithCSVDownload(resources, c)
}

func (handler *ApiHandler) DownloadInventoryCSVForView(c *gin.Context) {
	viewId := c.Param("viewId")

	view, err := handler.ctrl.GetView(c, viewId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resources := make([]models.Resource, 0)

	if len(view.Filters) == 0 {
		err := handler.db.NewRaw("SELECT * FROM resources").Scan(handler.ctx, &resources)
		if err != nil {
			logrus.WithError(err).Errorf("select failed")
		}
		respondWithCSVDownload(resources, c)
	}

	filterWithTags := false
	whereQueries := make([]string, 0)
	for _, filter := range view.Filters {
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
				c.JSON(http.StatusInternalServerError, gin.H{"error": "operation is invalid or not supported"})
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
				if handler.db.Dialect().Name() == dialect.SQLite {
					query = fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') IN (%s)))", key, strings.Join(filter.Values, ","))
				}
				whereQueries = append(whereQueries, query)
			case "NOT_CONTAINS":
			case "IS_NOT":
				for i := 0; i < len(filter.Values); i++ {
					filter.Values[i] = fmt.Sprintf("'%s'", filter.Values[i])
				}
				query := fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' NOT IN (%s)))", key, strings.Join(filter.Values, ","))
				if handler.db.Dialect().Name() == dialect.SQLite {
					query = fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') NOT IN (%s)))", key, strings.Join(filter.Values, ","))
				}
				whereQueries = append(whereQueries, query)
			case "IS_EMPTY":
				if handler.db.Dialect().Name() == dialect.SQLite {
					whereQueries = append(whereQueries, fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') = ''))", key))
				} else {
					whereQueries = append(whereQueries, fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' = ''))", key))
				}
			case "IS_NOT_EMPTY":
				if handler.db.Dialect().Name() == dialect.SQLite {
					whereQueries = append(whereQueries, fmt.Sprintf("((json_extract(value, '$.key') = '%s') AND (json_extract(value, '$.value') != ''))", key))
				} else {
					whereQueries = append(whereQueries, fmt.Sprintf("((res->>'key' = '%s') AND (res->>'value' != ''))", key))
				}
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "operation is invalid or not supported"})
				return
			}
		} else if filter.Field == "tags" {
			switch filter.Operator {
			case "IS_EMPTY":
				if handler.db.Dialect().Name() == dialect.SQLite {
					whereQueries = append(whereQueries, "json_array_length(tags) = 0")
				} else {
					whereQueries = append(whereQueries, "jsonb_array_length(tags) = 0")
				}
			case "IS_NOT_EMPTY":
				if handler.db.Dialect().Name() == dialect.SQLite {
					whereQueries = append(whereQueries, "json_array_length(tags) != 0")
				} else {
					whereQueries = append(whereQueries, "jsonb_array_length(tags) != 0")
				}
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "operation is invalid or not supported"})
				return
			}
		} else if filter.Field == "cost" {
			switch filter.Operator {
			case "EQUAL":
				cost, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "value should be a number"})
					return
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost = %f)", cost))
			case "BETWEEN":
				min, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "value should be a number"})
					return
				}
				max, err := strconv.ParseFloat(filter.Values[1], 64)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "value should be a number"})
					return
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost >= %f AND cost <= %f)", min, max))
			case "GREATER_THAN":
				cost, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "value should be a number"})
					return
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost > %f)", cost))
			case "LESS_THAN":
				cost, err := strconv.ParseFloat(filter.Values[0], 64)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "value should be a number"})
					return
				}
				whereQueries = append(whereQueries, fmt.Sprintf("(cost < %f)", cost))
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "value should be a number"})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "field is invalid or unsupported"})
			return
		}
	}

	whereClause := strings.Join(whereQueries, " AND ")

	if filterWithTags {
		query := fmt.Sprintf("SELECT id, resource_id, provider, account, service, region, name, created_at, fetched_at,cost, metadata, tags,link FROM resources CROSS JOIN jsonb_array_elements(tags) AS res WHERE %s ORDER BY id", whereClause)
		if len(view.Exclude) > 0 {
			s, _ := json.Marshal(view.Exclude)
			query = fmt.Sprintf("SELECT id, resource_id, provider, account, service, region, name, created_at, fetched_at,cost, metadata, tags,link FROM resources CROSS JOIN jsonb_array_elements(tags) AS res WHERE %s AND id NOT IN (%s) ORDER BY id", whereClause, strings.Trim(string(s), "[]"))
		}
		if handler.db.Dialect().Name() == dialect.SQLite {
			query = fmt.Sprintf("SELECT resources.id, resources.resource_id, resources.provider, resources.account, resources.service, resources.region, resources.name, resources.created_at, resources.fetched_at, resources.cost, resources.metadata, resources.tags, resources.link FROM resources CROSS JOIN json_each(tags) WHERE type='object' AND %s ORDER BY resources.id", whereClause)
			if len(view.Exclude) > 0 {
				s, _ := json.Marshal(view.Exclude)
				query = fmt.Sprintf("SELECT resources.id, resources.resource_id, resources.provider, resources.account, resources.service, resources.region, resources.name, resources.created_at, resources.fetched_at, resources.cost, resources.metadata, resources.tags, resources.link FROM resources CROSS JOIN json_each(tags) WHERE resources.id NOT IN (%s) AND type='object' AND %s ORDER BY resources.id", strings.Trim(string(s), "[]"), whereClause)
			}
		}

		err = handler.db.NewRaw(query).Scan(handler.ctx, &resources)
		if err != nil {
			logrus.WithError(err).Errorf("scan failed")
		}
	} else {
		query := fmt.Sprintf("SELECT * FROM resources WHERE %s ORDER BY id", whereClause)
		if len(view.Exclude) > 0 {
			s, _ := json.Marshal(view.Exclude)
			query = fmt.Sprintf("SELECT * FROM resources WHERE %s AND id NOT IN (%s) ORDER BY id", whereClause, strings.Trim(string(s), "[]"))
		}

		err = handler.db.NewRaw(query).Scan(handler.ctx, &resources)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	respondWithCSVDownload(resources, c)
}

func respondWithCSVDownload(resources []models.Resource, c *gin.Context) {
	file, err := os.Create("/tmp/export.csv")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create file at /tmp"})
		return
	}

	defer file.Close()
	defer os.Remove("/tmp/export.csv")

	fw := bufio.NewWriter(file)
	csvWriter := csv.NewWriter(fw)

	header := []string{"id", "provider", "account", "name", "service", "region", "tags", "cost", "metadata"}
	if err := csvWriter.Write(header); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not write csv"})
		return
	}

	for _, record := range resources {
		tags, err := json.Marshal(record.Tags)
		if err != nil {
			log.Fatalf("Could not marshal tags")
		}
		metadata, err := json.Marshal(record.Metadata)
		if err != nil {
			log.Fatalf("Could not marshal metadata")
		}

		row := []string{
			record.ResourceId, record.Provider, record.Account, record.Name, record.Service, record.Region, string(tags), fmt.Sprintf("%2.f", record.Cost), string(metadata),
		}
		if err := csvWriter.Write(row); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not write csv"})
			return
		}
	}
	fw.Flush()

	c.FileAttachment("/tmp/export.csv", "export.csv")
}
