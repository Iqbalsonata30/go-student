package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Create(http.ResponseWriter, *http.Request, httprouter.Params)
	Login(http.ResponseWriter, *http.Request, httprouter.Params)
}
