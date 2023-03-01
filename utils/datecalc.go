package utils

import "time"

func BeginningOfMonth(date time.Time) time.Time {
	beginningOfMonth, err := time.Parse("2006-01-02", date.AddDate(0, 0, -date.Day()+1).Format("2006-01-02"))
	if err != nil {
		return time.Now()
	}
	return beginningOfMonth
}
