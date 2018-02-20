package models

import "time"

type Cost struct {
	Start  time.Time
	End    time.Time
	Amount float64
	Unit   string
}
