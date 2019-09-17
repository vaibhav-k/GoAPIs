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
	router.HandleFunc("/student/{studentID}", controllers.GetStudent).Methods("GET")
	router.HandleFunc("/student/{studentID}", controllers.UpdateStudent).Methods("PUT")
	router.HandleFunc("/student/{studentID}", controllers.DeleteStudent).Methods("DELETE")

	return router
}

// Tests if all the students can be viewed
func TestStudentsEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	// Make the request
	request, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	RouterStudent().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	// Check if what we wanted is what we got
	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

// Tests if all the student whose valid ID is given can be viewed
func TestValidStudentIDEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	// Make the request
	request, _ := http.NewRequest("GET", "/student/1", nil)
	response := httptest.NewRecorder()
	RouterStudent().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	// Check if what we wanted is what we got
	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

// Tests if all the student whose invalid ID is given can be viewed
func TestInvalidStudentIDEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/student/1234", nil)
	response := httptest.NewRecorder()
	RouterStudent().ServeHTTP(response, request)

	// Make the request
	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	// Check if what we wanted is what we got
	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}

// Tests if all the student whose valid ID is given can be updated with the correct ID
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

	// Make the request
	request, _ := http.NewRequest("PUT", "/student/37", bytes.NewBuffer(jsonPerson))
	response := httptest.NewRecorder()
	RouterStudent().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	// Check if what we wanted is what we got
	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

// Tests if all the student whose valid ID is given can be updated with the wrong ID
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

	// Make the request
	request, _ := http.NewRequest("PUT", "/student/379", bytes.NewBuffer(jsonPerson))
	response := httptest.NewRecorder()
	RouterStudent().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	// Check if what we wanted is what we got
	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}

// Tests if all the student whose valid ID is given can be updated with the wrong first name
func TestInvalidFirstNameUpdateStudentEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	student := &models.Students{
		FirstName:      "",
		LastName:       "Jobs",
		EmailID:        "stevejobs@school.com",
		Password:       "student",
		Telephone:      "1092837465",
		ClassSectionID: 22,
	}
	jsonPerson, _ := json.Marshal(student)

	// Make the request
	request, _ := http.NewRequest("PUT", "/student/3", bytes.NewBuffer(jsonPerson))
	response := httptest.NewRecorder()
	RouterStudent().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	// Check if what we wanted is what we got
	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}

// Tests if all the student whose valid ID is given can be updated with the wrong email ID format
func TestInvalidEmailIDUpdateStudentEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	student := &models.Students{
		FirstName:      "Steve",
		LastName:       "Jobs",
		EmailID:        "stevejobsschoolcom",
		Password:       "student",
		Telephone:      "1092837465",
		ClassSectionID: 22,
	}
	jsonPerson, _ := json.Marshal(student)

	// Make the request
	request, _ := http.NewRequest("PUT", "/student/3", bytes.NewBuffer(jsonPerson))
	response := httptest.NewRecorder()
	RouterStudent().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	// Check if what we wanted is what we got
	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}

// Tests if all the student whose valid ID is given can be updated with the wrong password
func TestInvalidPasswordUpdateStudentEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	student := &models.Students{
		FirstName:      "Steve",
		LastName:       "Jobs",
		EmailID:        "stevejobs@school.com",
		Password:       "",
		Telephone:      "1092837465",
		ClassSectionID: 22,
	}
	jsonPerson, _ := json.Marshal(student)

	// Make the request
	request, _ := http.NewRequest("PUT", "/student/3", bytes.NewBuffer(jsonPerson))
	response := httptest.NewRecorder()
	RouterStudent().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	// Check if what we wanted is what we got
	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}

// Tests if all the student whose valid ID is given can be updated with the wrong telephone
func TestInvalidTelephoneUpdateStudentEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	student := &models.Students{
		FirstName:      "Steve",
		LastName:       "Jobs",
		EmailID:        "stevejobs@school.com",
		Password:       "student",
		Telephone:      "12345",
		ClassSectionID: 22,
	}
	jsonPerson, _ := json.Marshal(student)

	// Make the request
	request, _ := http.NewRequest("PUT", "/student/3", bytes.NewBuffer(jsonPerson))
	response := httptest.NewRecorder()
	RouterStudent().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	// Check if what we wanted is what we got
	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}

// Tests if all the student whose correct ID can be deleted
func TestValidIDDeleteStudentEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	// Make the request
	request, _ := http.NewRequest("DELETE", "/student/13", nil)
	response := httptest.NewRecorder()
	RouterStudent().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	// Check if what we wanted is what we got
	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

// Tests if all the student whose incorrect ID can be deleted
func TestInvalidIDDeleteStudentEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	// Make the request
	request, _ := http.NewRequest("DELETE", "/student/58765", nil)
	response := httptest.NewRecorder()
	RouterStudent().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	// Check if what we wanted is what we got
	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}
