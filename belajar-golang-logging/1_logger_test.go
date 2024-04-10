package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()

	logger.Println("Hello World")
	logger.Println("Hello World")
	logger.Println("Hello World")
}
