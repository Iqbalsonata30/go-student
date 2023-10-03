package service

import (
	"context"
	"database/sql"

	"github.com/iqbalsonata30/go-student/helper"
	"github.com/iqbalsonata30/go-student/model/domain"
	"github.com/iqbalsonata30/go-student/model/web"
	"github.com/iqbalsonata30/go-student/repository"
)

type UserServiceImpl struct {
	Repository repository.UserRepository
	DB         *sql.DB
}

func NewUserService(repo repository.UserRepository, db *sql.DB) UserService {
	return &UserServiceImpl{repo, db}
}

func (s *UserServiceImpl) Create(ctx context.Context, user web.UserRequest) (*web.CreateUserResponse, error) {
	req := domain.User{
		Username: user.Username,
		Password: user.Password,
	}
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}

	res, err := s.Repository.Save(ctx, tx, req)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return helper.EntityToCreateUserResponse(res), nil

}
