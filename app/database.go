package app

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresDB() (*PostgresDB, error) {
	if os.Getenv("DB_USER") == "" {
		return nil, fmt.Errorf("DB USER environment must be set")
	}
	if os.Getenv("DB_DATABASE") == "" {
		return nil, fmt.Errorf("Database name environment must be set")
	}
	if os.Getenv("DB_HOST") == "" {
		return nil, fmt.Errorf("DB Host environment must be set")
	}
	if os.Getenv("DB_PORT") == "" {
		return nil, fmt.Errorf("DB Port name environment must be set")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		return nil, fmt.Errorf("DB PASSWORD  environment must be set")
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresDB{
		DB: db,
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
            religion varchar(15) NOT NULL,
            created_at timestamp default CURRENT_TIMESTAMP,
            updated_at timestamp default CURRENT_TIMESTAMP
    );`
	_, err := s.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
