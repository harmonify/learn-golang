package main

import (
	"fmt"
	"reflect"
)

// reflect is often useful when building general libraries or utilitiess
// https://golang.org/pkg/reflect

type Sample struct {
	Name string `required:"true"`
}

type Person struct {
	Name    string `required:"true"`
	Address string `required:"true"`
	Email   string `required:"true"`
	Age     int    `required:"true"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("value type:", valueType)
	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Printf("%s.%s with type %s\n", valueType, valueField.Name, valueField.Type)
	}
}

func IsValid(value any) (result bool) {
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()

			// note: DON'T use the below code in production
			// the validation isn't perfect. it is served as an example
			switch f.Type.Kind() {
			case reflect.String:
				result = data != ""
			case reflect.Int:
				result = data != 0
			}
		}
	}
	return result
}

func main() {
	readField(Sample{"Eko"})

	fmt.Println("=========================================")

	person := Person{
		Name: "Eko",
		Address: "Jakarta",
		Email: "eko@example.com",
		Age: 30,
	}

	readField(person)
	
	fmt.Println(IsValid(person))
}
