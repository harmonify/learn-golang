package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

// sync.Once, allows us to run a function exactly one time within concurenccy programming
func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	counter := 0
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			once.Do(func() {
				counter++
			})
		}()
	}

	group.Wait()
	fmt.Println("Counter", counter)
}
