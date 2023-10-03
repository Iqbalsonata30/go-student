package service

import (
	"context"

	"github.com/iqbalsonata30/go-student/model/web"
)

type UserService interface {
	Create(context.Context, web.UserRequest) (*web.CreateUserResponse, error)
}
