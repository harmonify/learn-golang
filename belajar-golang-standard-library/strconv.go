package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("parse boolean to string =============================")
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(result)
	}

	fmt.Println("parse int to string ================================")
	result2, err2 := strconv.Atoi("1000")
	if err2 != nil {
		fmt.Println("Error", err2.Error())
	} else {
		fmt.Println(result2)
	}

	fmt.Println("parse string to int ================================")
	result3 := strconv.Itoa(999)
	fmt.Println(result3)

	result4 := strconv.FormatInt(999, 2)
	fmt.Println(result4)


}
