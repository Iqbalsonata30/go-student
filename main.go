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

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	log.Printf("server running on port %s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
