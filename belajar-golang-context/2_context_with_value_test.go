package belajar_golang_context

import (
	"context"
	"fmt"
	"testing"
)

func TestContextWithValue(t *testing.T) {
	// Contexts are immutable, its value cannot be changed
	contextA := context.Background()

	// Calling WithValue method will create a new child context instead of modifying the original context
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextF.Value("f")) // "F", found
	fmt.Println(contextF.Value("c")) // "C", found from its parent
	fmt.Println(contextF.Value("b")) // nil, it can only traverse to the up- most context
}
