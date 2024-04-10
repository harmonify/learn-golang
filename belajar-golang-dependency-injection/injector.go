//go:build wireinject
// +build wireinject

package main

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/constant"
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/middleware"
	"belajar-golang-restful-api/repository"
	"belajar-golang-restful-api/service"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var db = app.NewDB(
	constant.DATABASE_HOST,
	constant.DATABASE_PORT,
	constant.DATABASE_NAME,
	constant.DATABASE_USERNAME,
	constant.DATABASE_PASSWORD,
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	service.NewCategoryService,
	controller.NewCategoryController,
)

func InitializedServer() *http.Server {
	wire.Build(
		wire.Value(db),
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}
