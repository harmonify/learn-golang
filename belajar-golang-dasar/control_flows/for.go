package main

import "fmt"

func main() {
	fmt.Println("Simple for =====================================")
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	fmt.Println("Complete for (w/ init & post statement)=====================================")

	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println("Perulangan 2 yang ke", counter2)
	}

	fmt.Println("For range - iterate data in a collection (i.e. array, slice, map) =====================================")

	names := []string{"Eko", "Kurniawan", "Khannedy"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
	// substitute unused variable with underscore (_)
	for _, name := range names {
		fmt.Println(name)
	}

	fmt.Println("Break & continue =====================================")

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan break ke", i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan continue ke", i)
	}
}
