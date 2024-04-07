package main

import "fmt"

func main() {

	name := "Eko"

	if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Joko" {
		fmt.Println("Cok")
	} else {
		fmt.Println("Siape lu")
	}

	fmt.Println("If short statement =================================")

	if length := len(name); length > 5 {
		fmt.Println("Kok iso")
	} else {
		fmt.Println("Wkwk iyalah")
	}

}
