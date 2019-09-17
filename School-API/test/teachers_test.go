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

func RouterTeacher() *mux.Router {
	router := mux.NewRouter()
	fmt.Println("http://localhost:8080")

	router.HandleFunc("/teacher/{teacherID}", controllers.GetTeacher).Methods("GET")
	router.HandleFunc("/teachers", controllers.GetTeachers).Methods("GET")
	router.HandleFunc("/teachers", controllers.AddTeacher).Methods("POST")
	router.HandleFunc("/teacher/{teacherID}", controllers.UpdateTeacher).Methods("PUT")
	router.HandleFunc("/teacher/{teacherID}", controllers.DeleteTeacher).Methods("DELETE")

	return router
}

// Tests if all the teacher details can be viewed
func TestTeachersEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/teachers", nil)
	response := httptest.NewRecorder()
	RouterTeacher().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

// Tests if the teacher details can be viewed whose valid ID is provided
func TestValidTeacherIDEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/teacher/1", nil)
	response := httptest.NewRecorder()
	RouterTeacher().ServeHTTP(response, request)
	// assert.Equal(t, http.StatusOK, response.Code, "OK response is expected")

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

// Tests if the teacher details can be viewed whose invalid ID is provided
func TestInvalidTeacherIDEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/teacher/1234", nil)
	response := httptest.NewRecorder()
	RouterTeacher().ServeHTTP(response, request)
	// assert.Equal(t, http.StatusNoContent, response.Code, "No content response is expected")

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}

// Tests if the teacher details can be added whose valid details are provided
func TestPostTeacherEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	teacher := &models.Teachers{
		FirstName: "Bigg",
		LastName:  "Boss",
		EmailID:   "biggboss@sony.com",
		Password:  "teacher",
	}
	jsonPerson, _ := json.Marshal(teacher)

	request, _ := http.NewRequest("POST", "/teachers", bytes.NewBuffer(jsonPerson))
	response := httptest.NewRecorder()
	RouterTeacher().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

// Tests if the teacher details can be updated whose valid details are provided
func TestValidIDUpdateTeacherEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	teacher := &models.Teachers{
		FirstName: "Poorest",
		LastName:  "Man",
		EmailID:   "poorestman@teaching.com",
		Password:  "teacher",
	}
	jsonPerson, _ := json.Marshal(teacher)

	request, _ := http.NewRequest("PUT", "/teacher/11", bytes.NewBuffer(jsonPerson))
	response := httptest.NewRecorder()
	RouterTeacher().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

// Tests if the teacher details can be updated whose invalid ID is provided
func TestInvalidIDUpdateTeacherEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	teacher := &models.Teachers{
		FirstName: "Poorest",
		LastName:  "Man",
		EmailID:   "poorestman@teaching.com",
		Password:  "teacher",
	}
	jsonPerson, _ := json.Marshal(teacher)

	request, _ := http.NewRequest("PUT", "/teacher/111", bytes.NewBuffer(jsonPerson))
	response := httptest.NewRecorder()
	RouterTeacher().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}

// Tests if the teacher details can be updated whose invalid first name is provided
func TestInvalidFirstNameUpdateTeacherEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	teacher := &models.Teachers{
		FirstName: "",
		LastName:  "Gates",
		EmailID:   "billgatesschoolcom",
		Password:  "teacher",
	}
	jsonPerson, _ := json.Marshal(teacher)

	request, _ := http.NewRequest("PUT", "/teacher/3", bytes.NewBuffer(jsonPerson))
	response := httptest.NewRecorder()
	RouterTeacher().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}

// Tests if the teacher details can be updated whose invalid email ID is provided
func TestInvalidEmailIDUpdateTeacherEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	teacher := &models.Teachers{
		FirstName: "Bill",
		LastName:  "Gates",
		EmailID:   "billgatesschoolcom",
		Password:  "teacher",
	}
	jsonPerson, _ := json.Marshal(teacher)

	request, _ := http.NewRequest("PUT", "/teacher/3", bytes.NewBuffer(jsonPerson))
	response := httptest.NewRecorder()
	RouterTeacher().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}

// Tests if the teacher details can be updated whose invalid password is provided
func TestInvalidPasswordUpdateTeacherEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	teacher := &models.Teachers{
		FirstName: "Bill",
		LastName:  "Gates",
		EmailID:   "billgates@school.com",
		Password:  "",
	}
	jsonPerson, _ := json.Marshal(teacher)

	request, _ := http.NewRequest("PUT", "/teacher/3", bytes.NewBuffer(jsonPerson))
	response := httptest.NewRecorder()
	RouterTeacher().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}

// Tests if the teacher details can be deleted whose valid ID is provided
func TestValidIDDeleteTeacherEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("DELETE", "/teacher/13", nil)
	response := httptest.NewRecorder()
	RouterTeacher().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

// Tests if the teacher details can be deleted whose invalid ID is provided
func TestInvalidIDDeleteTeacherEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("DELETE", "/teacher/133", nil)
	response := httptest.NewRecorder()
	RouterTeacher().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}
