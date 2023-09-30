package service

import (
	"context"

	"github.com/iqbalsonata30/go-student/model/web"
)

type StudentService interface {
	Create(context.Context, web.StudentRequest) (*web.CreateStudentResponse, error)
	FindAll(context.Context) ([]web.StudentResponse, error)
	FindById(context.Context, string) (*web.StudentResponse, error)
	DeleteById(context.Context, string) error
}
