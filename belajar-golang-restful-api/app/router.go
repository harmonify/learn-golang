package app

import (
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/exception"
	"belajar-golang-restful-api/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		_, err := w.Write([]byte("Hello World"))
		helper.PanicIfError(err)
	})
	router.GET("/api/categories", categoryController.FindAll)
	router.POST("/api/categories", categoryController.Create)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
