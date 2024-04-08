package belajar_golang_goroutines

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"testing"
)

// Range channel allows us to ignore the data count that we would receive from a channel
// This is useful when we didn't know the specific data count we would receive from a goroutine that continuous sending data to the channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		n := int(math.Ceil(rand.Float64()*20) + 1) // 1-21
		for i := 0; i < n; i++ {
			channel <- "perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	// we don't need to sleep here, because this for loop will stop when the channel is closed
	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}
