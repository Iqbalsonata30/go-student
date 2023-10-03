package repository

import (
	"context"
	"database/sql"

	"github.com/iqbalsonata30/go-student/model/domain"
)

type UserRepository interface {
	Save(context.Context, *sql.Tx, domain.User) (*domain.User, error)
}
