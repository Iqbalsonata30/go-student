package exception

import (
	"net/http"

	"github.com/iqbalsonata30/go-student/helper"
	"github.com/iqbalsonata30/go-student/model/web"
)

type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}

func NotFoundPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiResp := web.ApiError{
			StatusCode: http.StatusNotFound,
			Error:      "404 Page not found ",
		}
		helper.JSONEncode(w, http.StatusNotFound, apiResp)
	}

}
