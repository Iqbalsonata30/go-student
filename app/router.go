package app

import (
	"github.com/iqbalsonata30/go-student/controller"
	"github.com/iqbalsonata30/go-student/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(controller controller.StudentController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/v1/students", controller.Create)

	router.NotFound = exception.NotFoundPage()
	router.PanicHandler = exception.ErrorHandler

	return router
}
