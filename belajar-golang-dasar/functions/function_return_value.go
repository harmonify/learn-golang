package main

import "fmt"

func main() {
	fmt.Println(getFullName("Eko", "Khannedy"))
}

func getHello(firstName string, lastName string) string {
	return "Hello " + firstName + " " + lastName
}
