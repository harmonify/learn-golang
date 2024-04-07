package main

import "fmt"

type Man struct {
	Name string
}

// It is recommended to always use pointer when referencing struct instance in its method
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	man := Man{"Eko"}

	fmt.Println(man)

	man.Married()

	fmt.Println(man)
}
