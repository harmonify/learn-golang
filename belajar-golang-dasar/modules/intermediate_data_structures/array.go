package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Eko"
	names[1] = "Kurniawan"
	names[2] = "Khannedy"
	fmt.Println(names)
	fmt.Println(names[0])

	values := [3]int{
		90,
		80,
	}
	fmt.Println(values)
	fmt.Println(values[2])
	fmt.Println(len(values))

	// This array length will be calculated and fixed to 10 elements
	others := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(others)
}
