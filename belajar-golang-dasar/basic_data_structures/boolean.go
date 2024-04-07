package main

import "fmt"

func main() {
	fmt.Println("Benar =", true)
	fmt.Println("Salah =", false)
	fmt.Println("Salah =", !true)
	fmt.Println("Benar =", !!true)
	fmt.Println("Benar =", 1 < 3 && 3 > 1)
	fmt.Println("Benar =", 1 > 3 || 3 > 1)
}
