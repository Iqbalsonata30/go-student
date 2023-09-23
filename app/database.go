package app

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db  *sql.DB
	ctx context.Context
}

func NewPostgresDB() (*PostgresDB, error) {
	connStr := "user=postgres  password=secret dbname=go_student sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	return &PostgresDB{
		db:  db,
		ctx: ctx,
	}, nil
}

func (s *PostgresDB) Init() error {
	query := `CREATE TABLE IF NOT EXISTS student(
            id uuid PRIMARY KEY,
            name varchar(100) NOT NULL,
            identity_number BIGINT UNIQUE NOT NULL,
            gender varchar(20) NOT NULL,
            major varchar(50) NOT NULL,
            class varchar(10) NOT NULL,
            religion varchar(15) NOT NULL
    );`
	_, err := s.db.ExecContext(s.ctx, query)
	if err != nil {
		return err
	}
	return nil
}
