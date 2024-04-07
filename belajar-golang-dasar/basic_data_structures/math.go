package main

import "fmt"

func main() {
	var (
		a = 10
		b = 10
		c = a + b
		d = a + b*c
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	var i = 10
	i += 10
	i -= 10
	i *= 10
	i /= 10
	i %= 3 // modulus
	fmt.Println(i)

	var j = 30
	j++
	fmt.Println(j)
	j--
	j--
	fmt.Println(j)
	fmt.Println(-j)
}
