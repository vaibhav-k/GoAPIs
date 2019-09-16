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

func TestSubjectsEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/subjects", nil)
	response := httptest.NewRecorder()
	RouterSubject().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestPostSubjectEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	subject := &models.SubjectAdd{
		SubjectID: 250,
		Title:     "Geometry",
		Classes:   []int{1, 2, 3, 4},
	}
	jsonSubject, _ := json.Marshal(&subject)

	request, _ := http.NewRequest("POST", "/subjects", bytes.NewBuffer(jsonSubject))
	response := httptest.NewRecorder()
	RouterSubject().ServeHTTP(response, request)

	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)

	assert.Equal(t, float64(http.StatusOK), resp["status_code"], "OK response is expected")
}
