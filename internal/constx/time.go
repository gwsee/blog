package constx

import "time"

var TimeLocation *time.Location

func init() {
	if TimeLocation == nil {
		TimeLocation = time.Local
	}
}
func Now() time.Time {
	return time.Now().In(TimeLocation)
}
