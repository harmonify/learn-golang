package main

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/middleware"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {
	server := InitializedServer()

	fmt.Println("Server running at http://localhost:3000")
	helper.PanicIfError(server.ListenAndServe())
}
