package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectDestination(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello Redirect")
}

func RedirectFrom(writer http.ResponseWriter, request *http.Request) {
	// logic
	http.Redirect(writer, request, "/redirect-destination", http.StatusTemporaryRedirect)
}

func RedirectOut(writer http.ResponseWriter, request *http.Request) {
	// logic
	http.Redirect(writer, request, "https://www.programmerzamannow.com/", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-destination", RedirectDestination)
	mux.HandleFunc("/redirect-out", RedirectOut)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
