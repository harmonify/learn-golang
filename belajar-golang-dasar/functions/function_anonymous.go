package main

import "fmt"

type Blacklist func(string) bool

func main() {
	blacklist := func(name string) bool {
		if name == "Eko" {
			return true
		} else {
			return false
		}
	}

	registerUser("Eko", blacklist)
	registerUser("Wendy", blacklist)
	registerUser("Budi", func(name string) bool {
		return name == "Anjing"
	})
}

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}
