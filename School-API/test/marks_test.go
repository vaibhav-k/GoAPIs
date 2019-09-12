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

func RouterMark() *mux.Router {
	router := mux.NewRouter()
	fmt.Println("http://localhost:8080")

	router.HandleFunc("/marks/{id}", controllers.GetMarks).Methods("GET")

	return router
}

func TestValidMarksEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/marks/59", nil)
	response := httptest.NewRecorder()
	RouterMark().ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "OK response is expected")
}

func TestInvalidMarksEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/marks/1234567890", nil)
	response := httptest.NewRecorder()
	RouterMark().ServeHTTP(response, request)
	assert.Equal(t, http.StatusNoContent, response.Code, "No response is expected")
}
