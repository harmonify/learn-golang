package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1 // copy value. By default, Golang pass variable by value.

	address2.City = "Bandung"
	fmt.Println(address1) // tidak berubah
	fmt.Println(address2) // berubah menjadi Bandung

	fmt.Println("====================================================")

	var address3 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address4 *Address = &address3 // pass by reference using pointer

	address4.City = "Bandung"
	fmt.Println(address3) // berubah
	fmt.Println(address4) // berubah
}
