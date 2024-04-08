package belajar_golang_context

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"testing"
	"time"
)

func CreateCounter(ctx context.Context, wg *sync.WaitGroup) chan int {
	destination := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				n := time.Duration(math.Ceil(rand.Float64()*200+50)) * time.Millisecond // 50-250ms
				time.Sleep(n)                                                           // simulasi slow
			}
		}
	}()

	return destination
}

// Context with cancel is excellent to help preventing goroutine leak.
// We can signal a goroutine that the context is done (or cancelled),
// then the goroutine can read it,
// and decide that its own process should be terminated.
func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	wg := &sync.WaitGroup{}

	destination := CreateCounter(ctx, wg)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	cancel() // mengirim sinyal cancel ke context

	wg.Wait()

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
