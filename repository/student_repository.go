package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/iqbalsonata30/go-student/model/domain"
)

type StudentRepository interface {
	Save(context.Context, *sql.Tx, domain.Student) (*domain.Student, error)
	FindById(context.Context, *sql.Tx, uuid.UUID) (*domain.Student, error)
	FindAll(context.Context, *sql.Tx) ([]domain.Student, error)
	UpdateById(context.Context, *sql.Tx, uuid.UUID, domain.Student) (*domain.Student, error)
	DeleteById(context.Context, *sql.Tx, uuid.UUID) error
}
