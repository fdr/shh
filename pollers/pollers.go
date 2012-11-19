package pollers

import (
	"fmt"
	"shh/mm"
	"time"
)

type PollerFunc func(now time.Time, measurements chan *mm.Measurement)

type MultiPoller map[string]PollerFunc

func NewMultiPoller() *MultiPoller {
	mp := MultiPoller(make(map[string]PollerFunc))
	return &mp
}

func (m *MultiPoller) RegisterPoller(name string, f PollerFunc) {
	(*m)[name] = f
}

func (m *MultiPoller) Poll(now time.Time, measurements chan *mm.Measurement) {
	for name, pollerFunc := range *m {
		measurements <- &mm.Measurement{
			now,
			fmt.Sprintf("ticking.%s", name),
			[]byte("true"),
		}

		go pollerFunc(now, measurements)
	}
}
