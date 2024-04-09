package main_test

import (
	"embed"
	"fmt"
	"testing"
)

// https://golang.org/pkg/path/#Match
//
//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if entry.IsDir() {
			continue
		}
		fmt.Println(entry.Name())
		file, _ := path.ReadFile("files/" + entry.Name())
		fmt.Println(string(file))
	}
}
