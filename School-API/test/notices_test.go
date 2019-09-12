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

func RouterNotice() *mux.Router {
	router := mux.NewRouter()
	fmt.Println("http://localhost:8080")

	router.HandleFunc("/notices/{id}", controllers.GetNotice).Methods("GET")
	router.HandleFunc("/notices", controllers.GetNotices).Methods("GET")

	return router
}

func TestValidNoticesEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/notices/1", nil)
	response := httptest.NewRecorder()
	RouterNotice().ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "OK response is expected")
}

func TestInvalidNoticesEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/notices/123456", nil)
	response := httptest.NewRecorder()
	RouterNotice().ServeHTTP(response, request)
	assert.Equal(t, http.StatusNoContent, response.Code, "No response is expected")
}

func TestNoticesEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/notices", nil)
	response := httptest.NewRecorder()
	RouterNotice().ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "OK response is expected")
}
