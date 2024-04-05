package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // pass by reference using pointer

	address2.City = "Bandung"
	fmt.Println(address1) // berubah
	fmt.Println(address2) // berubah

	fmt.Println("===============================================================")

	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"} // change all variables that are pointed at the same data as `address2` onto the new Address instance
	fmt.Println(address1)
	fmt.Println(address2)
}
