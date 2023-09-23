package repository

import (
	"context"

	"github.com/iqbalsonata30/go-student/model/domain"
)

type StudentRepository interface {
	Save(context.Context, *domain.Student) error
	FindById(context.Context, string) (*domain.Student, error)
}
