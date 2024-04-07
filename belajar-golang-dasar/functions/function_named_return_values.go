package main

import "fmt"

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println("Hello", firstName, middleName, lastName)
}

// The variable declaration happens on the function declaration
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Eko"
	middleName = "Kurniawan"
	lastName = "Khannedy"
	return firstName, middleName, lastName
}
