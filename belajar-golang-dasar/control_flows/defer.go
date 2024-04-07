package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("Running application")
}

func main() {
	runApplication()
}
