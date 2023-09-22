package app

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {
	connStr := "user=postgres  password=secret dbname=go_student sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
