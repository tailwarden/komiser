package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/tailwarden/komiser/models"
)

type AWSCtxKey uint8

const (
	CostexplorerKey AWSCtxKey = iota
)

func GetCostAndUsage(ctx context.Context, region string, svcName string) (float64, error) {
	total := 0.0
	costexplorerOutputList, ok := ctx.Value(CostexplorerKey).([]*costexplorer.GetCostAndUsageOutput)
	if !ok || costexplorerOutputList == nil {
		return 0, fmt.Errorf("incorrect costexplorerOutputList")
	}
	for _, costexplorerOutput := range costexplorerOutputList {
		for _, group := range costexplorerOutput.ResultsByTime {
			for _, v := range group.Groups {
				if v.Keys[0] == svcName {
					amt, err := strconv.ParseFloat(*v.Metrics["UnblendedCost"].Amount, 64)
					if err != nil {
						return 0, err
					}
					total += amt
				}
			}
		}
	}
	return total, nil
}

func ExtractResources(costexplorerOutputList []*costexplorer.GetCostAndUsageOutput) ([]models.Resource, error) {
	r := []models.Resource{}
	for _, costexplorerOutput := range costexplorerOutputList {
		for _, group := range costexplorerOutput.ResultsByTime {
			endInterval := group.TimePeriod.End
			createdAt, err := time.Parse("2006-01-02", *endInterval)
			if err != nil {
				return r, err
			}
			for _, v := range group.Groups {
				amt, err := strconv.ParseFloat(*v.Metrics["UnblendedCost"].Amount, 64)
				if err != nil {
					return r, err
				}
				r = append(r, models.Resource{
					Provider:  "AWS",
					Service:   v.Keys[0],
					Region:    v.Keys[1],
					Cost:      amt,
					CreatedAt: createdAt,
					FetchedAt: createdAt,
				})
			}
		}
	}
	return r, nil
}

type ProductEntry struct {
	Product struct {
		Attributes struct {
			Group              string `json:"group"`
			Operation          string `json:"operation"`
			GroupDescription   string `json:"groupDescription"`
			RequestDescription string `json:"requestDescription"`
			InstanceType       string `json:"instanceType"`
			InstanceTypeFamily string `json:"instanceTypeFamily"`
		} `json:"attributes"`
	} `json:"product"`
	Terms struct {
		OnDemand map[string]struct {
			PriceDimensions map[string]PriceDimensions `json:"priceDimensions"`
		} `json:"OnDemand"`
	} `json:"terms"`
}

type PriceDimensions struct {
	EndRange     string  `json:"endRange"`
	BeginRange   float64 `json:"beginRange,string"`
	PricePerUnit struct {
		USD float64 `json:"USD,string"`
	} `json:"pricePerUnit"`
}

func GetCost(pds []PriceDimensions, v float64) float64 {
	total := 0.0
	for _, pd := range pds {
		applicableRange := v
		if pd.BeginRange < v {
			if pd.EndRange != "Inf" {
				endRange, _ := strconv.ParseFloat(pd.EndRange, 64)
				if v > endRange {
					applicableRange = endRange
				}
			}
			total += (applicableRange - pd.BeginRange) * pd.PricePerUnit.USD
		}
	}
	return total
}

func GetPriceMap(pricingOutput *pricing.GetProductsOutput, field string) (map[string][]PriceDimensions, error) {
	priceMap := make(map[string][]PriceDimensions)

	if pricingOutput != nil && len(pricingOutput.PriceList) > 0 {
		for _, item := range pricingOutput.PriceList {
			price := ProductEntry{}
			err := json.Unmarshal([]byte(item), &price)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
			}

			var key string
			switch field {
			case "group":
				key = price.Product.Attributes.Group
			case "operation":
				key = price.Product.Attributes.Operation
			case "groupDescription":
				key = price.Product.Attributes.GroupDescription
			case "requestDescription":
				key = price.Product.Attributes.RequestDescription
			case "instanceType":
				key = price.Product.Attributes.InstanceType
			case "instanceTypeFamily":
				key = price.Product.Attributes.InstanceTypeFamily
			}

			unitPrices := []PriceDimensions{}
			for _, pd := range price.Terms.OnDemand {
				for _, p := range pd.PriceDimensions {
					unitPrices = append(unitPrices, p)
				}
			}

			priceMap[key] = unitPrices
		}
	}

	return priceMap, nil
}

func Int64PtrToFloat64(i *int64) float64 {
	if i == nil {
		return 0.0 // or any default value you prefer
	}
	return float64(*i)
}
