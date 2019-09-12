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

func RouterStudent() *mux.Router {
	router := mux.NewRouter()
	fmt.Println("http://localhost:8080")

	router.HandleFunc("/students", controllers.GetStudents).Methods("GET")
	router.HandleFunc("/students/{id}", controllers.GetStudent).Methods("GET")

	return router
}

func TestStudentsEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	RouterStudent().ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "OK response is expected")
}
