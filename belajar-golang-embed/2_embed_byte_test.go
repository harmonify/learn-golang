package main_test

import (
	_ "embed"
	"io/fs"
	"os"
	"testing"
)

//go:embed github.png
var logo []byte

func TestEmbedByte(t *testing.T) {
	err := os.WriteFile("github-new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
