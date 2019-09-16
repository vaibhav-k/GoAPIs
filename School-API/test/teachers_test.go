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

	router.HandleFunc("/teacher/{id}", controllers.GetTeacher).Methods("GET")
	router.HandleFunc("/teachers", controllers.GetTeachers).Methods("GET")
	router.HandleFunc("/teachers", controllers.AddTeacher).Methods("POST")
	router.HandleFunc("/teacher/{id}", controllers.UpdateTeacher).Methods("PUT")
	router.HandleFunc("/teacher/{id}", controllers.DeleteTeacher).Methods("DELETE")

	return router
}

func TestTeachersEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/teachers", nil)
	response := httptest.NewRecorder()
	RouterTeacher().ServeHTTP(response, request)
	// assert.Equal(t, 200, response.Code, "OK response is expected")

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

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

func TestPostTeacherEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	teacher := &models.Teachers{
		TeacherID: 16,
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

func TestUpdateTeacherEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	teacher := &models.Teachers{
		TeacherID: 15,
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

func TestDeleteTeacherEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("DELETE", "/teacher/13", nil)
	response := httptest.NewRecorder()
	RouterTeacher().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}
