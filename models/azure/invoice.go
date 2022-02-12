package azure

type Invoice struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
