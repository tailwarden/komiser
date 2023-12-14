package utils

import "time"

func BeginningOfMonth(date time.Time) time.Time {
	beginningOfMonth, err := time.Parse("2006-01-02", date.AddDate(0, 0, -date.Day()+1).Format("2006-01-02"))
	if err != nil {
		return time.Now()
	}
	return beginningOfMonth
}

func EndingOfMonth(date time.Time) time.Time {
	endingOfMonth, err := time.Parse("2006-01-02", date.AddDate(0, 1, -date.Day()+1).Format("2006-01-02"))
	if err != nil {
		return time.Now()
	}
	return endingOfMonth
}

func BeginningMonthsAgo(date time.Time, months int) time.Time {
	beginningOfTargetMonth, err := time.Parse("2006-01-02", date.AddDate(0, -months, -date.Day()+1).Format("2006-01-02"))
	if err != nil {
		return time.Now()
	}
	return beginningOfTargetMonth
}

func EndingOfLastMonth(date time.Time) time.Time {
	endingOfMonth, err := time.Parse("2006-01-02", date.AddDate(0, 0, -date.Day()+1).Format("2006-01-02"))
	if err != nil {
		return time.Now()
	}
	return endingOfMonth
}
