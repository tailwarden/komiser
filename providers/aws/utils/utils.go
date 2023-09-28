package utils

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/service/pricing"
)

type ProductEntry struct {
	Product struct {
		Attributes struct {
			Group     string `json:"group"`
			Operation string `json:"operation"`
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
				if v < endRange {
					applicableRange = v
				}
			}
			total += (applicableRange - pd.BeginRange) * pd.PricePerUnit.USD
		}
	}
	return total
}

func GetPriceMap(pricingOutput *pricing.GetProductsOutput) (map[string][]PriceDimensions, error) {
	priceMap := make(map[string][]PriceDimensions)

	if pricingOutput != nil && len(pricingOutput.PriceList) > 0 {
		for _, item := range pricingOutput.PriceList {
			price := ProductEntry{}
			err := json.Unmarshal([]byte(item), &price)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
			}

			group := price.Product.Attributes.Group
			unitPrices := []PriceDimensions{}
			for _, pd := range price.Terms.OnDemand {
				for _, p := range pd.PriceDimensions {
					unitPrices = append(unitPrices, p)
				}
			}

			priceMap[group] = unitPrices
		}
	}

	return priceMap, nil
}
