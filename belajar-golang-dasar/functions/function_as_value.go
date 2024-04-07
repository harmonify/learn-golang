package main

import "fmt"

func main() {
	goodbye := getGoodbye
	fmt.Println(goodbye)
	fmt.Println(goodbye("Eko"))
}

func getGoodbye(name string) string {
	return "Good Bye " + name
}
