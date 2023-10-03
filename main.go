package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/iqbalsonata30/go-student/app"
	"github.com/iqbalsonata30/go-student/controller"
	"github.com/iqbalsonata30/go-student/repository"
	"github.com/iqbalsonata30/go-student/service"
)

func main() {
	store, err := app.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	studentRepository := repository.NewRepositoryStudent()
	studentService := service.NewStudentService(studentRepository, store.DB, validate)
	studentController := controller.NewStudentContoller(studentService)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, store.DB)
	userController := controller.NewUserController(userService)

	router := app.NewRouter(studentController, userController)

	log.Println("server running on port:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
