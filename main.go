package main

import (
	"log"

	"github.com/iqbalsonata30/go-student/app"
)

func main() {
	db, err := app.NewDB()
	if err != nil {
		log.Fatal(err)
	}
}
