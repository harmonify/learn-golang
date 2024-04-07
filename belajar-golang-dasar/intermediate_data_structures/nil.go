package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

// Cannot, nil can only be used for these interfaces: interface, function, map, slice, pointer, channel
// func Impossible(name string) string {
// 	if name == "" {
// 		return nil
// 	} else {
// 		return name
// 	}
// }

func main() {
	data := NewMap("")
	fmt.Println(data) // empty
	if data == nil {
		fmt.Println("Data map masih kosong")
	} else {
		fmt.Println(data["name"])
	}
	fmt.Println(data["name"])
	fmt.Println("That statement didn't throw an error")
}
