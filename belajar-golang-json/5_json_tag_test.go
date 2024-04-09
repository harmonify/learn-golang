package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "P0001",
		Name:     "Apple Mac Book Pro",
		ImageURL: "http://example.com/image.png",
	}

	// When marshalling, the json package will read the struct tag to decide the resulting json attribute name
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	// Somehow, Golang's json package is case-insensitive when reading (or unmarshalling)
	// So, (id = ID = Id), and (ImageUrl = imageurl = imageUrl)
	// But, (ImageURL != image_url)
	jsonString := `{"id":"P0001","name":"Apple Mac Book Pro","image_url":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	json.Unmarshal(jsonBytes, product)
	fmt.Println(product)
}
