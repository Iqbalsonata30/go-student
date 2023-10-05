package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iqbalsonata30/go-student/controller"
	"github.com/iqbalsonata30/go-student/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(sc controller.StudentController, uc controller.UserController) *httprouter.Router {
	router := httprouter.New()
	registerRoutes(router, sc, uc)
	return router
}

func registerRoutes(router *httprouter.Router, sc controller.StudentController, uc controller.UserController) {
	router.POST("/api/v1/students", verifyToken(sc.Create))
	router.GET("/api/v1/students", sc.FindAll)
	router.GET("/api/v1/students/:id", sc.FindById)
	router.DELETE("/api/v1/students/:id", verifyToken(sc.DeleteById))
	router.PUT("/api/v1/students/:id", verifyToken(sc.UpdateById))

	router.POST("/api/v1/users", uc.Create)
	router.POST("/api/v1/login", uc.Login)

	router.NotFound = exception.NotFoundPage()
	router.PanicHandler = exception.ErrorHandler
}

func verifyToken(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		tokenString := r.Header.Get("X-API-Key")
		_, err := validateToken(tokenString)
		if err != nil {
			panic(exception.NewUnauthorizedError("Unauthorized"))
		}
		handler(w, r, p)
	}
}

func validateToken(token string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})

}
