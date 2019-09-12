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

func RouterAttd() *mux.Router {
	router := mux.NewRouter()
	fmt.Println("http://localhost:8080")

	router.HandleFunc("/attendance", controllers.GetAttendance).Methods("GET")
	router.HandleFunc("/attendances", controllers.GetAttendances).Methods("GET")

	return router
}

func TestValidAttendancesEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/attendance?date=02-Jan-20", nil)
	response := httptest.NewRecorder()
	RouterAttd().ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "OK response is expected")
}

func TestInvalidAttendancesEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/attendance?date=200-Jan-20", nil)
	response := httptest.NewRecorder()
	RouterAttd().ServeHTTP(response, request)
	assert.Equal(t, http.StatusNoContent, response.Code, "No content response is expected")
}

func TestAttendancesEndpoint(t *testing.T) {
	// Initialize the database connection
	models.InitDB()

	request, _ := http.NewRequest("GET", "/attendances", nil)
	response := httptest.NewRecorder()
	RouterAttd().ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "OK response is expected")
}
