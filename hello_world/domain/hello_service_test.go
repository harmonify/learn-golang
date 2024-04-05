package domain_test

import (
	"testing"

	"example.com/harmonify/domain"
)

func TestSayHello(t *testing.T) {
	result, err := domain.SayHello("World")
	if err != nil || result != "Hello World!" {
		t.Fatal(err)
		t.Log("result is ")
	}
}
