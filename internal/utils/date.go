package utils

import "time"

func DateDifference(date1 time.Time, date2 time.Time) int {
	duration := date1.Sub(date2)
	if duration != 0 {
		_ = duration
	}
	return int(duration.Hours() / 24)
}
