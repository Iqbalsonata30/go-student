package app

import (
	"database/sql"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/iqbalsonata30/go-student/controller"
	"github.com/iqbalsonata30/go-student/exception"
	"github.com/iqbalsonata30/go-student/helper"
	"github.com/iqbalsonata30/go-student/repository"
	"github.com/iqbalsonata30/go-student/service"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(db *sql.DB) *httprouter.Router {
	validate := validator.New(validator.WithRequiredStructEnabled())
	studentRepository := repository.NewRepositoryStudent()
	studentService := service.NewStudentService(studentRepository, db, validate)
	studentController := controller.NewStudentContoller(studentService)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db)
	userController := controller.NewUserController(userService)

	router := httprouter.New()
	studentRouter(router, studentController)
	userRouter(router, userController)

	router.NotFound = exception.NotFoundPage()
	router.PanicHandler = exception.ErrorHandler
	return router
}

func studentRouter(router *httprouter.Router, sc controller.StudentController) {
	router.POST("/api/v1/students", middleware(sc.Create))
	router.GET("/api/v1/students", sc.FindAll)
	router.GET("/api/v1/students/:id", sc.FindById)
	router.DELETE("/api/v1/students/:id", middleware(sc.DeleteById))
	router.PUT("/api/v1/students/:id", middleware(sc.UpdateById))

}

func userRouter(router *httprouter.Router, uc controller.UserController) {
	router.POST("/api/v1/users", uc.Create)
	router.POST("/api/v1/login", uc.Login)
}

func middleware(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		tokenString := r.Header.Get("X-API-Key")
		token, err := helper.ValidateToken(tokenString)
		if err != nil {
			panic(exception.NewUnauthorizedError(err.Error()))
		}
		if !token.Valid {
			panic(exception.NewUnauthorizedError("Unauthorized"))
		}
		handler(w, r, p)
	}
}
