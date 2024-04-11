package main

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/constant"
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/middleware"
	"belajar-golang-restful-api/repository"
	"belajar-golang-restful-api/service"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB(
		constant.DATABASE_HOST,
		constant.DATABASE_PORT,
		constant.DATABASE_NAME,
		constant.DATABASE_USERNAME,
		constant.DATABASE_PASSWORD,
	)
	defer db.Close()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	defer server.Close()

	fmt.Println("Server running at http://localhost:3000")
	helper.PanicIfError(server.ListenAndServe())
}
