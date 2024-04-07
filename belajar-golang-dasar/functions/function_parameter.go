package main

import "fmt"

func main() {
	sayHelloTo("Eko", "Khannedy")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
