package middleware

import (
	"net/http"
	// "github.com/iqbalsonata30/go-student/helper"
	// "github.com/iqbalsonata30/go-student/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// key := r.Header.Get("X-API-Key")
	// if key != "" {
	m.Handler.ServeHTTP(w, r)
	// }
	// apiResp := web.ApiError{
	// StatusCode: http.StatusUnauthorized,
	// Error:      "Unauthorized",
	// }
	// helper.JSONEncode(w, http.StatusUnauthorized, apiResp)
}
