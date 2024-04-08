package belajar_golang_goroutines

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

// Golang pass variable as value by default
// But for channel, it is passed as reference no matter what
func GiveMeResponse(channel chan string) {
	n := time.Duration(math.Ceil(rand.Float64()*2000) + 100) // 100-2100ms
	fmt.Printf("Waiting %dms...\n", n)
	time.Sleep(n * time.Millisecond)
	channel <- "Eko Kurniawan Khannedy"
	// fmt.Println("Selesai mengirim data ke channel")
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
}
