package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// Table Test concept:
// We as programmers provide a slice, which contains parameters and expectation for the unit test result
// And then we iterate the slice to create _dynamic_ sub tests

func TestTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorldEko",
			request:  "Eko",
			expected: "Hello Eko",
		},
		{
			name:     "HelloWorldKurniawan",
			request:  "Kurniawan",
			expected: "Hello Kurniawan",
		},
		{
			name:     "HelloWorldWendy",
			request:  "Wendy",
			expected: "Hello Wendy",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(
				t,
				test.expected,
				result,
				fmt.Sprintf("Result must be '%s'", test.expected),
			)
		})
	}
}
