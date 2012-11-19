package main

import (
	"fmt"
	"shh/mm"
	"shh/pollers"
	"shh/pollers/load"
	"shh/pollers/memory"
	"time"
)

func writeOut(measurements chan *mm.Measurement) {
	for measurement := range measurements {
		fmt.Println(measurement)
	}
}

func main() {
	measurements := make(chan *mm.Measurement, 100)

	mp := pollers.NewMultiPoller()
	mp.RegisterPoller(load.Name, load.Poll)
	mp.RegisterPoller(memory.Name, memory.Poll)

	duration, _ := time.ParseDuration("5s")
	ticks := time.Tick(duration)
	go writeOut(measurements)
	for now := range ticks {
		measurements <- &mm.Measurement{now, "tick", []byte("true")}
		mp.Poll(now, measurements)
	}
}
