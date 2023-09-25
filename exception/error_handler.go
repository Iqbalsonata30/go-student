package exception

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/iqbalsonata30/go-student/helper"
	"github.com/iqbalsonata30/go-student/model/web"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err any) {
	if validationError(w, r, err) {
		return
	}

	internalError(w, r, err)
}

func validationError(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		out := make([]web.Error, len(exception))
		for i, fe := range exception {
			out[i] = web.Error{Field: fe.Field(), Message: helper.WriteMsgForTag(fe.Tag())}
		}
		apiResp := web.ApiError{
			StatusCode: http.StatusBadRequest,
			Error:      out,
		}
		helper.JSONEncode(w, http.StatusBadRequest, apiResp)
		return true
	} else {
		return false
	}
}

func internalError(w http.ResponseWriter, r *http.Request, err any) {
	apiResp := web.ApiError{
		StatusCode: http.StatusInternalServerError,
		Error:      err,
	}
	helper.JSONEncode(w, http.StatusInternalServerError, apiResp)

}
