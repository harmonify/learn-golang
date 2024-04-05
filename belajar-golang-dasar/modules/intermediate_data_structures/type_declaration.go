package main

import "fmt"

func main() {
	type NoKtp string

	var ktpEko NoKtp = "11111111"
	fmt.Println(ktpEko)

	var contoh string = "222222222"
	contohKtp := NoKtp(contoh)
	fmt.Println(contohKtp)
}
