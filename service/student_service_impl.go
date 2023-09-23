package service

import (
	"context"
	"database/sql"

	"github.com/iqbalsonata30/go-student/helper"
	"github.com/iqbalsonata30/go-student/model/domain"
	"github.com/iqbalsonata30/go-student/model/web"
	"github.com/iqbalsonata30/go-student/repository"
)

type StudentServiceImpl struct {
	Repository repository.StudentRepositoryImpl
	DB         *sql.DB
}

func NewStudentService(repository repository.StudentRepositoryImpl, DB *sql.DB) StudentService {
	return &StudentServiceImpl{
		Repository: repository,
		DB:         DB,
	}
}

func (s *StudentServiceImpl) Create(ctx context.Context, req web.StudentRequest) (*web.StudentResponse, error) {
	student := domain.Student{
		Name:           req.Name,
		IdentityNumber: req.IdentityNumber,
		Gender:         req.Gender,
		Major:          req.Major,
		Class:          req.Class,
	}
	tx, err := s.DB.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	res, err := s.Repository.Save(ctx, tx, student)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return helper.EntityToResponse(res), nil
}
