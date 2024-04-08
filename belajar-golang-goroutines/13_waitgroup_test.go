package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	// group.Done() will add -1 to the waitgroup counter state
	defer group.Done()

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		// group.Add() should be called before starting the goroutine to avoid race condition
		group.Add(1)
		go RunAsynchronous(group)
	}

	// group.Wait() will wait until the counter state is 0
	// After this method is called, we must not call its .Add method after that
	// or it will be panic
	group.Wait()

	fmt.Println("Selesai")
}
