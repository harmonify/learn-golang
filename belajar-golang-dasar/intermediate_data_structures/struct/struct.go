package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, ", my name is", customer.Name)
}

func main() {
	var wendy Customer

	fmt.Println(wendy)

	wendy.Name = "Wendy"
	wendy.Address = "Indonesia"
	wendy.Age = 21
	fmt.Println(wendy)
	fmt.Println(wendy.Name)
	fmt.Println(wendy.Address)
	fmt.Println(wendy.Age)

	eko := Customer{
		Name:    "Eko",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(eko)

	budi := Customer{"Budi", "Indonesia", 30}
	fmt.Println(budi)

	wendy.sayHello("Agus")
	eko.sayHello("Agus")
	budi.sayHello("Agus")
}
