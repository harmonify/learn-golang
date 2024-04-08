package helper

import (
	"fmt"
	"os"
	"testing"
)

// Function named `TestMain` will be executed automatically by Golang
// We can use this to manage the unit test lifecycle
// We can only have one `TestMain` function per package
func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TESTS")

	// Run the unit tests
	code := m.Run()

	// after
	fmt.Println("AFTER UNIT TESTS")
	os.Exit(code)
}
