package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// logic web
		fmt.Fprint(writer, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	defer server.Close()

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
