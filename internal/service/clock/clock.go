package clock

import "time"

//go:generate mockgen -destination ./mock/clock.go -package mock . Clock

type Clock interface {
	Now() time.Time
}

type RealClock struct {
}

func NewRealClock() *RealClock {
	return &RealClock{}
}

func (r RealClock) Now() time.Time {
	return time.Now()
}
