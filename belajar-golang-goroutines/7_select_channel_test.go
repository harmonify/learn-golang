package belajar_golang_goroutines

import (
	"fmt"
	"testing"
)

// Select channel, allows us to receive data from several channels
// And spawn goroutine to process the received data
func TestSelectFastestChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	// This allows us to receive the fastest data from different channels
	// Similar to Promise.race method in JavaScript
	select {
	case data := <-channel1:
		fmt.Println("Data dari Channel 1", data)
	case data := <-channel2:
		fmt.Println("Data dari Channel 2", data)
	}

	fmt.Println("Selesai")
}

func TestSelectMultipleChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
		}
		counter++
		if counter == 2 {
			break
		}
	}

	fmt.Println("Selesai")
}
