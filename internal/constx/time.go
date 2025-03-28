package constx

import (
	"log"
	"time"
)

var TimeLocation *time.Location

func init() {
	if TimeLocation == nil {
		var err error
		TimeLocation, err = time.LoadLocation("Asia/Shanghai")
		if err != nil {
			log.Println("TimeLocation:", TimeLocation, err)
		} else {
			TimeLocation = time.Local
		}
	}
}
func Now() time.Time {
	return time.Now().In(TimeLocation)
}
func NowUnix() int64 {
	return time.Now().In(TimeLocation).Unix()
}
