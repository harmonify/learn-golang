package belajar_golang_goroutines

import (
	"fmt"
	"testing"
)

// Buffered channel allows for non-blocking process
// Only if the channel is not full
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Eko"
	channel <- "Kurniawan"
	channel <- "Khannedy"

	// Normal channel will not reach here because it is blocked until
	// other goroutine take the data

	fmt.Println("Selesai")
}
func TestBufferedChannel2(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Eko"
		channel <- "Kurniawan"
		channel <- "Khannedy"
		channel <- "Hehe"
		channel <- "There"
	}()

	fmt.Println(<-channel, <-channel, <-channel)
	fmt.Println("Selesai")
}
