package repository

import (
	"context"
	"database/sql"

	"github.com/iqbalsonata30/go-student/model/domain"
)

type StudentRepository interface {
	Save(context.Context, *sql.Tx, domain.Student) (*domain.Student, error)
	FindById(context.Context, *sql.Tx, string) (*domain.Student, error)
	FindAll(context.Context, *sql.Tx) (*[]domain.Student, error)
	Update(context.Context, *sql.Tx, string) (*domain.Student, error)
	Delete(context.Context, *sql.Tx, string) error
}
