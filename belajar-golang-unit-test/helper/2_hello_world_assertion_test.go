package helper

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// if assert is failed, it will call t.Fail
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Assert")
	assert.Equal(t, "Hello Assert", result, "Result must be 'Hello Assert'")
	t.Log("TestHelloWorldAssert done")
}

// if require is failed, it will call t.FailNow
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Require")
	require.Equal(t, "Hello Require", result, "Result must be 'Hello Require'")
	t.Log("TestHelloWorldRequire done") // will not be executed if this test is failed
}
