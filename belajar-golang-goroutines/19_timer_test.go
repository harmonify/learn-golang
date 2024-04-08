package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Timer will send the current time (after the specified delay has passed)
// to its channel
func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

// After method is a simplified version of NewTimer we use above
func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)
}

// AfterFunc method accepts delay and function as its parameters
// The function will be executed after the specified delay
// This behaves like setTimeout() in JavaScript
func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}

	group.Add(1)
	time.AfterFunc(5*time.Second, func() {
		defer group.Done()
		fmt.Println(time.Now())
	})

	fmt.Println(time.Now())

	group.Wait()
}
