package domain

import "fmt"

type HelloService struct {}

func SayHello(name string) (string, error) {
	result := "Hello " + name + "!"
	fmt.Println(result)
	return result, nil
}
