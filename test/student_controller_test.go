package test

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/iqbalsonata30/go-student/app"
	"github.com/iqbalsonata30/go-student/controller"
	"github.com/iqbalsonata30/go-student/repository"
	"github.com/iqbalsonata30/go-student/service"
)

func SetupPostgresql() *sql.DB {
	connStr := "user=postgres  password=secret dbname=go_student sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db

}

func SetupNewRouter(db *sql.DB) http.Handler {
	validate := validator.New(validator.WithRequiredStructEnabled())
	repository := repository.NewRepositoryStudent()

	service := service.NewStudentService(repository, db, validate)
	controller := controller.NewStudentContoller(service)
	router := app.NewRouter(controller)

	return router
}

func TruncateDB(db *sql.DB) {
	db.Exec("truncate student;")
}

func TestCreateStudentSuccess(t *testing.T) {
	db := SetupPostgresql()
	TruncateDB(db)
	router := SetupNewRouter(db)

	reqBody := strings.NewReader(`{   
        "name":"Iqbal Sonata",
        "identityNumber":2110127263323,
        "gender":"Male",
        "major":"Computer Science",
        "class":"3-PTK-1",
        "religion":"Islam"
    }`)
	req := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/students", reqBody)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	res := rec.Result()
	if res.Status != "201 Created" {
		t.Fatalf("the status code should've 201 Created but got : %v", res.Status)
	}
	body, _ := io.ReadAll(res.Body)
	var resBody map[string]any
	json.Unmarshal(body, &resBody)

	if int(resBody["statusCode"].(float64)) != 201 {
		t.Fatalf("the status code should've 201 but got : %d", int(resBody["statusCode"].(float64)))
	}
	if resBody["message"] != "Student has been added succesfully" {
		t.Fatalf(`message should've "Student has been added succesfully" but got : %v`, resBody["message"])
	}
	if resBody["data"].(map[string]any)["id"] == uuid.Invalid.String() {
		t.Fatal("the id is not valid")
	}
}

// unc TestCreateStudentFailed(t *testing.T) {
// db := SetupPostgresql()
// TruncateDB(db)
// router := SetupNewRouter(db)
