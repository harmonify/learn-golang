package main

import "fmt"

func main() {
	firstName, lastName := getFullName()
	fmt.Println("Hello", firstName, lastName)

	fmt.Println("ignore return value ==============================")

	firstName2, _ := getFullName()
	fmt.Println("Hello", firstName2)

}

func getFullName() (string, string) {
	return "Eko", "Khannedy"
}
