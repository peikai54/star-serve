package utils

import (
	"time"
)

func GetTodayEnd() time.Time {
	t := time.Now()
	addTime := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	return addTime
}
