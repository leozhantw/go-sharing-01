package dateutil

import "time"

func IsQuarterly(date time.Time) bool {
	return int(date.Month())%3 == 0
}
