package main

import "fmt"

func main() {
	// style 1
	var name string
	name = "Wendy"
	fmt.Println(name)
	name = "Surya"
	fmt.Println(name)

	// style 2
	var other string = "Eko"
	fmt.Println(other)

	// style 3
	otherName := "Budi"
	fmt.Println(otherName)

	// style 4
	var (
		firstName = "Wendy"
		lastName  = "Surya"
		fullName  string
	)
	fmt.Println(firstName + lastName)
	fullName = fmt.Sprintf("%s %s", firstName, lastName)
	fmt.Println(fullName)
}
