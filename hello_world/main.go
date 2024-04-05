//go:build (prod && ignore) || (dev && ignore) || test
// +build prod,ignore dev,ignore test

package main

import (
	"fmt" // A package in the Go standard library.
	// "io/ioutil" // Implements some I/O utility functions.
	// m "math"    // Math library with local alias m.
	// "net/http"  // Yes, a web server!
	// "os"        // OS functions like working with the file system
	// "strconv"   // String conversions.
	"example.com/harmonify/domain"
	"rsc.io/quote"
)

// A function definition. Main is special. It is the entry point for the
// executable program. Love it or hate it, Go uses brace brackets.
func main() {
	// Println outputs a line to stdout.
	// It comes from the package fmt.
	fmt.Println("Hello world!")

	domain.SayHello("Harmonify")

	fmt.Println(quote.Go())

	// Call another function within this package.
	// beyondHello()
}
