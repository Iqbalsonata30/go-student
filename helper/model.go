package helper

import (
	"github.com/iqbalsonata30/go-student/model/domain"
	"github.com/iqbalsonata30/go-student/model/web"
)

func EntityToCreateStudentResponse(student *domain.Student) *web.CreateStudentResponse {
	return &web.CreateStudentResponse{
		ID: student.ID,
	}
}
func EntityToCreateUserResponse(user *domain.User) *web.CreateUserResponse {
	return &web.CreateUserResponse{
		ID: user.ID,
	}

}

func EntityToResponse(student *domain.Student) *web.StudentResponse {
	return &web.StudentResponse{
		ID:             student.ID,
		Name:           student.Name,
		IdentityNumber: student.IdentityNumber,
		Gender:         student.Gender,
		Major:          student.Major,
		Class:          student.Class,
		Religion:       student.Religion,
		CreatedAt:      student.CreatedAt,
		UpdatedAt:      student.UpdatedAt,
	}

}

func EntityToResponses(students []domain.Student) []web.StudentResponse {
	studentResponses := make([]web.StudentResponse, 0)

	for _, student := range students {
		studentResponses = append(studentResponses, *EntityToResponse(&student))
	}
	return studentResponses

}
