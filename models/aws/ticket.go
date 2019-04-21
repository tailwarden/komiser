package aws

import "time"

type Ticket struct {
	Timestamp    time.Time `json:"timestamp"`
	Status       string    `json:"status"`
	SeverityCode string    `json:"severity"`
	CategoryCode string    `json:"category"`
	ServiceCode  string    `json:"service"`
}
