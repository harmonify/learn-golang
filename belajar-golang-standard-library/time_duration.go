package main

import (
	"fmt"
	"time"
)

func main() {
	var duration1 time.Duration = time.Second * 100
	var duration2 time.Duration = time.Millisecond * 10
	duration3 := duration1 - duration2

	fmt.Println(duration3)
	fmt.Printf("duration : %d\n", duration3)
}
