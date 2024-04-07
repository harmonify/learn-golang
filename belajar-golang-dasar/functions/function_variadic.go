package main

import "fmt"

func main() {
	total := sumAll(10, 10, 10, 10, 10)
	fmt.Println(total)

	fmt.Println("Using slice in variadic function =========================================================")
	numbers := []int{10, 10, 10}
	fmt.Println(sumAll(numbers...))
}

// Variadic function
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
