package gcp

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"cloud.google.com/go/bigquery"
	. "github.com/mlabouardy/komiser/models/gcp"
	"google.golang.org/api/iterator"
)

func (gcp GCP) CostInLastSixMonths() ([]Cost, error) {
	costs := make([]Cost, 0)

	query := fmt.Sprintf("SELECT invoice.month, (SUM(CAST(cost * 1000000 AS int64)) + SUM(IFNULL((SELECT SUM(CAST(c.amount * 1000000 AS int64)) FROM UNNEST(credits) c),0))) / 1000000 AS total FROM `%s` GROUP BY 1 ORDER BY 1 ASC LIMIT 6;", gcp.BigQueryDataset)

	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, strings.Split(gcp.BigQueryDataset, ".")[0])
	if err != nil {
		return costs, err
	}

	q := client.Query(query)
	it, err := q.Read(ctx)
	if err != nil {
		return costs, err
	}

	for {
		var values []bigquery.Value
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return costs, err
		}

		raw := values[0].(string)
		date := raw[:4] + "-" + raw[4:]

		costs = append(costs, Cost{
			Date: date,
			Cost: values[1].(float64),
		})
	}

	return costs, err
}

func (gcp GCP) MonthlyCostPerService() ([]CostPerService, error) {
	costs := make([]CostPerService, 0)

	query := fmt.Sprintf("SELECT service.description, invoice.month, currency, SUM(cost) as total FROM `%s` GROUP BY service.description, invoice.month, currency ORDER BY invoice.month;", gcp.BigQueryDataset)

	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, strings.Split(gcp.BigQueryDataset, ".")[0])
	if err != nil {
		return costs, err
	}

	q := client.Query(query)
	it, err := q.Read(ctx)
	if err != nil {
		return costs, err
	}

	temp := make(map[string][]Group, 0)
	for {
		var values []bigquery.Value
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return costs, err
		}

		service := values[0].(string)
		raw := values[1].(string)
		date := raw[:4] + "-" + raw[4:]
		currency := values[2].(string)
		cost := values[3].(float64)

		if temp[date] != nil {
			temp[date] = append(temp[date], Group{
				Cost:    cost,
				Unit:    currency,
				Service: service,
			})
		} else {
			temp[date] = make([]Group, 0)
			temp[date] = append(temp[date], Group{
				Cost:    cost,
				Unit:    currency,
				Service: service,
			})
		}
	}

	for date, groups := range temp {
		sort.Slice(groups, func(i, j int) bool {
			return groups[i].Cost > groups[j].Cost
		})

		costs = append(costs, CostPerService{
			Date:   date,
			Unit:   groups[0].Unit,
			Groups: groups,
		})
	}
	return costs, err
}
