package main

import "fmt"

func main() {
	const firstName = "Wendy"
	const lastName = "Surya"
	fmt.Println(firstName + " " + lastName)

	const (
		first = "Eko"
		last  = "Khannedy"
	)
	fmt.Println(first)
	fmt.Println(last)
}
