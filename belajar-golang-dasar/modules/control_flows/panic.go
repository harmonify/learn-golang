package main

import "fmt"

func endApp() {
	message := recover()
	fmt.Println("Terjadi panic:", message)
	fmt.Println("End app")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Ups an error")
	} else {
		fmt.Println("Running application")
	}
}

func main() {
	runApp(true)
}
