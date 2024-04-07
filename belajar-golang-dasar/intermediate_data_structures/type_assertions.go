package main

import "fmt"

func random() any {
	// return "OK"
	// return 1
	return true
}

func main() {
	var result any = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)
	// var resultInt int = result.(int)
	// fmt.Println(resultInt) // will cause a panic because wrong type conversion

	fmt.Println("===================================================")

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}
