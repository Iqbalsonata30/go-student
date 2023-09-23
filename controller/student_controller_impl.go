package controller

import (
	"log"
	"net/http"

	"github.com/iqbalsonata30/go-student/helper"
	"github.com/iqbalsonata30/go-student/model/web"
	"github.com/iqbalsonata30/go-student/service"
	"github.com/julienschmidt/httprouter"
)

type StudentControllerImpl struct {
	Service service.StudentService
}

func NewStudentContoller(service service.StudentService) StudentController {
	return &StudentControllerImpl{
		Service: service,
	}
}

func (c *StudentControllerImpl) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req := web.StudentRequest{}
	helper.BodyRequest(r, &req)
	student, err := c.Service.Create(r.Context(), req)
	if err != nil {
		log.Fatal(err)
	}
	res := web.ApiResponse{
		StatusCode: http.StatusCreated,
		Message:    "Student has been added sucesfully",
		Data:       student,
	}
	helper.JSONEncode(w, http.StatusCreated, res)
}
