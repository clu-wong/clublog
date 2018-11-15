package timex

import "time"

type NowTime struct {
	time.Time
}

func New(t time.Time) *NowTime{
	return &NowTime{t}
}