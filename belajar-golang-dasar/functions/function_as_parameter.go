package main

import "fmt"

func main() {
	sayHelloWithFilter("Eko", offensiveFilter)
	sayHelloWithFilter("Anjing", offensiveFilter)
	sayHelloWithFilter2("Anjing", offensiveFilter)
}

func offensiveFilter(word string) string {
	if word == "Anjing" {
		return "***"
	} else {
		return word
	}
}

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

type Filter func(string) string

func sayHelloWithFilter2(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}
