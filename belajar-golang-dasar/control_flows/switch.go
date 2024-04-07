package main

import "fmt"

func main() {

	name := "Eko"

	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
		fmt.Println(name)
	case "Joko":
		fmt.Println("Cok")
	default:
		fmt.Println("Siapa lo")
		// title := "The greatest eko" // cool, the scope didn't leak out like javascript
	}
	// fmt.Println(title)

	fmt.Println("Switch short statement ====================================")

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Kok iso?")
	case false:
		fmt.Println("Mantul")
	}

	fmt.Println("Switch without condition statement ====================================")

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("WKWKWK")
	default:
		fmt.Println("WOW")
	}
}
