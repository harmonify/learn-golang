package helper

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot be run on MacOS")
	}
	result := HelloWorld("Skip")
	assert.Equal(t, "Hello Skip", result, "Result must be 'Hello Skip'")
	t.Log("TestSkip done")
}
