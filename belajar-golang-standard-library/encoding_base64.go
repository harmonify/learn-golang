package main

import (
	"fmt"
	"encoding/base64"
)

func main() {
	value := "Eko Kurniawan Khannedy"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)
	
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(decoded)
		fmt.Println(string(decoded))
	}
}
