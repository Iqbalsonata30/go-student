package controller

import (
	"net/http"

	"github.com/iqbalsonata30/go-student/exception"
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
		exception.ErrorHandler(w, r, err)
		return
	}
	res := web.ApiResponse{
		StatusCode: http.StatusCreated,
		Message:    "Student has been added succesfully",
		Data:       student,
	}
	helper.JSONEncode(w, http.StatusCreated, res)
}

func (c *StudentControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	students, err := c.Service.FindAll(r.Context())
	if err != nil {
		exception.ErrorHandler(w, r, err)
		return
	}
	res := web.ApiResponse{
		StatusCode: http.StatusOK,
		Message:    "Success get all data students",
		Data:       students,
	}
	helper.JSONEncode(w, http.StatusOK, res)

}

func (c *StudentControllerImpl) FindById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	student, err := c.Service.FindById(r.Context(), id)
	if err != nil {
		exception.ErrorHandler(w, r, err)
		return
	}
	res := web.ApiResponse{
		StatusCode: http.StatusOK,
		Message:    "Success get data student",
		Data:       student,
	}
	helper.JSONEncode(w, http.StatusOK, res)
}
