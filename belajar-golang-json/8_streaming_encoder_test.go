package belajar_golang_json

import (
	"encoding/json"
	"os"
	"testing"
)

// Opposite of the JSON decoder,
// JSON encoder is useful to directly send JSON data using the `io.Writer`
// without having to store the JSON data in temporary string or []byte variable
func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Eko",
		MiddleName: "Kurniawan",
		LastName:   "Khannedy",
	}

	encoder.Encode(customer)
}
