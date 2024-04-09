package main_test

import (
	_ "embed"
	"fmt"
	"testing"
)

// There should not be any space between the comment tag '//' and 'go:embed'

//go:embed version.txt
var version string

func TestEmbedString(t *testing.T) {
	fmt.Println(version)
}
