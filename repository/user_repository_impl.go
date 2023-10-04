package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/iqbalsonata30/go-student/helper"
	"github.com/iqbalsonata30/go-student/model/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}

}

func (r *UserRepositoryImpl) Save(ctx context.Context, db *sql.Tx, req domain.User) (*domain.User, error) {
	query := `INSERT INTO user_account(id,username,password) VALUES($1,$2,$3);`
	saltedPassword, err := helper.HashAndSalted([]byte(req.Password))
	if err != nil {
		return nil, bcrypt.ErrPasswordTooLong
	}
	id := uuid.New()
	_, err = db.ExecContext(ctx, query, id, req.Username, saltedPassword)
	if err != nil {
		return nil, err
	}
	req.ID = id
	return &req, nil
}
