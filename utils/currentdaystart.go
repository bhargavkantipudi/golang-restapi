package utils

import "time"

func CurrentDayStart() int64 {

	tm := time.Now()
	year, month, day := tm.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Unix()
}
