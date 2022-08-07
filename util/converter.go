package util

import (
	"time"
)

func GetTimestamp() (timestamp string) {
	ts := time.Now()

	// utc -> kst
	loc, _ := time.LoadLocation("Asia/Seoul")
	kst := ts.In(loc)
	timestamp = kst.Format(time.RFC3339)

	return
}
