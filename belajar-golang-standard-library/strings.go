package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Eko Kurniawan", "Eko"))
	fmt.Println(strings.Split("Eko Kurniawan", " "))
	fmt.Println(strings.ToLower("Eko Kurniawan"))
	fmt.Println(strings.ToUpper("Eko Kurniawan"))
	fmt.Println(strings.ToTitle("eko kurniawan khannedy"))
	fmt.Println(strings.Trim("               Eko Kurniawan     Khannedy ", " "))
	fmt.Println(strings.ReplaceAll("Eko Kurniawan Eko Khannedy", "Eko", "Budi"))
}
