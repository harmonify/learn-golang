package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for index, arg := range args {
		fmt.Println("arg", index, ":", arg)
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error:", err.Error())
	}
}
