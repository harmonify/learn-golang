package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// Ticker is a representation of repeating events,
// if the ticker time is expired, it will send the current time to its channel.
// Ticker can be used to repeat call a function.
// Ticker behaves like setInterval function in JavaScript.
func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for time := range ticker.C {
		fmt.Println(time)
	}
}

// .Tick will directly returns channel for the time ticker
// This will run forever because we didn't call the .Stop method
func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}
}
