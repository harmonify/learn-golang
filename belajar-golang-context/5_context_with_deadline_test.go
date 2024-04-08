package belajar_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	// Context with deadline behaves similar like context with timeout
	// But we specify `time.Time` instead of `time.Duration`
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	wg := &sync.WaitGroup{}

	destination := CreateCounter(ctx, wg)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
	}

	wg.Wait()

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
