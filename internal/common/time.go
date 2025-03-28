package common

import "time"

func GetZeroAt(ts int64) int64 {
	t := time.Unix(ts, 0)
	zeroTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return zeroTime.Unix()
}

func GetNextZeroAt(ts int64) int64 {
	t := time.Unix(ts, 0).AddDate(0, 0, 1)
	zeroTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return zeroTime.Unix()
}
