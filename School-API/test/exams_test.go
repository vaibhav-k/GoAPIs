package test

import (
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

func RouterExam() *mux.Router {
	router := mux.NewRouter()
	fmt.Println("http://localhost:8080")

	router.HandleFunc("/exams/{id}", controllers.GetExam).Methods("GET")
	router.HandleFunc("/exams", controllers.GetExams).Methods("GET")
	return router
}
func TestExamsEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/exams", nil)
	response := httptest.NewRecorder()
	RouterExam().ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "OK response is expected")
}

func TestValidExamIDEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/exams/1", nil)
	response := httptest.NewRecorder()
	RouterExam().ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "Content response is expected")
}

func TestInvalidExamIDEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/exams/1234567890", nil)
	response := httptest.NewRecorder()
	RouterExam().ServeHTTP(response, request)
	assert.Equal(t, http.StatusNoContent, response.Code, "No content response is expected")
}
