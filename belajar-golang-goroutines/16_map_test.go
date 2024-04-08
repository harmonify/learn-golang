package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func AddToMap(wg *sync.WaitGroup, data *sync.Map, key string, value int) {
	defer wg.Done()
	data.Store(key, value)
}

// sync.Map is a struct similar to Go map,
// but it is safe to use in concurrency environment, like goroutine
// It has methods like .Store, .Load, .Delete, and .Range (to iterate the data).
func TestMap(t *testing.T) {
	data := &sync.Map{}
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go AddToMap(wg, data, fmt.Sprintf("Loop ke %d", i), i)
	}

	wg.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
