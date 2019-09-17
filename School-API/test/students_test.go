package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"../controllers"
	"../models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func RouterStudent() *mux.Router {
	router := mux.NewRouter()
	fmt.Println("http://localhost:8080")

	router.HandleFunc("/students", controllers.GetStudents).Methods("GET")
	router.HandleFunc("/student/{id}", controllers.GetStudent).Methods("GET")
	router.HandleFunc("/student/{id}", controllers.UpdateStudent).Methods("PUT")
	router.HandleFunc("/student/{id}", controllers.DeleteStudent).Methods("DELETE")

	return router
}

func TestStudentsEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	RouterStudent().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

func TestValidStudentIDEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/student/1", nil)
	response := httptest.NewRecorder()
	RouterStudent().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

func TestInvalidStudentIDEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/student/1234", nil)
	response := httptest.NewRecorder()
	RouterStudent().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}

func TestValidIDUpdateStudentEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	student := &models.Students{
		FirstName:      "Steve",
		LastName:       "Jobs",
		EmailID:        "stevejobs2@school.com",
		Password:       "student",
		Telephone:      "55555",
		ClassSectionID: 22,
	}
	jsonPerson, _ := json.Marshal(student)

	request, _ := http.NewRequest("PUT", "/student/37", bytes.NewBuffer(jsonPerson))
	response := httptest.NewRecorder()
	RouterStudent().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

func TestInvalidIDUpdateStudentEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	student := &models.Students{
		FirstName:      "Steve",
		LastName:       "Jobs",
		EmailID:        "stevejobs@school.com",
		Password:       "student",
		Telephone:      "55555",
		ClassSectionID: 22,
	}
	jsonPerson, _ := json.Marshal(student)

	request, _ := http.NewRequest("PUT", "/student/379", bytes.NewBuffer(jsonPerson))
	response := httptest.NewRecorder()
	RouterStudent().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}

func TestInvalidDetailsUpdateStudentEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	student := &models.Students{
		FirstName:      "",
		LastName:       "Jobs",
		EmailID:        "stevejobsschoolcom",
		Password:       "student",
		Telephone:      "1092837465",
		ClassSectionID: 22,
	}
	jsonPerson, _ := json.Marshal(student)

	request, _ := http.NewRequest("PUT", "/student/3", bytes.NewBuffer(jsonPerson))
	response := httptest.NewRecorder()
	RouterStudent().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}

func TestValidIDDeleteStudentEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("DELETE", "/student/13", nil)
	response := httptest.NewRecorder()
	RouterStudent().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

func TestInvalidIDDeleteStudentEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("DELETE", "/student/58765", nil)
	response := httptest.NewRecorder()
	RouterStudent().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}
