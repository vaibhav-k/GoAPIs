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

func RouterSubject() *mux.Router {
	router := mux.NewRouter()
	fmt.Println("http://localhost:8080")

	router.HandleFunc("/subjects", controllers.GetSubjects).Methods("GET")
	router.HandleFunc("/subjects", controllers.AddSubject).Methods("POST")

	return router
}

// Tests if all the sections can be viewed
func TestSubjectsEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	// Make the request
	request, _ := http.NewRequest("GET", "/subjects", nil)
	response := httptest.NewRecorder()
	RouterSubject().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	// Check if what we wanted is what we got
	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

// Tests if a section can be added with valid details
func TestValidPostSubjectEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	subject := &models.SubjectAdd{
		SubjectID: 250,
		Title:     "Geometry",
		Classes:   []int{1, 2, 3, 4},
	}
	jsonSubject, _ := json.Marshal(&subject)

	// Make the request
	request, _ := http.NewRequest("POST", "/subjects", bytes.NewBuffer(jsonSubject))
	response := httptest.NewRecorder()
	RouterSubject().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	// Check if what we wanted is what we got
	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}

// Tests if a section can be added with invalid details
func TestInvalidTitlePostSubjectEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	subject := &models.SubjectAdd{
		SubjectID: 250,
		Title:     "",
		Classes:   []int{1, 2, 3, 4},
	}
	jsonSubject, _ := json.Marshal(&subject)

	// Make the request
	request, _ := http.NewRequest("POST", "/subjects", bytes.NewBuffer(jsonSubject))
	response := httptest.NewRecorder()
	RouterSubject().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	// Check if what we wanted is what we got
	assert.Equal(t, float64(http.StatusNoContent), resp["status_code"], "No content response is expected")
}
