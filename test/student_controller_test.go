package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/iqbalsonata30/go-student/app"
	"github.com/iqbalsonata30/go-student/model/domain"
	"github.com/iqbalsonata30/go-student/repository"
)

func SetupPostgresql() *sql.DB {
	if os.Getenv("DB_USER") == "" {
		panic("DB USER environment must be set")
	}
	if os.Getenv("DB_DATABASE") == "" {
		panic("Database name environment must be set")
	}
	if os.Getenv("DB_HOST") == "" {
		panic("DB Host environment must be set")
	}
	if os.Getenv("DB_PORT") == "" {
		panic("DB Port name environment must be set")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		panic("DB PASSWORD  environment must be set")
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db

}

func SetupNewRouter(db *sql.DB) http.Handler {
	router := app.NewRouter(db)

	return router
}

func TruncateDB(db *sql.DB) {
	db.Exec("truncate student;")
}

func registerUser(db *sql.DB, router http.Handler) {
	reqBody := strings.NewReader(`{
    "username":"iqbalsonata1",
    "password":"secret"
  }`)
	req := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/users", reqBody)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
}

func getToken(db *sql.DB, router http.Handler) string {
	reqBody := strings.NewReader(`{
      "username":"iqbalsonata1",
      "password":"secret"
    }`)
	req := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/login", reqBody)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	res := rec.Result()
	body, _ := io.ReadAll(res.Body)
	var resBody map[string]any
	json.Unmarshal(body, &resBody)
	token := resBody["data"].(map[string]any)["token"]
	return token.(string)
}

func TestCreateStudent(t *testing.T) {
	t.Run("create student success", func(t *testing.T) {
		db := SetupPostgresql()
		TruncateDB(db)
		router := SetupNewRouter(db)
		registerUser(db, router)
		token := getToken(db, router)

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
		req.Header.Set("X-API-Key", token)
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
	})
	t.Run("create student with validation required", func(t *testing.T) {
		db := SetupPostgresql()
		TruncateDB(db)
		router := SetupNewRouter(db)
		registerUser(db, router)
		token := getToken(db, router)
		reqBody := strings.NewReader(`{}`)

		req := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/students", reqBody)
		rec := httptest.NewRecorder()
		req.Header.Set("X-API-Key", token)
		router.ServeHTTP(rec, req)
		res := rec.Result()
		if res.Status != "400 Bad Request" {
			t.Fatalf("status code should've  400 Bad Request but got : %v", res.Status)
		}

		body, _ := io.ReadAll(res.Body)
		var resBody map[string]any
		json.Unmarshal(body, &resBody)

		if int(resBody["statusCode"].(float64)) != 400 {
			t.Fatalf("status code should've  400 but got : %v", int(resBody["statusCode"].(float64)))
		}

		for _, v := range resBody["error"].([]any) {
			if v.(map[string]any)["message"] != "this field is required" {
				t.Fatalf("the validation should be this fields is required but got : %v", v.(map[string]any)["message"])
			}
		}

	})
}

func TestGetAllStudents(t *testing.T) {
	db := SetupPostgresql()
	router := SetupNewRouter(db)
	tx, _ := db.Begin()
	studentRepository := repository.NewRepositoryStudent()
	student, _ := studentRepository.Save(context.Background(), tx, domain.Student{
		Name:           "test",
		IdentityNumber: 2122,
		Gender:         "male",
		Major:          "test",
		Class:          "test",
		Religion:       "test",
	})
	tx.Commit()

	req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/students", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	res := rec.Result()
	if res.StatusCode != 200 {
		t.Fatalf("status code should've 200 but got : %v", res.StatusCode)
	}

	body, _ := io.ReadAll(res.Body)
	var resBody map[string]any
	json.Unmarshal(body, &resBody)

	if resBody["message"] != "Success get all data students" {
		t.Fatalf("message should be Success get all data students but got : %v", resBody["message"])
	}
	var students = resBody["data"].([]any)
	studentRes := students[0].(map[string]any)
	if studentRes["name"] != student.Name {
		t.Fatalf(fmt.Sprintf("name should be %v but got %v", studentRes["name"], student.Name))
	}
	if int(studentRes["identityNumber"].(float64)) != student.IdentityNumber {
		t.Fatalf(fmt.Sprintf("student's identityNumber should be %v but got %v", studentRes["identityNumber"], student.IdentityNumber))
	}
	if studentRes["gender"] != student.Gender {
		t.Fatalf(fmt.Sprintf("student's gender  should be %v but got %v", studentRes["gender"], student.Gender))
	}
	if studentRes["major"] != student.Major {
		t.Fatalf(fmt.Sprintf("student's major  should be %v but got %v", studentRes["major"], student.Major))
	}
	if studentRes["class"] != student.Class {
		t.Fatalf(fmt.Sprintf("student's class  should be %v but got %v", studentRes["class"], student.Class))
	}
	if studentRes["religion"] != student.Religion {
		t.Fatalf(fmt.Sprintf("student's religion  should be %v but got %v", studentRes["religion"], student.Religion))
	}
}

func TestFindStudentByID(t *testing.T) {
	t.Run("find by id succesfully", func(t *testing.T) {
		db := SetupPostgresql()
		TruncateDB(db)
		router := SetupNewRouter(db)
		tx, _ := db.Begin()
		studentRepository := repository.NewRepositoryStudent()
		student, _ := studentRepository.Save(context.Background(), tx, domain.Student{
			Name:           "test",
			IdentityNumber: 2122,
			Gender:         "male",
			Major:          "test",
			Class:          "test",
			Religion:       "test",
		})
		tx.Commit()
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:3000/api/v1/students/%v", student.ID), nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		res := rec.Result()
		if res.StatusCode != 200 {
			t.Fatalf("the status code should've 200 but got : %d", res.StatusCode)
		}
		body, _ := io.ReadAll(res.Body)
		var resBody map[string]any
		json.Unmarshal(body, &resBody)

		if resBody["message"] != "Success get data student" {
			t.Fatalf("message should be Success get data students but got : %v", resBody["message"])
		}
		if ok := json.Valid([]byte(body)); !ok {
			t.Fatal("the result data is not valid json")
		}
	})

	t.Run("find by id invalid uuid", func(t *testing.T) {
		db := SetupPostgresql()
		TruncateDB(db)
		router := SetupNewRouter(db)
		req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/students/asda", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		res := rec.Result()
		if res.StatusCode != 400 {
			t.Fatalf("the status code should've 400 but got : %d", res.StatusCode)
		}
		body, _ := io.ReadAll(res.Body)
		var resBody map[string]any
		json.Unmarshal(body, &resBody)

		if resBody["error"] != "Invalid student id" {
			t.Fatalf("message should be Success get data students but got : %v", resBody["error"])
		}
		if ok := json.Valid([]byte(body)); !ok {
			t.Fatal("the result data is not valid json")
		}
	})

	t.Run("find by id not found", func(t *testing.T) {
		db := SetupPostgresql()
		TruncateDB(db)
		router := SetupNewRouter(db)
		req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/students/22a5f8df-0460-4fa8-9db3-95cac91f6f86", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		res := rec.Result()
		if res.StatusCode != 404 {
			t.Fatalf("the status code should've 404 but got : %d", res.StatusCode)
		}
		body, _ := io.ReadAll(res.Body)
		var resBody map[string]any
		json.Unmarshal(body, &resBody)

		if resBody["error"] != "student is not found." {
			t.Fatalf("message should be Success get data students but got : %v", resBody["error"])
		}
		if ok := json.Valid([]byte(body)); !ok {
			t.Fatal("the result data is not valid json")
		}
	})
}

func TestDeleteStudentById(t *testing.T) {
	t.Run("Delete student sucesfully", func(t *testing.T) {
		db := SetupPostgresql()
		router := SetupNewRouter(db)
		registerUser(db, router)
		token := getToken(db, router)

		tx, _ := db.Begin()
		studentRepository := repository.NewRepositoryStudent()
		student, _ := studentRepository.Save(context.Background(), tx, domain.Student{
			Name:           "test",
			IdentityNumber: 2122,
			Gender:         "male",
			Major:          "test",
			Class:          "test",
			Religion:       "test",
		})
		tx.Commit()
		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:3000/api/v1/students/%s", student.ID), nil)
		rec := httptest.NewRecorder()
		req.Header.Set("X-API-Key", token)
		router.ServeHTTP(rec, req)
		res := rec.Result()
		if res.StatusCode != 200 {
			t.Fatalf("status code should've 200 but got : %d", res.StatusCode)
		}
		body, _ := io.ReadAll(res.Body)
		var resBody map[string]any
		json.Unmarshal(body, &resBody)
		if resBody["message"] != "Success delete the student" {
			t.Fatalf("message is not valid, your message : %s", resBody["message"])

		}
	})

	t.Run("delete by id invalid uuid", func(t *testing.T) {
		db := SetupPostgresql()
		TruncateDB(db)
		router := SetupNewRouter(db)
		registerUser(db, router)
		token := getToken(db, router)

		req := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/v1/students/asda", nil)
		rec := httptest.NewRecorder()
		req.Header.Set("X-API-Key", token)
		router.ServeHTTP(rec, req)
		res := rec.Result()
		if res.StatusCode != 400 {
			t.Fatalf("the status code should've 400 but got : %d", res.StatusCode)
		}
		body, _ := io.ReadAll(res.Body)
		var resBody map[string]any
		json.Unmarshal(body, &resBody)

		if resBody["error"] != "Invalid student id" {
			t.Fatalf("message should be Success delete the  student but got : %v", resBody["error"])
		}
		if ok := json.Valid([]byte(body)); !ok {
			t.Fatal("the result data is not valid json")
		}
	})

	t.Run("student id is  not found", func(t *testing.T) {
		db := SetupPostgresql()
		TruncateDB(db)
		router := SetupNewRouter(db)
		registerUser(db, router)
		token := getToken(db, router)

		req := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/v1/students/22a5f8df-0460-4fa8-9db3-95cac91f6f86", nil)
		rec := httptest.NewRecorder()
		req.Header.Set("X-API-Key", token)
		router.ServeHTTP(rec, req)
		res := rec.Result()
		if res.StatusCode != 404 {
			t.Fatalf("the status code should've 404 but got : %d", res.StatusCode)
		}
		body, _ := io.ReadAll(res.Body)
		var resBody map[string]any
		json.Unmarshal(body, &resBody)

		if resBody["error"] != "student is not found" {
			t.Fatalf("message should be Success delete the student but got : %v", resBody["error"])
		}
		if ok := json.Valid([]byte(body)); !ok {
			t.Fatal("the result data is not valid json")
		}
	})

}

func TestUpdateStudent(t *testing.T) {

}

func TestNotFoundPage(t *testing.T) {
	db := SetupPostgresql()
	TruncateDB(db)
	router := SetupNewRouter(db)
	reqBody := strings.NewReader(`{}`)
	req := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/stundetsts", reqBody)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	res := rec.Result()
	if res.Status != "404 Not Found" {
		t.Fatalf("status code should've 404 but got : %v", res.Status)
	}
	body, _ := io.ReadAll(res.Body)
	var resBody map[string]any
	json.Unmarshal(body, &resBody)

	if int(resBody["statusCode"].(float64)) != 404 {
		t.Fatalf("status code should've  404 but got : %v", int(resBody["statusCode"].(float64)))

	}
	if resBody["error"].(any) != "404 Page not found " {
		t.Fatal("message should be 404 Page not found ")
	}

}
