package main

import "fmt"

func main() {
	fmt.Println("1 - Map =====================================================================================")
	var person1 map[string]string = map[string]string{}
	person1["name"] = "Eko"
	person1["address"] = "Subang"

	person := map[string]string{
		"name":    "Eko",
		"address": "Subang",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)
	fmt.Println(person["whatisthis"]) // the value depends on the value type declaration, if string then "", if int then 0
	fmt.Println(len(person))

	fmt.Println("2 - Map advanced =====================================================================================")
	book := make(map[string]string)
	book["title"] = "Crypto Trading Guide"
	book["author"] = "Akademi Crypto"
	book["year"] = "2023"
	book["ups"] = "Wrong"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)

}
