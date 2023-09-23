package repository

import (
	"context"

	"github.com/iqbalsonata30/go-student/model/domain"
)

type StudentRepositoryImpl struct{}

func NewRepositoryStudent() StudentRepository {
	return &StudentRepositoryImpl{}
}

func (r *StudentRepositoryImpl) Save(ctx context.Context, student *domain.Student) error {
	return nil
}
func (r *StudentRepositoryImpl) FindById(ctx context.Context, id string) (*domain.Student, error) {
	return nil, nil
}
