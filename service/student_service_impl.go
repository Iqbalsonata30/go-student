package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/iqbalsonata30/go-student/exception"
	"github.com/iqbalsonata30/go-student/helper"
	"github.com/iqbalsonata30/go-student/model/domain"
	"github.com/iqbalsonata30/go-student/model/web"
	"github.com/iqbalsonata30/go-student/repository"
)

type StudentServiceImpl struct {
	Repository repository.StudentRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewStudentService(repository repository.StudentRepository, DB *sql.DB, validator *validator.Validate) StudentService {
	return &StudentServiceImpl{
		Repository: repository,
		DB:         DB,
		Validate:   validator,
	}
}

func (s *StudentServiceImpl) Create(ctx context.Context, req web.StudentRequest) (*web.CreateStudentResponse, error) {
	err := s.Validate.Struct(req)
	if err != nil {
		return nil, err
	}
	student := domain.Student{
		Name:           req.Name,
		IdentityNumber: req.IdentityNumber,
		Gender:         req.Gender,
		Major:          req.Major,
		Class:          req.Class,
		Religion:       req.Religion,
	}
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}
	res, err := s.Repository.Save(ctx, tx, student)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return helper.EntityToCreateResponse(res), nil
}

func (s *StudentServiceImpl) FindAll(ctx context.Context) ([]web.StudentResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}
	res, err := s.Repository.FindAll(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return helper.EntityToResponses(res), nil

}

func (s *StudentServiceImpl) FindById(ctx context.Context, id string) (*web.StudentResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}
	sID, err := uuid.Parse(id)
	if err != nil {
		panic(exception.NewBadRequestError("Invalid student id"))
	}
	res, err := s.Repository.FindById(ctx, tx, sID)
	if err != nil {
		tx.Rollback()
		panic(exception.NewNotFoundError(err.Error()))
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return helper.EntityToResponse(res), nil
}

func (s *StudentServiceImpl) DeleteById(ctx context.Context, id string) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	sID, err := uuid.Parse(id)
	if err != nil {
		panic(exception.NewBadRequestError("Invalid student id"))
	}
	err = s.Repository.DeleteById(ctx, tx, sID)
	if err != nil {
		tx.Rollback()
		panic(exception.NewNotFoundError(err.Error()))
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (s *StudentServiceImpl) UpdateById(ctx context.Context, id string, req web.StudentRequest) (*web.StudentResponse, error) {
	err := s.Validate.Struct(req)
	if err != nil {
		return nil, err
	}
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}
	student := domain.Student{
		Name:           req.Name,
		IdentityNumber: req.IdentityNumber,
		Gender:         req.Gender,
		Major:          req.Major,
		Class:          req.Class,
		Religion:       req.Religion,
	}

	sID, err := uuid.Parse(id)
	if err != nil {
		panic(exception.NewBadRequestError("Invalid student id"))
	}

	res, err := s.Repository.UpdateById(ctx, tx, sID, student)
	if err != nil {
		tx.Rollback()
		panic(exception.NewNotFoundError(err.Error()))
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return helper.EntityToResponse(res), nil

}
