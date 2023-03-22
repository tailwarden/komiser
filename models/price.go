package models

type PricingResult struct {
	Product struct {
		Sku string `json:"sku"`
	} `json:"product"`
	Terms map[string]map[string]map[string]map[string]struct {
		PricePerUnit struct {
			USD string `json:"USD"`
		} `json:"pricePerUnit"`
	} `json:"terms"`
}
