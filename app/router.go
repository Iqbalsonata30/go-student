package app

import (
	"net/http"

	"github.com/iqbalsonata30/go-student/controller"
	"github.com/iqbalsonata30/go-student/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(sc controller.StudentController, uc controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/v1/students", sc.Create)
	router.GET("/api/v1/students", sc.FindAll)
	router.GET("/api/v1/students/:id", sc.FindById)
	router.DELETE("/api/v1/students/:id", sc.DeleteById)
	router.PUT("/api/v1/students/:id", sc.UpdateById)

	router.POST("/api/v1/users", uc.Create)

	router.NotFound = exception.NotFoundPage()
	router.PanicHandler = exception.ErrorHandler

	return router
}

func handleFunc(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
