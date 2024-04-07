package main

import (
	"fmt"
	"regexp"
)

// Golang's regex use RE2 library
// Syntax: https://github.com/google/re2/wiki/Syntax

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("e[a-z]o")
	
	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("eKo"))

	fmt.Println(regex.FindAllString("eko edo egi ego e1o eto eKoezo", 10))
}
