package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Person struct {
	Name string
}

// A sync.Pool is used to avoid unnecessary allocations and reduce the amount of work the garbage collector has to do.
// When passing a value that is not a pointer to a function that accepts an interface, the value needs to be placed on the heap, which means an additional allocation
// This cancelled out the benefits of sync.Pool
// https://staticcheck.io/docs/checks/#SA6002
func TestPool(t *testing.T) {
	wg := sync.WaitGroup{}
	pool := sync.Pool{
		New: func() any {
			return &Person{"New"}
		},
	}

	pool.Put(&Person{"Eko"})
	pool.Put(&Person{"Kurniawan"})
	pool.Put(&Person{"Khannedy"})

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// There's a chance that the .Get() call will returns `nil` as
			// the pool is empty, in that case, specifying a function that
			// returns default value in the `New` property of the pool, could be beneficial.
			data := pool.Get()
			switch value := data.(type) {
			case *Person:
				fmt.Println("Loop", i, "- Person:", value.Name)
			default:
				fmt.Println("Loop", i, "- Unknown:", value)
			}
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	wg.Wait()
	fmt.Println("Selesai")
}
