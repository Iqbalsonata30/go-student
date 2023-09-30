package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type StudentController interface {
	Create(http.ResponseWriter, *http.Request, httprouter.Params)
	FindAll(http.ResponseWriter, *http.Request, httprouter.Params)
	FindById(http.ResponseWriter, *http.Request, httprouter.Params)
	DeleteById(http.ResponseWriter, *http.Request, httprouter.Params)
}
