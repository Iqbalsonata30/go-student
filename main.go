package main

import (
	"log"

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
}
