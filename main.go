package main

import (
	"log"
	"net/http"

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
	repository := repository.NewRepositoryStudent()
	service := service.NewStudentService(repository, store.DB)
	controller := controller.NewStudentContoller(service)
	router := app.NewRouter(controller)

	log.Fatal(http.ListenAndServe(":3000", router))
}
