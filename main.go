package main

import (
	"log"
	"net/http"

	"github.com/iqbalsonata30/go-student/app"
)

func main() {
	store, err := app.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	router := app.NewRouter(store.DB)
	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	log.Printf("server running on port %s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
