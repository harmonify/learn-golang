package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// JSON often comes in from the network or files
// A JSON Decoder is useful to handle this case,
// as we didn't need to store the JSON in a temporary variable
// we could use the `io.Reader` directly
func TestStreamDecoder(t *testing.T) {

	reader, _ := os.Open("Customer.json")
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)

}
