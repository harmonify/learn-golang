package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)
	var nilaiUnsigned16 uint16 = uint16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16) // Number overflow, int16 value range from -32768 to 32767
	fmt.Println(nilai16 + 3)
	fmt.Println(nilaiUnsigned16)

	name := "Wendy Surya"
	var n = name[0] // type is byte or uint8
	var nString = string(n)
	fmt.Println(name)
	fmt.Println(n)
	fmt.Println(nString)
}
