package main

import "fmt"

func main() {
	fmt.Println("11111 - slice ========================================================")

	// Slice = sub array
	// Slice size can change unlike array
	// Slice has 3 data, which are pointer, length, and capacity
	// Pointer, which point to the first element in the slice
	// Length, the slice length
	// Capacity, the slice capacity, where the length must not exceed the capacity

	months := [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}
	fmt.Println(cap(months)) // same as len(array)

	slice1 := months[4:6]
	fmt.Println(slice1)

	slice2 := months[:3]
	fmt.Println(slice2)

	var slice3 []string = months[3:]
	fmt.Println(slice3)

	slice4 := months[:]
	fmt.Println(slice4)

	fmt.Println("22222 - len() and cap() of slice ===============================================")

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	fmt.Println("33333 - append() ==============================================================")

	/**
	 * append(): make new slice by pushing it to the end of the slice,
	 * but if the slice capacity is already full,
	 * it will make a new array (what is the length?) to fit the new value.
	 * The array isn't accessible to us though.
	 */

	// Array
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	fmt.Println(days)

	// Slice
	daySlice1 := days[5:] // [Sabtu, Minggu]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days)      // [Senin, Selasa, Rabu, Kamis, Jumat, Sabtu Baru, Minggu Baru], the values inside `days` are changed because `daySlice1` only act as a pointer to `days`.
	fmt.Println(daySlice1) // [Sabtu Baru, Minggu Baru]

	daySlice2 := append(daySlice1, "Libur Baru") // [Sabtu Baru, Minggu Baru, Libur Baru]
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2) // [Ups, Minggu Baru, Libur Baru]
	fmt.Println(daySlice1) // [Sabtu Baru, Minggu Baru], the values in `daySlice1` are unchanged because the `daySlice2` is using a new array to fit the new value ("Ups")
	fmt.Println(days)      // [Senin, Selasa, Rabu, Kamis, Jumat, Sabtu Baru, Minggu Baru], unchanged as well

	fmt.Println("days", len(days))
	fmt.Println("daySlice1", len(daySlice1), cap(daySlice1)) // use array `days`
	fmt.Println("daySlice2", len(daySlice2), cap(daySlice2)) // use new array

	daySlice3 := append(daySlice2, "Wkwk1")
	daySlice4 := append(daySlice3, "Wkwk2")

	fmt.Println(daySlice2)
	fmt.Println(daySlice3)
	fmt.Println(daySlice4)
	fmt.Println("daySlice3", len(daySlice3), cap(daySlice3)) // use the same array as `daySlice2`
	fmt.Println("daySlice4", len(daySlice4), cap(daySlice4)) // use new array

	fmt.Println("44444 - make() ==============================================================")

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Eko"
	newSlice[1] = "Eko"
	// newSlice[2] = "Eko" // will cause an error (panic)

	fmt.Println("newSlice", newSlice, len(newSlice), cap(newSlice))

	newSlice2 := append(newSlice, "Eko")
	fmt.Println("newSlice2", newSlice2, len(newSlice2), cap(newSlice2)) // still use the same array as `newSlice`

	newSlice2[0] = "Budi"
	fmt.Println("newSlice2", newSlice2, len(newSlice2), cap(newSlice2))
	fmt.Println("newSlice", newSlice, len(newSlice), cap(newSlice))
	// wtf

	fmt.Println("55555 - copy() ==============================================================")
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println("fromSlice", fromSlice, len(fromSlice), cap(fromSlice))
	fmt.Println("toSlice", toSlice, len(toSlice), cap(toSlice))

	fmt.Println("66666 - array vs slice ==============================================================")
	// [n]string or [...]string is array type declaration, where n is an integer
	// []string is slice type declaration

	thisIsArray := [...]int{1, 2, 3, 4, 5}
	thisIsSlice := []int{1, 2, 3, 4, 5}

	fmt.Println("thisIsArray", thisIsArray, len(thisIsArray), cap(thisIsArray))
	fmt.Println("thisIsSlice", thisIsSlice, len(thisIsSlice), cap(thisIsSlice))

	// slice usage is more common than array
}
