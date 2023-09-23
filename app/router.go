package app

import (
	"github.com/iqbalsonata30/go-student/controller"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(controller controller.StudentController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/v1/students", controller.Create)
	return router
}
