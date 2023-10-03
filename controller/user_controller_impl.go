package controller

import (
	"net/http"

	"github.com/iqbalsonata30/go-student/exception"
	"github.com/iqbalsonata30/go-student/helper"
	"github.com/iqbalsonata30/go-student/model/web"
	"github.com/iqbalsonata30/go-student/service"
	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &UserControllerImpl{
		service: service,
	}
}

func (c *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req := web.UserRequest{}
	helper.BodyRequest(r, &req)
	user, err := c.service.Create(r.Context(), req)
	if err != nil {
		exception.ErrorHandler(w, r, err)
		return
	}
	apiResp := web.ApiResponse{
		StatusCode: http.StatusCreated,
		Message:    "Succesfully created user",
		Data:       user,
	}
	helper.JSONEncode(w, http.StatusCreated, apiResp)
}
