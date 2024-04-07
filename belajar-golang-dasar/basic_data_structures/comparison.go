package main

import "fmt"

func main() {
	var (
		a = 10
		b = 30
	)

	fmt.Println(a < b)
	fmt.Println(a > b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a == b)
	fmt.Println(a != b)

	fmt.Println("==================================================")

	var (
		name1 = "Eko"
		name2 = "Eko"
	)
	result := name1 == name2
	fmt.Println(result)
	fmt.Println("a" < "b")
}
