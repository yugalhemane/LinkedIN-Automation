package stealth

import (
	"math/rand"
	"time"
)

type TimingProfile struct {
	Min time.Duration
	Max time.Duration
}

func NewDefaultTiming() *TimingProfile {
	return &TimingProfile{
		Min: 500 * time.Millisecond,
		Max: 1800 * time.Millisecond,
	}
}

// Sleep simulates human think time
func (t *TimingProfile) Sleep() {
	delay := t.Min + time.Duration(rand.Int63n(int64(t.Max-t.Min)))
	time.Sleep(delay)
}
