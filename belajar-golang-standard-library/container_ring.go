// Circular list data structure
package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	const SIZE = 5
	var data *ring.Ring = ring.New(SIZE)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println(value)
	})
}
