package belajar_golang_context

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	// ctx, cancel := context.WithTimeout(parent, 2*time.Second)
	ctx, cancel := context.WithTimeoutCause(parent, 2*time.Second, errors.New("counter timeout"))
	// Although long running context will be cancelled,
	// when the timeout is reached,
	// This deferred cancel call will handle the case
	// where the context's process is done before the timeout
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
