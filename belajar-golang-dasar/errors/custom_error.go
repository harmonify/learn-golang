package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"id is empty"}
	}

	if id != "eko" {
		return &notFoundError{"id not found"}
	}

	// ok

	return nil
}

func main() {
	// err := SaveData("eko", nil) // success
	// err := SaveData("", nil) // validationError
	err := SaveData("budi", nil) // notFoundError
	if err != nil {
		// if short statement
		// <declaration>, <condition>
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("validation error:", validationErr.Error())
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println("not found error:", notFoundErr.Error())
		// } else {
		// 	fmt.Println("unknown error:", err.Error())
		// }

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation error:", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error:", finalError.Error())
		default:
			fmt.Println("unknown error:", finalError.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}
