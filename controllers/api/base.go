package api

import "time"

const (
	kTimeLayout  = "2006-01-02 15:04:05"
	kTimeLayout2 = "2006-01-02"
)

func TimeFormat(ts int64) string {
	curTime := time.Unix(ts, 0)
	return curTime.Format(kTimeLayout)
}

func TimeFormat2(ts int64) string {
	curTime := time.Unix(ts, 0)
	return curTime.Format(kTimeLayout2)
}
