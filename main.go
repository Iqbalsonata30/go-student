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
	repository := repository.NewRepositoryStudent()
	service := service.NewStudentService(repository, store.DB, validate)
	controller := controller.NewStudentContoller(service)
	router := app.NewRouter(controller)

	log.Println("server running on port:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
