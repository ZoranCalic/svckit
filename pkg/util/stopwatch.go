package util

import "time"

type Stopwatch struct {
	start time.Time
}

func NewStopwatch() *Stopwatch {
	return &Stopwatch{start: time.Now()}
}

func (s *Stopwatch) GetMs() float64 {
	return float64(time.Now().Sub(s.start).Nanoseconds()) / 1e6
}

func (s *Stopwatch) GetNs() int {
	return int(time.Now().Sub(s.start).Nanoseconds())
}

func (s *Stopwatch) Lap() float64 {
	ret := float64(time.Now().Sub(s.start).Nanoseconds()) / 1e6
	s.start = time.Now()
	return ret
}
