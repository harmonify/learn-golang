package main

import (
	"errors"
	"fmt"
	// "math"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		// return int(math.Inf(0)), errors.New("Pembagian dengan NOL")
		return 0, errors.New("Pembagian dengan NOL")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	// hasil, err := Pembagian(100, 10)
	hasil, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Hasil", hasil)
		fmt.Println("Error", err.Error())
	}
}
