package helper

import (
	"github.com/iqbalsonata30/go-student/model/domain"
	"github.com/iqbalsonata30/go-student/model/web"
)

func EntityToResponse(domain *domain.Student) *web.StudentResponse {
	return &web.StudentResponse{
		ID: domain.ID,
	}
}
