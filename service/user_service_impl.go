package service

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iqbalsonata30/go-student/exception"
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
		panic(exception.NewBadRequestError(err.Error()))
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return helper.EntityToCreateUserResponse(res), nil

}

func (s *UserServiceImpl) Authenticate(ctx context.Context, user web.UserRequest) (*web.UserLoginResponse, error) {
	req := domain.User{
		Username: user.Username,
		Password: user.Password,
	}
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}

	res, err := s.Repository.Authenticate(ctx, tx, req)
	if err != nil {
		tx.Rollback()
		panic(exception.NewBadRequestError(err.Error()))
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	token, err := s.createToken(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return helper.EntityToUserLoginResponse(res, token), nil
}

func (s *UserServiceImpl) createToken(user domain.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"expiresAt": time.Now().Add(time.Hour).Unix(),
		"username":  user.Username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
