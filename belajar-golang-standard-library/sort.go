package main

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

// We don't need to decorate the instance with * (as pointer)
// because a slice is already a pointer
func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type SortUserByName []User

func (s SortUserByName) Len() int {
	return len(s)
}

func (s SortUserByName) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

func (s SortUserByName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"Eko", 30},
		{"Budi", 35},
		{"Adit", 23},
		{"Joko", 25},
		{"Joko", 30},
		{"Joko", 28},
	}
	fmt.Println("Original:", users)

	sort.Sort(UserSlice(users))
	fmt.Println("Sort by age:", users)

	sort.Sort(SortUserByName(users))
	fmt.Println("Sort by name:", users)

	sort.Sort(sort.Reverse(SortUserByName(users)))
	fmt.Println("Reverse sort by name:", users)

	// Note: in many situations, the newer slices.SortFunc function is more ergonomic and runs faster.
	slices.SortFunc(users, func(a, b User) int {
		if n := cmp.Compare(a.Name, b.Name); n != 0 {
			return n
		}
		// If names are equal, order by age
		return cmp.Compare(a.Age, b.Age)
	})
	fmt.Println("Sort by name then age:", users)
}
